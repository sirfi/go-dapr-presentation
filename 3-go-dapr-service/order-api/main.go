package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/sirfi/go-dapr-presentation/3-go-dapr-service/models"
)

func createOrderHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in == nil {
		err = errors.New("invocation parameter required")
		return
	}

	order := models.OrderDto{}
	err = json.Unmarshal(in.Data, &order)
	if err != nil {
		return
	}

	client, err := dapr.NewClient()
	if err != nil {
		return
	}
	defer client.Close()

	if err = client.SaveState(ctx, "statestore", fmt.Sprintf("order_%v", order.Id.String()), in.Data, nil); err != nil {
		return
	}

	newOrderEvent := models.NewOrderEventDto{
		Id: order.Id,
	}

	if err = client.PublishEvent(ctx, "pubsub", "new_order", newOrderEvent); err != nil {
		return
	}

	out = &common.Content{}

	return
}

func main() {
	service := daprd.NewService(":8080")

	service.AddServiceInvocationHandler("/order/create", createOrderHandler)

	if err := service.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error: %v", err)
	}
}
