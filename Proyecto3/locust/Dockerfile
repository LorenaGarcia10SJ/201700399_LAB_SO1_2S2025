FROM python:3.11-slim
WORKDIR /app
COPY locustfile.py .
RUN pip install locust
EXPOSE 8089
CMD ["locust", "-f", "locustfile.py", "--host", "http://34.117.112.215"]
