# Atlas Operations Relay Specifications

## Overview

The operations relay for Atlas serves as an infrastructure layer facilitating communication between the auctioneer, solvers, and bundlers. The following specifications must be strictly implemented by the relay operator.

## APIs

2 APIs are exposed by the operations relay, please see the following documents for technical specifications.
- [REST API](./rest-api.yaml)
- [Websocket API](./ws-api.yaml)

## Lifecycle

1. The frontend calls `/userOperation` endpoint to submit a new user operation. The endpoint returns the `userOpHash` in case of success, or an error.
2. The relay broadcasts the user operation to solvers via websocket (solvers connect to the relay and subscribe to the `newUserOperations` topic). When the operation is broadcast, the relay must also start the auction timer for this operation.
3. Solvers send solver operations to the relay in response to the user operation (via REST or websocket API).
4. The relay stops collecting solver operations after the auction timer hits 500ms.
5. The frontend calls the `/solverOperations` endpoint, specifying the `userOpHash` as a parameter, to retrieve the solver operations tied to a particular user operation. If it calls this endpoint before the end of the auction duration (500ms), the endpoint returns an error. If the frontend sets the `wait` parameter to `true` when calling, the relay will hold the request until the auction’s completion then return the results.
6. The frontend calls `/bundleOperations` endpoint with the user, solver, and dApp operations. The endpoint returns a success message in case of success, or an error.
7. The relay forwards the data submitted to `/bundleOperations` to the appointed bundler via websocket (bundlers connect and authenticate themselves to the relay via a dedicated websocket channel).
8. The bundler generates an Atlas transaction, propagates it to the block producer, and sends back the transaction hash to the relay (it does not wait for the transaction to be included in a block).
9. If the bundler does not respond with a pending transaction hash after some specified period, the relay should send the the operations to one or more fallback bundlers. A list of these permitted bundlers can be queried from the Atlas contract.
10. The relay stocks the transaction hash, ready to serve a future call from the frontend.
11. The frontend calls `/bundleHash` endpoint, specifying the `userOpHash` in parameters to retrieve the Atlas transaction hash tied to a particular user operation. If it calls before the relay gets the transaction hash back from the bundler, the endpoint returns an error. If the frontend set the `wait` parameter to `true` when calling, the relay will hold the request until it gets the transaction hash back from the bundler, then return the result.

## Verification

### Static checks

When frontend submits a **user** **operation**, the relay should make the following checks, before simulating the operation:
- `To` field must be Atlas address
- `Gas` field should be <= 1,000,000 (default value, can be configured by the operator)
- `Deadline` should be higher than current block
- `Signature` field must be valid

When **solvers** submit their operations, the relay should make the following checks, before simulating the operations:
- `To` field must be Atlas address
- `Gas` field should be <= 1,000,000 (default value, can be configured by the operator)
- `MaxFeePerGas` field should be equal or higher than **userOp.maxFeePerGas**
- `Deadline` field should be equal or higher than **userOp.deadline**
- `Control` field should be equal to **userpOp.control**
- `UserOpHash` should match an existing **userOpHash** previously submitted
- `Signature` field must be valid

When frontend submits a bundle, the relay should re-run validations on the user operation, and make the following checks on the **dApp** **operation**, before simulating the bundle:
- `From` field should be equal to **userOp.sessionKey** only if **userOp.sessionKey** is set
- `To` field must be Atlas address
- `Gas` field should be <= 1,000,000 (default value, can be configured by the operator)
- `Deadline` field should be equal or higher than **userOp.deadline**
- `Control` field should be equal to **userOp.control**
- `UserOpHash` field should be the correct hash of the **userOp**
- `CallChainHash` field should be the correct hash of the **callChain**
- `Signature` field must be valid

### Simulations

Note that all simulations referenced here can be easily done by calling view functions on the Atlas entrypoint contract using the inputs from users and solvers.

- When a user operation is submitted it should be simulated. The result from the Atlas entrypoint will tell you if this operation successfully passes the userOp phase of simulation.
- When the relay receives a userOp, it should also generate a mock dapp operation that is associated with it.
- The relay must individually simulate each solver operation it receives using the associated mock dapp operation. The relay must read the bonded atlETH (escrowed gas) balance of the solver from the Atlas contract, and verify that the solver has 3x more atlETH than the simulated gas cost in ETH.
- When a bundle is submitted, it should be simulated before forwarding it to the appointed bundler.
