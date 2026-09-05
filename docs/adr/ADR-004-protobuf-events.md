# ADR-004: Protobuf for internal events; JSON for SBI interfaces

**Status:** Accepted
**Date:** 2026-09-04
**Phase:** C0

## Context

PulseGrid has three communication surfaces with different requirements:
the Service Based Interfaces (SBI) toward simulated external NFs, the
audit event channel toward Kafka, and (in extension E5) an internal
contract between `pcf-api` and an extracted `decision-engine`.

## Decision

| Situation | Format | Reason |
| --- | --- | --- |
| SBI interface with external NF | HTTP/JSON | 3GPP fidelity — TS 29.500/29.501 define it this way; not negotiable |
| Admin API | HTTP/JSON | Consumed by a browser |
| Audit events (Kafka) | **Protobuf** | Compile-time typed contract, schema evolution with compatibility rules enforceable in CI |
| Internal `decision-engine` contract (E5) | gRPC/Protobuf | Same argument as events, plus native deadlines and cancellation |

## Alternatives considered

**JSON for events too**, for consistency with the HTTP layer. Rejected:
JSON gives no compile-time verified contract and no automatable schema
compatibility rules. It would also miss the opportunity to practice
binary serialization and schema versioning — a distinct, complementary
skill to the mandatory JSON SBI layer. The apparent inconsistency (JSON
outside, Protobuf inside) is intentional: each protocol is chosen for
what's on the other end, not for aesthetic uniformity.

**Avro with a schema registry.** Rejected for now: it adds an extra
infrastructure service (the registry) that this lab doesn't need to make
its point. Protobuf already gives typed contracts without depending on
an extra component. This would be reconsidered only if a real need for
centralized schema governance across multiple producers appeared.

## Consequences

- `contracts/proto/` is the source of truth for audit events and for the
  `decision-engine` contract; it is validated in CI with backward
  compatibility checks (`buf breaking` or equivalent).
- A reviewer who sees gRPC exposed toward `smf-simulator` should read it
  as a design error, not modernization, this decision's table is the
  reference to avoid it.
- Kafka is never on the synchronous decision path: if the broker goes
  down, `pcf-api` keeps responding with the same JSON contract.
