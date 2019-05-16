package replicated

import (
	"strings"
	"time"
)

const DebugLogLevel = "debug"
const InfoLogLevel = "info"
const ErrorLogLevel = "error"
const WarningLogLevel = "warning"
const PanicLogLevel = "panic"

type LogEntry struct {
	// TODO: set default level to panic in struct
	level      string
	date       *time.Time
	file       string
	lineNumber string
	text       string
	format     int
}

// ParseReplicatedLogs parses logs in for replicated, replicated-ui and replicated-operator from log samples
// found in the wild.
func ParseReplicatedLogs(logs string) ([]LogEntry, error) {
	parsedLogs := make([]LogEntry, 0)
	for _, line := range strings.Split(logs, "\n") {
		words := strings.Split(line, " ")
		entry := newLogEntry()

		parseFormat1(0, line, words, &entry)

		if missingLogData(entry) {
			entry = newLogEntry()
			parseFormat2(0, line, words, &entry)
		}

		if missingLogData(entry) {
			entry = newLogEntry()
			parseFormat3(0, line, words, &entry)
		}

		if missingLogData(entry) {
			entry = newLogEntry()
			parseFormat4(0, line, words, &entry)
		}

		if entry.file != "" || entry.lineNumber != "" || entry.text != "" {
			parsedLogs = append(parsedLogs, entry)
		}
	}
	return parsedLogs, nil
}

func newLogEntry() LogEntry {
	return LogEntry{
		level: PanicLogLevel,
	}
}

func missingLogData(entry LogEntry) bool {
	return entry.date == nil || entry.text == "" || entry.level == PanicLogLevel
}

//INFO 2017/01/13 14:15:26 eventmanager/event_manager.go:30 No event listener registered for event "EventAppStatusRefresh"
func parseFormat1(offset int, line string, words []string, logEntry *LogEntry) {
	if offset > len(words) {
		return
	}

	// Error LEVEL
	if offset == 0 {
		isErrorLevel, level := isErrorLevel(words[0])
		if isErrorLevel {
			logEntry.level = level
			parseFormat1(offset+1, line, words, logEntry)
		}
	}

	// date and time
	if offset == 1 {
		// check if offset 1 is date and offset 2 is time
		if len(words) < 2 {
			return
		}

		date, err := time.Parse("2006/01/02 15:04:05", words[1]+" "+words[2])
		if err != nil {
			return
		}
		logEntry.date = &date

		parseFormat1(offset+2, line, words, logEntry)
	}

	// file and line number
	if offset == 3 {
		if !strings.Contains(words[3], ".go:") {
			return
		}
		parts := strings.Split(words[3], ":")
		logEntry.file = parts[0]
		logEntry.lineNumber = parts[1]
		parseFormat1(offset+1, line, words, logEntry)
	}

	// text
	if offset == 4 {
		startTextOffset := len(words[0]) + len(words[1]) + len(words[2]) + len(words[3]) + 4
		logEntry.text = line[startTextOffset:]
	}

	logEntry.format = 1
}

// 2016/10/18 17:37:44 Autodetected docker host ip 172.17.0.1
func parseFormat2(offset int, line string, words []string, logEntry *LogEntry) {
	// No level exists in this log format, default to info
	logEntry.level = InfoLogLevel

	if offset == 0 {
		// check if offset 0 is date and offset 1 is time
		if len(words) < 1 {
			return
		}

		date, err := time.Parse("2006/01/02 15:04:05", words[0]+" "+words[1])
		if err != nil {
			return
		}
		logEntry.date = &date

		parseFormat2(offset+2, line, words, logEntry)
	}

	if offset == 2 {
		startTextOffset := len(words[0]) + len(words[1]) + 2
		logEntry.text = line[startTextOffset:]
	}

	logEntry.format = 2
}

func isErrorLevel(i string) (bool, string) {
	switch strings.ToLower(i) {
	case "erro", "error":
		return true, ErrorLogLevel
	case "info":
		return true, InfoLogLevel
	case "warn":
		return true, WarningLogLevel
	case "debu", "debug":
		return true, DebugLogLevel
	default:
		return false, PanicLogLevel
	}
}

//ERRO 2017-01-13T14:15:28+00:00 [replicated-native] preflight.go:208 remove /host/etc/os-release: device or resource busy
func parseFormat3(offset int, line string, words []string, logEntry *LogEntry) {
	if offset > len(words) {
		return
	}

	// Error LEVEL
	if offset == 0 {
		isErrorLevel, level := isErrorLevel(words[0])
		if isErrorLevel {
			logEntry.level = level
			parseFormat3(offset+1, line, words, logEntry)
		}
	}

	// date and time
	if offset == 1 {
		// check if offset 1 is both the date and time
		if len(words) < 1 {
			return
		}

		date, err := time.Parse("2006-01-02T15:04:05+00:00", words[1])
		if err != nil {
			return
		}
		logEntry.date = &date

		parseFormat3(offset+1, line, words, logEntry)
	}

	// logging process
	if offset == 2 {
		if strings.HasPrefix(words[2], "[") && strings.HasSuffix(words[2], "]") {
			// TODO: do we need to record the log source?
			parseFormat3(offset+1, line, words, logEntry)
		}
	}

	// file and line number
	if offset == 3 {
		if !strings.Contains(words[3], ".go:") {
			return
		}
		parts := strings.Split(words[3], ":")
		logEntry.file = parts[0]
		logEntry.lineNumber = parts[1]
		parseFormat3(offset+1, line, words, logEntry)
	}

	// text
	if offset == 4 {
		startTextOffset := len(words[0]) + len(words[1]) + len(words[2]) + len(words[3]) + 4
		logEntry.text = line[startTextOffset:]
	}

	logEntry.format = 3
}

// Journald log format
//Feb 23 07:11:59 ip-172-31-10-33.ap-southeast-2.compute.internal docker[22585]: DEBU 2017-02-23T07:11:59+00:00 [replicated-native] events.go:75 Demuxer sending event "container-start" with suffix "7a0b0dfa8d8f4ed47b822cf1ecd920e3"
func parseFormat4(offset int, line string, words []string, logEntry *LogEntry) {
	if offset > len(words) {
		return
	}

	// date and time
	if offset == 0 {
		// check if we have enough words to match date and time
		if len(words) < 2 {
			return
		}

		date, err := time.Parse("Jan 02 15:04:05", words[0]+" "+words[1]+" "+words[2])
		if err != nil {
			return
		}
		logEntry.date = &date

		parseFormat4(offset+5, line, words, logEntry)
	}

	// Error LEVEL
	if offset == 5 {
		isErrorLevel, level := isErrorLevel(words[5])
		if isErrorLevel {
			logEntry.level = level
			parseFormat4(offset+2, line, words, logEntry)
		}
	}

	// logging process
	if offset == 7 {
		if strings.HasPrefix(words[7], "[") && strings.HasSuffix(words[7], "]") {
			// TODO: do we need to record the log source?
			parseFormat4(offset+1, line, words, logEntry)
		}
	}

	// file and line number
	if offset == 8 {
		if !strings.Contains(words[8], ".go:") {
			return
		}
		parts := strings.Split(words[8], ":")
		logEntry.file = parts[0]
		logEntry.lineNumber = parts[1]
		parseFormat4(offset+1, line, words, logEntry)
	}

	// text
	if offset == 9 {
		startTextOffset := strings.Index(line, words[9])
		logEntry.text = line[startTextOffset:]
	}

	logEntry.format = 4
}
