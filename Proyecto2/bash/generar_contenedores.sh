#!/bin/bash

# 3 imágenes que debes construir antes:
#   alto_ram
#   alto_cpu
#   bajo_consumo

IMAGENES=("alto_ram" "alto_cpu" "bajo_consumo")

for i in {1..10}
do
  IMG=${IMAGENES[$RANDOM % ${#IMAGENES[@]}]}
  docker run -d --name so1_test_$i $IMG
done
