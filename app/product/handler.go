package main

import (
	"context"

	"github.com/meilingluolingluo/gomall/app/product/biz/service"
	product "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/product"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

func (s *ProductCatalogServiceImpl) CreateProduct(ctx context.Context, req *product.CreateProductReq) (res *product.CreateProductResp, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *ProductCatalogServiceImpl) DeleteProduct(ctx context.Context, req *product.DeleteProductReq) (res *product.DeleteProductResp, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *ProductCatalogServiceImpl) UpdateProduct(ctx context.Context, req *product.UpdateProductReq) (res *product.UpdateProductResp, err error) {
	//TODO implement me
	panic("implement me")
}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}
