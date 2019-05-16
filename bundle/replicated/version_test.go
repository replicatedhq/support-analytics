package replicated

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionParse(t *testing.T) {
	// TODO: what happens to the text if there are multiple operators?
	versionText := `PRODUCT                VERSION     ID
replicated-ui          2.0.38      N/A
replicated-operator    2.0.36      N/A
replicated             2.0.1657    N/A
`
	versions, err := ParseVersion(versionText)
	assert.Nil(t, err)
	assert.Equal(t, "2.0.1657", versions.Replicated.String())
	assert.Equal(t, "2.0.38", versions.UI.String())
	assert.Equal(t, 1, len(versions.Operators))
	assert.Equal(t, "2.0.36", versions.Operators[0].String())
}
