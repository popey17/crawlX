# Crawler Learning Progress

## Current
- **Phase:** 1 — Sequential CLI crawler
- **Step:** 1.3 (in progress)

## Completed steps
### Phase 0 — Setup
- [x] 0.1–0.6 — toolchain, module, layout, Makefile, git, lint
- [~] Phase 0 checkpoint — skipped by user override

### Phase 1 — Sequential CLI
- [x] 1.1 — HTTP GET with timeout + close/drain body
- [x] 1.2 — context + User-Agent (`NewRequestWithContext` + `Do`)
- [ ] 1.3 — Parse HTML: title + anchor hrefs
- [ ] 1.4 — (not started)

## About me
- Level: intermediate on APIs; concurrency learned in theory, not built yet
- Pace preference: faster through Phase 0 basics
- OS / editor: Arch Linux / Cursor
- Go: `go1.27.1` via mise

## Concepts to revisit
- Concurrency (Phase 2)
- Dockerfiles (Phase 4)
- Close error vs primary error (named returns)

## Decisions (ADR-style)
- **Module path:** `crawlx`
- **Fetcher:** package-level client; ctx first; User-Agent `crawlx/0.1 (+local training)`; method param kept
- **Crawl targets:** own local test server / owned sites only

## Later
- Vercel (thin UI only)
- `internal/api` — Phase 3
- Docker — Phase 4
