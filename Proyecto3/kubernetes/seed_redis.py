import redis
import time
import random

# Conectar con Redis
r = redis.Redis(host='valkey-service', port=6379)

municipios = ["chinautla", "mixco", "chuarrancho"]
weather_conditions = ["sunny", "rain", "cloudy"]

for i in range(1, 21):  # 20 reportes de prueba
    key = f"report:{i}"
    r.hset(key, mapping={
        "municipality": random.choice(municipios),
        "temperature": random.randint(20, 35),
        "humidity": random.randint(50, 90),
        "weather": random.choice(weather_conditions),
        "timestamp": int(time.time())
    })
    time.sleep(0.1)

print("Datos insertados correctamente.")
