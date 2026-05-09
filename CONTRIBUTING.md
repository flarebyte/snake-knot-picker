# Contributing

## Development workflow

- `make format`: format Go and design-meta TypeScript examples
- `make test`: run all Go unit/integration tests
- `make test-fixtures`: run fixture-backed regression tests
- `make test-race`: run Go tests with race detector
- `make lint`: run Go vet/golangci-lint and design-meta lint checks
- `make typecheck-ts`: type-check design-meta TypeScript examples
- `make coverage`: generate full Go coverage report
- `make coverage-critical`: coverage summary focused on parser/compiler/argv/validators

## Notes

- Runtime input contract is tokenized argv (`[]string`), not a single command string.
- Inline flag assignment syntax (`--key=value`) is intentionally rejected.
- For repeatable flags, use tokenized pairs such as `[]string{"--add", "a", "--add", "b"}`.
