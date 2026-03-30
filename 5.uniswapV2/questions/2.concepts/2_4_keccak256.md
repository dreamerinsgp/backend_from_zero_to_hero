# Q1：What is `keccak256`?

## Short answer

**`keccak256`** is a **cryptographic hash function** (Ethereum uses the **Keccak-256** variant). In Solidity it appears as a **global function**:

- **Input:** arbitrary `bytes` (often produced with `abi.encode`, `abi.encodePacked`, etc.).
- **Output:** a fixed **32-byte** value (`bytes32`) that looks random but is **deterministic**: same input ⇒ same hash.

It is **one-way** (you cannot recover the input from the hash) and **collision-resistant** in practice. On-chain it is cheap enough to use everywhere you need a **compact fingerprint** or **unpredictable-looking** fixed-size ID.

---

## In Uniswap V2 Factory (`salt`)

```29:30:contracts/UniswapV2Factory.sol
        bytes32 salt = keccak256(abi.encodePacked(token0, token1));
        assembly {
```

Here **`keccak256(abi.encodePacked(token0, token1))`** builds the **`CREATE2` `salt`**:

- **`token0`** and **`token1`** are **20-byte addresses** (canonical sorted order).
- **`abi.encodePacked`** concatenates them **without** extra padding between elements (unlike `abi.encode`), giving a **compact** 40-byte preimage.
- **`keccak256`** maps that preimage to a **single `bytes32` salt**.

Because the salt is **fully determined** by `(token0, token1)`, the **`CREATE2`** address for the Pair is **deterministic** for that token pair (given the same Factory, same Pair bytecode). See [2_2_deterministic_pair_deployment.md](2_2_deterministic_pair_deployment.md).

---

## Related uses in Ethereum (context)

| Use | Idea |
|-----|------|
| **CREATE2 salt** | Uniswap: bind deploy address to `(token0, token1)`. |
| **Storage keys** | e.g. `mapping` slots use keccak in the EVM’s layout rules. |
| **Commitments** | hide a value until reveal (hash the secret). |

---

## Related

- [2_3_createpool_logic.md](2_3_createpool_logic.md)  
- [milestone2.md](milestone2.md)
