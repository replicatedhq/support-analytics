package replicated

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/blang/semver"
	"github.com/replicatedcom/support-analytics/issue"
)

type Versions struct {
	Replicated semver.Version
	Operators  []semver.Version
	UI         semver.Version
}

// ParseVersion get the versions of the replicated daemon, ui and all operators
func ParseVersion(txt string) (*Versions, error) {
	versions := &Versions{}

	lines := strings.Split(txt, "\n")
	for _, line := range lines {
		s := bufio.NewScanner(strings.NewReader(line))
		s.Split(bufio.ScanWords)
		if !s.Scan() {
			continue
		}
		product := s.Text()
		if product == "PRODUCT" || !s.Scan() {
			continue
		}
		versionText := s.Text()
		version, err := semver.Parse(versionText)
		if err != nil {
			log.Fatalf("Cannot parse version: %s", err)
			continue
		}
		switch strings.ToLower(product) {
		case "replicated":
			versions.Replicated = version
		case "replicated-ui":
			versions.UI = version
		case "replicated-operator":
			// TODO: figure out what it looks like with multiple operators, currently it overwrites with the last operator
			versions.Operators = append(versions.Operators, version)
		}
	}

	return versions, nil
}

func (v *Versions) CheckUIVersionValid(i issue.Issues) error {
	if v.Replicated.LTE(v.UI) {
		i.AppendIssue(issue.Error(fmt.Sprintf("Replicated UI version is invalid replicated=%s ui=%s", v.Replicated.String(), v.UI.String())))
		return nil
	}
	if !v.Replicated.EQ(v.UI) {
		i.AppendIssue(issue.Error(fmt.Sprintf("Replicated UI version does not match daemon; replicated=%s ui=%s", v.Replicated.String(), v.UI.String())))
	}
	return nil
}

func (v *Versions) CheckOperatorVersionsValid(i issue.Issues) error {
	for _, operator := range v.Operators {
		if operator.LT(v.Replicated) {
			i.AppendIssue(issue.Error(fmt.Sprintf("Replicated operator version must be greater than replicated; replicated=%s ui=%s", v.Replicated.String(), operator.String())))
		}
		if !v.Replicated.EQ(v.UI) {
			i.AppendIssue(issue.Error(fmt.Sprintf("Replicated Operator does not match daemon; replicated=%s ui=%s", v.Replicated.String(), v.UI.String())))
		}
	}
	return nil
}
