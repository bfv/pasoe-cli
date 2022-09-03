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
	"os"
	"strconv"

	"github.com/bfv/pasoe-cli/model"
)

type ListAgentSessionParams struct {
	Pid               int
	Treshold          int
	AboveTresholdOnly bool
	Verbose           bool
}

type KillAgentSessionParams struct {
	Apps      []string
	SessionId int
	AgentId   string
	Pid       int
	Forced    bool
	Threshold int // in MiB
	KillAll   bool
}

func ListAgentSessions(inst PasInstance, apps []string, params ListAgentSessionParams) {

	if len(apps) == 0 {
		apps = getApplicationNames(inst)
	}

	iterateAgents(inst, apps, func(app string, agent model.Agent) {

		agentPid, _ := strconv.Atoi(agent.Pid)
		if params.Pid > -1 && params.Pid != agentPid {
			return
		}

		headerDisplayed := false
		iterateAgentSessions(inst, app, agent, func(app string, as model.AgentSession) {
			if params.AboveTresholdOnly && as.SessionMemory < (params.Treshold*1000000) {
				return
			}

			if !headerDisplayed {
				fmt.Printf("[%v] sessions for agent: %v (pid: %v)\n", app, agent.AgentId, agent.Pid)
				headerDisplayed = true
			}

			color := ""
			if params.Treshold > -1 && as.SessionMemory > (params.Treshold*1024*1024) {
				color = red
			}
			fmt.Printf("  session: %v, %vmem: %v%v [%v]", as.SessionId, color, ByteCountIEC(as.SessionMemory), reset, as.SessionState)
			if params.Verbose {
				fmt.Printf("    (start: %v, threadId: %v, connectionId: %v, external state: %v)", as.StartTime, as.ThreadId, as.ConnectionId, as.SessionExternalState)
			}
			fmt.Println()
		})
	})
}

func KillAgentSessions(inst PasInstance, apps []string, params KillAgentSessionParams) {

	if len(apps) == 0 {
		apps = getApplicationNames(inst)
	}

	iterateAgents(inst, apps, func(app string, agent model.Agent) {
		iterateAgentSessions(inst, app, agent, func(app string, as model.AgentSession) {
			if evalKill(params, app, agent, as) {
				killSession(inst, app, agent, as.SessionId, params) // candidate for Go routine?
			}
		})
	})
}

func killSession(inst PasInstance, app string, agent model.Agent, sessionId int, params KillAgentSessionParams) error {

	fmt.Printf("[%v] kill session %v on agent %v", app, sessionId, agent.Pid)
	untrappable := "0"
	if params.Forced {
		untrappable = "1"
	}

	path := fmt.Sprintf("/oemanager/applications/%v/agents/%v/sessions/%v?terminateOpt=%v", app, agent.AgentId, sessionId, untrappable)
	res, err := doRequest("DELETE", inst, path)

	if err == nil && res.StatusCode == 200 {
		fmt.Printf("  OK\n")
	} else {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	return nil
}

func evalKill(params KillAgentSessionParams, app string, agent model.Agent, as model.AgentSession) bool {

	//fmt.Printf("eval %v on agentId %v\n", agent.Pid, as.SessionId)
	agentPid, _ := strconv.Atoi(agent.Pid)
	kill := true

	if params.Pid > 0 && params.Pid != agentPid {
		kill = false
	}

	if len(params.Apps) > 0 && !contains(params.Apps, app) {
		kill = false
	}

	kill = kill && ((agent.AgentId == params.AgentId && (as.SessionId == params.SessionId)) ||
		(agentPid == params.Pid && (as.SessionId == params.SessionId || params.KillAll)) ||
		(contains(params.Apps, app) && params.KillAll) ||
		(agent.Pid == params.AgentId && params.KillAll) ||
		(params.Threshold > 0 && as.SessionMemory > params.Threshold*1024*1024) ||
		(params.Pid <= 0 && len(params.Apps) == 0 && params.KillAll))

	return kill

}

func iterateAgentSessions(inst PasInstance, app string, agent model.Agent, f func(app string, as model.AgentSession)) {

	path := fmt.Sprintf("/oemanager/applications/%v/agents/%v/sessions", app, agent.AgentId)
	res, err := doRequest("GET", inst, path)

	if err == nil && res.StatusCode == 200 {
		r1, err := readJson(res)
		if err != nil {
			printError(err)
		}
		agentSessions, _ := extractAgentSessions(r1)
		for _, agentSession := range agentSessions {
			f(app, agentSession)
		}
	} else if res.StatusCode == 404 {
		fmt.Printf("app '%v' not found\n", app)
		os.Exit(1)
	}

}
