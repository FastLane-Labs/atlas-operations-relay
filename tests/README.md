# Developer notes

All EOAs and contracts used here have been deployed/integrated/funded on Sepolia for the sole purpose of the integration tests.
Changing even one of those values will most likely break the tests.

For reference, here are the necessary components.

## EOAs

### User
Has at least 1e12 tokenB, as defined in the demo swap intent data, and has approved Atlas to spend it (permit69).

### Solver
Has bonded at least 0.5 ATLETH on Atlas.

## Contracts

### Atlas, AtlasVerification, Simulator
Those 3 contracts are dependent on each others.

### SwapIntentDAppControl
Is integrated and dependent on Atlas.

### SimpleRfqSolver
Has at least 0.01 ETH, and 200e12 tokenA, as defined in the demo swap intent.