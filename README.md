# Cronx

[![Build Status](https://github.com/focela/cronx/workflows/Goreleaser/badge.svg)](https://github.com/focela/cronx/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/focela/cronx)](https://goreportcard.com/report/github.com/focela/cronx)
[![Go Reference](https://pkg.go.dev/badge/github.com/focela/cronx?status.svg)](https://pkg.go.dev/github.com/focela/cronx?tab=doc)
[![Release](https://img.shields.io/github/release/focela/cronx.svg?style=flat-square)](https://github.com/focela/cronx/releases)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

A cron job scheduler written in Go that executes commands at specified intervals with signal handling.

## Features

- **Cron Expression Support**: Standard cron expressions with descriptor support (@daily, @weekly, etc.)
- **Signal Handling**: Graceful shutdown on SIGINT and SIGTERM signals
- **Cross-Platform**: Runs on Linux, macOS, and Windows
- **Command Execution**: Execute commands with arguments and output redirection
- **Concurrency Safe**: Context-based cancellation and WaitGroup synchronization
- **Error Handling**: Proper error logging with context

## Installation

### Pre-built Binaries

Download releases from [GitHub Releases](https://github.com/focela/cronx/releases):

```bash
# Linux
wget https://github.com/focela/cronx/releases/latest/download/cronx_linux_amd64.tar.gz
tar -xzf cronx_linux_amd64.tar.gz
sudo mv cronx /usr/local/bin/

# macOS
wget https://github.com/focela/cronx/releases/latest/download/cronx_darwin_amd64.tar.gz
tar -xzf cronx_darwin_amd64.tar.gz
sudo mv cronx /usr/local/bin/

# Windows
# Download cronx_windows_amd64.zip and extract
```

### From Source

```bash
git clone https://github.com/focela/cronx.git
cd cronx
go build -o cronx .
```

## Quick Start

Basic example:

```bash
# Using installed binary
cronx "* * * * *" echo "Hello World"

# Or run directly with Go
go run . "* * * * *" echo "Test"
```

## Usage

```bash
cronx [schedule] [command] [args ...]
```

### Common Use Cases

```bash
# Database backups
cronx "0 2 * * *" backup-database

# Log rotation
cronx "@daily" rotate-logs

# Health checks
cronx "*/5 * * * *" health-check

# Data synchronization
cronx "0 */6 * * *" sync-data

# Cleanup tasks
cronx "@weekly" cleanup-temp-files
```

### Cron Expression Format

```
┌───────────── minute (0 - 59)
│ ┌───────────── hour (0 - 23)
│ │ ┌───────────── day of month (1 - 31)
│ │ │ ┌───────────── month (1 - 12)
│ │ │ │ ┌───────────── day of week (0 - 6) (Sunday to Saturday)
│ │ │ │ │
* * * * *
```

### Descriptors

- `@yearly` or `@annually`: Run once a year
- `@monthly`: Run once a month
- `@weekly`: Run once a week
- `@daily` or `@midnight`: Run once a day
- `@hourly`: Run once an hour

## Signal Handling

Cronx handles the following signals:

- **SIGINT** (Ctrl+C): Stops the scheduler and waits for running jobs to complete
- **SIGTERM**: Same as SIGINT, used for process termination

## Development

### Prerequisites

- Go 1.21 or later
- Git

### Building

```bash
git clone https://github.com/focela/cronx.git
cd cronx
go build -o cronx .
```

### Testing

```bash
# Manual testing
go run . "* * * * *" echo "Test"

# Test signal handling
go run . "@hourly" echo "Test" # Press Ctrl+C to test shutdown

# Note: Automated test suite will be developed in future releases
```

## Documentation

See the [API documentation on go.dev](https://pkg.go.dev/github.com/focela/cronx).

## Dependencies

- [robfig/cron/v3](https://github.com/robfig/cron) - Cron expression parsing and scheduling

## License

Licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE) for details.

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details on how to get started.

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test your changes manually
5. Submit a pull request

Please also read our [Code of Conduct](CODE_OF_CONDUCT.md) before contributing.

## Security

For security vulnerabilities, please email opensource@focela.com instead of using the issue tracker.

## Support

For issues and questions, please use the [GitHub Issues](https://github.com/focela/cronx/issues) page.

## Troubleshooting

### Common Issues

#### Command not found
```bash
# If you get "command not found" error
export PATH=$PATH:/usr/local/bin
# or add to your shell profile (.bashrc, .zshrc, etc.)
echo 'export PATH=$PATH:/usr/local/bin' >> ~/.bashrc
```

#### Permission denied
```bash
# If you get permission denied when moving to /usr/local/bin
sudo mv cronx /usr/local/bin/
# or install to user directory
mkdir -p ~/.local/bin
mv cronx ~/.local/bin/
export PATH=$PATH:~/.local/bin
```

#### Invalid schedule error
```bash
# Check your cron expression format
# Valid: "0 9 * * 1" (Monday at 9 AM)
# Invalid: "0 9 * * 8" (day 8 doesn't exist)
# Use cron expression validator online if needed
```

#### Signal handling issues
```bash
# If signals aren't handled properly, check:
# 1. Process is running in foreground (not background with &)
# 2. Terminal supports signal forwarding
# 3. No other signal handlers are interfering
```

### Debug Mode

For debugging, you can run with verbose output:

```bash
# Enable Go race detector
go run -race . "* * * * *" echo "Test"

# Run with Go debug flags
GODEBUG=gctrace=1 go run . "* * * * *" echo "Test"
```

### Getting Help

If you encounter issues not covered here:

1. Check the [GitHub Issues](https://github.com/focela/cronx/issues)
2. Search existing issues for similar problems
3. Create a new issue with:
   - OS and architecture
   - Go version (`go version`)
   - Cron expression used
   - Error message or unexpected behavior
   - Steps to reproduce

## Roadmap

- [ ] Configuration file support
- [ ] Logging to file
- [ ] Metrics collection
- [ ] Health check endpoint
- [ ] Docker image distribution
