// order-service/handler.go

package main

import(
	"context"
	"fmt"
	"sync"
	"time"
	"sort"
  	"net/smtp"
	"html/template"
	"math"
	"os"
	"bytes"
	"strconv"

	pb "github.com/charles-hashdak/cleartoo-services/order-service/proto/order"
	cartPb "github.com/charles-hashdak/cleartoo-services/cart-service/proto/cart"
	catalogPb "github.com/charles-hashdak/cleartoo-services/catalog-service/proto/catalog"
	userPb "github.com/charles-hashdak/cleartoo-services/user-service/proto/user"
)

type handler struct{
	repository
	cartClient cartPb.CartService
	catalogClient catalogPb.CatalogService
	userClient userPb.UserService
	addOrderMutex sync.Mutex
	updateOrderStatusMutex sync.Mutex
	transactionMutex sync.Mutex
}

func (s *handler) Order(ctx context.Context, req *pb.OrderRequest, res *pb.OrderResponse) error {
	s.addOrderMutex.Lock()
	defer s.addOrderMutex.Unlock()
	orderId, err := s.repository.Order(ctx, MarshalOrderRequest(req))

	if err != nil{
		fmt.Println(err)
		return err
	}

	res.Added = true

	_, err = s.cartClient.EmptyCart(ctx, &cartPb.GetRequest{
		UserId: req.Order.UserId,
	})

	if err != nil{
		return err
	}

	for _, item := range req.Order.Products {
		_, err = s.catalogClient.Unavailable(ctx, &catalogPb.Product{
			Id: item.Id,
		})

		if err != nil{
			return err
		}
	}

	_, err = s.userClient.SendNotification(ctx, &userPb.Notification{
		UserId: req.Order.Products[0].OwnerId,
		Title: "New order!",
		Body: "Check your sales!",
	})

	if err != nil{
		return err
	}

	userRes, err := s.userClient.Get(ctx, &userPb.User{
		Id: req.Order.UserId,
	})
	if err != nil {
		return err
	}

	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")

	// Receiver email address.
	to := []string{
		userRes.User.Email,
	}

	// smtp server configuration.
	smtpHost := "us2.smtp.mailhostbox.com"
	smtpPort := "25"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, err := template.ParseFiles("/var/templates/emails/order_confirmation.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer

	headers := "MIME-version: 1.0;\nContent-Type: text/html;"

	body.Write([]byte(fmt.Sprintf("From: Cleartoo <no_reply@cleartoo.co.th>\r\nTo: "+userRes.User.Email+"\r\nSubject: Order received!\r\n%s\n\n", headers)))

	t.Execute(&body, struct {
		Username 		string
		Date 			string
		ShippingMethod 	string
		PaymentMethod 	string
		SubTotal 		string
		ShippingFees 	string
		Taxes 			string
		Total 			string
		OrderNumber		string
	}{
		Username: userRes.User.Username,
		Date: time.Now().Format("2006-01-02 15:04:05"),
		ShippingMethod: req.Order.ShippingMethod,
		PaymentMethod: req.Order.PaymentMethod,
		SubTotal: strconv.FormatFloat(float64(req.Order.SubTotal), 'f', -1, 32),
		ShippingFees: strconv.FormatFloat(float64(req.Order.ShippingFees), 'f', -1, 32),
		Taxes: strconv.FormatFloat(float64(req.Order.Taxes), 'f', -1, 32),
		Total: strconv.FormatFloat(float64(req.Order.Total), 'f', -1, 32),
		OrderNumber: orderId[len(orderId)-4:],
	})

	// Sending email.
	mailErr := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if mailErr != nil {
		return mailErr
	}
	
	return nil
}

func (s *handler) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest, res *pb.UpdateOrderStatusResponse) error {
	s.updateOrderStatusMutex.Lock()
	defer s.updateOrderStatusMutex.Unlock()
	status, err := s.repository.UpdateOrderStatus(ctx, MarshalUpdateOrderStatusRequest(req))
	if err != nil {
		return err
	}
	if status == "canceled" {
		order, err := s.repository.GetSingleOrder(ctx, MarshalGetSingleRequest(&pb.GetSingleRequest{
			OrderId: req.OrderId,
		}))
		if err != nil {
			return err
		}
		for _, product := range order.Products {
			_, err = s.catalogClient.Available(ctx, &catalogPb.Product{
				Id: product.ID.Hex(),
			})
			if err != nil{
				return err
			}
		}
	}
	res.Status = status
	res.Updated = true
	return nil
}

func (s *handler) UpdateOrderShippingStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest, res *pb.UpdateOrderStatusResponse) error {
	status, err := s.repository.UpdateOrderShippingStatus(ctx, MarshalUpdateOrderStatusRequest(req))
	if err != nil {
		return err
	}
	res.Status = status
	res.Updated = true
	return nil
}

