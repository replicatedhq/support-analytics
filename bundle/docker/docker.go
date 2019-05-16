package docker

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
)

type DockerInfo struct {
	ServerVersion string `json:"ServerVersion"`
}

func ParseDockerInfo(jsonBlob []byte) (*DockerInfo, error) {
	info := DockerInfo{}
	err := json.Unmarshal(jsonBlob, &info)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &info, nil
}
