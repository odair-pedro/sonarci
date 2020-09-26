package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"sonarci/sonar/sonarrestapi"
	"strings"
)

const (
	flagProjects = "projects"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for SonarQube projects",
	Long:  "Search and retrieve information about the specified SonarQube projects.",
	Run:   search,
}

func init() {
	searchCmd.Flags().StringP(flagProjects, "p", "", "SonarQube projects key. Eg: my-sonar-project | my-sonar-project-1,my-sonar-project-2")
	_ = rootCmd.MarkFlagRequired(flagProjects)
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

	api := sonarrestapi.NewApi(pFlags.Server, pFlags.Token, pFlags.Timeout)
	results, err := api.SearchProjects(projects)
	if err != nil {
		log.Fatalln("Failure to search projects: ", err)
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

func padRight(str string, sufix string, count int) string {
	return str + strings.Repeat(sufix, count-len(str))
}
