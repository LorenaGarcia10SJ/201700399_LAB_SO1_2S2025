# Proyecto 3 - Tweets del Clima (Skeleton)
Carnet: 201700399

Este repositorio contiene una **estructura mínima funcional** para desarrollar el Proyecto 3:
- API de ingreso en **Rust** (Actix Web)
- Servicios en **Go** (writers/consumers) mínimos
- Archivo **.proto** para gRPC
- **Locust** para generación de carga
- **Dockerfiles** y **YAMLs** de Kubernetes de ejemplo
- README con pasos básicos para continuar

> Esta plantilla está pensada para que la completes con lógica de Kafka, RabbitMQ y Valkey,
> y la publiques en tu Zot Registry. Revisa los archivos y adapta los parámetros de red/registry.

## Estructura
```
proyecto3/
 ├── api_rust/
 ├── go_client/
 ├── go_writer_kafka/
 ├── go_writer_rabbitmq/
 ├── consumers/
 ├── locust/
 ├── proto/
 ├── server-go/
 └── kubernetes/
```

# Creación de Cluster
![alt text](/Proyecto3/Documentacion/img/cluster.png)

## Pasos rápidos (local)
1. Abrir en VS Code:
   ```bash
   code /mnt/data/proyecto3_skeleton
   ```
2. Construir imagenes:
   ```bash
   cd api_rust
   docker build -t api-rust:v1 .

   cd server-go
      protoc \
   --plugin=protoc-gen-go=/home/oem/go/bin/protoc-gen-go \
   --plugin=protoc-gen-go-grpc=/home/oem/go/bin/protoc-gen-go-grpc \
   --go_out=go_grpc_server/proto \
   --go-grpc_out=go_grpc_server/proto \
   weathertweet.proto

   sudo docker build -t server-go:v1 .
   ```
3. Subir a Zot:
   ```bash
   docker tag api-rust:v1 hypostatic-curvier-izabella.ngrok-free.dev/api-rust:v1
   docker push hypostatic-curvier-izabella.ngrok-free.dev/api-rust:v1

   sudo docker tag server-go:latest hypostatic-curvier-izabella.ngrok-free.dev/server-go:v1
   docker push hypostatic-curvier-izabella.ngrok-free.dev/server-go:v1
   ``` 
4. Aplicar manifests:
   ```bash
   kubectl apply -f ingress.yaml
   kubectl apply -f server-go.yaml
   kubectl apply -f api_rust-deployment.yaml
   kubectl apply -f kubernetes/

   kubectl get all
   ```

   ### Ver logs
   kubectl logs pod/api-rust-6bb94dc7cd-zp8m4 

### Verificar ip de ingress
kubectl get ingress

   curl -X POST 34.117.112.215/clima \
   -H "Content-Type: application/json" \
   -d '{"municipality":4,"temperature":25,"humidity":80,"weather":2}'

5. Ejecutar Locust apuntando al Ingress o NodePort.


## Asignación de municipio (según carnet 201700399)
Último dígito: 9 -> **chinautla**

---

# 
Revisa `docs/` para notas adicionales.

# Engrok
hypostatic-curvier-izabella.ngrok-free.dev

docker tag api-rust:v1 hypostatic-curvier-izabella.ngrok-free.dev/api-rust:v1
docker push hypostatic-curvier-izabella.ngrok-free.dev/api-rust:v1

docker tag server-go:v1 hypostatic-curvier-izabella.ngrok-free.dev/server-go:v1
docker push hypostatic-curvier-izabella.ngrok-free.dev/server-go:v1

