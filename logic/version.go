package logic

import (
	_ "embed"
	"fmt"
	"runtime"
	"strings"
)

//go:embed version.txt
var vStr string

func DisplayVersion() {
	thisVersion := strings.Split(strings.Split(vStr, "\n")[0], " ")[1]
	fmt.Printf("\n%v\n", vStr)
	latest, _ := GetLatestRelease()
	if thisVersion != latest.Name {
		fmt.Printf("%vupdate available: %v as of %v%v\n", red, latest.Name, latest.ReleasedAt, reset)
		if runtime.GOOS == "linux" {
			fmt.Printf("  %v%v%v\n", yellow, latest.LinuxArchiveURL, reset)
		} else {
			fmt.Printf("  %v%v%v\n", yellow, latest.WindowsArchiveURL, reset)
		}
	}

	fmt.Println()
}
