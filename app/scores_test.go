package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScores(t *testing.T) {
	game := Game{}
	yes, err := game.GetScores("NBA", "2025-2-8")

	t.Log("payload", yes[0])

	if err != nil {
		t.Error()
	}

	assert.Equal(t, yes[0].Team, "Pacers")
	assert.Equal(t, yes[0].Op, "Lakers")

}
