package service

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mohashari/catalyst-test/mock"
	"github.com/mohashari/catalyst-test/model"
	"github.com/mohashari/catalyst-test/repository"
	"github.com/mohashari/catalyst-test/utils"
)

func Test_service_CreateOrder(t *testing.T) {
	type fields struct {
		repo *repository.Repository
	}
	type args struct {
		ctx context.Context
		req OrderRequest
	}
	now := time.Now()

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	customerModel := model.Customer{
		ID: 1,
	}
	productModel := model.Product{
		ID:    1,
		Price: 10000,
	}
	orderDetails := []model.OrderDetail{}
	orderDetails = append(orderDetails, model.OrderDetail{
		Product:  productModel,
		Quantity: 10,
		Amount:   100000,
	})
	orderModel := model.Order{
		Customer:     customerModel,
		OrderDate:    now,
		CreatedAt:    now,
		Amount:       100000,
		OrderDetails: orderDetails,
	}

	mockUtil := mock.NewMockUtils(ctrl)
	mockUtil.EXPECT().TimeNow().Return(now)
	utils.SetUtils(mockUtil)

	mockCustomer := mock.NewMockCustomerRepo(ctrl)
	mockCustomer.EXPECT().GetByID(ctx, int64(1)).Return(customerModel, nil)
	mockCustomer.EXPECT().GetByID(ctx, int64(1)).Return(model.Customer{}, sql.ErrNoRows)
	mockCustomer.EXPECT().GetByID(ctx, int64(1)).Return(customerModel, nil)

	mockProduct := mock.NewMockProductRepo(ctrl)
	mockProduct.EXPECT().GetByID(ctx, int64(1)).Return(productModel, nil)
	mockProduct.EXPECT().GetByID(ctx, int64(1)).Return(model.Product{}, sql.ErrNoRows)

	mockOrder := mock.NewMockOrderRepo(ctrl)
	mockOrder.EXPECT().Insert(ctx, orderModel).Return(int64(1), nil)

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp DefaultResponse
		wantErr  bool
	}{
		{
			name: "Test_service_CreateOrder",
			fields: fields{
				repo: &repository.Repository{
					CustomerRepo: mockCustomer,
					OrderRepo:    mockOrder,
					ProductRepo:  mockProduct,
				},
			},
			args: args{
				ctx: ctx,
				req: OrderRequest{
					CustomerID: 1,
					OrderProducts: []OrderProduct{
						{
							ProductID: 1,
							Quantity:  10,
						},
					},
				},
			},
			wantResp: DefaultResponse{
				Message: success,
				Data:    int64(1),
			},
			wantErr: false,
		},

		{
			name: "Test_service_CreateOrder error get customer",
			fields: fields{
				repo: &repository.Repository{
					CustomerRepo: mockCustomer,
					OrderRepo:    mockOrder,
					ProductRepo:  mockProduct,
				},
			},
			args: args{
				ctx: ctx,
				req: OrderRequest{
					CustomerID: 1,
					OrderProducts: []OrderProduct{
						{
							ProductID: 1,
							Quantity:  10,
						},
					},
				},
			},
			wantResp: DefaultResponse{},
			wantErr:  true,
		},

		{
			name: "Test_service_CreateOrder error get product",
			fields: fields{
				repo: &repository.Repository{
					CustomerRepo: mockCustomer,
					OrderRepo:    mockOrder,
					ProductRepo:  mockProduct,
				},
			},
			args: args{
				ctx: ctx,
				req: OrderRequest{
					CustomerID: 1,
					OrderProducts: []OrderProduct{
						{
							ProductID: 1,
							Quantity:  10,
						},
					},
				},
			},
			wantResp: DefaultResponse{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			gotResp, err := s.CreateOrder(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("service.CreateOrder() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestOrderProduct_calculateAmount(t *testing.T) {
	type fields struct {
		ProductID int64
		Quantity  int
	}
	type args struct {
		amount float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderProduct{
				ProductID: tt.fields.ProductID,
				Quantity:  tt.fields.Quantity,
			}
			if got := o.calculateAmount(tt.args.amount); got != tt.want {
				t.Errorf("OrderProduct.calculateAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderRequest_Valid(t *testing.T) {
	type fields struct {
		CustomerID    int64
		OrderProducts []OrderProduct
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderRequest{
				CustomerID:    tt.fields.CustomerID,
				OrderProducts: tt.fields.OrderProducts,
			}
			if err := o.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("OrderRequest.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetOrderDetailByID(t *testing.T) {
	type fields struct {
		repo *repository.Repository
	}
	type args struct {
		ctx context.Context
		id  int64
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var orderDetails []model.OrderDetail

	customerModel := model.Customer{
		ID: 1,
	}
	brand := model.Brand{
		ID: 1,
	}

	product := model.Product{
		ID:    1,
		Brand: brand,
	}

	orderDetails = append(orderDetails, model.OrderDetail{
		Product: product,
	})
	ordermodel := model.Order{
		ID:           1,
		Customer:     customerModel,
		OrderDetails: orderDetails,
	}

	orderMock := mock.NewMockOrderRepo(ctrl)
	orderMock.EXPECT().GetByID(ctx, int64(1)).Return(ordermodel, nil)

	customerMock := mock.NewMockCustomerRepo(ctrl)
	customerMock.EXPECT().GetByID(ctx, int64(1)).Return(customerModel, nil)

	orderDetailMock := mock.NewMockOrderDetailRepo(ctrl)
	orderDetailMock.EXPECT().GetDetailByOrderID(ctx, int64(1)).Return(orderDetails, nil)

	productMock := mock.NewMockProductRepo(ctrl)
	productMock.EXPECT().GetByID(ctx, int64(1)).Return(product, nil)

	brandMock := mock.NewMockBrandRepo(ctrl)
	brandMock.EXPECT().GetByID(ctx, int64(1)).Return(brand, nil)

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp DefaultResponse
		wantErr  bool
	}{
		{
			name: "Test_service_GetOrderDetailByID",
			fields: fields{
				repo: &repository.Repository{
					CustomerRepo:    customerMock,
					OrderRepo:       orderMock,
					BrandRepo:       brandMock,
					ProductRepo:     productMock,
					OrderDetailRepo: orderDetailMock,
				},
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantResp: DefaultResponse{
				Message: success,
				Data: model.Order{
					ID: 1,
					Customer: model.Customer{
						ID: 1,
					},
					OrderDetails: orderDetails,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			gotResp, err := s.GetOrderDetailByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetOrderDetailByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("service.GetOrderDetailByID() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
