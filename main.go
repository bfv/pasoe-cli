package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	inst := PasInstance{protocol: "http", host: "localhost", port: 8810}

	res, err := doRequest("GET", inst, "/oemanager/applications")

	if err != nil {
		fmt.Println(err)
	} else {
		oeres, _ := readJson(res)
		apps, _ := extractApplicationNames(oeres)
		killAllAgents(apps, inst)
	}

}

func killAllAgents(apps []string, inst PasInstance) {

	for _, app := range apps {

		path := fmt.Sprintf("/oemanager/applications/%v/agents", app)
		res, err := doRequest("GET", inst, path)
		if err == nil {

			r1, err := readJson(res)
			if err != nil {
				printError(err)
			}

			agents, _ := extractAgents(r1)
			for _, agent := range agents {
				res2, _ := doRequest("DELETE", inst, fmt.Sprintf("/oemanager/applications/%v/agents/%v", app, agent.AgentdId))
				if res2.StatusCode == 200 {
					fmt.Printf("[%v] agent killed: %v (pid: %v)\n", app, agent.AgentdId, agent.Pid)
				} else {
					fmt.Printf("error killing agent [%v] %v (pid: %v)", app, agent.AgentdId, agent.Pid)
				}
			}
		} else {
			printError(err)
		}
	}
}

func doRequest(verb string, inst PasInstance, path string) (*http.Response, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%v://%v:%v%v", inst.protocol, inst.host, inst.port, path)
	req, _ := http.NewRequest(verb, url, nil)
	req.SetBasicAuth("tomcat", "tomcat")
	res, err := client.Do(req)
	return res, err
}

func printError(err error) {
	fmt.Printf("error: %v\n", err)
}

func readJson(res *http.Response) (Response, error) {

	defer res.Body.Close()

	resp := Response{}
	err := json.NewDecoder(res.Body).Decode(&resp)

	return resp, err
}

func extractApplicationNames(res Response) ([]string, error) {

	var appResponse ApplicationResponse
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

func extractAgents(res Response) ([]Agent, error) {

	var agentResponse AgentsResponse
	var agents []Agent

	err := json.Unmarshal(res.Result, &agentResponse)
	if err != nil {
		printError(err)
	} else {
		agents = append(agents, agentResponse.Agents...)
	}
	return agents, err
}

// attempt to make this generic
/*
func parseResponse[R ApplicationResponse | AgentsResponse, V Application | Agent](res Response) ([]V, error) {
	var array []V
	var jsonResponse R

	err := json.Unmarshal(res.Result, &jsonResponse)
	if err != nil {
		printError(err)
	} else {
		metaValue := reflect.ValueOf(jsonResponse).Elem()
		if metaValue.FieldByName("Agents") != (reflect.Value{}) {
			array = append(array, metaValue.FieldByName("Agents")...)
		}

	}

	return array, nil
}
*/
