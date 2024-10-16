package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScores(t *testing.T) {
	game := Game{}
	yes, err := game.GetScores("NFL", "2024-10-14")

	t.Log("payload", yes[0])

	if err != nil {
		t.Error()
	}

	assert.Equal(t, yes[0].Team, "Bills")
	assert.Equal(t, yes[0].Op, "Jets")

}
