package sonyflake

import (
	"testing"
	"time"
)

func TestNewAndNextID(t *testing.T) {
	cfg := DefaultConfig()
	cfg.MachineID = 42
	gen, err := New(cfg)
	if err != nil {
		t.Fatalf("New err = %v", err)
	}
	id, err := gen.NextID()
	if err != nil || id == 0 {
		t.Fatalf("NextID = %d, %v", id, err)
	}
}

func TestGenerator_NilNextID(t *testing.T) {
	var g *Generator
	if _, err := g.NextID(); err == nil {
		t.Error("expected error for nil generator")
	}
}

func TestConfig_Validate(t *testing.T) {
	if err := (Config{MachineID: 1, StartTime: DefaultConfig().StartTime}).Validate(); err != nil {
		t.Errorf("Validate(valid) = %v", err)
	}
	if err := (Config{MachineID: 1}).Validate(); err == nil {
		t.Error("expected start time error")
	}
	if err := (Config{MachineID: 0, StartTime: DefaultConfig().StartTime}).Validate(); err == nil {
		t.Error("expected machine id error")
	}
	if err := (Config{MachineID: MaxMachineID, StartTime: DefaultConfig().StartTime}).Validate(); err != nil {
		t.Errorf("Validate(max machine id) = %v", err)
	}
}

func TestNew_InvalidConfig(t *testing.T) {
	_, err := New(Config{})
	if err == nil {
		t.Error("expected error")
	}
}

func TestNew_StartTimeAhead(t *testing.T) {
	cfg := DefaultConfig()
	cfg.MachineID = 1
	cfg.StartTime = time.Now().Add(24 * time.Hour)
	if _, err := New(cfg); err == nil {
		t.Error("expected error for future start time")
	}
}

func TestNextID_UniqueIncreasing(t *testing.T) {
	cfg := DefaultConfig()
	cfg.MachineID = MaxMachineID
	gen, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	var prev uint64
	for i := 0; i < 1000; i++ {
		id, err := gen.NextID()
		if err != nil {
			t.Fatal(err)
		}
		if id <= prev {
			t.Fatalf("id %d not greater than %d", id, prev)
		}
		prev = id
	}
}
