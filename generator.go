package sonyflake

import (
	"errors"
	"fmt"
	"time"

	sonyflake "github.com/sony/sonyflake/v2"
)

// MaxMachineID is the upper bound for Sonyflake machine IDs (uint16).
const MaxMachineID = 65535

// Config holds Sonyflake settings. MachineID must be supplied by the application.
type Config struct {
	// StartTime is the Sonyflake epoch; IDs are unique after this instant.
	StartTime time.Time
	// MachineID is the 1..65535 worker identifier within a datacenter.
	MachineID uint16
}

// DefaultConfig returns default start time (2025-01-01 UTC). Set MachineID before New.
func DefaultConfig() Config {
	return Config{
		StartTime: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
	}
}

// ConfigDefault is an alias for DefaultConfig (naming consistency across go-* libs).
func ConfigDefault() Config { return DefaultConfig() }

// Validate checks machine ID and start time.
func (c Config) Validate() error {
	if c.MachineID < 1 || c.MachineID > MaxMachineID {
		return errors.New("sonyflake: machine id must be between 1 and 65535")
	}
	if c.StartTime.IsZero() {
		return errors.New("sonyflake: start time required")
	}
	return nil
}

// Generator produces distributed uint64 IDs.
type Generator struct {
	flake *sonyflake.Sonyflake
}

// New creates a Generator from Config.
func New(cfg Config) (*Generator, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	machineID := cfg.MachineID
	setting := sonyflake.Settings{
		StartTime: cfg.StartTime,
		MachineID: func() (int, error) {
			return int(machineID), nil
		},
		CheckMachineID: func(id int) bool {
			return id >= 1 && id <= MaxMachineID
		},
	}
	flake, err := sonyflake.New(setting)
	if err != nil {
		return nil, fmt.Errorf("sonyflake initialization failed: %w", err)
	}
	return &Generator{flake: flake}, nil
}

// NextID returns the next unique ID. sonyflake/v2 yields int64 (always
// non-negative); it is exposed as uint64 to keep this library's API stable.
func (g *Generator) NextID() (uint64, error) {
	if g == nil || g.flake == nil {
		return 0, fmt.Errorf("sonyflake is not initialized")
	}
	id, err := g.flake.NextID()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}
