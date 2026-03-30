# Milestone 3 — Reserves, balances, mint, burn

[← Milestone 2](milestone2.md) · [Roadmap](roadmap.md) · [Next: Milestone 4 →](milestone4.md)

## Core concepts in this milestone

### A. AMM economics (subset)

| Concept | Notes |
|--------|--------|
| **Reserves vs ERC20 balances** | `getReserves()` vs `IERC20.balanceOf(pair)` — can differ until `_update` / `mint` / `burn` / `swap` / `sync`. |
| **Liquidity as shares** | LP `totalSupply`; `mint` / `burn` use pro-rata math vs reserves. |
| **First mint / `MINIMUM_LIQUIDITY`** | First depositor: `sqrt(amount0 * amount1) - MINIMUM_LIQUIDITY` LP to user; `MINIMUM_LIQUIDITY` locked to `address(0)` (anti-manipulation). |

### C. Pair: liquidity (full)

| Concept | Notes |
|--------|--------|
| **Transfer then mint** | `mint` uses increments `balance - reserve` for each token — no `transferFrom` from `msg.sender`. |
| **Add liquidity** | `mint(to)` — mints LP to `to`. |
| **Remove liquidity** | `burn(to)` — burns LP held by Pair, sends underlying to `to` (Router usually sends LP to Pair first). |

### H. Supporting libraries (subset)

| Concept | Notes |
|--------|--------|
| **`sqrt` for liquidity** | [`Math.sol`](../../contracts/libraries/Math.sol) — used in `mint` for first liquidity. |
| **Overflow-safe arithmetic** | [`SafeMath.sol`](../../contracts/libraries/SafeMath.sol) — used throughout Pair. |

### `_update` (intro)

| Concept | Notes |
|--------|--------|
| **Syncing reserves** | After liquidity events, `_update` writes `reserve0`, `reserve1`, `blockTimestampLast` (TWAP inputs updated here — full TWAP math in Milestone 5). |

## Read / do

- **Read:** `mint`, `burn`, `_update` in [`UniswapV2Pair.sol`](../../contracts/UniswapV2Pair.sol); [`Math.sol`](../../contracts/libraries/Math.sol).
- **Do (optional):** `yarn check:liquidity`, `yarn add:liquidity` — [Add liquidity](../1.environment/6.add_liquidity.md).

## Related inventory sections

- **A** (partial): reserves vs balances, liquidity shares, `MINIMUM_LIQUIDITY`.
- **C** (full), **H** (partial).
