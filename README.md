# Sonar CI
A simple CLI for help you integrate SonarQube inspections with CI pipelines.

[![build](https://github.com/odair-pedro/sonarci/workflows/build/badge.svg)](https://github.com/odair-pedro/sonarci/actions?query=workflow%3ABuild)
[![Coverage Status](https://coveralls.io/repos/github/odair-pedro/sonarci/badge.svg?branch=master)](https://coveralls.io/github/odair-pedro/sonarci?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/odair-pedro/sonarci)](https://goreportcard.com/report/github.com/odair-pedro/sonarci)

[![GitHub version](https://badge.fury.io/gh/odair-pedro%2Fsonarci.svg)](https://github.com/odair-pedro/sonarci/releases/latest)
[![Open Source Helpers](https://www.codetriage.com/odair-pedro/sonarci/badges/users.svg)](https://www.codetriage.com/odair-pedro/sonarci)

## Installation

### Automated
Those who want to get started quickly and easily may install using the following command.
``` 
curl -sSL https://raw.githubusercontent.com/odair-pedro/sonarci/master/install-local.sh | bash
````

### From source
Those who want to install from source code using the following command (on root repo directory):

``` 
make
``` 

###### Note: SonarCI will be installed on current directory

---

## Usage

```
./sonarci --help
```

```
SonarCI is a CLI library for help you integrate and use SonarQube inspections.

Usage:
  sonarci [command]

Available Commands:
  help           Help about any command
  search         Search for SonarQube projects
  server-version Get SonarQube server version
  validate       Validate quality gate status

Flags:
  -h, --help            help for sonarci
  -s, --server string   SonarQube server address
  -t, --timeout int     Timeout in milliseconds. Default value is 30000 ms
  -o, --token string    Authentication Token
  -v, --version         version for sonarci

Use "sonarci [command] --help" for more information about a command.

```

### Search Projects Command
```
./sonarci search --help
```

```
Search and retrieve information about the specified SonarQube projects.

Usage:
  sonarci search [flags]

Flags:
  -h, --help              help for search
  -p, --projects string   SonarQube projects key. Eg: my-sonar-project | my-sonar-project-1,my-sonar-project-2

Global Flags:
  -s, --server string   SonarQube server address
  -t, --timeout int     Timeout in milliseconds. Default value is 30000 ms
  -o, --token string    Authentication Token
```

### Server Version Command
```
./sonarci server-version --help
```

```
Get SonarQube server version

Usage:
  sonarci server-version [flags]

Flags:
  -h, --help   help for server-version

Global Flags:
  -s, --server string   SonarQube server address
  -t, --timeout int     Timeout in milliseconds. Default value is 30000 ms
  -o, --token string    Authentication Token
```

## Building from source
On root directory, run the command:
```
make build-linux
``` 
or 
```
make build-windows
```

## Looking for examples?
Soon...


