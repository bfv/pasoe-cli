/*
Copyright © 2022 Bronco Oostermeyer <dev@bfv.io>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
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
