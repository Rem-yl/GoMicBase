package service

import (
	"Store/StoreServ/pb"
	"Store/internal"
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

func TestGetBrandByName(t *testing.T) {
	client, err := internal.GetStoreServClient()
	if err != nil {
		log.Panicln(err.Error())
	}

	req := &pb.BrandNameRequest{
		Name: "Rem",
	}

	resp, err := client.GetBrandByName(context.Background(), req)
	if err != nil {
		log.Panicln(err.Error())
	}

	log.Printf("Brand Found Name: %s, Logo: %s", resp.Name, resp.Logo)
}

func TestGetBrandById(t *testing.T) {
	client, err := internal.GetStoreServClient()
	if err != nil {
		log.Panicln(err.Error())
	}

	req := &pb.BrandIdRequest{
		Id: 2,
	}

	resp, err := client.GetBrandById(context.Background(), req)
	if err != nil {
		log.Panicln(err.Error())
	}

	log.Printf("Brand Found Name: %s, Logo: %s", resp.Name, resp.Logo)
}

func TestUpdateBrandById(t *testing.T) {
	client, err := internal.GetStoreServClient()
	if err != nil {
		log.Panicln(err.Error())
	}

	req := &pb.UpdateBrandRequest{
		Id:   2,
		Name: "Rem",
		Logo: "hello, Rem",
	}
	resp, err := client.UpdateBrandById(context.Background(), req)
	if err != nil {
		log.Panicln(err.Error())
	}

	log.Printf("Update Brand Name: %s, Logo: %s", resp.Name, resp.Logo)
}

func TestDeleteBrandById(t *testing.T) {
	client, err := internal.GetStoreServClient()
	if err != nil {
		log.Panicln(err.Error())
	}

	req := &pb.BrandIdRequest{
		Id: 1,
	}

	resp, err := client.DeleteBrandById(context.Background(), req)
	if err != nil {
		log.Panicln(err.Error())
	}
	if resp.Ok {
		log.Printf("Delete Brand Id: %d", req.Id)
	} else {
		log.Println("Delete Failed")
	}
}
