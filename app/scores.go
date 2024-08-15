package app

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"
)

type Scores struct {
	Team    string
	Op      string
	Score   string
	OpScore string
	Winner  string
}

func getScores(league string, date string) []Scores {
	var team = ""
	var op = ""
	var score = ""
	var op_score = ""
	full_score := []Scores{}
	client := &http.Client{}
	url := fmt.Sprintf("https://statmilk.bleacherreport.com/api/scores/schedules?league=%s&date=%s", league, date)
	req, _ := http.NewRequest("GET", url, nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error getting response")
	}
	reader, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	game_progress := gjson.Get(string(reader), "game_groups.0.name").String()
	if game_progress == "Completed" {
		//gjson sytnax, to understand sytnax
		data := gjson.Get(string(reader), "game_groups.0.games")
		for _, obj := range data.Array() {
			obj_to_string := obj.String()
			team = gjson.Get(obj_to_string, "team_one.name").String()
			op = gjson.Get(obj_to_string, "team_two.name").String()
			score = gjson.Get(obj_to_string, "team_one.score").String()
			op_score = gjson.Get(obj_to_string, "team_two.score").String()
			team_two_score, err := strconv.Atoi(op_score)

			if err != nil {
				// ... handle error
				panic(err)
			}

			team_one_score, err := strconv.Atoi(score)

			if err != nil {
				// ... handle error
				panic(err)
			}

			winner := ""

			if team_one_score > team_two_score {
				winner = team
			} else {
				winner = op
			}
			scores := Scores{
				Team:    team,
				Op:      op,
				Score:   score,
				OpScore: op_score,
				Winner:  winner,
			}
			full_score = append(full_score, scores)
		}

	}

	return full_score
}
