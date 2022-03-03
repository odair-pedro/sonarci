
![cover](https://user-images.githubusercontent.com/26411431/156560959-55277853-6983-49ef-a598-e8e24774edf8.png)

A simple CLI to help you integrate Sonar inspections with CI pipelines.

[![build](https://github.com/odair-pedro/sonarci/workflows/build/badge.svg)](https://github.com/odair-pedro/sonarci/actions?query=workflow%3ABuild)
[![Coverage Status](https://coveralls.io/repos/github/odair-pedro/sonarci/badge.svg?branch=master)](https://coveralls.io/github/odair-pedro/sonarci?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/odair-pedro/sonarci)](https://goreportcard.com/report/github.com/odair-pedro/sonarci)

[![GitHub version](https://img.shields.io/github/v/release/odair-pedro/sonarci?label=version&color=brightgreen)](https://github.com/odair-pedro/sonarci/releases/latest)
[![Docker Version](https://img.shields.io/docker/v/odairpedro/sonarci-cli?label=docker%20version)](https://hub.docker.com/r/odairpedro/sonarci-cli)
[![Open Source Helpers](https://www.codetriage.com/odair-pedro/sonarci/badges/users.svg)](https://www.codetriage.com/odair-pedro/sonarci)

---

## What does it do?
Sor far, basically, SonarCI can decorate your pull requests with a quality report (and also other cool things :smiley:)

**Like that:**

### Quality Report

![Status](https://img.shields.io/badge/status-FAILED-red?style=for-the-badge)

| Metric | Rating |
|-|-|
| New Coverage |![Status](https://img.shields.io/badge/49.56%25-FAILED-red?style=for-the-badge)|
| New Duplicated Lines Density |![Status](https://img.shields.io/badge/0.00%25-SUCCESS-brightgreen?style=for-the-badge)|
| New Code Smells |[![Status](https://img.shields.io/badge/0-SUCCESS-brightgreen?style=for-the-badge)]()|
| Reliability |![Status](https://img.shields.io/badge/SUCCESS-brightgreen?style=for-the-badge)|
| Security |![Status](https://img.shields.io/badge/SUCCESS-brightgreen?style=for-the-badge)|
| Maintainability |![Status](https://img.shields.io/badge/SUCCESS-brightgreen?style=for-the-badge)|

**Or something simpler, if you want:**

![image](https://user-images.githubusercontent.com/26411431/156571196-396843ca-c3ae-4f63-813f-c91a135fed68.png)


---

## Installation

To get started quickly and easily install use the following commands.

### Global

To install globally, run the following command:
``` 
curl -sSL https://raw.githubusercontent.com/odair-pedro/sonarci/master/install.sh | sudo sh
````

Note: If the command sonarci fails after installation, check your path. You can also create a symbolic link to /usr/bin or any other directory in your path.

Example:
```
sudo ln -s /usr/local/bin/sonarci /usr/bin/sonarci
```

### Local

To install locally in current directory, run the following command:
``` 
curl -sSL https://raw.githubusercontent.com/odair-pedro/sonarci/master/install-local.sh | sh
````

### From source
To build and install locally from source code use the following command (on root directory):

``` 
make install
``` 

Note: SonarCI will be installed on current directory

---

## Runing with Docker

If you prefer to use SonarCI CLI with Docker, just run the following command:

```
docker run odairpedro/sonarci-cli sonarci [command]
```

or the following command for help and usage:

```
docker run odairpedro/sonarci-cli sonarci --help
```

See more in https://hub.docker.com/repository/docker/odairpedro/sonarci-cli

---
## Usage

```
sonarci --help
```

```
SonarCI is a CLI library for help you integrate and use SonarQube inspections.

Usage:
  sonarci [command]

Available Commands:
  decorate       Decorate pull request with the quality gate report
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

### Command: Decorate

Decorate pull request with the SonarQube's quality gate report.

```
sonarci decorate --help
```

```
Decorate pull request with the SonarQube's quality gate report.

Usage:
  sonarci decorate [flags]

Flags:
  -h, --help                  help for decorate
  -p, --project string        SonarQube projects key
  -r, --pull-request string   Pull request ID

Global Flags:
  -s, --server string   SonarQube server address
  -t, --timeout int     Timeout in milliseconds. Default value is 30000 ms
  -o, --token string    Authentication Token
```

For more detail about how to configure pull request decoration, read the section [Environment Variables](#environment-variables)

### Command: Search projects

Search and retrieve information about the specified SonarQube projects.

```
sonarci search --help
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

### Comand: Server version

Get SonarQube server version.

```
sonarci server-version --help
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

### Command: Validate

Validate a branch or pull request status on SonarQube.

```
sonarci validate --help
```

```
Validate a branch or pull request status on SonarQube.

Usage:
  sonarci validate [command]

Available Commands:
  branch      Validate branch status
  pr          Validate pull request status

Flags:
  -h, --help   help for validate

Global Flags:
  -s, --server string   SonarQube server address
  -t, --timeout int     Timeout in milliseconds. Default value is 30000 ms
  -o, --token string    Authentication Token
```

#### Branch validation

Validate a branch status on SonarQube.

```
sonarci validate branch --help
```

```
Validate a branch status on SonarQube.

Usage:
  sonarci validate branch [branch name] [flags]

Flags:
  -h, --help             help for branch
  -p, --project string   SonarQube projects key

Global Flags:
  -s, --server string   SonarQube server address
  -t, --timeout int     Timeout in milliseconds. Default value is 30000 ms
  -o, --token string    Authentication Token
```

#### Pull request validation

Validate a pull request status on SonarQube.

```
sonarci validate pr --help
```

```
Validate a pull request status on SonarQube.

Usage:
  sonarci validate pr [pull request id] [flags]

Flags:
  -d, --decorate         Decorate a pull request with quality gate results
  -h, --help             help for pr
  -p, --project string   SonarQube projects key

Global Flags:
  -s, --server string   SonarQube server address
  -t, --timeout int     Timeout in milliseconds. Default value is 30000 ms
  -o, --token string    Authentication Token
```

### Environment Variables

For pull request decoration you need to set the following environment variables:

|Name|Type|Description|
|-|-|-|
|SONARCI_DECORATION_TYPE|String|Allowed values: azrepos or github|
|SONARCI_DECORATION_PROJECT|String|Project URI|
|SONARCI_DECORATION_REPOSITORY|String|Repository name|
|SONARCI_DECORATION_TOKEN|String|PAT|

---

## Looking for examples?

### Azure Pipeline

You can use the following example to integrate a pipeline pull request analyzer with SonarCI

```yaml
trigger:
- none

pool:
  vmImage: 'ubuntu-latest'

steps:
- script: |
    curl -sSL https://raw.githubusercontent.com/odair-pedro/sonarci/master/install.sh | sudo sh

    export SONARCI_DECORATION_TYPE="azrepos"
    export SONARCI_DECORATION_PROJECT="$(AZDEVEOPS-COMPANY-NAME)/$(System.TeamProject)"
    export SONARCI_DECORATION_REPOSITORY="$(Build.Repository.Name)"
    export SONARCI_DECORATION_TOKEN=$(AZDEVEOPS-PAT)

    sonarci --version
    sonarci -s "$(SONAR-SERVER)" -o $(SONAR-TOKEN) decorate -r $(System.PullRequest.PullRequestId) -p "$(SONAR-PROJECT-KEY)"
  displayName: 'SonarCI'

```



### GitHub Action
Coming soon ...
