package main

import (
	_ "embed"
	"fmt"
)

//go:generate ./gen_version.sh

//go:embed version.txt
var version = "unknown version"

func main() {
fmt.Println("version:", version)
}
