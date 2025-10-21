from locust import HttpUser, task, between
import random

MUNICIPALITIES = ["mixco","guatemala","amatitlan","chinautla"]
WEATHERS = ["sunny","cloudy","rainy","foggy"]

class WeatherTweetUser(HttpUser):
    wait_time = between(0.5, 2)

    @task
    def send_tweet(self):
        data = {
            "municipality": random.choice(MUNICIPALITIES),
            "temperature": random.randint(15, 35),
            "humidity": random.randint(30, 95),
            "weather": random.choice(WEATHERS)
        }
        self.client.post("/tweet", json=data)
