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
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bfv/pasoe-cli/model"
)

func printError(err error) {
	fmt.Printf("error: %v\n", err)
}

func readJson(res *http.Response) (model.Response, error) {

	defer res.Body.Close()

	resp := model.Response{}
	err := json.NewDecoder(res.Body).Decode(&resp)

	return resp, err
}

func extractApplications(res model.Response) ([]model.Application, error) {

	var appResponse model.ApplicationResponse
	var apps []model.Application

	err := json.Unmarshal(res.Result, &appResponse)
	if err != nil {
		printError(err)
	} else {
		apps = append(apps, appResponse.Applications...)
	}
	return apps, err
}

func extractApplicationNames(res model.Response) ([]string, error) {

	var appResponse model.ApplicationResponse
	var apps []string

	err := json.Unmarshal(res.Result, &appResponse)
	if err != nil {
		printError(err)
	} else {
		for _, app := range appResponse.Applications {
			apps = append(apps, app.Name)
		}
	}
	return apps, err
}

func extractAgents(res model.Response) ([]model.Agent, error) {

	var agentResponse model.AgentsResponse
	var agents []model.Agent

	err := json.Unmarshal(res.Result, &agentResponse)
	if err != nil {
		printError(err)
	} else {
		agents = append(agents, agentResponse.Agents...)
	}
	return agents, err
}

func extractAgentSessions(res model.Response) ([]model.AgentSession, error) {
	var agentSessionResponse model.AgentSessionsReponse
	var agentSessions []model.AgentSession

	err := json.Unmarshal(res.Result, &agentSessionResponse)
	if err != nil {
		printError(err)
	} else {
		agentSessions = append(agentSessions, agentSessionResponse.AgentsSessions...)
	}

	return agentSessions, err
}
