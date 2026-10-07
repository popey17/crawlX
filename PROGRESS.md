# Crawler Learning Progress

## Current
- **Phase:** 0 — Setup
- **Step:** Phase 0 checkpoint (quiz)

## Completed steps
- [x] 0.1 — Verify Go toolchain and workspace
- [x] 0.2 — Initialize Go module (`crawlx`)
- [x] 0.3 — Project layout (`cmd/api` + `internal/{crawler,fetcher,parser,storage,config}`)
- [x] 0.4 — Makefile (`run`, `build`, `test`, `vet`, `fmt`, `.PHONY`, `bin/api`)
- [x] 0.5 — Git init + `.gitignore`
- [x] 0.6 — Linter baseline (golangci-lint v2, `make lint` clean)

## About me
- Level: intermediate on APIs; concurrency learned in theory, not built yet
- Pace preference: faster through Phase 0 basics
- OS / editor: Arch Linux / Cursor
- Go: `go1.27.1` via mise (`linux/amd64`)
- Docker: used containers; has not authored own Dockerfiles yet

## Concepts to revisit
- Concurrency (practice-heavy starting Phase 2)
- Writing Dockerfiles (Phase 4)
- Makefile `.PHONY`
- golangci-lint v2 config (`version: "2"`, `linters.default: none`)

## Decisions (ADR-style)
- **Module path:** `crawlx`
- **Binary entrypoint:** `cmd/api` with `package main`
- **Build output:** `bin/api`
- **Hosting:** Docker/long-running preferred; Vercel parked
- **Docker in Phase 0:** skipped — Phase 4
- **Linter:** golangci-lint v2 with govet/errcheck/staticcheck/ineffassign/unused

## Later
- Vercel deploy — thin UI only if desired
- `internal/api` package — Phase 3
