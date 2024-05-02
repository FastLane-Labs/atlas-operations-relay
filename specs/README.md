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

## Error codes

### 1xxx (Relay's internal error)

**10xx** **(Relay's internal error):**
- 1000: server internal error

### 2xxx (Operation related error)

**20xx** **(User operation error):**
- 2001: user operation has invalid signature
- 2002: user operation's 'to' field must be atlas contract address
- 2003: user operation's deadline exceeded
- 2004: failed to compute user operation hash
- 2005: failed to compute user proof hash
- 2006: user operation has invalid signature
- 2007: user operation's gas limit exceeded

**21xx** **(Solver operation error):**
- 2100: solver operation's 'to' field must be atlas contract address
- 2101: solver operation's gas limit exceeded
- 2102: solver operation's maxFeePerGas must be equal or higher the user operation
- 2103: solver operation's deadline exceeded or lower than user operation's
- 2104: solver operation's dApp control does not match the user operation's
- 2105: failed to compute solver proof hash
- 2106: solver operation has invalid signature
- 2107: failed to compute solver operation hash

**22xx** **(DApp operation error):**
- 2200: dApp operation's 'from' field does not match user operation's session key
- 2201: dApp operation's 'to' field must be atlas contract address
- 2202: dApp operation's deadline exceeded or lower than user operation's
- 2203: dApp operation's dApp control does not match the user operation's
- 2204: dApp operation's user operation hash does not match the user operation's
- 2205: dApp operation's call chain hash is invalid
- 2206: failed to compute dApp proof hash
- 2207: dApp operation has invalid signature
- 2208: dApp operation's gas limit exceeded

### 3xxx (API and server related errors)

**30xx** **(API error):**
- 3000: malformed request
- 3001: malformed json
- 3002: invalid parameter
- 3003: server corrupted data
- 3004: invalid operation hash
- 3005: invalid address
- 3006: invalid timestamp
- 3007: expired signature
- 3008: bad signature (decode/recover error)
- 3009: signature mismatch

**31xx** **(Relay error):**
- 3100: failed to forward bundle
- 3101: failed to get dapp signatories
- 3102: failed to get atlEth bonded balance
- 3103: relay is offline

**32xx** **(Server error):**
- 3200: bundler is offline
- 3201: clogged connection
- 3202: bundling failure

### 4xxx (Auction related error)

**40xx** **(Auction validation error):**
- 4000: auction for this user operation has already started
- 4001: auction not found
- 4002: user operation failed simulation
- 4003: solver operation failed simulation
- 4004: not enough atlEth bonded balance

**41xx** **(Auction other error):**
- 4100: auction for this user operation has already ended
- 4101: auction for this user operation is ongoing
- 4102: solver is already participating in this auction
- 4103: solver operation not found

### 5xxx (Bundle related error)

**50xx** **(Bundle validation error):**
- 5000: bundle for this user operation has already been submitted
- 5001: bundle not found
- 5002: bundle failed simulation

**51xx** **(Bundle other error):**
- 5100: bundle has not been bundled yet
- 5101: bundle has already been bundled
- 5102: bundle has already errored
