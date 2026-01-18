# Contributing to safeconv

Thank you for your interest in contributing to safeconv!

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/safeconv.git`
3. Create a feature branch: `git checkout -b feature/your-feature`

## Development Setup

```bash
# Install dependencies (none required - stdlib only)
go mod download

# Run tests
go test -v ./...

# Run tests with race detection
go test -race ./...

# Run linter
golangci-lint run

# Check test coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Guidelines

### Code Style

- Follow standard Go conventions (`gofmt`, `goimports`)
- Run `golangci-lint run` before submitting
- Keep functions small and focused

### Testing

- Write table-driven tests
- Test boundary conditions (min/max values, zero, negative)
- Test error cases explicitly
- Maintain or improve test coverage (currently 97.8%)

### Commits

- Use clear, descriptive commit messages
- Reference issues where applicable (`Fixes #123`)
- Keep commits focused on a single change

### Pull Requests

1. Ensure all tests pass: `go test -race ./...`
2. Ensure linting passes: `golangci-lint run`
3. Update CHANGELOG.md under `[Unreleased]`
4. Provide a clear description of changes
5. Link related issues

## Adding New Conversions

When adding new conversion functions:

1. Follow the existing naming pattern: `{FromType}To{ToType}`
2. Add corresponding `Must*` variant in `must.go`
3. Add comprehensive tests covering:
   - Zero value
   - Positive in-range values
   - Boundary values (exactly at min/max)
   - Overflow cases (just past max)
   - Underflow cases (just past min)
   - For floats: NaN, +Inf, -Inf
4. Update documentation in `doc.go` if needed
5. Update README.md function tables

## Reporting Issues

- Check existing issues first
- Include Go version (`go version`)
- Include OS and architecture
- Provide minimal reproduction code

## Questions?

Open an issue with the `question` label.
