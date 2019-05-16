package os

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCentOSDetection(t *testing.T) {
	centOS := `NAME="CentOS Linux"
VERSION="7 (Core)"
ID="centos"
ID_LIKE="rhel fedora"
VERSION_ID="7"
PRETTY_NAME="CentOS Linux 7 (Core)"
ANSI_COLOR="0;31"
CPE_NAME="cpe:/o:centos:centos:7"
HOME_URL="https://www.centos.org/"
BUG_REPORT_URL="https://bugs.centos.org/"

CENTOS_MANTISBT_PROJECT="CentOS-7"
CENTOS_MANTISBT_PROJECT_VERSION="7"
REDHAT_SUPPORT_PRODUCT="centos"
REDHAT_SUPPORT_PRODUCT_VERSION="7"
`
	os, err := Init(centOS)
	assert.Nil(t, err)
	assert.Equal(t, "centos", os.ID)
	assert.Equal(t, "7.0.0", os.Version.String())
}

func TestUbuntuOSDetection(t *testing.T) {
	centOS := `NAME="Ubuntu"
VERSION="14.04.5 LTS, Trusty Tahr"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 14.04.5 LTS"
VERSION_ID="14.04"
HOME_URL="http://www.ubuntu.com/"
SUPPORT_URL="http://help.ubuntu.com/"
BUG_REPORT_URL="http://bugs.launchpad.net/ubuntu/"
`
	os, err := Init(centOS)
	assert.Nil(t, err)
	assert.Equal(t, "ubuntu", os.ID)
	assert.Equal(t, "14.4.0", os.Version.String())
}

func TestFedoraOSDetection(t *testing.T) {
	centOS := `NAME="Red Hat Enterprise Linux Server"
VERSION="7.2 (Maipo)"
ID="rhel"
ID_LIKE="fedora"
VERSION_ID="7.2"
PRETTY_NAME="Red Hat Enterprise Linux Server 7.2 (Maipo)"
ANSI_COLOR="0;31"
CPE_NAME="cpe:/o:redhat:enterprise_linux:7.2:GA:server"
HOME_URL="https://www.redhat.com/"
BUG_REPORT_URL="https://bugzilla.redhat.com/"

REDHAT_BUGZILLA_PRODUCT="Red Hat Enterprise Linux 7"
REDHAT_BUGZILLA_PRODUCT_VERSION=7.2
REDHAT_SUPPORT_PRODUCT="Red Hat Enterprise Linux"
REDHAT_SUPPORT_PRODUCT_VERSION="7.2"
`
	os, err := Init(centOS)
	assert.Nil(t, err)
	assert.Equal(t, "rhel", os.ID)
	assert.Equal(t, "7.2.0", os.Version.String())
}

func TestDebianOSDetection(t *testing.T) {
	centOS := `PRETTY_NAME="Debian GNU/Linux 8 (jessie)"
NAME="Debian GNU/Linux"
VERSION_ID="8"
VERSION="8 (jessie)"
ID=debian
HOME_URL="http://www.debian.org/"
SUPPORT_URL="http://www.debian.org/support"
BUG_REPORT_URL="https://bugs.debian.org/"
`
	os, err := Init(centOS)
	assert.Nil(t, err)
	assert.Equal(t, "debian", os.ID)
	assert.Equal(t, "8.0.0", os.Version.String())
}

func TestOracleLinuxOSDetection(t *testing.T) {
	centOS := `NAME="Oracle Linux Server"
VERSION="6.8"
ID="ol"
VERSION_ID="6.8"
PRETTY_NAME="Oracle Linux Server 6.8"
ANSI_COLOR="0;31"
CPE_NAME="cpe:/o:oracle:linux:6:8:server"
HOME_URL="https://linux.oracle.com/"
BUG_REPORT_URL="https://bugzilla.oracle.com/"

ORACLE_BUGZILLA_PRODUCT="Oracle Linux 6"
ORACLE_BUGZILLA_PRODUCT_VERSION=6.8
ORACLE_SUPPORT_PRODUCT="Oracle Linux"
ORACLE_SUPPORT_PRODUCT_VERSION=6.8
`
	os, err := Init(centOS)
	assert.Nil(t, err)
	assert.Equal(t, "ol", os.ID)
	assert.Equal(t, "6.8.0", os.Version.String())
}
