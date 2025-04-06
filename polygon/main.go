package main

import (
	"context"
	"log"
	"os"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {
	c := polygon.New(os.Getenv("polygon"))

	params := models.ListTickersParams{}.
		WithMarket(models.AssetClass("fx")).
		WithActive(true).
		WithOrder(models.Order("asc")).
		WithLimit(100).
		WithSort(models.Sort("ticker"))

	iter := c.ListTickers(context.Background(), params)

	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}
}
