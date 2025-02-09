from fastapi import Request, FastAPI
from fastapi.concurrency import run_in_threadpool
from src.scores import Game

app = FastAPI()

@app.post('/scores')
async def scores(request: Request):
    body = await request.json()
    league = body.get("league")
    date = body.get("date")

    game_instance = Game()
    
    # Call the synchronous get_scores method in a threadpool to avoid blocking
    scores_data = await run_in_threadpool(game_instance.get_scores, league, date)
    if not scores_data:
        return {"code": 1, "message": "League does not exist or date has no data"}
    
    return {"code":0, "data": scores_data}
    