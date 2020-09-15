package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"sonarci/sonar"
	"time"
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
	_ = rootCmd.MarkPersistentFlagRequired(flagProjects)
}

func search(cmd *cobra.Command, args []string) {
	_ = args
	projects, _ := cmd.Flags().GetString(flagProjects)
	if !validateFlag(flagProjects, projects) {
		return
	}

	persistentFlags := cmd.Parent().PersistentFlags()
	server, _ := persistentFlags.GetString(flagServer)
	if !validateFlag(flagToken, server) {
		return
	}

	token, _ := persistentFlags.GetString(flagToken)
	if !validateFlag(flagToken, token) {
		return
	}

	timeout, _ := persistentFlags.GetInt(flagTimout)
	if timeout == 0 {
		timeout = timeoutDefault
	}

	api := sonar.NewApi(server, token, time.Duration(timeout)*time.Millisecond)
	results, err := api.SearchProjects(projects)
	if err != nil {
		log.Fatalln("Failure to search projects: ", err)
	}

	for r := range results {
		log.Println("\nProject Id   : ", r.Id)
		log.Println("Project Name : ", r.Name)
		log.Println("Project Key  : ", r.Key)
		log.Println("Organization : ", r.Organization)
		log.Println("Visibility   : ", r.Visibility)
	}
}
