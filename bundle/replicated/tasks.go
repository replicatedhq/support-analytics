package replicated

import (
	"regexp"
	"sort"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

type Tasks struct {
	TimeRange *time.Duration
}

type SortableTime struct {
	theTimeField time.Time
}

type byTime []SortableTime

func (x byTime) Len() int      { return len(x) }
func (x byTime) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x byTime) Less(i, j int) bool {
	return x[i].theTimeField.Before(x[j].theTimeField)
}

// ParseTasks parses the tasks file to determine how long the install has been running for
func ParseTasks(txt string) (*Tasks, error) {
	tasks := &Tasks{}

	var timestamps byTime
	lines := strings.Split(txt, "\n")

	for _, line := range lines {
		// Normalize the line by replacing four or more spaces with a split character
		re := regexp.MustCompile(`\s{4,}`)
		line = re.ReplaceAllString(line, "\t")

		// Use the fields function to break apart the line on the split character
		fields := strings.FieldsFunc(line, pipe)
		if len(fields) < 3 {
			continue
		}

		timestampString := fields[2]
		if timestampString == "STARTED" {
			continue
		}

		timestamp, err := time.Parse("2006-01-02 15:04:05 MST", timestampString)
		if err != nil {
			log.Error(err)
			continue
		}

		timestamps = append(timestamps, SortableTime{theTimeField: timestamp})
	}

	log.Debugf("Tasks timestamps found %d", len(timestamps))
	if len(timestamps) >= 2 {
		sort.Sort(timestamps)
		now := timestamps[0]
		then := timestamps[len(timestamps)-1]
		d := then.theTimeField.Sub(now.theTimeField)
		tasks.TimeRange = &d
	}

	return tasks, nil
}

func pipe(r rune) bool {
	return r == '\t'
}
