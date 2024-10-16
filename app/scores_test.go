package app

import (
	"testing"
)

func TestScores(t *testing.T) {
	game := Game{}
	yes, err := game.GetScores("NFL", "2024-10-14")

	t.Log("payload", yes)

	if err == nil {
		t.Error()
	}

}
