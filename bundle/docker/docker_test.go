package docker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseDockerInfo(t *testing.T) {
	json := `{
  "ID": "TWV7:TPPS:TERH:3HXS:4UEJ:2FMF:FG7J:PUUD:N4EX:FRDN:JMZT:PIT3",
  "Containers": 29,
  "ContainersRunning": 15,
  "ContainersPaused": 0,
  "ContainersStopped": 14,
  "Images": 13,
  "Driver": "devicemapper",
  "DriverStatus": [
    [
      "Pool Name",
      "docker-253:1-2359942-pool"
    ],
    [
      "Pool Blocksize",
      "65.54 kB"
    ],
    [
      "Base Device Size",
      "10.74 GB"
    ],
    [
      "Backing Filesystem",
      "xfs"
    ],
    [
      "Data file",
      "/dev/loop0"
    ],
    [
      "Metadata file",
      "/dev/loop1"
    ],
    [
      "Data Space Used",
      "7.358 GB"
    ],
    [
      "Data Space Total",
      "107.4 GB"
    ],
    [
      "Data Space Available",
      "83.43 GB"
    ],
    [
      "Metadata Space Used",
      "9.081 MB"
    ],
    [
      "Metadata Space Total",
      "2.147 GB"
    ],
    [
      "Metadata Space Available",
      "2.138 GB"
    ],
    [
      "Udev Sync Supported",
      "true"
    ],
    [
      "Deferred Removal Enabled",
      "false"
    ],
    [
      "Deferred Deletion Enabled",
      "false"
    ],
    [
      "Deferred Deleted Device Count",
      "0"
    ],
    [
      "Data loop file",
      "/var/lib/docker/devicemapper/devicemapper/data"
    ],
    [
      "Metadata loop file",
      "/var/lib/docker/devicemapper/devicemapper/metadata"
    ],
    [
      "Library Version",
      "1.02.107-RHEL7 (2015-10-14)"
    ]
  ],
  "SystemStatus": null,
  "Plugins": {
    "Volume": [
      "local"
    ],
    "Network": [
      "bridge",
      "null",
      "host"
    ],
    "Authorization": null
  },
  "MemoryLimit": true,
  "SwapLimit": true,
  "KernelMemory": true,
  "CpuCfsPeriod": true,
  "CpuCfsQuota": true,
  "CPUShares": true,
  "CPUSet": true,
  "IPv4Forwarding": true,
  "BridgeNfIptables": false,
  "BridgeNfIp6tables": false,
  "Debug": false,
  "OomKillDisable": true,
  "ExperimentalBuild": false,
  "NFd": 88,
  "NGoroutines": 159,
  "SystemTime": "2016-12-09T14:05:22.972725167+08:00",
  "ExecutionDriver": "",
  "LoggingDriver": "json-file",
  "CgroupDriver": "cgroupfs",
  "NEventsListener": 2,
  "KernelVersion": "3.10.0-327.22.2.el7.x86_64",
  "OperatingSystem": "CentOS Linux 7 (Core)",
  "OSType": "linux",
  "Architecture": "x86_64",
  "IndexServerAddress": "https://index.docker.io/v1/",
  "RegistryConfig": {
    "InsecureRegistryCIDRs": [
      "127.0.0.0/8"
    ],
    "IndexConfigs": {
      "docker.io": {
        "Name": "docker.io",
        "Mirrors": null,
        "Secure": true,
        "Official": true
      }
    },
    "Mirrors": null
  },
  "NCPU": 4,
  "MemTotal": 16658821120,
  "DockerRootDir": "/var/lib/docker",
  "HttpProxy": "",
  "HttpsProxy": "",
  "NoProxy": "",
  "Name": "iZ2ze9lpubc768vzhoyyxiZ",
  "Labels": null,
  "ServerVersion": "1.11.2",
  "ClusterStore": "",
  "ClusterAdvertise": "",
  "Swarm": {
    "NodeID": "",
    "NodeAddr": "",
    "LocalNodeState": "",
    "ControlAvailable": false,
    "Error": "",
    "RemoteManagers": null,
    "Nodes": 0,
    "Managers": 0,
    "Cluster": {
      "ID": "",
      "Version": {},
      "CreatedAt": "0001-01-01T00:00:00Z",
      "UpdatedAt": "0001-01-01T00:00:00Z",
      "Spec": {
        "Orchestration": {},
        "Raft": {
          "ElectionTick": 0,
          "HeartbeatTick": 0
        },
        "Dispatcher": {},
        "CAConfig": {},
        "TaskDefaults": {},
        "EncryptionConfig": {
          "AutoLockManagers": false
        }
      }
    }
  }
}`

	info, err := ParseDockerInfo([]byte(json))
	assert.Nil(t, err)
	assert.Equal(t, "1.11.2", info.ServerVersion)
}
