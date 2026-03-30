# Deterministic pair deployment (Uniswap V2 Factory)

## Q1：What can we learn from deterministic pair deployment?

**Interview angle:** It is not just a trick—it encodes **how Uniswap scales integrations**.


| Lesson                      | Explanation                                                                                                                                                                                                                              |
| --------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Predictable addresses**   | Before any `createPair` tx, anyone can **compute** the Pair address from **factory address + token addresses + init code** (same chain). UIs and routers do not need a centralized “pair registry” API to know where the pool will live. |
| **One pool per pair**       | Factory enforces a **canonical ordering** (`token0 < token1`) and a **single** deployment per `(token0, token1)`. Determinism + uniqueness go together: the **salt** is derived only from those two addresses.                           |
| **Composable integrations** | Other contracts (aggregators, routers, analytics) can **derive** `pair` off-chain or in-contract using the same rule, then call `swap` / `mint` on that address.                                                                         |
| **Cross-ecosystem pattern** | **CREATE2 + init code hash + salt** is the standard Ethereum pattern for **counterfactual** or **precomputable** contract addresses (Uniswap is a prime example).                                                                        |


---

## Q2：Why is it deterministic?

**Short answer:** Because the Factory deploys the Pair with the `**CREATE2`** opcode, not ordinary `**CREATE**`.


| Factor                | Role                                                                                                                                                                                         |
| --------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `**CREATE2**`         | Deployed address depends only on `**msg.sender` (Factory)**, `**salt`**, and `**keccak256(init_code)**`—not on **nonce**. So the address is **fixed** for a given (factory, salt, bytecode). |
| **Fixed bytecode**    | `type(UniswapV2Pair).creationCode` is the **same** for every Pair deployment (same compiler output / implementation).                                                                        |
| **Fixed salt scheme** | `salt = keccak256(abi.encodePacked(token0, token1))` — **one** salt per **unordered** token pair (after sorting into `token0`, `token1`).                                                    |


With `**CREATE`**, address depends on deployer **nonce**, so you could not predict the Pair address from tokens alone. `**CREATE2`** removes that randomness.

**Code anchor:** `[UniswapV2Factory.sol](../../contracts/UniswapV2Factory.sol)` — `createPair` uses `type(UniswapV2Pair).creationCode` and `create2(..., salt)`.

---

## Q3：How is this feature guaranteed (what must hold)?

**On-chain (protocol guarantees):**

1. **Always deploy via this Factory’s `createPair`** — the only path that uses the agreed `**CREATE2` + salt + Pair init code** for that deployment.
2. **Canonical token order** — `token0` / `token1` are **sorted** before `salt` and deployment, so the same two tokens **always** yield the **same** salt and **one** pool.
3. **No duplicate** — `getPair[token0][token1] == address(0)` before deploy, so the deterministic address is **not** reused for a second pool.
4. **Immutable Pair logic** — all Pairs share the **same** `creationCode` from the **same** `UniswapV2Pair` contract definition at Factory compile time.

**Off-chain / integrator side (if you recompute the address):**

- Use the **same** Factory address, **same** chain, **same** `UniswapV2Pair` **init code hash** (and the standard **CREATE2** address formula) and **same** `salt = keccak256(abi.encodePacked(token0, token1))`.  
- If any of these differ (wrong factory, wrong bytecode version, wrong token order), the **computed** address will **not** match the real Pair.

**What does *not* guarantee determinism across chains:** the **same** `(token0, token1)` yields the **same** Pair **only if** the **Factory address** and **bytecode / init code hash** match that chain’s deployment. Tokens are often **different addresses** per chain anyway.

---

## Related

- [milestone2.md](milestone2.md)  
- [2_1_purpose.md](2_1_purpose.md)

