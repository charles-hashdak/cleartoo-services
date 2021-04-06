// order-service/main.go

package main

import(
	"context"
	"fmt"

	pb "github.com/charles-hashdak/cleartoo-services/order-service/proto/order"
	cartPb "github.com/charles-hashdak/cleartoo-services/cart-service/proto/cart"
)

type handler struct{
	repository
	cartClient cartPb.CartService
}

func (s *handler) Order(ctx context.Context, req *pb.OrderRequest, res *pb.OrderResponse) error {

	fmt.Println(req)
	err := s.repository.Order(ctx, MarshalOrderRequest(req))

	if err != nil{
		fmt.Println(err)
		return err
	}

	res.Added = true

	_, err2 := s.cartClient.EmptyCart(ctx, &cartPb.GetRequest{
		UserId: req.UserId,
	})

	if err2 != nil{
		return err2
	}

	return nil
}

func (s *handler) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest, res *pb.UpdateOrderStatusResponse) error {
	status, err := s.repository.UpdateOrderStatus(ctx, MarshalUpdateOrderStatusRequest(req))
	if err != nil {
		return err
	}
	res.Status = status
	res.Updated = true
	return nil
}

func (s *handler) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest, res *pb.CancelOrderResponse) error {
	err := s.repository.CancelOrder(ctx, MarshalCancelOrderRequest(req))
	if err != nil {
		return err
	}
	res.Canceled = true
	return nil
}

func (s *handler) GetOrders(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {
	orders, err := s.repository.GetOrders(ctx, MarshalGetRequest(req))
	if err != nil {
		return err
	}
	res.Orders = UnmarshalOrderCollection(orders)
	return nil
}

func (s *handler) GetInTransitOrders(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {
	orders, err := s.repository.GetInTransitOrders(ctx, MarshalGetRequest(req))
	if err != nil {
		return err
	}
	res.Orders = UnmarshalOrderCollection(orders)
	return nil
}

func (s *handler) GetSales(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {
	orders, err := s.repository.GetOrders(ctx, MarshalGetRequest(req))
	if err != nil {
		return err
	}
	res.Orders = UnmarshalOrderCollection(orders)
	return nil
}

func (s *handler) GetSingleOrder(ctx context.Context, req *pb.GetSingleRequest, res *pb.GetSingleResponse) error {
	order, err := s.repository.GetSingleOrder(ctx, MarshalGetSingleRequest(req))
	if err != nil {
		return err
	}
	res.Order = UnmarshalOrder(order)
	return nil
}

func (s *handler) GetWallet(ctx context.Context, req *pb.GetWalletRequest, res *pb.GetWalletResponse) error {
	wallet, err := s.repository.GetWallet(ctx, MarshalGetWalletRequest(req))
	if err != nil {
		return err
	}
	res.Wallet = UnmarshalWallet(wallet)
	return nil
}

func (s *handler) InitializeWallet(ctx context.Context, req *pb.InitializeWalletRequest, res *pb.InitializeWalletResponse) error {
	err := s.repository.InitializeWallet(ctx, MarshalInitializeWalletRequest(req))
	if err != nil {
		return err
	}
	res.Added = true
	return nil
}

func (s *handler) UpdateWallet(ctx context.Context, req *pb.UpdateWalletRequest, res *pb.UpdateWalletResponse) error {

	err := s.repository.UpdateWallet(ctx, MarshalUpdateWalletRequest(req))

	if err != nil{
		return err
	}

	res.Edited = true

	return nil
}

func (s *handler) AddTransaction(ctx context.Context, req *pb.AddTransactionRequest, res *pb.AddTransactionResponse) error {

	err := s.repository.AddTransaction(ctx, MarshalAddTransactionRequest(req))

	if err != nil{
		return nil
	}

	getWalletReq := &pb.GetWalletRequest{
		WalletId: req.Transaction.WalletId,
	}
	wallet, err2 := s.repository.GetWallet(ctx, MarshalGetWalletRequest(getWalletReq))

	if err2 != nil{
		return err2
	}

	wallet.Balance = wallet.Balance + req.Transaction.Amount
	updateWalletReq := &pb.UpdateWalletRequest{
		Wallet: UnmarshalWallet(wallet),
	}
	err3 := s.repository.UpdateWallet(ctx, MarshalUpdateWalletRequest(updateWalletReq))

	if err3 != nil{
		return err3
	}

	res.Added = true

	return nil
}