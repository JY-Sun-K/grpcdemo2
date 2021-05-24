package services

import (
	"context"
)

type ProdService struct {

}
//服务具体实现
func(s *ProdService) GetProdStock(ctx context.Context, req *ProdRequest)(*ProdResponse, error){
	stock :=req.ProdId
	if req.ProdArea==ProdAreas_A {
		stock=30
	}else if req.ProdArea==ProdAreas_B {
		stock=31
	}else if req.ProdArea==ProdAreas_C {
		stock=50
	}
	return &ProdResponse{ProdStock:stock},nil
}

func(s *ProdService) GetProdStocks(ctx context.Context, qs *QuerySize)(*ProdResponseList, error){
	return &ProdResponseList{Prodres: []*ProdResponse{
		{ProdStock: 25},
		{ProdStock: 26},
		{ProdStock: 27},
		{ProdStock: 28},
	}},nil
}

func (s *ProdService) GetProdInfo(ctx context.Context,req  *ProdRequest) (*ProdModel, error) {
	ret:=ProdModel{
		ProdId:        101,
		ProdName:      "测试商品",
		ProdPrice:     20.5,
	}
	return &ret,nil
}