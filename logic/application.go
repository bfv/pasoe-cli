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
	"net/url"

	"github.com/bfv/pasoe-cli/model"
)

func getApplicationNames(inst PasInstance) []string {

	var apps []string

	res, err := doRequest("GET", inst, "/oemanager/applications")
	if err != nil {
		fmt.Println(err)
	} else {
		oeres, _ := readJson(res)
		apps, _ = extractApplicationNames(oeres)
	}

	return apps
}

func GetApplications(inst PasInstance) []model.Application {

	var apps []model.Application

	res, err := doRequest("GET", inst, "/oemanager/applications")
	if err != nil {
		fmt.Println(err)
	} else {
		oeres, _ := readJson(res)
		apps, _ = extractApplications(oeres)
	}

	return apps
}

func ListApplications(inst PasInstance, verbose bool) {

	apps := GetApplications(inst)

	for _, app := range apps {

		fmt.Printf("%v\n", app.Name)

		if verbose {
			for _, webapp := range app.Webapps {
				fmt.Printf("  %v (%v, %v, %v)\n", webapp.Name, getPort(webapp.URI), getPort(webapp.SecureURI), webapp.State)
				for _, transport := range webapp.Transports {
					if transport.State != "ENABLED" {
						continue
					}
					fmt.Printf("    %-4s: %v (%v)\n", transport.Name, getPath(transport.URI), transport.Status)
				}
			}
		}
	}

}

func getPath(tUrl string) string {
	u, _ := url.Parse(tUrl)
	return u.Path
}

func getPort(tUrl string) string {
	u, _ := url.Parse(tUrl)
	return u.Port()
}
