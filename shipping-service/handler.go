// shipping-service/handler.go

package main

import(
	"context"

	pb "github.com/charles-hashdak/cleartoo-services/shipping-service/proto/shipping"
	_ "github.com/pkg/errors"
)


type handler struct {
	repository
}

func (s *handler) CreateAddress(ctx context.Context, req *pb.Address, res *pb.CreateAddressResponse) error {

	// Save our Address
	if err := s.repository.CreateAddress(ctx, MarshalAddress(req)); err != nil {
		return err
	}

	res.Created = true
	res.Address = req
	return nil
}

func (s *handler) GetAddresses(ctx context.Context, req *pb.GetRequest, res *pb.GetAddressesResponse) error {
	addresses, err := s.repository.GetAddresses(ctx, MarshalGetRequest(req))
	if err != nil {
		return err
	}
	res.Addresses = UnmarshalAddressCollection(addresses, req.UserId)
	return nil
}

func (s *handler) GetAddress(ctx context.Context, req *pb.GetRequest, res *pb.Address) error {
	address, err := s.repository.GetAddress(ctx, MarshalGetRequest(req))
	if err != nil {
		return err
	}
	res = UnmarshalAddress(address)
	return nil
}

func (s *handler) GetCountries(ctx context.Context, req *pb.Request, res *pb.GetCountriesResponse) error {
	countries, err := s.repository.GetCountries(ctx, MarshalRequest(req))
	if err != nil {
		return err
	}
	res.Countries = UnmarshalCountries(countries)
	return nil
}

func (s *handler) GetCities(ctx context.Context, req *pb.GetRequest, res *pb.GetCitiesResponse) error {
	cities, err := s.repository.GetCities(ctx, MarshalGetRequest(req))
	if err != nil {
		return err
	}
	res.Cities = UnmarshalCities(cities)
	return nil
}

func (s *handler) GetAddAddressData(ctx context.Context, req *pb.Request, res *pb.GetAddAddressDataResponse) error {
	countries, err := s.repository.GetCountries(ctx, MarshalRequest(req))
	if err != nil {
		return err
	}
	res.Countries = UnmarshalCountries(countries)
	var getReq = new(pb.GetRequest)
	cities, err := s.repository.GetCities(ctx, MarshalGetRequest(getReq))
	if err != nil {
		return err
	}
	res.Cities = UnmarshalCities(cities)
	return nil;
}