package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/replicatedcom/preflight/log"
	"github.com/spf13/cobra"
)

type ListAppsResponse struct {
	Apps []*App `json:"apps"`
}

type AppDetailsResponse struct {
	App      AppDetails   `json:"app"`
	Releases []AppRelease `json:"releases"`
}

type AppRelease struct {
	Config          string     `json:"-"`
	Sequence        int        `json:"sequence"`
	Created         time.Time  `json:"created"`
	CurrentChannels []*Channel `json:"current_channels"`
}

type Release struct {
	Sequence     int    `json:"sequence"`
	Label        string `json:"label"`
	ReleaseNotes string `json:"release_notes"`
	Yaml         string `json:"yaml"`
}

type Channel struct {
	Id             string   `json:"id"`
	Name           string   `json:"name"`
	Position       int      `json:"position"`
	CurrentRelease *Release `json:"current_release"`
}

type App struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Team         *Team     `json:"team"`
	ReleaseCount int       `json:"release_count"`
	LicenseCount int       `json:"license_count"`
	Created      time.Time `json:"created"`
	Modified     time.Time `json:"modified"`
	IsTest       bool      `json:"is_test"`
	WebHookURL   string    `json:"web_hook_url"`
}

type Team struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Activated   bool      `json:"is_activated"`
	HasContract bool      `json:"has_contract"`
	Members     []*Member `json:"members"`
}

type Member struct {
	ID       string    `json:"id"`
	Email    string    `json:"email"`
	Created  time.Time `json:"created"`
	ReadOnly string    `json:"read_only"`
}

type AppDetails struct {
	App      *App       `json:"app"`
	Channels []*Channel `json:"channels"`
}

var AppYamlCommand = &cobra.Command{
	Use:   "appyaml",
	Short: "Download YAML",
	Long:  `Download YAML for analysis`,
	Run: func(cmd *cobra.Command, args []string) {
		authToken := cmd.Flag("auth").Value.String()

		// Input is a local file calls apps.json which can be gotten via the admin-api and
		// https://adminapi.replicated.com/v1/apps

		appBytes, err := ioutil.ReadFile("apps.json")
		if err != nil {
			log.Fatal(err)
		}

		var apps ListAppsResponse
		err = json.Unmarshal(appBytes, &apps)
		if err != nil {
			log.Fatal(err)
		}

		for _, app := range apps.Apps {
			fmt.Printf("App %s Team %s ID %s\n", app.Name, app.Team.Name, app.Id)
			dumpApp(authToken, app.Id)
		}
	},
}

func dumpApp(authToken string, appId string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://adminapi.replicated.com/v1/app/"+appId, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+authToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()

	var app AppDetailsResponse
	err = json.Unmarshal([]byte(newStr), &app)
	if err != nil {
		log.Fatal(err)
	}

	for _, channel := range app.App.Channels {
		if channel.CurrentRelease != nil {
			os.MkdirAll("data/"+appId+"/"+channel.Name, 0777)
			fmt.Printf("    writing YAML %s\n", channel.Name)
			ioutil.WriteFile("data/"+appId+"/"+channel.Name+"/release.yaml", []byte(channel.CurrentRelease.Yaml), 0644)
		}
	}
}

func init() {
	AppYamlCommand.PersistentFlags().StringP("auth", "a", "", "auth bearer token")
}
