package bundle

import (
	"fmt"
	"io/ioutil"
	"strings"
	//"syscall"
	//"time"

	"github.com/replicatedcom/support-analytics/bundle/docker"
	"github.com/replicatedcom/support-analytics/bundle/os"
	repl "github.com/replicatedcom/support-analytics/bundle/replicated"
	"github.com/replicatedcom/support-analytics/issue"

	log "github.com/Sirupsen/logrus"
	//"github.com/dustin/go-humanize"
)

type Analyze struct {
	BasePath string
}

func (a *Analyze) AnalyzeSupportBundle() error {
	i := issue.Issues{}

	// TODO: add analysis for replicated-logs, look for errors
	// TODO: add analysis for replicated-logs, look for timestamps that go backwards

	// TODO: move the version analysis into its own file
	// TODO: look at the docker file for health check
	// TODO: look at all containers and those that are not marked as ephemeral must have good exit codes

	versionFile := a.absPath("daemon/replicated/replicated-versions.txt")
	versionBytes, err := ioutil.ReadFile(versionFile)
	if err != nil {
		log.Fatal(err)
		return err
	}
	v, err := repl.ParseVersion(string(versionBytes))
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Debugf("Found replicated version %s", v.Replicated.String())
	log.Debugf("Found replicated-ui version %s", v.UI.String())
	for _, operator := range v.Operators {
		log.Debugf("Found replicated-operator version %s", operator.String())
	}

	//var st syscall.Stat_t
	//if err := syscall.Stat(a.absPath("daemon/etc/os-release"), &st); err != nil {
	//	log.Error(err)
	//}
	//ctime := time.Unix(int64(st.Birthtimespec.Sec), int64(st.Birthtimespec.Nsec))
	//relativeTimeToCreate := humanize.Time(ctime)

	osFile := a.absPath("daemon/etc/os-release")
	osFileBytes, err := ioutil.ReadFile(osFile)
	if err != nil {
		log.Error(err)
		return err
	}
	osData, err := os.Init(string(osFileBytes))
	if err != nil {
		log.Error(err)
		return err
	}

	dockerFile := a.absPath("daemon/docker/docker_info.json")
	dockerFileBytes, err := ioutil.ReadFile(dockerFile)
	if err != nil {
		log.Error(err)
		return err
	}
	dockerInfo, err := docker.ParseDockerInfo(dockerFileBytes)
	if err != nil {
		log.Error(err)
		return err
	}

	tasksFile := a.absPath("daemon/replicated/tasks.txt")
	tasksFileBytes, err := ioutil.ReadFile(tasksFile)
	if err != nil {
		log.Error(err)
		return err
	}
	tasksInfo, err := repl.ParseTasks(string(tasksFileBytes))
	if err != nil {
		log.Error(err)
		return err
	}

	// Verify that the ui and operators are higher than replicated
	v.CheckUIVersionValid(i)
	v.CheckOperatorVersionsValid(i)

	fmt.Print("Environment\n")
	fmt.Printf("  Replicated version %s\n", v.Replicated)
	fmt.Printf("  OS %s %s\n", osData.ID, osData.Version.String())
	fmt.Printf("  Docker %s\n", dockerInfo.ServerVersion)

	timeRange := "unknown"
	if tasksInfo.TimeRange != nil {
		timeRange = tasksInfo.TimeRange.String()
	}
	fmt.Printf("  Uptime %s\n", timeRange)

	//fmt.Printf("  Taken %s\n", relativeTimeToCreate)

	// TODO: include if its a proxy install

	// TODO: determine the timestamp of the logs, print the errors in the last 2 hours

	// Print errors
	fmt.Print("\nErrors\n")
	printIssues(i.BySeverity(issue.ERROR))

	// Print warnings
	fmt.Print("\nWarnings\n")
	printIssues(i.BySeverity(issue.ERROR))

	return nil
}

func printIssues(issues []issue.Issue) {
	if len(issues) == 0 {
		fmt.Print("  * None found\n")
		return
	}
	for _, iss := range issues {
		fmt.Printf("  * %s\n", iss.Description)
	}
}

func (a *Analyze) absPath(path string) string {
	return strings.Join([]string{a.BasePath, path}, "/")
}
