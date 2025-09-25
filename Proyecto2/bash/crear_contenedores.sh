#!/bin/bash
echo "Creando contenedores de prueba..."

# Contenedores de bajo consumo
for i in {1..3}; do
    docker run -d --name low_$i busybox sleep 3600
done

# Contenedores de alto consumo
for i in {1..2}; do
    docker run -d --name high_$i stress --vm 1 --vm-bytes 100M --timeout 60s
done
