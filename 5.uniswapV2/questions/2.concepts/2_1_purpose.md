# Q1：What is the purpose of the Factory contract in Uniswap V2?

## Short answer (interview ~30 seconds)

The **Factory** is the **canonical deployer and registry** for Uniswap V2 **Pairs**. It is the **single place** that:

1. **Creates** a new **`UniswapV2Pair`** for a given **two ERC20s** (with **one pool per unordered pair** — no duplicates).
2. **Deploys pairs deterministically** via **`CREATE2`**, so everyone agrees on the **Pair address** for `(token0, token1)` without an off-chain database.
3. **Initializes** each Pair (`initialize(token0, token1)`) and **records** it in **`getPair`** and **`allPairs`**.
4. **Holds protocol-level governance** for the **protocol fee switch**: **`feeTo`** and **`feeToSetter`** (who may receive protocol fees and who can update those settings).

So: **Factory = “where pools are born + how we find them + who can turn on protocol fees.”** The actual swapping and liquidity math live in **Pair**, not Factory.

---

## What each responsibility means

| Responsibility | Why it matters |
|----------------|----------------|
| **`createPair(tokenA, tokenB)`** | Permissionless creation of the AMM pool for two tokens; sorts into **`token0 < token1`** so order of arguments does not create two pools. |
| **No duplicate pools** | `getPair[token0][token1] == 0` before deploy — **one** Pair per token pair. |
| **`CREATE2` + salt** | Pair address = **function of (factory, token0, token1, bytecode)** — predictable for routers, UIs, and integrations. |
| **`initialize`** | Pair’s storage (`token0`, `token1`) is set **only** by Factory right after deploy (Pair enforces `msg.sender == factory` on `initialize`). |
| **`getPair` / `allPairs`** | **Discovery**: list pools and resolve address for two tokens. |
| **`feeTo` / `feeToSetter`** | Optional **protocol fee** on liquidity growth (used by Pair’s `_mintFee` when `feeTo != 0`). |

---

## What Factory does *not* do

- **No swaps, no liquidity math** — that is **`UniswapV2Pair`** (`mint`, `burn`, `swap`).
- **No user-facing “add liquidity” UX** — that is typically **Router** (periphery).

---

## Code anchor

See [`UniswapV2Factory.sol`](../../contracts/UniswapV2Factory.sol): `createPair`, `getPair`, `allPairs`, `setFeeTo`, `setFeeToSetter`.

---

## Related

- Milestone 2: [milestone2.md](milestone2.md)
