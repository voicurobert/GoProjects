package main

import (
	"fmt"
	"testing"

	"github.com/voicurobert/GoProjects/microservices/client/client"
	"github.com/voicurobert/GoProjects/microservices/client/client/products"
)

func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := products.NewListProductsParams()

	prods, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v", prods)

	t.Logf("%v", prods.GetPayload()[0])
}
