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
	"log"
	"os"

	"github.com/bfv/pasoe-cli/model"
)

func AddAgents(inst PasInstance, apps []string, n int) {

	if len(apps) == 0 {
		apps = getApplicationNames(inst)
	}

	if len(apps) != 1 {
		log.Fatal("can't determine which application to add agents to, use --app parameter")
	}

	path := fmt.Sprintf("/oemanager/applications/%v/addAgent", apps[0])
	for n > 0 {
		doRequest("POST", inst, path)
		n--
	}
}

type KillAgentParams struct {
	Apps   []string
	Number int
	Ids    []string
}

func KillAgents(inst PasInstance, params KillAgentParams) {

	if len(params.Apps) == 0 {
		params.Apps = getApplicationNames(inst)
	}

	iterateAgents(inst, params.Apps, func(app string, agent model.Agent) {

		if len(params.Ids) > 0 && !contains(params.Ids, agent.AgentdId) {
			return
		}

		res2, _ := doRequest("DELETE", inst, fmt.Sprintf("/oemanager/applications/%v/agents/%v", app, agent.AgentdId))
		if res2.StatusCode == 200 {
			fmt.Printf("[%v] agent killed: %v (pid: %v)\n", app, agent.AgentdId, agent.Pid)
		} else {
			fmt.Printf("error killing agent [%v] %v (pid: %v)", app, agent.AgentdId, agent.Pid)
			os.Exit(1)
		}

		params.Number--
		if params.Number == 0 {
			os.Exit(0)
		}
	})
}

func ListAgents(inst PasInstance, apps []string) {

	if len(apps) == 0 {
		apps = getApplicationNames(inst)
	}

	iterateAgents(inst, apps, func(app string, agent model.Agent) {
		fmt.Printf("[%v] agent: %v (pid: %v)\n", app, agent.AgentdId, agent.Pid)
	})
}

func iterateAgents(inst PasInstance, apps []string, f func(app string, agents model.Agent)) error {

	var err error

	for _, app := range apps {

		path := fmt.Sprintf("/oemanager/applications/%v/agents", app)
		res, err := doRequest("GET", inst, path)
		if err == nil && res.StatusCode == 200 {
			r1, err := readJson(res)
			if err != nil {
				printError(err)
			}
			agents, _ := extractAgents(r1)
			for _, agent := range agents {
				f(app, agent)
			}
		} else if res.StatusCode == 404 {
			fmt.Printf("app '%v' not found\n", app)
			os.Exit(1)
		}
	}
	return err
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
