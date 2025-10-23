package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "server_go/go_grpc_server/proto"

	"github.com/redis/go-redis/v9"

	"google.golang.org/grpc"
)

// Crear cliente global de Redis
var ctx = context.Background()
var rdb *redis.Client

type server struct {
	pb.UnimplementedWeatherTweetServiceServer
}

func (s *server) SendTweet(ctx context.Context, req *pb.WeatherTweetRequest) (*pb.WeatherTweetResponse, error) {
	data := fmt.Sprintf("municipality:%v temperature:%v humidity:%v weather:%v",
		req.Municipality, req.Temperature, req.Humidity, req.Weather)

	// Guardar en Redis
	err := rdb.LPush(ctx, "weather_tweets", data).Err()
	if err != nil {
		fmt.Printf("❌ Error guardando en Redis: %v\n", err)
		return &pb.WeatherTweetResponse{Status: "Error guardando en Redis ❌"}, nil
	}

	fmt.Printf("✅ Go gRPC recibió y guardó en Redis: %s\n", data)
	return &pb.WeatherTweetResponse{Status: "Guardado en Redis ✅"}, nil
}

func main() {
	// Conexión a Redis
	rdb = redis.NewClient(&redis.Options{
		Addr: "valkey-service:6379",
	})

	// Probar conexión
	err := rdb.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("❌ No se pudo conectar a Valkey (Redis): %v", err)
	}
	fmt.Println("✅ Conectado a Valkey (Redis) correctamente")

	// Iniciar servidor gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Error escuchando: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterWeatherTweetServiceServer(grpcServer, &server{})

	fmt.Println("🚀 Go gRPC server escuchando en :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Error en Serve: %v", err)
	}
}

/* package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/redis/go-redis/v9"
	pb "server_go/go_grpc_server/proto"

	"google.golang.org/grpc"
)

// Crear cliente global de Redis
var ctx = context.Background()
var rdb *redis.Client

type server struct {
	pb.UnimplementedWeatherTweetServiceServer
}

func (s *server) SendTweet(ctx context.Context, req *pb.WeatherTweetRequest) (*pb.WeatherTweetResponse, error) {
	data := fmt.Sprintf("municipality:%v temperature:%v humidity:%v weather:%v",
		req.Municipality, req.Temperature, req.Humidity, req.Weather)

	// Guardar en Redis
	err := rdb.LPush(ctx, "weather_tweets", data).Err()
	if err != nil {
		fmt.Printf("❌ Error guardando en Redis: %v\n", err)
		return &pb.WeatherTweetResponse{Status: "Error guardando en Redis ❌"}, nil
	}

	fmt.Printf("✅ Go gRPC recibió y guardó en Redis: %s\n", data)
	return &pb.WeatherTweetResponse{Status: "Guardado en Redis ✅"}, nil
}

func main() {
	// Conexión a Redis
	rdb = redis.NewClient(&redis.Options{
		Addr: "valkey-service:6379",
	})

	// Probar conexión
	err := rdb.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("❌ No se pudo conectar a Valkey (Redis): %v", err)
	}
	fmt.Println("✅ Conectado a Valkey (Redis) correctamente")

	// Iniciar servidor gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Error escuchando: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterWeatherTweetServiceServer(grpcServer, &server{})

	fmt.Println("🚀 Go gRPC server escuchando en :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Error en Serve: %v", err)
	}
} */
