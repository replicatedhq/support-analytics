package os

import (
	"strings"

	"bytes"
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	"github.com/blang/semver"
)

type OS struct {
	ID      string
	Version semver.Version
}

type OSDetect struct {
	ID         string
	Version_ID string
}

func Init(osReleaseString string) (*OS, error) {
	osReleaseString = makeIntoToml(osReleaseString)
	osDetect := OSDetect{}
	_, err := toml.Decode(osReleaseString, &osDetect)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	version, err := semver.Parse(makeSemanticVersion(osDetect.Version_ID))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &OS{
		ID:      osDetect.ID,
		Version: version,
	}, nil
}
func makeIntoToml(str string) string {
	var lines []string
	for _, line := range strings.Split(str, "\n") {
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			// TODO: ensure the right hand side is enclosed in quotes
			if !strings.Contains(parts[1], "\"") {
				parts[1] = "\"" + parts[1] + "\""
			}
			lines = append(lines, strings.Join(parts, "="))
		} else {
			lines = append(lines, line)
		}
	}
	return strings.Join(lines, "\n")
}

func makeSemanticVersion(version string) string {
	versionArray := strings.Split(version, ".")
	buildReturn := []string{"0", "0", "0"}
	for index, versionSlice := range versionArray {
		buildReturn[index] = removeZeroPadding(versionSlice)
	}
	return strings.Join(buildReturn, ".")
}

// Remove prepended zero's, so 04 becomes 4
func removeZeroPadding(str string) string {
	var buffer bytes.Buffer
	foundNonZeroChar := false
	for _, ch := range str {
		if string(ch) == "0" && !foundNonZeroChar {
			continue
		}
		foundNonZeroChar = true
		buffer.WriteString(string(ch))
	}
	return buffer.String()
}
