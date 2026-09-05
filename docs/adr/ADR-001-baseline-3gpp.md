# ADR-001: 3GPP Release 17 baseline for the Core

**Status:** Accepted
**Date:** 2026-09-04
**Phase:** C0

## Context

PulseGrid needs a stable 3GPP specification baseline to model
`Npcf_SMPolicyControl` and the common data structures (`Supi`, `Snssai`,
`ProblemDetails`). The release landscape at the time of this decision:

| Release | Status | Relevance |
| --- | --- | --- |
| Release 17 | Frozen in 2022; still receiving maintenance versions | Most widely deployed 5G Core baseline in live networks; abundant public reference material |
| Release 18 | Frozen; first 5G-Advanced release | Current vendor implementation target |
| Release 19 | Frozen December 2025 (TSG plenary, Baltimore); corrections only | Last release dedicated entirely to 5G evolution |
| Release 20 | Open; 5G-Advanced maintenance plus first 6G studies | Not an implementation target |

A frozen release is not a frozen document: the TS 29.512 archive shows
Release 16, 17, 18 and 19 packages all updated as recently as March 2026.

## Decision

Use **3GPP Release 17** as the Core's educational baseline. Working
reference, pinned explicitly:

- TS 23.501 / 23.502 / 23.503 — 5GS architecture, PCF-SMF procedure
  context, Policy and Charging Control concepts.
- TS 29.512, archive package `29512-hi0` / v17.18.0 (checked 2026-08-26)
  — `Npcf_SMPolicyControl` reference.
- TS 29.500 / 29.501 / 29.571 — SBI foundations: HTTP/2+JSON realization,
  API design principles, common data types.
- TS 33.501 — security architecture (documented in the Core, implemented
  in extension E10).

## Alternatives considered

**Release 18.** Rejected as the Core baseline: it is the current vendor
target, but that means less public reference material and fewer open
implementations to validate modelled behavior against. The project's
value is fidelity to the most widely deployed baseline, not the newest
one.

**Release 19.** Rejected: only frozen in December 2025, corrections only.
Not enough matured reference material for an educational lab.

## Consequences

- No 5G-Advanced-only functionality (Release 18+) is modelled in the
  Core; if an extension needs it, it is documented in its own
  `docs/3gpp-profile/PROFILE.md` with its own version pin.
- The version pin is re-verified at the start of every phase that
  touches a contract, and at release. A newer maintenance version does
  not force an advance — but if the pin is not advanced, the fact that it
  was reviewed and retained is recorded.
- The repository never copies specification text; it only records the
  selected version, the implemented subset, and the project's own
  contracts.
