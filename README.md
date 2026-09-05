# PulseGrid — PCF 5G
Distributed Go policy engine for 5G session control (3GPP Rel-17) utilizing an event-driven architecture, Redis/Valkey cache-aside, and PostgreSQL transactional outbox pattern.

## What this is

- A lab, not a production PCF. The Core implements a scoped subset of
  `Npcf_SMPolicyControl` against **3GPP Release 17** (see
  `docs/adr/ADR-001-baseline-3gpp.md` for why this release and not a
  newer one), talking to a simulated SMF — not real network equipment.
- A deliberately staged project: the Core has its own release gate and
  is fully presentable on its own. Extensions add depth in a specific
  direction; they never make the Core "more finished" than it already is.

## What this is not

- Not a certified or spec-complete 3GPP PCF implementation.
- Not built for production traffic, real subscriber data, or every
  interface a real PCF exposes — only the ones the Core scopes in.

## The Mandatory Core

Three deployables, one modular monolith and two independent workers,
coordinating only through PostgreSQL and Kafka, never through a direct
call to one another (see `docs/adr/ADR-002-modular-monolith.md`):

| Component | Role |
| --- | --- |
| `pcf-api` (`cmd/pcf`) | The SM Policy Control API. Internally modular (`internal/policy`, `internal/session`), one process externally. |
| `outbox-relay` (`cmd/relay`) | Claims outbox rows and publishes domain events to Kafka. |
| `audit-worker` (`cmd/audit`) | Consumes those events and writes the audit trail. |

PostgreSQL is the source of truth for associations, decisions, and
idempotency; Redis is a fail-open cache with no correctness role (see
`docs/adr/ADR-003-postgresql-source-of-truth.md`). SBI traffic is
HTTP/JSON per 3GPP; internal events are Protobuf (see
`docs/adr/ADR-004-protobuf-events.md`).

## Optional Extensions

Not all of them — the plan is explicit that you pick **one route**, not
a catalog, based on what the portfolio is meant to demonstrate:

| Extension | Adds | Best route for |
| --- | --- | --- |
| E1 — AM/AF policies | Mobility and application-triggered policy, alongside SM | Telecom / 5G domain roles |
| E2 — NRF, UDR, CHF integration | Service discovery, subscriber data lookup, spending-limit control | Telecom / 5G domain roles |
| E3 — Durable notifications | Guaranteed third-party callback delivery via NATS JetStream | Backend / distributed systems roles |
| E4 — Policy Catalog admin | An operations console and Admin API over catalog data | Full-stack / product roles |
| E5 — Decision-engine extraction | Measured gRPC extraction, with a keep-or-revert report | Backend / distributed systems roles |
| E6 — Policy copilot | Read-only AI research assistant over operational data | Full-stack / product roles |
| E7 — Kubernetes local | Helm chart, `kind` profile, backlog-based autoscaling | Platform / DevOps / SRE roles |
| E8 — OCI portability | Documented diff report running the same release on OCI | Platform / DevOps / SRE roles |
| E9 — CI portability | Same Makefile targets ported to Jenkins | Platform / DevOps / SRE roles |
| E10 — Security | OAuth2 and mTLS across service-to-service calls | Backend / distributed systems roles |

Two rules that apply to any extension chosen: the Core is never degraded
to accommodate one, and one finished extension is worth more than three
half-done ones.

## Roadmap

| Phase | Focus |
| --- | --- |
| F1 — Monolithic MVP (C0–C3) | A thin, complete, end-to-end system |
| F2 — Modularization & reliability (C4–C8) | Idempotency, outbox, observability, `v1.0.0-core` release |
| F3 — Measured extraction | One extension from the table above |
| F4 — Local Kubernetes | Declarative deployment, no cloud cost |
| F5 — Cloud | Real operation, real cost discipline, verified teardown |

The Core is presentable at the end of F2, independent of anything after
it.

## Quickstart

```bash
go build ./...   # compiles the three binaries
make validate    # fmt + vet + build + test — the same target CI runs
make dev         # Postgres/Kafka/Redis + the three binaries (from Phase C1 onward)
```

## Documentation map

| Document | What it covers |
| --- | --- |
| `docs/adr/` | Architecture decisions, with alternatives considered and rejected |
| `docs/PROJECT_PLAN.md` | The full scope contract: Core spec, extensions, phases, gates |
| `CONTRIBUTING.md` | How to build, test, and submit changes |

## Status

Currently in Phase C0 (scope and architecture). See `docs/adr/` for the
decisions made so far.

## License & Disclaimers

### License
This project is licensed under the [Apache License 2.0](LICENSE).

### Educational & 3GPP Disclaimer
* **Educational Purpose:** **PulseGrid PCF 5G** is an independent, non-commercial portfolio project developed strictly for educational, research, and technical demonstration purposes[cite: 1]. It is not a 3GPP-certified product and is not intended for commercial production use[cite: 1].
* **Intellectual Property Notice:** 3GPP™ specifications (including TS 29.512 and TS 23.501) and associated trademarks are the intellectual property of the 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC)[cite: 1]. 
* **Independent Implementation:** This repository contains original code and schema representations inspired by open 3GPP architectural standards[cite: 1]. It does not distribute proprietary 3GPP specification texts or copyrighted source materials[cite: 1].
