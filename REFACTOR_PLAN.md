# Refactor Plan

This plan keeps behavior stable while improving structure and maintainability.

## Step 1: Add a tracked refactor plan

- Add this `REFACTOR_PLAN.md` file.
- Use it as the execution checklist for atomic commits.

## Step 2: Extract profile file writing into a reusable package

- Create `internal/output` with shared JSON and `.info` writing helpers.
- Refactor `cmd/main/main.go` to use helpers instead of duplicated write logic.
- Keep generated output format and paths identical.

## Step 3: Add baseline unit tests for utility logic

- Add tests for:
  - `GetLayerHeight`
  - `GetNozzleSize`
  - `EllipticalExtrusionRate`
- Ensure tests document expected parsing/math behavior and run in CI.

## Step 4: Clean `prusa_profiles` loader to remove debug leftovers

- Remove dead JSON decode / pretty logging behavior.
- Make response handling explicit and deterministic.
- Keep public API stable (`Load(version string) (Result, error)`).

## Step 5: Split `process` package logic by concern (non-functional)

- Move big single-file logic into focused files:
  - model/types
  - generation flow
  - inheritance/system lookup
  - speed helpers
- Preserve behavior and JSON output.

## Step 6: Final verification pass

- Run `gofmt` on touched files.
- Run `go test ./...`.
- Capture any follow-up cleanup items if needed.
