package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	flagProjects = "projects"
)

func NewSearchCmd() *cobra.Command {
	searchCmd := &cobra.Command{
		Use:   "search",
		Short: "Search for SonarQube projects",
		Long:  "Search and retrieve information about the specified SonarQube projects.",
		Run:   search,
	}

	searchCmd.Flags().StringP(flagProjects, "p", "", "SonarQube projects key. Eg: my-sonar-project | my-sonar-project-1,my-sonar-project-2")
	_ = searchCmd.MarkFlagRequired(flagProjects)

	return searchCmd
}

func search(cmd *cobra.Command, args []string) {
	_ = args

	projects, _ := cmd.Flags().GetString(flagProjects)
	if !validateFlag(flagProjects, projects) {
		return
	}

	pFlags := getPersistentFlagsFromCmd(cmd)
	if pFlags == nil {
		return
	}

	api := createSonarApi(pFlags.Server, pFlags.Token, pFlags.Timeout)
	results, err := api.SearchProjects(projects)
	if err != nil {
		log.Fatal("Failure to search projects: ", err)
	}

	const colSize = 30
	log.Printf("%s | %s | %s | %s | %s", padRight("project-id ", "-", colSize),
		padRight("project-name ", "-", colSize), padRight("project-key ", "-", colSize),
		padRight("organization ", "-", colSize), padRight("visibility ", "-", colSize))

	for r := range results {
		log.Printf("%s | %s | %s | %s | %s", padRight(r.Id, " ", colSize), padRight(r.Name, " ", colSize),
			padRight(r.Key, " ", colSize), padRight(r.Organization, " ", colSize), padRight(r.Visibility, " ", colSize))
	}
}

func padRight(str string, suffix string, length int) string {
	return (str + strings.Repeat(suffix, length-len(str)))[:length]
}
