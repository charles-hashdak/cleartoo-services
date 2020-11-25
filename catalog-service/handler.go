// catalog-service/handler.go

package main

import(
	"context"

	pb "github.com/charles-hashdak/cleartoo-services/catalog-service/proto/catalog"
	cartPb "github.com/charles-hashdak/cleartoo-services/cart-service/proto/cart"
	_ "github.com/pkg/errors"
)


type handler struct {
	repository
	cartClient cartPb.CartService
}

func (s *handler) CreateProduct(ctx context.Context, req *pb.Product, res *pb.CreateProductResponse) error {

	// Save our product
	if err := s.repository.CreateProduct(ctx, MarshalProduct(req)); err != nil {
		return err
	}

	res.Created = true
	res.Product = req
	return nil
}

func (s *handler) GetProducts(ctx context.Context, req *pb.GetRequest, res *pb.GetProductsResponse) error {
	products, err := s.repository.GetProducts(ctx, MarshalGetRequest(req))
	if err != nil {
		return err
	}
	res.Products = UnmarshalProductCollection(products, req.UserId)
	return nil
}

func (s *handler) GetProduct(ctx context.Context, req *pb.GetRequest, res *pb.Product) error {
	product, err := s.repository.GetProduct(ctx, MarshalGetRequest(req))
	if err != nil {
		return err
	}
	res2, err2 := s.cartClient.IsInCart(ctx, &cartPb.IsInCartRequest{
		UserId: req.UserId,
		ProductId: req.ProductId,
	})
	if err2 != nil {
		return err2
	}
	product.InCart = res2.In
	res = UnmarshalProduct(product, req.UserId)
	return nil
}

func (s *handler) Wish(ctx context.Context, req *pb.GetRequest, res *pb.WishResponse) error {
	if err := s.repository.Wish(ctx, MarshalGetRequest(req)); err != nil {
		return err
	}

	res.Wished = true
	return nil
}

func (s *handler) GetWishes(ctx context.Context, req *pb.GetRequest, res *pb.GetProductsResponse) error {
	products, err := s.repository.GetWishes(ctx, MarshalGetRequest(req))
	if err != nil {
		return err
	}
	res.Products = UnmarshalProductCollection(products, req.UserId)
	return nil
}

func (s *handler) GetSizes(ctx context.Context, req *pb.GetRequest, res *pb.GetSizesResponse) error {
	sizes, err := s.repository.GetSizes(ctx, MarshalGetRequest(req))
	if err != nil {
		return err
	}
	res.Sizes = UnmarshalSizeCollection(sizes)
	return nil
}

func (s *handler) GetCategories(ctx context.Context, req *pb.GetRequest, res *pb.GetCategoriesResponse) error {
	categories, err := s.repository.GetCategories(ctx, MarshalGetRequest(req))
	if err != nil {
		return err
	}
	res.Categories = UnmarshalCategoryCollection(categories)
	return nil
}

func (s *handler) GetBrands(ctx context.Context, req *pb.Request, res *pb.GetBrandsResponse) error {
	brands, err := s.repository.GetBrands(ctx, MarshalRequest(req))
	if err != nil {
		return err
	}
	res.Brands = UnmarshalBrandCollection(brands)
	return nil
}

func (s *handler) GetColors(ctx context.Context, req *pb.Request, res *pb.GetColorsResponse) error {
	colors, err := s.repository.GetColors(ctx, MarshalRequest(req))
	if err != nil {
		return err
	}
	res.Colors = UnmarshalColorCollection(colors)
	return nil
}

func (s *handler) GetConditions(ctx context.Context, req *pb.Request, res *pb.GetConditionsResponse) error {
	conditions, err := s.repository.GetConditions(ctx, MarshalRequest(req))
	if err != nil {
		return err
	}
	res.Conditions = UnmarshalConditionCollection(conditions)
	return nil
}

func (s *handler) GetAddProductData(ctx context.Context, req *pb.Request, res *pb.GetAddProductDataResponse) error {
	var gendersReq = new(pb.GetRequest)
	var gendersFilter = new(pb.Filter)
	gendersFilter.Key = "parent_id"
	gendersFilter.Condition = "$exists"
	gendersFilter.Value = "false"
	gendersReq.Filters = append(gendersReq.Filters, gendersFilter)
	categories, err := s.repository.GetCategories(ctx, MarshalGetRequest(gendersReq))
	if err != nil {
		return err
	}
	res.Genders = UnmarshalCategoryCollection(categories)
	conditions, err := s.repository.GetConditions(ctx, MarshalRequest(req))
	if err != nil {
		return err
	}
	res.Conditions = UnmarshalConditionCollection(conditions)
	colors, err := s.repository.GetColors(ctx, MarshalRequest(req))
	if err != nil {
		return err
	}
	res.Colors = UnmarshalColorCollection(colors)
	brands, err := s.repository.GetBrands(ctx, MarshalRequest(req))
	if err != nil {
		return err
	}
	res.Brands = UnmarshalBrandCollection(brands)
	materials, err := s.repository.GetMaterials(ctx, MarshalRequest(req))
	if err != nil {
		return err
	}
	res.Materials = UnmarshalMaterialCollection(materials)
	return nil
}