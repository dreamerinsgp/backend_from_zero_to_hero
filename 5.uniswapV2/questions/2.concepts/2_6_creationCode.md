
On **Etherscan**, what you see under **“Logs”** for a transaction are **EVM logs** — almost always from **`emit` in Solidity** (events). Things that **do not** show up there:

- **`console.log` / Foundry console** — only in local test output, not on-chain.
- **`log` in assembly** — low-level `LOG` opcodes also produce logs, but Solidity `emit` is the normal way.

### What to do so something appears on Etherscan

1. **Define and emit an event** in your contract (e.g. temporary debug in `createPair`):

```solidity
event DebugPairBytecode(uint256 codeLength, bytes32 codeHash);

emit DebugPairBytecode(bytecode.length, keccak256(bytecode));
```

2. **Deploy / call** the contract so that line runs in a real transaction.

3. Open the **transaction hash** on Etherscan → tab **“Logs”** (or **“Event Logs”**).

4. You’ll see topics/data for your event. If the contract is **verified**, Etherscan usually **decodes** the event name and fields. If not verified, you still see raw topics/data and may need to decode manually.

### Tips

- Prefer **small** payloads: full `bytecode` in an event is huge and awkward; **length + `keccak256(bytecode)`** is enough to confirm which artifact you used.
- **`indexed`** parameters (up to 3) become topics and are easier to filter; non-indexed go in **data**.

**Bottom line:** use **`emit`** (events). That is the standard way to get **visible logs on Etherscan** for a successful transaction.