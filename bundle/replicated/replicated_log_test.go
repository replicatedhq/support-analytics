package replicated

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogFormat1(t *testing.T) {
	LogSnippet := `2016/10/18 17:37:44 Autodetected docker host ip 172.17.0.1
2016/10/18 17:37:44 Environment variable 'DAEMON_REGISTRY_ENDPOINT' is not set, using implicit ip address from 'DAEMON_ENDPOINT'
2016/10/18 17:37:44 No autoconfig file found. This is a new operator.`

	parsedLogs, err := ParseReplicatedLogs(LogSnippet)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(parsedLogs))

	entry := parsedLogs[0]
	assert.Equal(t, InfoLogLevel, entry.level)
	assert.NotNil(t, entry.date)
	assert.Equal(t, "Autodetected docker host ip 172.17.0.1", entry.text)
	assert.Equal(t, 2, entry.format)
}

func TestLogFormat2(t *testing.T) {
	LogSnippet := `INFO 2017/01/13 14:15:26 daemon/daemon.go:90 Starting replicated daemon v2.4.1
INFO 2017/01/13 14:15:26 crypto/local_cipher.go:33 Existing key+nonce loaded from filesystem
INFO 2017/01/13 14:15:26 eventmanager/event_manager.go:30 No event listener registered for event "EventAppStatusRefresh"
INFO 2017/01/13 14:15:28 eventmanager/event_manager.go:30 No event listener registered for event "RELOAD_AGENT_CLIENT_CA"`

	parsedLogs, err := ParseReplicatedLogs(LogSnippet)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(parsedLogs))

	entry := parsedLogs[0]
	assert.Equal(t, InfoLogLevel, entry.level)
	assert.NotNil(t, entry.date)
	assert.Equal(t, "daemon/daemon.go", entry.file)
	assert.Equal(t, "90", entry.lineNumber)
	assert.Equal(t, "Starting replicated daemon v2.4.1", entry.text)
	assert.Equal(t, 1, entry.format)
}

func TestLogFormat3(t *testing.T) {
	LogSnippet := `ERRO 2017-01-13T14:15:28+00:00 [replicated-native] preflight.go:208 remove /host/etc/os-release: device or resource busy`

	parsedLogs, err := ParseReplicatedLogs(LogSnippet)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parsedLogs))

	entry := parsedLogs[0]
	assert.Equal(t, 3, entry.format)
	assert.Equal(t, ErrorLogLevel, entry.level)
	assert.NotNil(t, entry.date)
	assert.Equal(t, "preflight.go", entry.file)
	assert.Equal(t, "208", entry.lineNumber)
	assert.Equal(t, "remove /host/etc/os-release: device or resource busy", entry.text)
}

func TestLogFormat4(t *testing.T) {
	LogSnippet := `Feb 23 07:11:59 ip-172-31-10-33.ap-southeast-2.compute.internal docker[22585]: DEBU 2017-02-23T07:11:59+00:00 [replicated-native] events.go:75 Demuxer sending event "container-start" with suffix "7a0b0dfa8d8f4ed47b822cf1ecd920e3"
Feb 23 07:08:01 ip-172-31-10-33.ap-southeast-2.compute.internal systemd[1]: replicated-operator.service stopping timed out. Terminating.`
	// TODO: make the last line parsable, it contains a standard preamble but and the replicated log is different to parse then systemd output

	parsedLogs, err := ParseReplicatedLogs(LogSnippet)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parsedLogs))

	entry := parsedLogs[0]
	assert.Equal(t, 4, entry.format)
	assert.Equal(t, DebugLogLevel, entry.level)
	assert.NotNil(t, entry.date)
	assert.Equal(t, "events.go", entry.file)
	assert.Equal(t, "75", entry.lineNumber)
	assert.Equal(t, `Demuxer sending event "container-start" with suffix "7a0b0dfa8d8f4ed47b822cf1ecd920e3"`, entry.text)
}
