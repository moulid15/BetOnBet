package app

import (
	"testing"
)

func TestScores(t *testing.T) {
	yes := getScores("NFL", "2024-8-7")

	t.Log("helloooo", yes)

	if yes == nil {
		t.Error()
	}

}
