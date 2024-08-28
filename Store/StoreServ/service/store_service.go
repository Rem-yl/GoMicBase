package service

import (
	"Store/StoreServ/database"
	"Store/StoreServ/pb"
	"context"
	"errors"
	"log"

	share "github.com/GoMicBase/Share"
)

type StoreService struct {
	pb.UnimplementedStoreServiceServer
}

func (server *StoreService) CreateNewBrand(ctx context.Context, req *pb.CreateNewBrandRequest) (resp *pb.BrandResponse, err error) {
	db := database.MysqlDB
	var brand database.Brand

	// 判断是否存在
	result := db.Where("name=?", req.Name).First(&brand)
	if result.RowsAffected != 0 {
		log.Printf("Brand Exists: %s", req.Name)
		return nil, errors.New(share.BrandExisted)
	}

	brand = database.Brand{
		Name: req.Name,
		Logo: req.Logo,
	}

	result = db.Create(&brand)

	if result.Error != nil {
		log.Printf("Error creating brand: %v", result.Error)
		return nil, result.Error
	}

	resp = &pb.BrandResponse{
		Name: req.Name,
		Logo: req.Logo,
	}

	return resp, nil
}
