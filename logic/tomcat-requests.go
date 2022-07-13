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
	"fmt"
	"net/http"
	"os"
	"time"
)

func doRequest(verb string, inst PasInstance, path string) (*http.Response, error) {

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	url := fmt.Sprintf("%v://%v:%v%v", inst.Protocol, inst.Host, inst.Port, path)
	req, err := http.NewRequest(verb, url, nil)
	if err != nil {
		printError(err)
		os.Exit(1)
	}

	req.SetBasicAuth(inst.User, inst.Password)
	res, err := client.Do(req)
	if err != nil {
		printError(err)
		os.Exit(1)
	}

	if res.StatusCode == 401 {
		fmt.Printf("HTTP error: %v, reason: unauthorized, path: %v\n", res.StatusCode, path)
		os.Exit(1)
	}

	if res.StatusCode != 200 {
		fmt.Printf("HTTP error: %v, path: %v\n", res.StatusCode, path)
		os.Exit(1)
	}

	return res, err
}
