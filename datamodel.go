package main

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
	Name   string `json:"name"`
	Type   string `json:"type"`
	OEType string `json:"oetype"`
}

type AgentsResponse struct {
	Agents []Agent `json:"agents"`
}

type Agent struct {
	AgentdId string `json:"agentId"`
	Pid      string `json:"pid"`
	State    string `json:"state"`
}
