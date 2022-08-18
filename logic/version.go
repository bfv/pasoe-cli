package logic

import (
	_ "embed"
	"fmt"
)

//go:embed version.txt
var vStr string

func DisplayVersion() {
	fmt.Printf("\n%v\n", vStr)
}
