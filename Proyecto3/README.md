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

   sudo docker build -t server-go .
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
   kubectl logs pod/api-rust-c844676fb-75hwv

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