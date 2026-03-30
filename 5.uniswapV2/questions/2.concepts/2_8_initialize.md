# Q1：What is the role of `initialize` on the Pair?

## Short answer

After **`CREATE2`** deploys a new **`UniswapV2Pair`**, the Pair’s storage for **`token0`** and **`token1`** is still **empty**. **`initialize(_token0, _token1)`** runs **once** to **write which two ERC20s** this pool trades.

Only the **Factory** may call it (`msg.sender == factory`). That ties each deployed Pair to **exactly one** canonical token pair and prevents anyone else from repointing the pool to malicious tokens.

---

## Why not set tokens in the `constructor`?

Every Pair instance uses the **same** compiled bytecode (`type(UniswapV2Pair).creationCode`). The **constructor** only sets **`factory = msg.sender`** (the Factory address):

```61:63:contracts/UniswapV2Pair.sol
    constructor() public {
        factory = msg.sender;
    }
```

Token addresses **differ per pool**, but **`CREATE2` uses identical init code** for all deployments, so you **cannot** pass `(token0, token1)` into the constructor without changing that pattern. Uniswap V2 uses **two steps**: deploy (constructor runs), then **`initialize`** with the specific tokens.

---

## What `initialize` does in code

```65:70:contracts/UniswapV2Pair.sol
    // called once by the factory at time of deployment
    function initialize(address _token0, address _token1) external {
        require(msg.sender == factory, 'UniswapV2: FORBIDDEN'); // sufficient check
        token0 = _token0;
        token1 = _token1;
    }
```

- **`token0` / `token1`** — used everywhere in **`mint`**, **`burn`**, **`swap`**, **`skim`**, **`sync`**, etc.

---

## Relation to Factory `createPair`

```solidity
IUniswapV2Pair(pair).initialize(token0, token1);
getPair[token0][token1] = pair;
```

1. **`initialize`** — configures the new Pair contract.  
2. **`getPair[...] = pair`** — registers that address in the Factory (discovery), separate from initialization.

---

## Related

- [2_3_createpool_logic.md](2_3_createpool_logic.md)  
- [milestone2.md](milestone2.md)
