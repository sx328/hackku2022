import fastapi.encoders
from fastapi import FastAPI, status, Response
from typing import Optional
from pydantic import BaseModel
import web3

w3 = Web3(Web3.HTTPProvider('http://127.0.0.1:8545'))

app = FastAPI()


class GenerateModel(BaseModel):
    fullName: str
    buy: bool
    signature: str




async def generateDocument(GenerateModel) -> dict:
    # Generate a name of the file based on the hash of the sig

    return {"location": "location"}

async def getlisting() -> dict:
    listings = {}


@app.post("/generate")
async def generate(request: GenerateModel):
    doc = generateDocument(request)
    return fastapi.encoders.jsonable_encoder(doc)


@app.get("/listings")
async def getlistings(request):
    for i in getListing():


