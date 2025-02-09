import requests
from datetime import datetime as dt
import datetime
import pytz
class Game:
    def get_scores(self, league, date):
        """
        Retrieves game scores for a given league and date.
        
        :param league: The league identifier as a string.
        :param date: The date in YYYY-MM-DD format.
        :return: A list of box score dictionaries.
        """
        time_now = dt.now().astimezone()
        total_seconds = int(time_now.utcoffset().total_seconds())
        url = f"https://statmilk.bleacherreport.com/api/scores/carousel?league={league}&date={date}&tz={total_seconds}"
        try:
            response = requests.get(url)
            response.raise_for_status()
        except requests.RequestException as e:
            print("Error getting response:", e)
            raise
        
        # Parse the JSON response
        data = response.json()
        
        # Access the game groups and check for the appropriate group
        game_groups = data.get("game_groups", [])
        if not game_groups:
            print("No game groups found in the response.")
            return []
        
        full_score_complete = []
        full_score_in_progress = []
        for group in game_groups:
            game_progress = group.get("name", "")
            # Only process games if the group status is "Completed"
            if game_progress == "Completed":
                games_list = group.get("games", [])
                
                for game in games_list:
                    utc_time_str = str(game.get("game_date", ""))
                    print(utc_time_str)
                    utc_dt = dt.fromisoformat(utc_time_str[:-1] + '+00:00')
                    local_dt = utc_dt.astimezone(pytz.timezone('US/Eastern')).strftime("%Y-%m-%d %I:%M:%S %p")
                    # print("UTC time:   ", utc_dt)
                    # print("Local time: ", local_dt)
                    team = game.get("team_one", {}).get("name", "")
                    print(team,":local: ", local_dt)
                    op = game.get("team_two", {}).get("name", "")
                    homeScore = game.get("team_one", {}).get("score", "")
                    awayScore = game.get("team_two", {}).get("score", "")
                    
                    # Convert scores to integers for comparison
                    try:
                        team_one_score = int(homeScore)
                        team_two_score = int(awayScore)
                    except (ValueError, TypeError) as e:
                        print("Error converting scores to integers:", e)
                        raise
                    
                    # Determine the winner based on the scores
                    winner = team if team_one_score > team_two_score else op
                    
                    # Create a box score dictionary similar to pb.BoxScore in Go
                    box_score = {
                        "away": team,
                        "home": op,
                        "awayScore": homeScore,
                        "homeScore": awayScore,
                        "Winner": winner,
                        "date": local_dt
                    }
                    full_score_complete.append(box_score)
            else:
                games_list = group.get("games", [])
                
                for game in games_list:
                    utc_time_str = str(game.get("game_date", ""))
                    print(utc_time_str)
                    utc_dt = dt.fromisoformat(utc_time_str[:-1] + '+00:00')
                    local_dt = utc_dt.astimezone(pytz.timezone('US/Eastern')).strftime("%Y-%m-%d %I:%M:%S %p")
                    
                    team = game.get("team_one", {}).get("name", "")
                    print(team,":local: ", local_dt)
                    op = game.get("team_two", {}).get("name", "")
                    homeScore = game.get("team_one", {}).get("score", "")
                    awayScore = game.get("team_two", {}).get("score", "")
                    homeOdds = game.get("team_one", {}).get("odds", "")
                    awayOdds = game.get("team_two", {}).get("odds", "")
                    
                    # Convert scores to integers for comparison
                    try:
                        team_one_score = int(homeScore)
                        team_two_score = int(awayScore)
                    except (ValueError, TypeError) as e:
                        print("Error converting scores to integers:", e)
                        raise
                    
                    # Create a box score dictionary similar to pb.BoxScore in Go
                    box_score = {
                        "away": team,
                        "home": op,
                        "awayScore": homeScore,
                        "homeScore": awayScore,
                        "homeOdds": homeOdds,
                        "awayOdds": awayOdds,
                        "date": local_dt
                    }
                    full_score_in_progress.append(box_score)
        
        return {"inProgress": full_score_in_progress, "completed": full_score_complete}

# Example usage:
if __name__ == "__main__":
    game = Game()
    league = "NBA"          # Replace with the desired league
    date = "2025-02-09"     # Replace with the desired date in YYYY-MM-DD format
    scores = game.get_scores(league, date)
    print(scores)