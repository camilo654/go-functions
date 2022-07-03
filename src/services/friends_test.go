package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreFriends(t *testing.T) {
	assert.True(t, AreFriends("tokyo", "kyoto"))
}

func TestAreFriends_False(t *testing.T) {
	assert.False(t, AreFriends("tokyo", "koyto"))
}

func TestAreFriends_DifferentSize_False(t *testing.T) {
	assert.False(t, AreFriends("tokyo", "kyto"))
}
