# Milestone 6 — Flash swap, skim, sync, reentrancy

[← Milestone 5](milestone5.md) · [Roadmap](roadmap.md) · [Next: Milestone 7 →](milestone7.md)

## Core concepts in this milestone

### D. Pair: trading (flash)

| Concept | Notes |
|--------|--------|
| **Flash swap** | If `data.length > 0`, calls `IUniswapV2Callee(to).uniswapV2Call(...)` — callee must return owed tokens to Pair before `swap` ends. |

### F. Safety and hygiene (full)

| Concept | Notes |
|--------|--------|
| **Reentrancy guard** | `lock` modifier — `unlocked` flag around external calls. |
| **Safe ERC20 transfer** | `_safeTransfer` — handles non-standard ERC20 return values. |
| **`skim` / `sync`** | `skim`: send excess balance over reserves to `to`; `sync`: set reserves = balances (oracle / rescue use cases). |

## Read / do

- **Read:** `swap` callback branch; `skim`, `sync`, `lock`, `_safeTransfer` in [`UniswapV2Pair.sol`](../../contracts/UniswapV2Pair.sol); [`IUniswapV2Callee`](../../contracts/interfaces/IUniswapV2Callee.sol).

## Related inventory sections

- **D** (flash swap remainder).
- **F** (full).
