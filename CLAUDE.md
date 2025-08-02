# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

bactl is a command-line tool for controlling battery charging on Apple Silicon MacBooks running macOS Tahoe. It provides fine-grained control over battery charging by interfacing with the System Management Controller (SMC) through the gosmc library.

## Architecture

### Core Components

- **main.go**: Primary CLI application entry point
- **Dependencies**: Uses github.com/charlie0129/gosmc library for SMC operations

### Key Functionality

The application provides battery charging control through SMC key `CHTE` (Tahoe charging control):
- **Status checking**: Read current charging state
- **Enable charging**: Set charging to enabled state (0x00000000)
- **Disable charging**: Set charging to disabled state (0x01000000)
- **Compatibility check**: Verify Tahoe firmware support

### Command Structure

```bash
bactl status    # Check current charging state
bactl enable    # Enable battery charging
bactl disable   # Disable battery charging
bactl check     # Check if charging control is supported
```

## Development Commands

### Build
```bash
make                    # Build bactl binary to bin/bactl
make build             # Same as above
go build main.go       # Build without make
```

### Installation
```bash
make install           # Install to /usr/local/bin/bactl
make uninstall         # Remove from /usr/local/bin/bactl
```

### Testing
No automated tests - requires real hardware with Tahoe firmware and root privileges.

## Important Implementation Details

### Platform Requirements
- **macOS Tahoe** with Apple Silicon required
- **Root privileges** needed for SMC access
- **Tahoe firmware** detection via CHTE key availability

### SMC Integration
- Uses gosmc library for SMC communication
- Targets specific SMC key: `CHTE` (`main.go:16`)
- 4-byte data format for charging state control

### Error Handling
- Platform validation (Darwin only)
- SMC connection error handling
- Firmware compatibility checking
- Invalid command detection

### Data Format
- **Enabled state**: `[]byte{0x00, 0x00, 0x00, 0x00}` (`main.go:20`)
- **Disabled state**: `[]byte{0x01, 0x00, 0x00, 0x00}` (`main.go:21`)
- All operations require exactly 4 bytes

## Dependencies

### External Libraries
- `github.com/charlie0129/gosmc`: SMC communication library
- Standard Go libraries: flag, fmt, log, os, runtime, bytes

### Build Requirements
- Go 1.17+
- macOS development environment
- Access to IOKit framework (via gosmc dependency)