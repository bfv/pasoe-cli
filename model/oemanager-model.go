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
package model

import "encoding/json"

type Response struct {
	Operation     string          `json:"operation"`
	Outcome       string          `json:"succes"`
	Result        json.RawMessage `json:"result"`
	ErrorMessage  string          `json:"errmsg"`
	VersionString string          `json:"versionStr"`
	VersionNr     int             `json:"versionNo"`
}

type ApplicationResponse struct {
	Applications []Application `json:"Application"`
}

type Application struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	Webapps     []Webapp `json:"webapps"`
	OEType      string   `json:"oetype"`
}

type Webapp struct {
	Name        string      `json:"name"`
	URI         string      `json:"uri"`
	SecureURI   string      `json:"securedUri"`
	Description string      `json:"description"`
	Type        string      `json:"type"`
	State       string      `json:"state"`
	Transports  []Transport `json:"transports"`
}

type Transport struct {
	Name        string `json:"name"`
	URI         string `json:"uri"`
	SecureURI   string `json:"securedUri"`
	Description string `json:"description"`
	State       string `json:"state"`
	Status      string `json:"status"`
	OEType      string `json:"oetype"`
}

type AgentsResponse struct {
	Agents []Agent `json:"agents"`
}

type Agent struct {
	AgentdId string `json:"agentId"`
	Pid      string `json:"pid"`
	State    string `json:"state"`
}

type AgentSessionsReponse struct {
	AgentsSessions []AgentSession `json:"AgentSession"`
}

type AgentSession struct {
	SessionId            int    `json:"SessionId"`
	SessionState         string `json:"SessionState"`
	StartTime            string `json:"StartTime"`
	EndTime              string `json:"EndTime"`
	ThreadId             int    `json:"ThreadId"`
	ConnectionId         int    `json:"ConnectionId"`
	SessionExternalState int    `json:"SessionExternalState"`
	SessionMemory        int    `json:"SessionMemory"`
}
