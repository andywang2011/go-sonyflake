package sonyflake_test

import (
	"fmt"

	"github.com/andywang2011/go-sonyflake"
)

func ExampleNew() {
	cfg := sonyflake.DefaultConfig()
	cfg.MachineID = 42

	gen, err := sonyflake.New(cfg)
	if err != nil {
		fmt.Println("init error")
		return
	}

	id, err := gen.NextID()
	if err != nil {
		fmt.Println("next error")
		return
	}
	fmt.Println(id > 0)
	// Output: true
}
