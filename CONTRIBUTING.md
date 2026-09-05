# Contributing to PulseGrid

This is currently a solo project, but it follows the same discipline a
team project would. The process itself is part of what this repo is
meant to demonstrate.

## Build and test

```bash
go build ./...   # compile
make fmt         # check formatting
make vet         # static analysis
make test        # unit tests
make validate    # all of the above, in the order CI runs them
```

Run `make validate` before every commit. CI runs the exact same target —
if it's green locally, it's green in CI.

## Branching and commits

- Branch names describe phase or intent, not initials or dates:
  `phase/c0-scaffold`, `feat/policy-evaluate`,
  `fix/outbox-relay-retry`.
- Commit messages: a short imperative summary line — `feat: add SM
  Policy Association create endpoint`, `docs: ADR-005 language
  convention`. One logical change per commit. Phase C0 was committed as
  scaffold first, then one ADR per commit — that granularity is the
  pattern to follow.
- Even solo, open a pull request per phase or module rather than
  pushing straight to `main`. The PR description is where reasoning
  lives when the commit message isn't enough room — it's the visible
  evidence of process a reviewer looks for.

## Architecture changes

Any change to process boundaries, the storage model, event format, or
the 3GPP version baseline requires a new or superseding ADR in
`docs/adr/`, not just a code change. Use the existing ADRs as the
expected format: Context, Decision, Alternatives considered,
Consequences.

## Scope discipline

Adding Redis, Kafka, Kubernetes, or full observability before Phase C4
is the most common way a project like this stalls. If a change seems to
need one of those before its phase, that's a signal to re-check the
phase plan, not to add the dependency early.
