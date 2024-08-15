package app

import (
	"testing"
)

func TestScores(t *testing.T) {
	yes := getScores("MLB", "2024-8-11")

	t.Log("payload", yes)

	if yes == nil {
		t.Error()
	}

}
