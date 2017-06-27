package n8s

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestN8s(t *testing.T) {
	assert.Equal(t, N8s("william"), "w5m", "handle single word")
	assert.Equal(t, N8s("william gough"), "w10h", "handle more than one word")
	assert.Equal(t, N8s("2"), "", "handle numbers")
}
