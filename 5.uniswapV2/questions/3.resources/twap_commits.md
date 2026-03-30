# Q1：Uniswap V2 Core 里和 TWAP / 预言机累加器相关的 commit

说明：下列来自本仓库 **`git log`** 中与 **TWAP、oracle、price cumulative、accumulator** 等关键词或实现强相关的提交（以 **upstream `uniswap-v2-core`** 历史为准；若你 fork 过，hash 应对齐官方仓库）。

在本地复现搜索：

```bash
git log --oneline --all --grep='oracle\|TWAP\|twap\|cumulative\|accumulator' -i
git log --oneline --all -S 'price0CumulativeLast' -- contracts/
```

---

## 核心提交（按时间）

| Commit | 日期 | 说明 |
|--------|------|------|
| `6dc7342` | 2019-10-02 | **`add twap accumulator`** — 引入基于时间的累加器思路（TWAP 的数据源）。 |
| `2c5a1af` | 2019-10-30 | `add uint128 where applicable` — 与储备/打包相关，和后续 oracle 槽位优化同一阶段。 |
| `3aa5371` | 2019-11-21 | **`mock out price accumulator`** — 测试里 mock 累加器行为。 |
| `3bb8a6d` | 2019-11-25 | **`compact storing for blocknumber and oracle data`** — 压缩存储、oracle 数据布局。 |
| `da6ba92` | 2019-11-26 | **`clean up names and implementation of oracle stuff`** — 命名与 oracle 实现整理（说明里提到 UQ 定点库更新）。 |
| `c2e415e` | — | Merge PR #6（`dan-dev`），包含上述早期 oracle 工作。 |
| `c4accff` | — | Merge PR #20（`ha-dev`），包含 `3bb8a6d` / `da6ba92` 等。 |
| `baa8ccb` | 2020-01-02 | **`fix swapped price accumulators`** — 修复两个方向累加器弄反的问题。 |
| `a55aa4b` | 2020-01-23 | **`block.number -> block.timestamp`** — TWAP 用 **时间** 积分，用时间戳而非区块号（配合 `tweak timing tests`）。 |
| `585ee2e` | 2020-01-22 | `tweak and add explanatory comments to libraries` — 含 **`UQ112x112`** 等库的注释（与价格编码相关）。 |
| `cfc9eec` | 2020-01-30 | **`add more robust tests for price{0,1}CumulativeLast`** — 针对累加变量的测试加强。 |

---

## 可能相关（弱相关 / 测试 / 周边）

| Commit | 说明 |
|--------|------|
| `7bf29e9` | `contracts: split swapInput into two functions` — 交换逻辑拆分，与累加器同一代码演进期，**非**直接改 TWAP。 |
| `ddbb5e0` | `make invariantLast public` — 与 **`kLast`/协议费** 可见性有关，**不是** TWAP 累加器本身。 |

---

## 当前代码落点

- 累加与更新：[`UniswapV2Pair.sol`](../../contracts/UniswapV2Pair.sol) 中 **`_update`**、`price0CumulativeLast` / `price1CumulativeLast`。  
- 定点数学：[`UQ112x112.sol`](../../contracts/libraries/UQ112x112.sol)。

若需 **完整 diff**，在本地执行：`git show <commit> -- contracts/UniswapV2Pair.sol`。
