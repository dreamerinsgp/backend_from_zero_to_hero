---
name: API Auto Framework Roadmap
overview: Create a comprehensive roadmap document for implementing an API automation framework from scratch, placed at `/home/ubuntu/dex_full/auto_test_framework/auto_framework/ROADMAP.md`. The roadmap will guide you through phases from project setup to CI integration, aligned with patterns from the existing api_auto reference.
todos: []
isProject: false
---

# API Auto Framework Roadmap Plan

## Objective

Create a single roadmap document at `[/home/ubuntu/dex_full/auto_test_framework/auto_framework/ROADMAP.md](/home/ubuntu/dex_full/auto_test_framework/auto_framework/ROADMAP.md)` that guides implementation of an API automation framework from scratch. The `auto_framework` directory will be created if it does not exist.

## Roadmap Content Structure

The roadmap will be a phased, actionable guide covering:

### Phase 1: Project Foundation (Day 1)

- Create project structure: `core/`, `tests/`, `data/`, `reports/`
- Set up `requirements.txt`: pytest, requests, PyYAML, pytest-html
- Add `pytest.ini` with markers (smoke, regression) and test discovery
- Implement `core/config.py`: BASE_URL, TIMEOUT from env vars
- Write a minimal `test_health.py` that hits `/health` and asserts status
- Verify: `pytest tests/test_health.py -v` passes against demo server

### Phase 2: HTTP Client Layer (Day 2)

- Implement `core/client.py`: `new_session()`, `get_client(base_url)` using `requests.Session`
- Add `conftest.py` with `base_url` and `client` fixtures (session/function scope)
- Refactor test_health to use `client` fixture
- Add `tests/test_users.py` for `/users` and `/users/<id>`
- Verify: smoke tests run without hardcoded URLs

### Phase 3: Data-Driven Testing (Day 3)

- Implement `core/loader.py`: `load_yaml(rel_path)` reading from `data/`
- Create `data/users.yaml` with login test cases (email, password, expect_status)
- Use `@pytest.mark.parametrize("case", load_yaml(...))` in `tests/test_login.py`
- Add negative cases: missing email, invalid JSON
- Verify: parametrized tests cover multiple scenarios from YAML

### Phase 4: Interface Association (Day 4)

- Add `auth_token` fixture in `conftest.py`: call `/login`, extract token, `scope="session"`
- Implement `tests/test_me.py`: use `auth_token` in `Authorization: Bearer <token>` header
- Document the pattern: fixture dependency injection for chained API flows
- Verify: authenticated endpoint tests pass with session-scoped token

### Phase 5: Reporting and Execution (Day 5)

- Add pytest-html: `--html=report.html --self-contained-html`
- Add retry for flaky tests (optional): `pytest-rerunfailures`
- Document CLI: `BASE_URL=xxx pytest -m smoke -v` vs `-m regression`
- Verify: HTML report generated with pass/fail summary

### Phase 6: CI Integration (Day 6)

- Add `.github/workflows/api-smoke.yml` (or equivalent) for smoke on push
- Add `.github/workflows/api-regression.yml` for full suite
- Document: env vars, artifact upload for reports
- Verify: workflow runs in CI against deployed/staged environment

### Phase 7: Extensions (Ongoing)

- Optional: pytest-xdist for parallel execution
- Optional: coverage for API code paths
- Optional: JSONPath/jmespath for complex response extraction
- Optional: environment-specific config files (dev/test/prod)

## Target Directory Layout

```
auto_test_framework/auto_framework/
├── ROADMAP.md          # This roadmap document
├── README.md           # (Optional) Quick start
└── (future implementation follows roadmap)
```

## Reference Material

The roadmap will reference patterns from the existing `[api_auto](/home/ubuntu/dex_full/auto_test_framework/api_auto)` implementation (client, config, loader, conftest) and the design notes in `[quetions/4.自动化测试的框架设计.md](/home/ubuntu/dex_full/auto_test_framework/api_auto/quetions/4.自动化测试的框架设计.md)`, but will be written as a standalone "from scratch" guide so you can implement each phase independently.

## Deliverable

A single markdown file: `[ROADMAP.md](/home/ubuntu/dex_full/auto_test_framework/auto_framework/ROADMAP.md)` containing the full phased roadmap with:

- Clear phase boundaries and time estimates
- Concrete file paths and code snippets where helpful
- Verification steps per phase
- Architecture diagram (mermaid) showing config → client/loader → tests flow

