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
