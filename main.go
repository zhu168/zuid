package main

import (
	"fmt"
)

func main() {
	//If you need to set the timestamp of the Snowflake algorithm in advance,
	//the default is 2023-08-30 00:00:00 UTC
	Epoch = uint64(1682908800000)
	// Create a ZUID instance and pass in the worker node ID
	zuid, err := NewZUID(1)
	if err != nil {
		fmt.Println("Error creating ZUID instance:", err)
		return
	}
	for i := 0; i < 10; i++ {
		//Generate a zuid
		id := zuid.NextIDSimple()
		fmt.Println("ZUID:", id, len(id))
	}
}
