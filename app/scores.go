package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
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
	var dat map[string][]map[string]interface{}
	var games map[string][]map[string][]interface{}
	reader, _ := io.ReadAll(res.Body)
	json.Unmarshal(reader, &dat)
	json.Unmarshal(reader, &games)
	if dat["game_groups"][0]["name"] == "Completed" {
		data := games["game_groups"][0]["games"]

		for _, obj := range data {
			game := obj.(map[string]interface{})
			team = game["team_one"].(map[string]interface{})["name"].(string)
			op = game["team_two"].(map[string]interface{})["name"].(string)
			score = game["team_one"].(map[string]interface{})["score"].(string)
			op_score = game["team_two"].(map[string]interface{})["score"].(string)
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
