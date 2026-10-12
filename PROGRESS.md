# Crawler Learning Progress

## Current
- **Phase:** 1 — Sequential CLI crawler
- **Step:** 1.12 (in progress)

## Completed steps
### Phase 0 — Setup
- [x] 0.1–0.6
- [~] Phase 0 checkpoint — skipped

### Phase 1 — Sequential CLI
- [x] 1.1–1.9 — core crawl + depth/parent
- [x] 1.10 — CLI flags
- [x] 1.11 — `enqueLinks` + `appendResult` extracted
- [ ] 1.12 — Issues report (bad pages + referrers from edges)

## Decisions (ADR-style)
- Per-fetch timeout inside loop
- Next: refactor before more features (Crawl too large)

## Later
- defer cancel in loop → explicit cancel per iteration
- Edge dedupe, httptest, robots, rate limit
- Phase 2 concurrency
