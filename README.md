# Sonar CI
A simple CLI for help you integrate SonarQube inspections with your CI pipelines.

[![GitHub](https://img.shields.io/github/license/odair-pedro/sonarci)](https://github.com/odair-pedro/sonarci/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/odair-pedro/sonarci)](https://goreportcard.com/report/github.com/odair-pedro/sonarci)
[![Open Source Helpers](https://www.codetriage.com/odair-pedro/sonarci/badges/users.svg)](https://www.codetriage.com/odair-pedro/sonarci)

## Usage
```
sonarci --help
```

#### Output
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

# Report

## Quality Report
![Status](https://img.shields.io/badge/Status-Failed-Red?style=for-the-badge)

---

[![Status](https://img.shields.io/badge/B-Security%20rating%20on%20New%20Code-red?style=for-the-badge)](https://www.google.com)
[![Status](https://img.shields.io/badge/22.5%25-coverage%20on%20new%20code-red?style=for-the-badge)](https://www.google.com)
[![Status](https://img.shields.io/badge/22.1%25-coverage-red?style=for-the-badge)](https://www.google.com)

### Reliability
| Bugs | New Bugs |
|-|-|
|![Status](https://img.shields.io/badge/A-0-Green?style=for-the-badge)|![Status](https://img.shields.io/badge/A-0-Green?style=for-the-badge)|

### Security
| Vulnerabilities | Security Hotspots | New Vulnerabilities | New Security Hotspots |
|-|-|-|-|
|![Status](https://img.shields.io/badge/B-1-brightgreen?style=for-the-badge)|19|![Status](https://img.shields.io/badge/C-10-yellow?style=for-the-badge)|29