func (s *handler) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest, res *pb.CancelOrderResponse) error {
	s.updateOrderStatusMutex.Lock()
	defer s.updateOrderStatusMutex.Unlock()
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
	layout := "2006-01-02 15:04:05"
	sort.Slice(orders, func(i, j int) bool {
		prevUpdatedAt, _ := time.Parse(layout, orders[i].UpdatedAt)
		nextUpdatedAt, _ := time.Parse(layout, orders[j].UpdatedAt)
	    return prevUpdatedAt.After(nextUpdatedAt)
	})
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

func (s *handler) GetOrdersByStatus(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {
	orders, err := s.repository.GetOrdersByStatus(ctx, MarshalGetRequest(req))
	if err != nil {
		return err
	}
	res.Orders = UnmarshalOrderCollection(orders)
	return nil
}

func (s *handler) GetSales(ctx context.Context, req *pb.GetRequest, res *pb.GetResponse) error {
	orders, err := s.repository.GetSales(ctx, MarshalGetRequest(req))
	if err != nil {
		return err
	}
	layout := "2006-01-02 15:04:05"
	sort.Slice(orders, func(i, j int) bool {
		prevUpdatedAt, _ := time.Parse(layout, orders[i].UpdatedAt)
		nextUpdatedAt, _ := time.Parse(layout, orders[j].UpdatedAt)
	    return prevUpdatedAt.After(nextUpdatedAt)
	})
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

func (s *handler) CreateOffer(ctx context.Context, req *pb.CreateOfferRequest, res *pb.CreateOfferResponse) error {

	// Save our offer
	if err := s.repository.CreateOffer(ctx, MarshalCreateOfferRequest(req)); err != nil {
		return err
	}

	order, err4 := s.repository.GetSingleOrder(ctx, MarshalGetSingleRequest(&pb.GetSingleRequest{
		OrderId: req.OrderId,
	}))
	if err4 != nil {
		return err4
	}

	senderRes, err2 := s.userClient.Get(ctx, &userPb.User{
		Id: order.UserID,
	})
	if err2 != nil {
		return err2
	}

	_, err3 := s.userClient.SendNotification(ctx, &userPb.Notification{
		UserId: order.Products[0].OwnerID,
		Title: "New reimbursement request from "+senderRes.User.Username+"!",
		Body: "",
	})
	if err3 != nil {
		return err3
	}

	res.Created = true
	res.Offer = req.Offer
	return nil
}

func (s *handler) Withdraw(ctx context.Context, req *pb.WithdrawRequest, res *pb.WithdrawResponse) error {
	s.transactionMutex.Lock()
	defer s.transactionMutex.Unlock()
	if err := s.repository.Withdraw(ctx, MarshalWithdrawRequest(req)); err != nil {
		return err
	}

	userRes, err := s.userClient.Get(ctx, &userPb.User{
		Id: req.UserId,
	})
	if err != nil {
		return err
	}

	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")

	// Receiver email address.
	to := []string{
		"admin@cleartooapp.com",
	}

	// smtp server configuration.
	smtpHost := "us2.smtp.mailhostbox.com"
	smtpPort := "25"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, err := template.ParseFiles("/var/templates/emails/withdraw_request.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer

	headers := "MIME-version: 1.0;\nContent-Type: text/html;"

	body.Write([]byte(fmt.Sprintf("From: Cleartoo <no_reply@cleartoo.co.th>\r\nTo: admin@cleartooapp.com\r\nSubject: Withdraw request!\r\n%s\n\n", headers)))

	t.Execute(&body, struct {
		Username 		string
		Amount 			string
		BankName 		string
		BankAccount 	string
	}{
		Username: userRes.User.Username,
		Amount: strconv.FormatFloat(float64(req.Amount), 'f', -1, 32),
		BankAccount: req.BankAccount,
		BankName: req.BankName,
	})

	// Sending email.
	mailErr := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if mailErr != nil {
		return mailErr
	}

	res.Created = true
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

func (s *handler) GetWallet(ctx context.Context, req *pb.GetWalletRequest, res *pb.GetWalletResponse) error {
	wallet, err := s.repository.GetWallet(ctx, MarshalGetWalletRequest(req))
	if err != nil {
		return err
	}
	wallet.Balance = float32(math.Round(float64(wallet.Balance*10))/10)
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
	s.transactionMutex.Lock()
	defer s.transactionMutex.Unlock()
	err := s.repository.UpdateWallet(ctx, MarshalUpdateWalletRequest(req))

	if err != nil{
		return err
	}

	res.Edited = true

	return nil
}

func (s *handler) AddTransaction(ctx context.Context, req *pb.AddTransactionRequest, res *pb.AddTransactionResponse) error {
	s.transactionMutex.Lock()
	defer s.transactionMutex.Unlock()
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