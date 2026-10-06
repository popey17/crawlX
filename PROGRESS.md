# Crawler Learning Progress

## Current
- **Phase:** 0 — Setup
- **Step:** 0.5 (in progress)

## Completed steps
- [x] 0.1 — Verify Go toolchain and workspace
- [x] 0.2 — Initialize Go module (`crawlx`)
- [x] 0.3 — Project layout (`cmd/api` + `internal/{crawler,fetcher,parser,storage,config}`)
- [x] 0.4 — Makefile (`run`, `build`, `test`, `vet`, `fmt`, `.PHONY`, `bin/api`)
- [ ] 0.5 — Git init + `.gitignore`
- [ ] 0.6 — Linter baseline (golangci-lint)

## About me
- Level: intermediate on APIs; concurrency learned in theory, not built yet
- Pace preference: faster through Phase 0 basics
- OS / editor: Arch Linux / Cursor
- Go: `go1.27.1` via mise (`linux/amd64`)
- Docker: used containers; has not authored own Dockerfiles yet
- Repo: init git in this folder

## Concepts to revisit
- Concurrency (practice-heavy starting Phase 2)
- Writing Dockerfiles (Phase 4)
- `package main` vs directory name (`cmd/api`)
- Makefile `.PHONY` (why command targets aren't files)

## Open questions / TODOs
- _(none)_

## Decisions (ADR-style)
- **Module path:** `crawlx` — local learning name
- **Go version in go.mod:** `1.27.1`
- **Binary entrypoint:** `cmd/api` with `package main`
- **Build output:** `bin/api` via Makefile
- **Hosting:** Docker/long-running preferred; Vercel parked
- **Git:** initialize in this folder (step 0.5)

## Later
- Vercel deploy — only if we add a thin UI/status layer later
- `internal/api` package — add in Phase 3 when REST appears
- Docker skeleton — Phase 4 (not Phase 0)
