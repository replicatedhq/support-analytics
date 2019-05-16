package issue

type Issues struct {
	issues []Issue
}

type Issue struct {
	Description string
	Severity    int
}

func (a *Issues) AppendIssue(i Issue) {
	a.issues = append(a.issues, i)
}

const ERROR = 2
const WARNING = 1

func Error(desc string) Issue {
	return Issue{
		Severity:    ERROR,
		Description: desc,
	}
}

func Warning(desc string) Issue {
	return Issue{
		Severity:    WARNING,
		Description: desc,
	}
}
func (i *Issues) BySeverity(severity int) []Issue {
	var ret []Issue
	for _, issue := range i.issues {
		if issue.Severity == severity {
			ret = append(ret, issue)
		}
	}
	return ret
}