# Arquitectura General
                ┌─────────────┐
                │  Locust     │   ← Genera tráfico (simula tweets)
                └─────┬───────┘
                      │
                      ▼
        ┌─────────────────────────┐
        │ Ingress NGINX (K8s)     │ ← Entrada al cluster
        └──────────┬──────────────┘
                   ▼
        ┌─────────────────────────┐
        │ API REST (Rust)         │ ← Recibe tweets por HTTP
        └──────────┬──────────────┘
                   ▼
        ┌─────────────────────────┐
        │ Go gRPC Server          │ ← Procesa el tweet recibido
        └───────┬────────────────┘
                ▼
     ┌────────────────────┐   ┌────────────────────┐
     │ Kafka Writer (Go)  │   │ RabbitMQ Writer(Go)│ ← Publican mensajes
     └────────────────────┘   └────────────────────┘
                │                      │
                ▼                      ▼
     ┌────────────────────┐   ┌────────────────────┐
     │ Kafka Consumer (Go)│   │ RabbitMQ Consumer  │ ← Leen y procesan
     └──────────┬─────────┘   └──────────┬─────────┘
                ▼                      ▼
                   ┌────────────────┐
                   │ Valkey (DB RAM)│ ← Guarda los datos del clima
                   └────────────────┘
                            │
                            ▼
                   ┌────────────────┐
                   │ Grafana        │ ← Visualiza los datos
                   └────────────────┘

# LOCUST
Ejecutar locust en modo web
locust -f locustfile.py --host=http://34.117.112.215

http://localhost:8089/?tab=charts

### Construir y cargar
docker build -t locust_local:latest .
docker run --rm -p 8089:8089 locust_local:latest

### Ver ip externa de grafana
kubectl get service grafana-service

### Abrir en navegador

http://34.29.148.237:3000/

Usuario: admin
Contraseña: admin


# Correr Zot
sudo docker run -d \
  --name zot \
  -p 5000:5000 \
  -v /home/lorecreativa12/zot-data:/var/lib/registry \
  ghcr.io/project-zot/zot-linux-amd64:latest

## Ejecutar Locust
export API_PATH=/clima
export MUNICIPALITY_OVERRIDE=chinautla
export PAYLOAD_RANDOM=true

locust -f locustfile.py --host=http://34.117.112.215 --headless -u 10 -r 5 -t 1m


kubectl logs deployment/server-go-deploy -f
# Confirmación de escalado automatico
kubectl get hpa -w

# Agregar Valkey
Comando para desplegar:

kubectl apply -f kubernetes/valkey-deployment.yaml

# Grafana
kubectl get service grafana-service
34.29.148.237
http://34.29.148.237:3000

Usuario: admin
Contraseña: adminlore

# Construir imagenes y subir a zot
# Writer RabbitMQ
cd go_writer_rabbitmq
docker build -t writer-rabbitmq:v1 .
docker tag writer-rabbitmq:v1 hypostatic-curvier-izabella.ngrok-free.dev/writer-rabbitmq:v1
docker push hypostatic-curvier-izabella.ngrok-free.dev/writer-rabbitmq:v1

# Writer Kafka
cd ../go_writer_kafka
docker build -t writer-kafka:v1 .
docker tag writer-kafka:v1 hypostatic-curvier-izabella.ngrok-free.dev/writer-kafka:v1
docker push hypostatic-curvier-izabella.ngrok-free.dev/writer-kafka:v1

# Desplegar en Kubernetes

Desde la raíz del proyecto:

kubectl apply -f kubernetes/valkey-deployment.yaml
kubectl apply -f kubernetes/grafana-deployment.yaml
kubectl apply -f kubernetes/go-writer-rabbitmq.yaml
kubectl apply -f kubernetes/go-writer-kafka.yaml

# Verifica que todos estén corriendo:

kubectl get pods

# Reconstruye la imagen:

docker build -t api-rust:v2 .
docker tag api-rust:v2 hypostatic-curvier-izabella.ngrok-free.dev/api-rust:v2
docker push hypostatic-curvier-izabella.ngrok-free.dev/api-rust:v2

# Actualiza el deployment:

kubectl set image deployment/api-rust api-rust=hypostatic-curvier-izabella.ngrok-free.dev/api-rust:v2

# Entrar al pod de Valkey y hacer:

kubectl exec -it deployment/valkey -- sh
valkey-cli
LRANGE weather_tweets 0 -1