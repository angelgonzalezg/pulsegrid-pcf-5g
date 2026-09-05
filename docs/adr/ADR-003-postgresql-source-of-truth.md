# ADR-003: PostgreSQL as the source of truth; Redis as a disposable cache

**Status:** Accepted
**Date:** 2026-09-04
**Phase:** C0

## Context

The Core needs to guarantee three correctness properties under real
concurrency: a single active policy association per session, idempotency
keys that survive client retries, and atomicity between a state change
and the event that announces it (outbox pattern). A decision cache is
also wanted to reduce latency.

## Decision

**PostgreSQL is the only source of truth** for: associations, decisions
(immutable and versioned), idempotency keys plus request hash, and outbox
rows. **Redis/Valkey only caches decisions**, in cache-aside mode, and its
failure can **never** affect the correctness of the system — only its
latency (*fail open*).

Hard implementation rule: no network call to Redis/Valkey or Kafka happens
inside an open PostgreSQL transaction. The outbox event is inserted in
the same transaction as the state change it describes.

## Alternatives considered

**Storing idempotency or active-association state in Redis/Valkey**, for speed.
Rejected: the day Redis/Valkey restarts, client retries that relied on that key
to avoid duplication silently create duplicate associations. Whatever
protects correctness lives in the source of truth, no exceptions, a
cache your correctness depends on isn't a cache, it's a poorly built
database.

**Event sourcing as the primary model** (rebuilding state from the event
log). Rejected for the Core: it would add projection and rebuild
machinery that isn't necessary to demonstrate the PCF contract — CRUD
with optimistic concurrency plus outbox already teaches atomicity and
idempotency without that extra cost. Not ruled out as a future
extension if justified.

## Consequences

- Optimistic concurrency (version number) protects concurrent updates to
  the same association; an update against a stale version fails instead
  of overwriting.
- `outbox-relay` claims rows with `SKIP LOCKED`, never with a full-table
  exclusive lock.
- Integration tests must include a "Redis down" scenario: same
  responses, only slower. If that scenario changes the result, the
  design violated this decision.
