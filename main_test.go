package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordLength(t *testing.T) {
	pwd := GeneratePassword(20, upper+lower+digits)
	assert.Equal(t, 20, len(pwd))
}
