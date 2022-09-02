# pasoe-cli

A standalone CLI tool for managing agents, no (nodejs, Java, etc) runtime needed. The primary aim for now is to list, add and kill agents.
PREREQUISITE: `oemanager` must be installed.

Example:

```
C:\dev\go\pasoe-cli\build>pasoe agent list
[crm] agent: UWGqyGbZSHG0n0KahvDpRA (pid: 12560)
[order] agent: KldZgISgT56OTeocwxQ6RQ (pid: 20112)
[order] agent: R-SAT9dwQeax3Qwn5Tm2Xg (pid: 24852)
```
(this pasoe instance has two ABL applications, crm & order)

More:
```
pasoe agent add --app crm -n 2
pasoe agen kill --app order --all
```

For help on the commands:\
`pasoe --help`

There some global option to this CLI tool:
```
A CLI tool for PAS instance (agent) management

Usage:
  pasoe [command]

Available Commands:
  agent        list/add/kill agents
  agentsession A command handling agent sessions
  application  Command for action related to (OEABL) applications
  completion   Generate the autocompletion script for the specified shell
  help         Help about any command
  version      Displays the version number

Flags:
      --config string     config file (default is $HOME/.pasoe-cli.yaml)
  -h, --help              help for pasoe
      --host string       hostname (default "localhost")
  -p, --password string   password (default "tomcat")
      --port int          port (default 8810)
      --protocol string   protocol, http or https (default "http")
  -u, --user string       username (default "tomcat")

Use "pasoe [command] --help" for more information about a command.
```

By: Bronco Oostermeyer, 2022
License: GPL3
