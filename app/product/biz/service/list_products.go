package service

import (
	"context"

	"github.com/meilingluolingluo/gomall/app/product/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/app/product/model"
	product "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)
	c, err := categoryQuery.GetProductByCategoryName(req.CategoryName)
	resp = &product.ListProductsResp{}
	for _, v1 := range c {
		for _, v2 := range v1.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(v2.ID),
				Name:        v2.Name,
				Description: v2.Description,
				Picture:     v2.Picture,
				Price:       v2.Price,
			})
		}
	}

	return resp, nil
}
