package service

import (
	"context"
	"fmt"
	"log"

	"github.com/mohashari/catalyst-test/model"
	"github.com/mohashari/catalyst-test/utils"
)

//OrderRequest ...
type OrderRequest struct {
	CustomerID    int64          `json:"customer_id"`
	OrderProducts []OrderProduct `json:"order_products"`
}

//OrderProduct ...
type OrderProduct struct {
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

func (o *OrderProduct) calculateAmount(amount float64) float64 {
	return amount * float64(o.Quantity)
}

//Valid ...
func (o *OrderRequest) Valid() error {
	if o.CustomerID <= 0 {
		return fmt.Errorf("customer id required")
	}
	if len(o.OrderProducts) <= 0 {
		return fmt.Errorf("order product min 1 product")
	} else {
		for _, orderProduct := range o.OrderProducts {
			if orderProduct.ProductID <= 0 {
				return fmt.Errorf("order product id required")
			}
			if orderProduct.Quantity <= 0 {
				return fmt.Errorf("order product quantity min 1")
			}
		}
	}
	return nil
}

func (s *service) CreateOrder(ctx context.Context, req OrderRequest) (resp DefaultResponse, err error) {

	if err := req.Valid(); err != nil {
		log.Println("level: ", "err ", "method: ", "valid req create order ", "message: ", err.Error())
		return resp, err
	}
	var amount float64

	customer, err := s.repo.CustomerRepo.GetByID(ctx, req.CustomerID)
	if err != nil {
		log.Println("level: ", "err ", "method: ", "get customer ", "message: ", err.Error())
		return resp, fmt.Errorf("failed get customer")
	}
	orderDetails := make([]model.OrderDetail, 0)

	for _, reqProduct := range req.OrderProducts {
		product, err := s.repo.ProductRepo.GetByID(ctx, reqProduct.ProductID)
		if err != nil {
			log.Println("level: ", "err ", "method: ", "get product by id ", "message: ", err.Error())
			return resp, fmt.Errorf("failed get product")
		}
		orderDetail := model.OrderDetail{
			Product:  product,
			Amount:   reqProduct.calculateAmount(product.Price),
			Quantity: reqProduct.Quantity,
		}
		orderDetails = append(orderDetails, orderDetail)
		amount = amount + orderDetail.Amount
	}

	now := utils.GetUtils().TimeNow()

	id, err := s.repo.OrderRepo.Insert(ctx, model.Order{
		Customer:     customer,
		OrderDate:    now,
		CreatedAt:    now,
		OrderDetails: orderDetails,
		Amount:       amount,
	})
	if err != nil {
		log.Println("level: ", "err ", "method: ", "insert order ", "message: ", err.Error())
		return resp, fmt.Errorf("failed to insert order")
	}

	return DefaultResponse{
		Message: success,
		Data:    id,
	}, nil
}

func (s *service) GetOrderDetailByID(ctx context.Context, id int64) (resp DefaultResponse, err error) {
	if id <= 0 {
		err = fmt.Errorf("id required")
		log.Println("level: ", "err ", "method: ", "validation get order by id ", "message: ", err.Error())
		return resp, err
	}
	order, err := s.repo.OrderRepo.GetByID(ctx, id)
	if err != nil {
		log.Println("level: ", "err ", "method: ", "get order by id ", "message: ", err.Error())
		return resp, fmt.Errorf("failed to get order")
	}

	customer, err := s.repo.CustomerRepo.GetByID(ctx, order.Customer.ID)
	if err != nil {
		log.Println("level: ", "err ", "method: ", "get customer ", "message: ", err.Error())
		return resp, fmt.Errorf("failed get customer")
	}

	orderDetails, err := s.repo.OrderDetailRepo.GetDetailByOrderID(ctx, order.ID)
	if err != nil {
		log.Println("level: ", "err ", "method: ", "get order detail ", "message: ", err.Error())
		return resp, fmt.Errorf("failed get order detail")
	}

	var details []model.OrderDetail

	for _, orderDetail := range orderDetails {
		product, err := s.repo.ProductRepo.GetByID(ctx, orderDetail.Product.ID)
		if err != nil {
			log.Println("level: ", "err ", "method: ", "get product ", "message: ", err.Error())
			return resp, fmt.Errorf("failed get product")
		}

		brand, err := s.repo.BrandRepo.GetByID(ctx, product.Brand.ID)
		if err != nil {
			log.Println("level: ", "err ", "method: ", "get brand by id ", "message: ", err.Error())
			return resp, err
		}
		product.Brand = brand
		orderDetail.Product = product
		details = append(details, orderDetail)
	}

	order.Customer = customer
	order.OrderDetails = details

	return DefaultResponse{
		Message: success,
		Data:    order,
	}, nil
}
