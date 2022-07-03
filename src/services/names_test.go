package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNames(t *testing.T) {
	namesStr := "pablo,pedro"
	expectedInt := 2
	expectedArr := []string{"pablo", "pedro"}

	resultInt, resultArr := GetNames(namesStr)

	assert.Equal(t, expectedInt, resultInt)
	assert.Equal(t, expectedArr, resultArr)
}
