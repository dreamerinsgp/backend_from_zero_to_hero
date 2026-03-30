# Milestone 5 — TWAP and protocol fee

[← Milestone 4](milestone4.md) · [Roadmap](roadmap.md) · [Next: Milestone 6 →](milestone6.md)

## Core concepts in this milestone

### E. Oracle / TWAP (full)

| Concept | Notes |
|--------|--------|
| **Time-weighted cumulative price** | `price0CumulativeLast`, `price1CumulativeLast` — integrals of price over time for off-chain TWAP. |
| **Fixed-point math** | [`UQ112x112.sol`](../../contracts/libraries/UQ112x112.sol) — Q112.112 format for price ratios. |
| **Per-block update** | `_update` accumulates prices when time elapses and reserves are non-zero; uses `blockTimestampLast`. |

### A. AMM economics (subset)

| Concept | Notes |
|--------|--------|
| **Protocol fee on growth** | `_mintFee`: when `feeTo != 0`, mint LP to `feeTo` proportional to \(\sqrt{k}\) growth vs `kLast`. |

### Factory + Pair linkage

| Concept | Notes |
|--------|--------|
| **`feeTo` / `kLast`** | Pair reads `IUniswapV2Factory(factory).feeTo()`; `kLast` tracks \(k\) after liquidity operations when fee switch is on. |

## Read / do

- **Read:** `_update` (TWAP branch), `_mintFee`, and `kLast` handling in `mint` / `burn` / `swap` in [`UniswapV2Pair.sol`](../../contracts/UniswapV2Pair.sol); [`UQ112x112.sol`](../../contracts/libraries/UQ112x112.sol); Factory `feeTo` in [`UniswapV2Factory.sol`](../../contracts/UniswapV2Factory.sol).

## Related inventory sections

- **E** (full).
- **A** (partial): protocol fee.
- **B** (partial): `feeTo` as governance output.
