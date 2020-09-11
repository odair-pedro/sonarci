package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for SonarQube project",
	Long:  "Search and retrieve information about the specified SonarQube project.",
	Run:   search,
}

func search(cmd *cobra.Command, args []string) {
	server, err := cmd.Parent().PersistentFlags().GetString("server")
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
		return
	}

	persistentFlags := cmd.Parent().PersistentFlags()
	username, _ := persistentFlags.GetString("username")
	project, _ := persistentFlags.GetString("project")

	log.Printf("Search was called with server: %s", server)
	respBuff, _ := getResponse(server, username, project)
	log.Print(string(respBuff))

	respModel := &wrapper{}
	json.Unmarshal(respBuff, respModel)

	fmt.Printf("\nProject information")
	fmt.Printf("\nId: %s", respModel.Components[0].Id)
	fmt.Printf("\nOrganization: %s", respModel.Components[0].Organization)
	fmt.Printf("\nName: %s", respModel.Components[0].Name)
	fmt.Printf("\nKey: %s", respModel.Components[0].Key)
	fmt.Printf("\nVisibility: %s\n", respModel.Components[0].Visibility)
}

func getResponse(server string, username string, project string) ([]byte, error) {
	client := &http.Client{Timeout: 30 * time.Second}

	url := fmt.Sprintf("%s/api/projects/search?projects=%s", server, project)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	authHeader := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:", username)))
	req.Header.Add("Authorization", "Basic "+authHeader)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	log.Printf("Status code: %d", resp.StatusCode)

	//json.NewDecoder(resp.Body).Decode(result)

	buff, _ := ioutil.ReadAll(resp.Body)
	result := &wrapper{}
	json.Unmarshal(buff, result)
	return buff, nil
}

type wrapper struct {
	Components []component `json:"components"`
}

type component struct {
	Id           string `json:"id"`
	Organization string `json:"organization"`
	Key          string `json:"key"`
	Name         string `json:"name"`
	Visibility   string `json:"visibility"`
}
