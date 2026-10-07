# Crawler Learning Progress

## Current
- **Phase:** 1 — Sequential CLI crawler
- **Step:** 1.2 (in progress)

## Completed steps
### Phase 0 — Setup
- [x] 0.1 — Verify Go toolchain and workspace
- [x] 0.2 — Initialize Go module (`crawlx`)
- [x] 0.3 — Project layout (`cmd/api` + `internal/{...}`)
- [x] 0.4 — Makefile
- [x] 0.5 — Git init + `.gitignore`
- [x] 0.6 — golangci-lint v2
- [~] Phase 0 checkpoint — **skipped by user override**

### Phase 1 — Sequential CLI
- [x] 1.1 — HTTP GET with timeout + close/drain body
- [ ] 1.2 — context + User-Agent on requests
- [ ] 1.3 — (not started)

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
- **Binary:** `cmd/api` (`package main`)
- **Fetcher:** package-level `http.Client` with Timeout; Close error only if primary err is nil
- **Crawl targets:** own local test server / owned sites only

## Later
- Vercel (thin UI only)
- `internal/api` — Phase 3
- Docker — Phase 4
