# Copilot instructions

This repository contains small, independent Go utility packages. Keep work scoped to the relevant package and avoid broad refactors.

## Key repo guidance

- Follow the project rules in the root AGENTS.md file.
- Prefer package-local changes over cross-cutting edits.
- Keep Go code idiomatic and compatible with the existing API style.
- Validate with the project test command after changes.

## Common validation

```bash
go test ./...
```

Or:

```bash
make test
```

## Package layout

The main folders are organized by responsibility, including `aws`, `bufio`, `byte`, `csv`, `digest`, `file`, `http`, `math`, `memory`, `retry`, and `time`.

Stay aligned with the repo’s package-first structure and keep changes consistent with neighboring code and tests.
