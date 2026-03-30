# Q1：What is **storage**?

In the EVM, **storage** is each contract’s **persistent** key–value store: 2²⁵⁶ slots of **32-byte** words, keyed by slot index. Values **survive** after the transaction ends and are part of **chain state**.

- **Reads/writes cost gas** (storage is expensive).
- **Solidity state variables** (`token0`, `token1`, `reserve0`, mappings, etc.) live in **storage** unless declared with other data locations.

So when we say `initialize` “stores the token pair,” we mean it writes **`token0` and `token1` into that Pair contract’s storage slots** — permanent until overwritten.

---

# Q2：Difference between **contract storage** and **stack / memory** when executing `CREATE2`?

The EVM separates **where data lives** and **how long it lasts**:

| Concept | What it is | Persists after tx? |
|--------|------------|---------------------|
| **Stack** | Small stack of **32-byte words** used by opcodes during execution. | No — gone when the call frame ends. |
| **Memory** | A **byte array** (`memory` in Solidity) used during a single execution (e.g. holding the **init bytecode** slice for `CREATE2`). | No — **volatile**; only for this message call. |
| **Storage** | Each **contract address** has its own persistent slots. | **Yes** — this is “on-chain state.” |

### During `create2(0, add(bytecode, 32), mload(bytecode), salt)` in the Factory

- **`bytecode`** lives in the Factory’s **memory** (temporary): the compiler placed `type(UniswapV2Pair).creationCode` there.
- **`CREATE2`** reads that **memory range** to get **`init_code`**, deploys a **new contract**, and returns the **new address**.
- The **Factory’s** stack holds operands like **`salt`** and intermediate values — all **ephemeral**.

### After deployment

- The **new Pair** has its **own storage** (initially: whatever the **constructor** wrote, e.g. `factory = msg.sender`).
- **`initialize`** then writes **`token0` / `token1`** into **that Pair’s storage** — not into Factory memory from the `CREATE2` line.

So: **`CREATE2` uses memory (and stack) only to perform deployment.** **`initialize` uses the Pair’s storage** to record the tokens. They are **different layers**; nothing in “Pair token addresses” is stored in the **Factory’s** `bytecode` memory slice after the call finishes.

---

## Related

- [2_8_initialize.md](2_8_initialize.md)
