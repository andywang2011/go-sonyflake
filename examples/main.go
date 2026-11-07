// Example: application supplies MachineID; library only generates IDs.
package main

import (
	"fmt"
	"log"

	libsony "github.com/andywang2011/go-sonyflake"
)

func main() {
	cfg := libsony.DefaultConfig()
	cfg.MachineID = 1 // in production: from env, DB, or coordinator

	gen, err := libsony.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 3; i++ {
		id, err := gen.NextID()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id)
	}
}
