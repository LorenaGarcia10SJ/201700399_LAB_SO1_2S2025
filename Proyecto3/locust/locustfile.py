# locustfile.py
from locust import HttpUser, task, between, events
import random, os, json, time

# Mapeo de municipios (según proto)
# 1 = mixco, 2 = guatemala, 3 = amatitlan, 4 = chinautla
MUNICIPALITIES = {
    "mixco": 1,
    "guatemala": 2,
    "amatitlan": 3,
    "chinautla": 4
}

WEATHERS = [1, 2, 3, 4]  # sunny, cloudy, rainy, foggy

# Opciones:
# - MUNICIPALITY_OVERRIDE: nombre de municipio para forzar (ej: chinautla)
# - PAYLOAD_RANDOM: "true"/"false" si quieres datos aleatorios
# - PATH: endpoint (por defecto /clima)
MUNICIPALITY_OVERRIDE = os.environ.get("MUNICIPALITY_OVERRIDE", "").lower()
PAYLOAD_RANDOM = os.environ.get("PAYLOAD_RANDOM", "true").lower() == "true"
PATH = os.environ.get("PATH", "/clima")

def build_payload():
    if MUNICIPALITY_OVERRIDE and MUNICIPALITY_OVERRIDE in MUNICIPALITIES:
        muni = MUNICIPALITIES[MUNICIPALITY_OVERRIDE]
    else:
        # default: distribuir según el enunciado del carnet (opcional)
        # Si quieres forzar chinautla (último dígito 9 -> chinautla), pon override.
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
    wait_time = between(0.2, 1.0)   # ajusta para mayor o menor frecuencia

    @task
    def send_tweet(self):
        payload = build_payload()
        headers = {"Content-Type": "application/json"}
        self.client.post(PATH, json=payload, headers=headers)
