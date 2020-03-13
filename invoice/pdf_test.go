package invoice

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestPdf_GetIoReader(t *testing.T) {
	pdf := Pdf{Content: "dGVzdDEyMw=="}
	reader := pdf.GetReader()

	bytes, err := ioutil.ReadAll(reader)
	require.NoError(t, err)

	assert.Equal(t, []byte("test123"), bytes)
}
