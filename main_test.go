package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Args = []string{"RoommateScan", "-h", "192.168.1.100-200", "-p", "80,443"}
	main()
}
