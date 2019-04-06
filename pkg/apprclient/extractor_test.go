package apprclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtract(t *testing.T) {
	factory := NewClientFactory()

	options := Options{
		Source: "https://quay.io/cnr",
	}
	client, err := factory.New(options)
	require.NoError(t, err)

	result, err := client.RetrieveBlob("akashem/bundles", "1.7.1")
	require.NoError(t, err)
	require.NotNil(t, result)

	e := &extractor{}
	_, err = e.Extract(result.RawYAML)

	assert.NoError(t, err)
}
