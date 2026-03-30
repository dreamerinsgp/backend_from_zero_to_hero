# Milestone 1 — Setup and how to read the code

[← Roadmap](roadmap.md) · [Next: Milestone 2 →](milestone2.md)

This milestone is about **tooling and mental model**, not yet the AMM math.

## Core concepts in this milestone

| Topic | Concepts |
|--------|-----------|
| **Scope** | **v2-core** vs **v2-periphery** (Router lives outside this repo; core = Factory + Pair + LP ERC20 + libraries). |
| **Repository layout** | [`contracts/`](../../contracts/) — main contracts; [`contracts/interfaces/`](../../contracts/interfaces/) — ABIs for external callers; [`contracts/libraries/`](../../contracts/libraries/) — `SafeMath`, `Math`, `UQ112x112`; [`contracts/test/`](../../contracts/test/) — test ERC20. |
| **Build output** | `yarn compile` → artifacts under `build/` (needed for scripts and tests). |
| **Testing** | `yarn test` — unit tests against a local mock chain (see [environment](../1.environment/1.environment.md) for Node version notes). |
| **High-level map** | What Factory, Pair, and `UniswapV2ERC20` each do — [Core features](../1.environment/4.core_features_of_this_contract.md). |

## Read / do

- [Environment](../1.environment/1.environment.md)
- [How to build](../1.environment/2.how_to_build_the_contract.md)
- Skim [Core features](../1.environment/4.core_features_of_this_contract.md)

## No AMM-specific inventory yet

Sections **A–H** in [roadmap](roadmap.md) start applying from **Milestone 2** onward.
