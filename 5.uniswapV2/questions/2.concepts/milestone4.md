# Milestone 4 — Swap and the 0.3% fee

[← Milestone 3](milestone3.md) · [Roadmap](roadmap.md) · [Next: Milestone 5 →](milestone5.md)

## Core concepts in this milestone

### A. AMM economics (subset)


| Concept                            | Notes                                                                                                                                     |
| ---------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| **Constant product** x \cdot y = k | Enforced after fee via adjusted balances in `swap` (see K check).                                                                         |
| **0.3% swap fee**                  | Implemented via `balance{0,1}Adjusted` and `amount{0,1}In.mul(3)` — equivalent to **997/1000** on input in `getAmountOut` math off-chain. |


### D. Pair: trading (except flash — Milestone 6)


| Concept                       | Notes                                                               |
| ----------------------------- | ------------------------------------------------------------------- |
| **Optimistic output**         | Tokens sent out to `to` **before** final balance check.             |
| **K invariant with fee**      | Post-trade balances must satisfy inequality with `1000**2` scaling. |
| `**amount0In` / `amount1In`** | Derived from balance delta vs old reserves minus outputs.           |
| `**INVALID_TO**`              | `to` must not be `token0` or `token1`.                              |


### Off-chain pairing


| Concept            | Notes                                                                                                                                     |
| ------------------ | ----------------------------------------------------------------------------------------------------------------------------------------- |
| `**getAmountOut**` | \text{amountIn} \times 997 \times \text{reserveOut} / (\text{reserveIn} \times 1000 + \text{amountIn} \times 997) — matches on-chain fee. |


## Read / do

- **Read:** Full `swap` in `[UniswapV2Pair.sol](../../contracts/UniswapV2Pair.sol)`.
- **Do (optional):** [Swap](../1.environment/7.swap.md), `yarn swap`.

## Related inventory sections

- **A** (partial): constant product, 0.3% fee.
- **D** (partial): optimistic output, K check, `INVALID_TO` — **flash swap** deferred to Milestone 6.

