# Q1：What are **v2-core** and **v2-periphery**, respectively? (interview-style)

## Short answer you can give in ~30 seconds

**v2-core** is the **minimal on-chain protocol**: the Factory that deploys Pairs, each Pair (the AMM pool + LP token), and small libraries. It defines **how liquidity and swaps work** at the lowest level.

**v2-periphery** is a **separate package of helper contracts**—mainly the **Router** and **libraries**—that sit **in front of** core. It makes **users’ and integrators’ lives easier**: wrapping ETH, multi-hop swaps, deadlines, slippage checks, and batched “approve + swap” flows, without changing the core math.

---

## What interviewers often want to hear

### v2-core

- **Role:** “Source of truth” for the protocol. Anyone can call it; it’s **permissionless** at the Pair level (aside from Factory governance knobs like `feeToSetter`).
- **What’s inside (typical):**
  - **Factory** — `createPair`, `getPair`, fee governance.
  - **Pair** — `mint` / `burn` / `swap`, reserves, TWAP accumulators, protocol-fee hook.
  - **LP token** — ERC20 representing pool shares, often with `permit`.
- **Design choice:** **Small, auditable surface.** Core avoids pulling tokens from users directly in many paths; e.g. `mint` and `swap` reason about **balances vs reserves**, so **Router** handles `transferFrom` and composition.

### v2-periphery

- **Role:** **Optional convenience and safety layer.** Not required for the protocol to exist; advanced users or other contracts could interact with core only (as in minimal scripts).
- **What’s inside (typical):**
  - **Router** (`UniswapV2Router02`) — `addLiquidity`, `removeLiquidity`, `swapExactTokensForTokens`, ETH↔WETH, etc.
  - **Supporting contracts** — e.g. migrators, possibly test helpers, depending on repo version.
- **Why it exists:** Core is **low-level** (easy to misuse: wrong ratio, no deadline, non-atomic steps). Periphery **bundles** operations, enforces **deadlines** and **minimum amounts** (slippage), and uses **WETH** for native ETH.

### Relationship (one sentence)

**Core defines the AMM; periphery defines the standard way apps and users interact with it safely and ergonomically.**

---

## Follow-up angles (if they dig deeper)

| Topic | Point |
|--------|--------|
| **Deployment** | Factory address is fixed per chain; Router is also deployed separately—dApps point users at Router, which calls core. |
| **Upgrades** | Core is **not** upgraded in place by Uniswap in the wild; new versions are new deployments. Periphery can be redeployed without changing Pair math. |
| **This repo** | `uniswap-v2-core` is **only** core; Router code lives in **`uniswap-v2-periphery`**. |

---

## Related in this repo

- Scope summary: [roadmap.md](roadmap.md)  
- Milestone 7 (periphery reading): [milestone7.md](milestone7.md)
