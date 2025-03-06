package service

import (
	"context"
	"log"

	"github.com/meilingluolingluo/gomall/app/product/biz/dal/mysql"
	"github.com/meilingluolingluo/gomall/app/product/model"
	product "github.com/meilingluolingluo/gomall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	products, err := productQuery.SearchProducts(req.Query)
	var result []*product.Product
	for _, v := range products {
		result = append(result, &product.Product{
			Id:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
		})
	}
	log.Printf("result = %v", result)
	return &product.SearchProductsResp{
		Results: result,
	}, err
}
