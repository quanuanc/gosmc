# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

gosmc is a Go library that provides access to Apple's System Management Controller (SMC) on Mac systems through C bindings. The library enables reading and writing SMC keys for hardware monitoring and control.

## Architecture

### Core Components

- **smc.go**: Main library implementation with CGO bindings to Apple's IOKit framework
- **type.go**: Defines the `Conn` interface for SMC operations
- **mock.go**: Mock implementation for testing without actual hardware access
- **cmd/main.go**: CLI application for SMC operations
- **smc.c/smc.h**: C implementation for low-level SMC communication

### Key Interfaces

The `Conn` interface (`type.go:3-13`) defines the contract for SMC operations:
- `Open()`: Establishes SMC connection
- `Close()`: Closes SMC connection  
- `Read(key string)`: Reads SMC key (4-character keys only)
- `Write(key string, value []byte)`: Writes to SMC key

### Data Flow

1. Client creates connection via `New()` or `NewMockConn()`
2. Opens connection with `Open()`
3. Performs read/write operations on 4-character SMC keys
4. Closes connection with `Close()`

## Development Commands

### Build
```bash
make                    # Build CLI binary to bin/gosmc
make build             # Same as above
go build ./cmd         # Build without make
```

### Testing
```bash
go test -v             # Run all tests with verbose output
go test               # Run tests quietly
```

**Note**: Tests interact with real SMC hardware and require root access on macOS. Tests will fail on non-Mac systems or without proper permissions.

### CLI Usage
```bash
bin/gosmc -k CH0B                # Read from SMC key CH0B
bin/gosmc -k CH0B -v 02         # Write hex value 02 to CH0B
```

## Important Implementation Details

### CGO Integration
- Uses `#cgo LDFLAGS: -framework IOKit` for macOS system access
- C code handles low-level SMC communication
- Go code provides safe interface with proper memory management

### SMC Key Constraints
- All SMC keys must be exactly 4 characters (`smc.go:52`, `smc.go:70`)
- Keys are case-sensitive
- Common keys: CH0B/CH0C (battery charging), PSTR (power), etc.

### Data Types
- Raw bytes are primary data format
- CLI supports hex encoding/decoding
- Float decoding available for "flt " data type keys (`cmd/main.go:47-51`)

### Error Handling
- Custom errors: `ErrKeyLength`, `ErrNoData` (`smc.go:23-25`)
- C function return codes mapped to Go errors
- Mock implementation provides predictable error scenarios

## Testing Strategy

The codebase uses testify/assert for testing. The main test (`smc_test.go:9-27`) requires:
- macOS system with SMC access
- Root privileges for write operations
- Real hardware (not virtual machines)

For development without hardware access, use `NewMockConn()` instead of `New()`.