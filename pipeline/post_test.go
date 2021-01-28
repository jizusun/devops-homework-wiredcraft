package pipeline

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newPost(t *testing.T) {
	actual := newPost()
	isEndedWithMd := strings.HasSuffix(actual.fileName, ".md")
	assert.True(t, isEndedWithMd, actual.fileName)
	assert.Len(t, actual.fileName, 28)
}
