# Milestone 7 — LP token + periphery (outside this repo)

[← Milestone 6](milestone6.md) · [Roadmap](roadmap.md)

## Core concepts in this milestone

### G. LP token (ERC20 + Permit) (full)

| Concept | Notes |
|--------|--------|
| **LP as ERC20** | [`UniswapV2ERC20.sol`](../../contracts/UniswapV2ERC20.sol) — `name`, `symbol`, `decimals`, transfers, allowances. |
| **Permit (EIP-2612 style)** | `permit(owner, spender, value, deadline, v, r, s)` — gasless approval via signature; `DOMAIN_SEPARATOR`, `PERMIT_TYPEHASH`, `nonces`. |

### Beyond v2-core: Periphery

| Concept | Notes |
|--------|--------|
| **Router** | `UniswapV2Router02` — atomic `addLiquidity`, `removeLiquidity`, `swapExact*`, multi-hop paths, ETH ↔ WETH, deadlines. |
| **Why it matters** | Core `mint`/`burn`/`swap` are low-level; production UIs use Router for safety and UX. |

## Read / do

- **Read:** [`UniswapV2ERC20.sol`](../../contracts/UniswapV2ERC20.sol); then Uniswap **v2-periphery** source or official docs for `UniswapV2Router02`.

## Related inventory sections

- **G** (full).
- Periphery is **not** in section **A–H** of [roadmap](roadmap.md) — it is the external stack on top of core.
