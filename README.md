# Sonar CI
A simple CLI for help you integrate SonarQube inspections with your CI pipelines.

[![Build](https://github.com/odair-pedro/sonarci/workflows/Build/badge.svg?branch=master)](https://github.com/odair-pedro/sonarci/actions?query=workflow%3ABuild)
[![Coverage Status](https://coveralls.io/repos/github/odair-pedro/sonarci/badge.svg?branch=master)](https://coveralls.io/github/odair-pedro/sonarci?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/odair-pedro/sonarci)](https://goreportcard.com/report/github.com/odair-pedro/sonarci)

[![GitHub](https://img.shields.io/github/license/odair-pedro/sonarci)](https://github.com/odair-pedro/sonarci/blob/master/LICENSE)
[![Open Source Helpers](https://www.codetriage.com/odair-pedro/sonarci/badges/users.svg)](https://www.codetriage.com/odair-pedro/sonarci)

## Usage

```
./sonarci --help
```

```
SonarQubeFast is a CLI library for help you integrate and use SonarQube inspections.

Usage:
  sonarqubeci [command]

Available Commands:
  help        Help about any command
  search      Search for SonarQube projects
  version     Get SonarQube server version

Flags:
  -h, --help            help for sonarqubeci
  -s, --server string   SonarQube server address
  -t, --timeout int     Timeout in milliseconds. Default value is 30000 ms
  -o, --token string    Authentication Token

Use "sonarqubeci [command] --help" for more information about a command.

```

### Search Command
```
./sonarci search --help
```

```
Search and retrieve information about the specified SonarQube projects.

Usage:
  sonarqubeci search [flags]

Flags:
  -h, --help              help for search
  -p, --projects string   SonarQube projects key. Eg: my-sonar-project | my-sonar-project-1,my-sonar-project-2

Global Flags:
  -s, --server string   SonarQube server address
  -t, --timeout int     Timeout in milliseconds. Default value is 30000 ms
  -o, --token string    Authentication Token

```

## Looking for examples?
Soon...


