# locustfile.py
from locust import HttpUser, task, between
import random, os, json

MUNICIPALITIES = {
    "mixco": 1,
    "guatemala": 2,
    "amatitlan": 3,
    "chinautla": 4
}

WEATHERS = [1, 2, 3, 4]

MUNICIPALITY_OVERRIDE = os.environ.get("MUNICIPALITY_OVERRIDE", "").lower()
PAYLOAD_RANDOM = os.environ.get("PAYLOAD_RANDOM", "true").lower() == "true"
API_PATH = os.environ.get("API_PATH", "/clima")  # <-- cambio aquí

def build_payload():
    if MUNICIPALITY_OVERRIDE and MUNICIPALITY_OVERRIDE in MUNICIPALITIES:
        muni = MUNICIPALITIES[MUNICIPALITY_OVERRIDE]
    else:
        muni = random.choice(list(MUNICIPALITIES.values()))
    if PAYLOAD_RANDOM:
        payload = {
            "municipality": muni,
            "temperature": random.randint(15, 35),
            "humidity": random.randint(30, 95),
            "weather": random.choice(WEATHERS)
        }
    else:
        payload = {"municipality": muni, "temperature":25, "humidity":60, "weather":1}
    return payload

class WeatherUser(HttpUser):
    wait_time = between(0.2, 1.0)

    @task
    def send_tweet(self):
        payload = build_payload()
        headers = {"Content-Type": "application/json"}
        self.client.post(API_PATH, json=payload, headers=headers)  # <-- cambio aquí también
