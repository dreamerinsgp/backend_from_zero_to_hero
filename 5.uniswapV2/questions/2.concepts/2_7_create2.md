reference: https://eips.ethereum.org/EIPS/eip-1014


That line is **inline Yul** calling the **`CREATE2`** opcode. It **deploys** a new `UniswapV2Pair` contract at an address computed from **this Factory**, **`salt`**, and the **init code** (Pair creation bytecode)—exactly what EIP‑1014 describes.

### What the whole line does

- **`pair :=`** — Result of `CREATE2` is the **new contract address** (or `0` on failure). That value is assigned to Solidity’s `pair` variable.
- **`create2(...)`** — One EVM instruction: “create a contract from memory, with this salt, using CREATE2 address rule.”

### Meaning of each argument (maps to the spec)

| Argument | In your code | Meaning (EIP‑1014 stack order) |
|----------|----------------|----------------------------------|
| **1 — endowment** | `0` | **Wei** sent to the new contract on creation. Uniswap uses **`0`** (no ETH; the Pair doesn’t need a balance to be created). |
| **2 — memory_start** | `add(bytecode, 32)` | **Offset in memory** where **`init_code`** starts. `bytecode` is Solidity `bytes memory`: the **first 32 bytes** hold the **length**; the **actual bytecode** starts at **`bytecode + 32`**. So this skips the length word and points at the real init code. |
| **3 — memory_length** | `mload(bytecode)` | **Length** of `init_code` in bytes — read from the first word of that `bytes` blob (same as `bytecode.length` in Solidity). |
| **4 — salt** | `salt` | **`bytes32`** used in \(\text{keccak256}(0xff \,\|\, \text{this} \,\|\, \text{salt} \,\|\, \text{keccak256(init\_code)})\) to fix the **deployed address**. Here `salt = keccak256(abi.encodePacked(token0, token1))`. |

So: **endowment = 0**, **memory slice = the Pair creation bytecode**, **salt = per‑pair salt**.

### How this ties to the EIP text you quoted

- **“Behaves like `CREATE` except the address formula”** — Same idea as `CREATE`: run **`init_code`**, deploy runtime code at the computed address. The **difference** is only **how the address is chosen** (CREATE2 uses **factory + salt + init code hash**, not **nonce**).
- **`0xff ++ address ++ salt ++ keccak256(init_code)`** — **`address`** is the **Factory** (`msg.sender` of `CREATE2`), **`salt`** is your 32‑byte value, **`init_code`** is what’s in memory at `add(bytecode, 32)` for length `mload(bytecode)`.
- **Extra hash cost** — EVM charges for hashing **`init_code`** (and the final mix); you don’t see that in Solidity except as **gas** on the tx.

### Tiny memory detail (why `add(bytecode, 32)`)

Solidity stores **`bytes memory`** as:

- `[bytecode … bytecode+31]` → **length** (32 bytes, big-endian)
- `[bytecode+32 …]` → **raw bytes**

So **`CREATE2` must read init code from `bytecode + 32`**, not from `bytecode`.

---

**One line:** `create2(0, add(bytecode, 32), mload(bytecode), salt)` means **“deploy `UniswapV2Pair` from this bytecode in memory, send 0 wei, use this salt so the new Pair’s address is deterministic per `(token0, token1)`.”**