package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/grpc"
	"github.com/redis/go-redis/v9"
	"github.com/sirfi/go-dapr-presentation/3-go-dapr-grpc-service/models"
)

func createOrderHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in == nil {
		err = errors.New("invocation parameter required")
		return
	}

	if in.Data == nil {
		err = errors.New("invocation data required")
		return
	}

	if in.ContentType != "application/json" {
		err = errors.New("invocation content type must be application/json")
		return
	}

	order := models.OrderDto{}
	err = json.Unmarshal(in.Data, &order)
	if err != nil {
		return
	}

	if err = redisClient.Set(ctx, fmt.Sprintf("order_%v", order.Id), order, 0).Err(); err != nil {
		return
	}

	newOrderEvent := models.NewOrderEventDto{
		Id: order.Id,
	}

	if err = daprClient.PublishEvent(ctx, "pubsub", "new_order", newOrderEvent); err != nil {
		return
	}

	out = &common.Content{}

	return
}

var daprClient dapr.Client
var redisClient *redis.Client

func main() {
	var err error
	daprClient, err = dapr.NewClient()
	if err != nil {
		return
	}
	defer daprClient.Close()

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer redisClient.Close()

	daprService, err := daprd.NewService(":50052")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	daprService.AddServiceInvocationHandler("order/create", createOrderHandler)

	if err = daprService.Start(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
