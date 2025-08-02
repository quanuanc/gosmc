package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/charlie0129/gosmc"
)

const (
	// ChargingKey3 is used for Tahoe firmware versions (macOS Tahoe)
	ChargingKey3 = "CHTE"
)

var (
	enabledBytes  = []byte{0x00, 0x00, 0x00, 0x00}
	disabledBytes = []byte{0x01, 0x00, 0x00, 0x00}
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "bactl - Battery Charging Control for Apple Silicon MacBooks (macOS Tahoe)\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  bactl status    Check current charging state\n")
		fmt.Fprintf(os.Stderr, "  bactl enable    Enable battery charging\n")
		fmt.Fprintf(os.Stderr, "  bactl disable   Disable battery charging\n")
		fmt.Fprintf(os.Stderr, "  bactl check     Check if charging control is supported\n")
		fmt.Fprintf(os.Stderr, "  bactl help      Show this help message\n\n")
		fmt.Fprintf(os.Stderr, "Note: This tool requires root privileges and only works on macOS Tahoe with Apple Silicon.\n")
	}

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Validate platform
	if runtime.GOOS != "darwin" {
		log.Fatal("bactl only supports macOS")
	}

	command := flag.Arg(0)

	c := gosmc.New()
	err := c.Open()
	if err != nil {
		log.Fatalf("Failed to open SMC connection: %v", err)
	}
	defer c.Close()

	switch command {
	case "status":
		err := checkStatus(c)
		if err != nil {
			log.Fatalf("Failed to check charging status: %v", err)
		}
	case "enable":
		err := enableCharging(c)
		if err != nil {
			log.Fatalf("Failed to enable charging: %v", err)
		}
		fmt.Println("Battery charging enabled")
	case "disable":
		err := disableCharging(c)
		if err != nil {
			log.Fatalf("Failed to disable charging: %v", err)
		}
		fmt.Println("Battery charging disabled")
	case "check":
		supported := isChargingControlCapable(c)
		if supported {
			fmt.Println("Charging control is supported (Tahoe firmware detected)")
		} else {
			fmt.Println("Charging control is not supported (Tahoe firmware not detected)")
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", command)
		flag.Usage()
		os.Exit(1)
	}
}

func isChargingControlCapable(c gosmc.Conn) bool {
	// Try to read the Tahoe charging key to check if it's available
	_, err := c.Read(ChargingKey3)
	return err == nil
}

func isChargingEnabled(c gosmc.Conn) (bool, error) {
	v, err := c.Read(ChargingKey3)
	if err != nil {
		return false, err
	}

	if len(v.Bytes) != 4 {
		return false, fmt.Errorf("unexpected data length: got %d bytes, expected 4", len(v.Bytes))
	}

	return bytes.Equal(v.Bytes, enabledBytes), nil
}

func checkStatus(c gosmc.Conn) error {
	if !isChargingControlCapable(c) {
		return fmt.Errorf("charging control not supported - requires macOS Tahoe with Apple Silicon")
	}

	enabled, err := isChargingEnabled(c)
	if err != nil {
		return err
	}

	if enabled {
		fmt.Println("Battery charging is ENABLED")
	} else {
		fmt.Println("Battery charging is DISABLED")
	}

	return nil
}

func enableCharging(c gosmc.Conn) error {
	if !isChargingControlCapable(c) {
		return fmt.Errorf("charging control not supported - requires macOS Tahoe with Apple Silicon")
	}

	return c.Write(ChargingKey3, enabledBytes)
}

func disableCharging(c gosmc.Conn) error {
	if !isChargingControlCapable(c) {
		return fmt.Errorf("charging control not supported - requires macOS Tahoe with Apple Silicon")
	}

	return c.Write(ChargingKey3, disabledBytes)
}
