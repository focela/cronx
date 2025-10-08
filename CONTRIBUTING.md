# Contributing to Cronx

Thank you for your interest in contributing to Cronx! This document provides guidelines and information for contributors.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Making Changes](#making-changes)
- [Pull Request Process](#pull-request-process)
- [Coding Standards](#coding-standards)
- [Testing](#testing)
- [Documentation](#documentation)
- [Issue Reporting](#issue-reporting)

## Code of Conduct

This project follows the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/your-username/cronx.git
   cd cronx
   ```
3. **Add the upstream remote**:
   ```bash
   git remote add upstream https://github.com/focela/cronx.git
   ```

## Development Setup

### Prerequisites

- Go 1.21 or later
- Git

### Building

```bash
# Build the project
go build -o cronx .

# Run the application
go run . "* * * * *" echo "Test"
```

## Making Changes

### Branch Naming

Use descriptive branch names following these patterns:

- `feature/description` - New features
- `fix/description` - Bug fixes
- `docs/description` - Documentation updates
- `refactor/description` - Code refactoring
- `test/description` - Test-related improvements (when available)

Examples:
- `feature/add-config-file-support`
- `fix/signal-handling-race-condition`
- `docs/update-installation-guide`

### Commit Messages

Follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Test-related changes (when automated tests are available)
- `chore`: Maintenance tasks

Examples:
```
feat(scheduler): add configuration file support
fix(signal): handle race condition in signal handling
docs(readme): update installation instructions
```

## Pull Request Process

### Before Submitting

1. **Update your fork**:
   ```bash
   git fetch upstream
   git checkout main
   git merge upstream/main
   ```

2. **Create a feature branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```

3. **Make your changes** following the coding standards

4. **Test your changes manually**:
   ```bash
   go build .
   go run . "* * * * *" echo "Test your changes"
   ```

5. **Update documentation** if needed

### Submitting a Pull Request

1. **Push your branch**:
   ```bash
   git push origin feature/your-feature-name
   ```

2. **Create a Pull Request** on GitHub with:
   - Clear title following conventional commits
   - Detailed description of changes
   - Reference to any related issues
   - Screenshots or examples if applicable

3. **Fill out the PR template**:
   ```markdown
   ## What this PR does / why we need it
   Brief description of what this PR does and why it's needed.

   ## Changes
   - List of specific changes made
   - Add feature X
   - Fix issue Y
   - Update configuration Z

   ## Staging PRs
   - N/A (or list related staging PRs)

   ## Testing
   - [ ] Manual testing performed
   - [ ] Test results and verification steps
   - [x] Specific test completed

   ## Notes
   - Additional context or considerations
   - Breaking changes (if any)

   ## Related
   - Closes #XXX
   - Fixes #XXX
   ```

### Review Process

1. **Automated checks** must pass (CI/CD, linting)
2. **Code review** by maintainers
3. **Address feedback** and update PR as needed
4. **Merge** after approval

## Coding Standards

### Go Code Style

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` and `goimports` for formatting
- Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use meaningful variable and function names
- Add godoc comments for exported functions

### Code Organization

- Keep functions small and focused
- Use early returns for error handling
- Avoid deep nesting
- Group related functionality together

### Error Handling

```go
// Good
if err := doSomething(); err != nil {
    return fmt.Errorf("failed to do something: %w", err)
}

// Avoid
if err != nil {
    // Handle error
}
```

### Comments

- Use English for all comments
- Write complete sentences ending with periods
- Start exported identifiers with the identifier name
- Explain "why" not "what" in inline comments

## Testing

### Manual Testing

```bash
# Test basic functionality
go run . "* * * * *" echo "Test"

# Test with different cron expressions
go run . "*/5 * * * *" echo "Every 5 minutes"
go run . "@daily" echo "Daily task"

# Test signal handling
go run . "* * * * *" echo "Test" # Then press Ctrl+C
```

### Automated Testing

Automated test suite will be developed in future releases:

```bash
# Run all tests (when available)
go test ./...

# Run tests with coverage
go test -cover ./...
```

## Design Decisions

### Architecture Principles

- **Single Responsibility**: Each function has one clear purpose
- **Error Handling**: Use early return pattern with proper error logging
- **Signal Handling**: Graceful shutdown on SIGINT/SIGTERM signals
- **Cross-Platform**: Support Linux, macOS, and Windows
- **Minimal Dependencies**: Only essential external dependencies

### Code Style Decisions

- **Package Structure**: Single main package for CLI tool simplicity
- **Function Naming**: Use descriptive names that explain behavior
- **Error Messages**: Clear, actionable error messages for users
- **Logging**: Use fmt.Printf for user-facing output, avoid complex logging
- **Concurrency**: Use WaitGroup for job synchronization

### Technical Choices

- **Cron Library**: robfig/cron/v3 for cron expression parsing and scheduling
- **Signal Handling**: os/signal package for cross-platform signal support
- **Command Execution**: os/exec for subprocess management
- **Error Handling**: os.Exit(1) for fatal errors, early return for recoverable errors
- **Documentation**: godoc format for all exported identifiers

### Future Considerations

- **Configuration**: YAML/JSON config files for complex setups
- **Logging**: Structured logging with levels and file output
- **Metrics**: Prometheus metrics for monitoring
- **Health Checks**: HTTP endpoint for health monitoring
- **Docker**: Container distribution

## Documentation

### README Updates

- Update relevant sections when adding features
- Include examples for new functionality
- Update installation instructions if needed
- Add troubleshooting section for common issues

### Code Documentation

- Add godoc comments for new exported functions
- Update package documentation if needed
- Include examples in function comments when helpful
- Include examples in comments when helpful

### Release Notes

- Document user-facing changes in pull requests
- Highlight breaking changes clearly
- Include migration notes when applicable

## Issue Reporting

### Before Creating an Issue

1. **Search existing issues** to avoid duplicates
2. **Check if it's already fixed** in the latest version
3. **Gather information** about your environment

### Creating an Issue

Use the appropriate issue template and include:

- **Clear title** describing the problem
- **Detailed description** of the issue
- **Steps to reproduce** the problem
- **Expected vs actual behavior**
- **Environment information** (OS, Go version, etc.)
- **Screenshots or logs** if applicable

### Issue Labels

- `bug`: Something isn't working
- `enhancement`: New feature or request
- `documentation`: Improvements or additions to documentation
- `question`: Further information is requested
- `help wanted`: Extra attention is needed

## Release Process

### Versioning

We follow [Semantic Versioning](https://semver.org/):
- `MAJOR`: Incompatible API changes
- `MINOR`: Backward-compatible functionality additions
- `PATCH`: Backward-compatible bug fixes

### Release Notes

- Include all user-facing changes
- Highlight breaking changes
- Mention new features and improvements
- List bug fixes

## Getting Help

- **GitHub Issues**: For bug reports and feature requests
- **Discussions**: For questions and general discussion
- **Email**: opensource@focela.com for security issues

## Recognition

Contributors will be recognized in:
- Release notes
- CONTRIBUTORS.md file
- GitHub contributors page

Thank you for contributing to Cronx.
