package main

import (
	_ "embed"
	"fmt"
)

//go:generate ./gen_version.sh

//go:embed version.txt
var version string

func main() {
fmt.Println("version:", version)
}
