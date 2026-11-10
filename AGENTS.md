# AGENTS.md

This repository is a small Go utility library organized as independent package folders such as `aws`, `bufio`, `csv`, `file`, `http`, `math`, `memory`, `retry`, and `time`.

## Source of truth

- This file is the canonical repo-wide instruction for automated agents and contributors.
- The project README and GitHub Copilot guidance are supporting documentation and should stay aligned with these rules.
- If there is ever a conflict, follow the instructions in this file.

## Project scope

- Keep changes small and focused on the relevant package.
- Prefer package-local edits over filesystem-wide refactors.
- Maintain compatibility with the library’s existing public APIs unless a change is explicitly requested.
- Preserve the small-library structure: each folder is intended to be a self-contained Go package.

## Go conventions

- Write idiomatic Go code that matches the surrounding package style.
- Prefer clear, minimal abstractions over broad generalization.
- Keep error handling explicit and operationally useful.
- Prefer package-level tests that validate the real behavior of the affected code.

## Validation

Before finalizing changes, run the relevant package or repo-level test command. The project’s standard verification command is:

```bash
go test ./...
```

The repo’s Makefile also exposes a test target:

```bash
make test
```

Use the narrowest command that checks the changed behavior; if the change is package-local, a focused package test is acceptable, but repo-level verification is preferred when practical.

## Change expectations

- Do not make unrelated cleanup changes in the same patch.
- If a fix requires behavior changes, document them clearly in the commit or PR description.
- Respect the existing package organization and naming patterns.
- When adding a new utility, keep it consistent with neighboring packages and their documentation style.

## Notes

- This repo contains compatibility helpers and legacy functions that should not be removed casually.
- Prefer incremental, backwards-compatible improvements over breaking API changes.
- The repo is intentionally modular and should remain easy to navigate by package.
