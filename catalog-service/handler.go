// catalog-service/handler.go

package main

import(
	"context"
	"fmt"

	pb "github.com/charles-hashdak/cleartoo-services/catalog-service/proto/catalog"
	userPb "github.com/charles-hashdak/cleartoo-services/user-service/proto/user"
	cartPb "github.com/charles-hashdak/cleartoo-services/cart-service/proto/cart"
	_ "github.com/pkg/errors"
)


type handler struct {
	repository
	userClient userPb.UserService
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

func (s *handler) EditProduct(ctx context.Context, req *pb.Product, res *pb.EditProductResponse) error {

	// Save our product
	if err := s.repository.EditProduct(ctx, MarshalProduct(req)); err != nil {
		return err
	}

	res.Edited = true
	res.Product = req
	return nil
}

func (s *handler) CreateOffer(ctx context.Context, req *pb.CreateOfferRequest, res *pb.CreateOfferResponse) error {

	// Save our offer
	if err := s.repository.CreateOffer(ctx, MarshalCreateOfferRequest(req)); err != nil {
		return err
	}

	senderRes, err2 := s.userClient.Get(ctx, &userPb.User{
		Id: req.Offer.UserId,
	})
	if err2 != nil {
		return err2
	}

	products, err4 := s.repository.GetProducts(ctx, MarshalGetRequest(&pb.GetRequest{
		Filters: []*pb.Filter{{
			Key: "_id",
			Condition: "$eq",
			Value: req.ProductId,
			Hex: true,
		}},
	}), s.userClient)
	if err4 != nil {
		return err4
	}

	_, err3 := s.userClient.SendNotification(ctx, &userPb.Notification{
		UserId: products[0].Owner.OwnerID,
		Title: "New offer from "+senderRes.User.Username+"!",
		Body: "",
	})
	if err3 != nil {
		return err3
	}

	res.Created = true
	res.Offer = req.Offer
	return nil
}

func (s *handler) EditOffer(ctx context.Context, req *pb.Offer, res *pb.EditOfferResponse) error {

	// Save our offer
	if err := s.repository.EditOffer(ctx, MarshalOffer(req)); err != nil {
		return err
	}

	res.Edited = true
	res.Offer = req
	return nil
}

func (s *handler) GetProducts(ctx context.Context, req *pb.GetRequest, res *pb.GetProductsResponse) error {
	products, err := s.repository.GetProducts(ctx, MarshalGetRequest(req), s.userClient)
	if err != nil {
		return err
	}
	for _, product := range products {
		inCartRes, err2 := s.cartClient.IsInCart(ctx, &cartPb.IsInCartRequest{
			UserId: req.UserId,
			ProductId: product.ID.Hex(),
		})
		if err2 != nil {
			fmt.Println(err2)
			return err2
		}
		product.InCart = inCartRes.In
		if len(product.Offers) > 0 {
			for _, offer := range product.Offers {
				if offer.UserID == req.UserId && offer.Status == "accepted" {
					product.Price = offer.Amount
				}
			}
		}
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

func (s *handler) GetGenders(ctx context.Context, req *pb.Request, res *pb.GetGendersResponse) error {
	genders, err := s.repository.GetGenders(ctx, MarshalRequest(req))
	if err != nil {
		return err
	}
	res.Genders = UnmarshalGenderCollection(genders)
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
	genders, err := s.repository.GetGenders(ctx, MarshalRequest(req))
	if err != nil {
		return err
	}
	res.Genders = UnmarshalGenderCollection(genders)
	var categoriesReq = new(pb.GetRequest)
	categories, err := s.repository.GetCategories(ctx, MarshalGetRequest(categoriesReq))
	if err != nil {
		return err
	}
	res.Categories = UnmarshalCategoryCollection(categories)
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