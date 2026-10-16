# Design: Colored CLI output + ASCII banner (crawlx)

**Date:** 2026-10-17  
**Status:** Approved (minimal implementation)  
**Scope:** Display only — no crawl logic changes; keep diffs small  

## Goal

Make the CLI easier to scan: ASCII banner at start, colored live crawl lines, colored summary and Issues. Matches option **B** from brainstorming.

## Non-goals

- Full “Done” ASCII art box (option C) — out of scope  
- Changing crawl / concurrency behavior  
- Forcing color when stdout is piped  

## Approach

Tiny `internal/ui` package using raw ANSI escape codes (no new third-party color dependency). Disable colors when stdout is not a TTY.

## Behavior

### Banner (startup, `cmd/api`)

- Print a compact ASCII wordmark for `CRAWLX` (cyan / bold when color on).  
- Print config: `seed`, `max-pages`, `max-depth`, `workers` (dim).  

### Live crawl lines (`appendResult` in crawler)

Keep existing fields (`#n`, depth, status, err, url, parent). Color by result:

| Condition | Color |
|-----------|--------|
| `err != ""` or status ≥ 400 | Red |
| status 300–399 | Yellow |
| status 200–299, no err | Green |
| `#n` / depth labels | Dim or cyan (secondary) |

### Summary (`main`)

`Done pages=… edges=… error=…`  
- If `error == 0`: green  
- Else: keep pages/edges normal; show error count in red  

### Issues (`GetReference`)

- Header “Issues” bold/red  
- Bad URL lines red  
- Referrer list dim  

### TTY / pipe safety

If stdout is not a terminal, all helpers return plain strings (no escapes).

## API sketch (`internal/ui`)

```text
Enabled() bool          // stdout is TTY
Banner() string         // ASCII art
Cyan/Bold/Dim/Green/Yellow/Red(s string) string
StatusLine(n, depth, status int, err, url, parent string) string  // optional helper
```

## Files to touch

| File | Change |
|------|--------|
| `internal/ui/color.go` | ANSI helpers + TTY detect |
| `internal/ui/banner.go` | ASCII banner |
| `cmd/api/main.go` | Banner + colored summary |
| `internal/crawler/crawler.go` | Colored `appendResult` + `GetReference` only |

## Acceptance

1. Interactive run shows banner + colored lines + colored Done/Issues.  
2. `go run ./cmd/api ... | cat` (or redirect) → no ANSI junk.  
3. Crawl behavior unchanged (same pages/edges counts for same flags).  
4. `go build ./...` succeeds.  

## Open decisions

None — colors and layout fixed as above unless review requests tweaks.
