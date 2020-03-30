package main

import (
	"fmt"
	"os"
	"testing"
)

func TestDemo(t *testing.T) {
	fmt.Println(os.Getenv("GOPROXY"))
}
