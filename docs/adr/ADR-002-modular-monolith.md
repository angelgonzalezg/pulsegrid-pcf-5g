# ADR-002: Modular monolith instead of microservices from day one

**Status:** Accepted
**Date:** 2026-09-04
**Phase:** C0

## Context

PulseGrid simulates a PCF with three Core processes (`pcf-api`,
`outbox-relay`, `audit-worker`). The architectural question is whether those
three processes, and the two internal domains, `policy` and `session` should be born as independent services or as a single process with
code boundaries.

## Decision

`pcf-api` is a **single-process modular monolith**: `internal/policy` and
`internal/session` are separate domain packages inside the same binary,
with no network calls between them. `outbox-relay` and `audit-worker` are
separate processes, but they coordinate only through the outbox table
and through Kafka, never through a direct call to one another.

## Alternatives considered

**Microservices from day 1** (one per domain: policy, session,
notifications). Rejected for this reason:

- *Technical:* the entire time budget would be
  consumed by coordination infrastructure (N pipelines, a compatibility
  matrix, local Kubernetes just to be able to work) before a single
  policy rule was functioning. Microservices mostly solve an
  organizational problem, independent teams deploying without
  coordinating, which doesn't exist with a single developer.

**A single binary for the API and the workers.** Rejected: `outbox-relay`
scales with backlog size and `audit-worker` scales with Kafka throughput
— neither metric is related to `pcf-api`'s requests per second. Mixing
them into one binary would hide exactly the pattern being demonstrated
(the outbox as a coordination boundary, not a direct call).

## Consequences

- Each domain module (`policy`, `session`) has a written extraction
  policy: it becomes an independent service only if it meets several of
  these criteria, distinct scaling signal, distinct lifecycle, distinct
  failure domain, clean data boundary with no shared transaction.
- Phase C5/E5 runs a two-week experiment measuring whether extracting
  `decision-engine` into a separate gRPC service was worth it. Reverting
  the extraction if the report doesn't justify it is a valid outcome, not
  a failure.
- No module is extracted in anticipation of scale that doesn't exist yet.
