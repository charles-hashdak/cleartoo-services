// catalog-service/handler.go

package main

import(
	"context"

	pb "github.com/charles-hashdak/cleartoo-services/catalog-service/proto/catalog"
	"github.com/pkg/errors"
)


type handler struct {
	repository
}

func (s *handler) CreateProduct(ctx context.Context, req *pb.Product, res *pb.CreateProductResponse) error {

	// Save our product
	if err = s.repository.Create(ctx, MarshalProduct(req)); err != nil {
		return err
	}

	res.Created = true
	res.Product = req
	return nil
}

// GetConsignments -
func (s *handler) GetProducts(ctx context.Context, req *pb.GetProductsRequest, res *pb.GetProductsResponse) error {
	products, err := s.repository.GetProducts(ctx, req)
	if err != nil {
		return err
	}
	res.Products = UnmarshalProductCollection(products)
	return nil
}