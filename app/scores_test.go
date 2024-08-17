package app

import (
	"testing"
)

func TestScores(t *testing.T) {
	game := Game{}
	yes, err := game.GetScores("MLB", "2024-8-11")

	t.Log("payload", yes)

	if err == nil {
		t.Error()
	}

}
