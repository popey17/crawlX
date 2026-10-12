# Crawler Learning Progress

## Current
- **Phase:** 2 — Concurrency
- **Step:** 2.1 (in progress)

## Completed steps
### Phase 0 — Setup
- [x] 0.1–0.6
- [~] checkpoint skipped

### Phase 1 — Sequential CLI
- [x] 1.1–1.12 — full sequential crawler + issues/referrers
- [x] Phase 1 checkpoint — passed (~4/5) with small depth wording fix

### Phase 2 — Concurrency
- [ ] 2.1 — Goroutines basics + ownership questions + race demo

## Concepts to revisit
- Depth = link hops from seed (not “page count”)
- 3xx + CheckRedirect (later)
- Rate limit / robots (Phase 2+)

## Decisions (ADR-style)
- Keep non-default ports in normalize
- Issues from pages + referrers from edges
