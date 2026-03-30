# Milestone 2 — Factory and pair lifecycle

[← Milestone 1](milestone1.md) · [Roadmap](roadmap.md) · [Next: Milestone 3 →](milestone3.md)

## Core concepts in this milestone

From the [roadmap](roadmap.md) inventory:

### B. Factory layer (full)

| Concept | Notes |
|--------|--------|
| **Deterministic pair deployment** | `CREATE2` with bytecode + `salt = keccak256(abi.encodePacked(token0, token1))` — same `(token0, token1)` always yields same Pair address. |
| **`token0` / `token1` ordering** | `token0 < token1` by address (`tokenA < tokenB` in Factory) — every pair and API uses this order. |
| **Registry** | `getPair(token0, token1)` (both directions), `allPairs[]`, `allPairsLength`, `PairCreated` event. |
| **Governance** | `feeTo`, `feeToSetter`; `setFeeTo`, `setFeeToSetter` (only `feeToSetter`). |
| **One-time pair init** | After deploy, Factory calls `pair.initialize(token0, token1)` — only Factory may call `initialize` (checked in Pair). |

### Pair bootstrapping (minimal)

| Concept | Notes |
|--------|--------|
| **Pair `constructor`** | Sets `factory = msg.sender` (the Factory). |
| **`initialize`** | Writes `token0`, `token1`; callable only by `factory`. |

## Read / do

- **Read:** [`UniswapV2Factory.sol`](../../contracts/UniswapV2Factory.sol), `initialize` in [`UniswapV2Pair.sol`](../../contracts/UniswapV2Pair.sol).
- **Do (optional):** `yarn deploy`, `yarn createpair` — [test createPair](../1.environment/5.test_createPair.md).

## Related inventory sections

- **B** only (plus Pair constructor / `initialize`). Sections **A, C–H** come in later milestones.
