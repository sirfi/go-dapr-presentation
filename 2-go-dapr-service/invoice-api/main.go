package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/redis/go-redis/v9"
	"github.com/sirfi/go-dapr-presentation/2-go-dapr-service/models"
)

func newOrderSubscriptionHandler(ctx context.Context, e *common.TopicEvent) (bool, error) {
	if e == nil {
		return false, errors.New("topic event is nil")
	}
	newOrderEvent := models.NewOrderEventDto{}
	err := e.Struct(&newOrderEvent)
	if err != nil {
		return false, err
	}

	order := models.OrderDto{}
	if err = redisClient.Get(ctx, fmt.Sprintf("order_%v", newOrderEvent.Id.String())).Scan(&order); err != nil {
		return false, err
	}

	log.Printf("invoice created. id: %v", order.Id)
	return false, nil
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

	daprService := daprd.NewService(":8081")

	newOrderSubscription := common.Subscription{
		PubsubName: "pubsub",
		Topic:      "new_order",
		Route:      "/invoice/create",
	}
	daprService.AddTopicEventHandler(&newOrderSubscription, newOrderSubscriptionHandler)

	if err = daprService.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error: %v", err)
	}
}
