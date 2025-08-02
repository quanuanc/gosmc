# bactl - Battery Charging Control for Apple Silicon MacBooks

A command-line tool for controlling battery charging on Apple Silicon MacBooks running macOS Tahoe, built using the gosmc library.

## Overview

`bactl` provides fine-grained control over battery charging on supported Apple Silicon MacBooks. This tool directly interfaces with the System Management Controller (SMC) to enable or disable battery charging, which can be useful for:

- Extending battery lifespan by preventing overcharging
- Managing charging during extended desktop use
- Battery health optimization

## Requirements

- macOS Tahoe with Apple Silicon
- Root privileges (sudo access)
- Compatible MacBook model with Tahoe firmware

## Installation

Build and install from source:

```bash
make bactl      # Build the binary
make install    # Install to /usr/local/bin/bactl
```

Or build manually:

```bash
go build -ldflags="-s -w" -o bin/bactl cmd/bactl.go
sudo cp bin/bactl /usr/local/bin/bactl
sudo chmod +x /usr/local/bin/bactl
```

## Usage

### Check Charging Status
```bash
sudo bactl status
```
Shows whether battery charging is currently enabled or disabled.

### Enable Battery Charging
```bash
sudo bactl enable
```
Allows the battery to charge normally.

### Disable Battery Charging
```bash
sudo bactl disable
```
Prevents the battery from charging while plugged in.

### Check Compatibility
```bash
sudo bactl check
```
Verifies if your system supports charging control (detects Tahoe firmware).

## Examples

```bash
# Check if your system supports charging control
$ sudo bactl check
Charging control is supported (Tahoe firmware detected)

# Check current charging status
$ sudo bactl status
Battery charging is ENABLED

# Disable charging for extended desktop use
$ sudo bactl disable
Battery charging disabled

# Re-enable charging when needed
$ sudo bactl enable
Battery charging enabled
```

## Uninstallation

```bash
make uninstall
```

## Technical Details

- Uses SMC key `CHTE` (Tahoe charging control)
- Enabled state: `0x00000000`
- Disabled state: `0x01000000`
- Requires direct SMC access through IOKit framework

## Limitations

- Only works on Apple Silicon MacBooks with Tahoe firmware
- Requires root privileges for SMC access
- Changes are temporary and reset on system restart
- Not compatible with Intel-based Macs or older firmware versions

## Safety Notes

- This tool directly modifies SMC values
- Use responsibly and understand the implications
- Battery charging will resume automatically after system restart
- Monitor battery levels when charging is disabled
