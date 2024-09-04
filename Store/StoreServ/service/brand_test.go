package service

import (
	"Store/StoreServ/internal"
	"Store/StoreServ/pb"
	"context"
	"log"
	"testing"
)

func TestCreateNewBrand(t *testing.T) {
	client, err := internal.GetStoreServClient()
	if err != nil {
		log.Panicln(err.Error())
	}

	req := &pb.CreateNewBrandRequest{
		Name: "ram",
		Logo: "hello, ram",
	}
	resp, err := client.CreateNewBrand(context.Background(), req)
	if err != nil {
		log.Panicln(err.Error())
	}

	log.Printf("Create New Brand: Name: %s, Logo: %s", resp.Name, resp.Logo)
}
