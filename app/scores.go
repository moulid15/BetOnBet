package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	pb "github.com/moulid15/BetOnBet/proto"
	"github.com/tidwall/gjson"
)

// type Scores struct {
// 	Team    string
// 	Op      string
// 	Score   string
// 	OpScore string
// 	Winner  string
// }

type Game struct {
}

func (g Game) GetScores(league string, date string) ([]*pb.BoxScore, error) {
	var team = ""
	var op = ""
	var score = ""
	var op_score = ""
	full_score := []*pb.BoxScore{}
	client := &http.Client{}
	url := fmt.Sprintf("https://statmilk.bleacherreport.com/api/scores/schedules?league=%s&date=%s", league, date)
	req, _ := http.NewRequest("GET", url, nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error getting response")
	}
	reader, err := io.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
		return nil, err
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
				log.Println(err)

				return nil, err
			}

			team_one_score, err := strconv.Atoi(score)

			if err != nil {
				// ... handle error
				log.Println(err)
				return nil, err
			}

			winner := ""

			if team_one_score > team_two_score {
				winner = team
			} else {
				winner = op
			}
			scores := pb.BoxScore{
				Team:    team,
				Op:      op,
				Score:   score,
				OpScore: op_score,
				Winner:  winner,
			}
			full_score = append(full_score, &scores)
		}

	}

	return full_score, nil
}
