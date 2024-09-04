package service

import (
	"Store/StoreServ/database"
	"Store/StoreServ/pb"
	"context"
	"errors"
	"log"

	share "github.com/GoMicBase/Share"
)

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

func (server *StoreService) GetBrandByName(ctx context.Context, req *pb.BrandNameRequest) (resp *pb.BrandResponse, err error) {
	db := database.MysqlDB
	var brand database.Brand

	result := db.Where("name=?", req.Name).First(&brand)
	if result.RowsAffected == 0 {
		log.Printf("Brand Name: %s Not Found!", req.Name)
		return nil, errors.New(share.ErrBrandNotFound)
	}

	resp = &pb.BrandResponse{
		Name: brand.Name,
		Logo: brand.Logo,
	}

	return resp, nil
}

func (server *StoreService) GetBrandById(ctx context.Context, req *pb.BrandIdRequest) (resp *pb.BrandResponse, err error) {
	db := database.MysqlDB
	var brand database.Brand

	result := db.First(&brand, req.Id)
	if result.RowsAffected == 0 {
		log.Printf("Brand Not Found, Id: %d", req.Id)
		return nil, errors.New(share.ErrBrandNotFound)
	}

	resp = &pb.BrandResponse{
		Name: brand.Name,
		Logo: brand.Logo,
	}

	return resp, nil
}

func (server *StoreService) UpdateBrandById(ctx context.Context, req *pb.UpdateBrandRequest) (resp *pb.BrandResponse, err error) {
	db := database.MysqlDB
	var brand database.Brand

	// 检查是否有这个数据
	result := db.First(&brand, req.Id)
	if result.RowsAffected == 0 {
		log.Printf("Brand Not Found, Id: %d", req.Id)
		return nil, errors.New(share.ErrBrandNotFound)
	}

	brand = database.Brand{
		Name: req.Name,
		Logo: req.Logo,
	}

	result.Updates(brand)

	resp = &pb.BrandResponse{
		Name: req.Name,
		Logo: req.Logo,
	}

	return resp, nil
}

func (server *StoreService) DeleteBrandById(ctx context.Context, req *pb.BrandIdRequest) (resp *pb.CheckResponse, err error) {
	db := database.MysqlDB
	var brand database.Brand

	// 检查是否有这个数据
	result := db.First(&brand, req.Id)
	if result.RowsAffected == 0 {
		log.Printf("Brand Not Found, Id: %d", req.Id)
		return nil, errors.New(share.ErrBrandNotFound)
	}

	result.Delete(&brand)

	resp = &pb.CheckResponse{
		Ok: true,
	}

	return resp, nil
}
