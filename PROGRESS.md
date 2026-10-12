# Crawler Learning Progress

## Current
- **Phase:** 1 — Sequential CLI crawler
- **Step:** Phase 1 checkpoint

## Completed steps
### Phase 0 — Setup
- [x] 0.1–0.6
- [~] Phase 0 checkpoint — skipped

### Phase 1 — Sequential CLI
- [x] 1.1–1.9 — fetch, parse, normalize, queue/seen, depth, first_parent
- [x] 1.10 — CLI flags
- [x] 1.11 — refactor helpers (`enqueueLinks`, `appendResult`)
- [x] 1.12 — Issues + referrers; port normalize fixed (keep non-default ports)

## Decisions (ADR-style)
- Depth = hops from seed; maxDepth/maxPage 0 = unlimited
- Edges always recorded; enqueue only if new + same host + depth OK
- Issues: status >= 300 or Error != "" (3xx needs no-follow later)
- Per-fetch timeout + cancel() after Fetch

## Later
- http.Client CheckRedirect for real 3xx
- robots.txt, rate limit, edge dedupe map
- httptest tests
- Phase 2 concurrency
