from fastapi import Request, FastAPI
from fastapi.concurrency import run_in_threadpool
from src.scores import Game
import yaml
from pathlib import Path

app = FastAPI()

def custom_openapi():
    """
    Load the OpenAPI schema from an external YAML file.
    If the file is not found, fallback to the automatically generated schema.
    """
    if app.openapi_schema:
        return app.openapi_schema

    # Path to your swagger YAML file (adjust if necessary)
    openapi_path = Path("swagger.yaml")
    if openapi_path.exists():
        with open(openapi_path, "r") as f:
            openapi_schema = yaml.safe_load(f)
        app.openapi_schema = openapi_schema
        return app.openapi_schema

app.openapi = custom_openapi
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

    