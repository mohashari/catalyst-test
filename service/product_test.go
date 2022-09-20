package service

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mohashari/catalyst-test/mock"
	"github.com/mohashari/catalyst-test/model"
	"github.com/mohashari/catalyst-test/repository"
)

func Test_service_CreateProduct(t *testing.T) {
	type fields struct {
		repo *repository.Repository
	}
	type args struct {
		ctx context.Context
		req ProductCreateReq
	}

	ctx := context.Background()

	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()

	mockBrand := mock.NewMockBrandRepo(mockctrl)
	mockBrand.EXPECT().GetByID(ctx, int64(1)).Return(model.Brand{}, nil)
	mockBrand.EXPECT().GetByID(ctx, int64(1)).Return(model.Brand{}, sql.ErrNoRows)

	mockProduct := mock.NewMockProductRepo(mockctrl)
	mockProduct.EXPECT().Insert(ctx, model.Product{
		Name:     "mouse",
		Price:    10000,
		Quantity: 20,
	}).Return(int64(1), nil)

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp DefaultResponse
		wantErr  bool
	}{
		{
			name: "Test_service_CreateProduct positif case",
			fields: fields{
				repo: &repository.Repository{
					BrandRepo:   mockBrand,
					ProductRepo: mockProduct,
				},
			},
			args: args{
				ctx: ctx,
				req: ProductCreateReq{
					BrandID:  1,
					Name:     "mouse",
					Price:    10000,
					Quantity: 20,
				},
			},
			wantResp: DefaultResponse{
				Message: success,
				Data:    int64(1),
			},
			wantErr: false,
		},
		{
			name: "Test_service_CreateProduct error get brand id",
			fields: fields{
				repo: &repository.Repository{
					BrandRepo:   mockBrand,
					ProductRepo: mockProduct,
				},
			},
			args: args{
				ctx: ctx,
				req: ProductCreateReq{
					BrandID:  1,
					Name:     "mouse",
					Price:    10000,
					Quantity: 20,
				},
			},
			wantResp: DefaultResponse{},
			wantErr:  true,
		},
		{
			name:   "Test_service_CreateProduct error validation",
			fields: fields{},
			args: args{
				ctx: ctx,
				req: ProductCreateReq{
					BrandID:  1,
					Name:     "mouse",
					Price:    10000,
					Quantity: 0,
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
			gotResp, err := s.CreateProduct(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("service.CreateProduct() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_service_GetProductByID(t *testing.T) {
	type fields struct {
		repo *repository.Repository
	}
	type args struct {
		ctx context.Context
		id  int64
	}

	ctx := context.Background()
	brandModel := model.Brand{
		ID:   1,
		Name: "cata",
	}
	productModel := model.Product{
		ID:       1,
		Brand:    brandModel,
		Name:     "mouse",
		Price:    10000,
		Quantity: 10,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProduct := mock.NewMockProductRepo(ctrl)
	mockProduct.EXPECT().GetByID(ctx, int64(1)).Return(productModel, nil)
	mockProduct.EXPECT().GetByID(ctx, int64(1)).Return(model.Product{}, sql.ErrNoRows)

	mockBrand := mock.NewMockBrandRepo(ctrl)
	mockBrand.EXPECT().GetByID(ctx, int64(1)).Return(brandModel, nil)

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp DefaultResponse
		wantErr  bool
	}{
		{
			name: "Test_service_GetProductByID positif case",
			fields: fields{
				repo: &repository.Repository{
					BrandRepo:   mockBrand,
					ProductRepo: mockProduct,
				},
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantResp: DefaultResponse{
				Message: success,
				Data:    productModel,
			},
			wantErr: false,
		},

		{
			name: "Test_service_GetProductByID error product no rows",
			fields: fields{
				repo: &repository.Repository{
					BrandRepo:   mockBrand,
					ProductRepo: mockProduct,
				},
			},
			args: args{
				ctx: ctx,
				id:  1,
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
			gotResp, err := s.GetProductByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetProductByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("service.GetProductByID() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_service_GetProductByBrandID(t *testing.T) {
	type fields struct {
		repo *repository.Repository
	}
	type args struct {
		ctx context.Context
		id  int64
	}

	ctx := context.Background()
	var products []model.Product

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProduct := mock.NewMockProductRepo(ctrl)
	mockProduct.EXPECT().GetByBrandID(ctx, int64(1)).Return(products, nil)
	mockProduct.EXPECT().GetByBrandID(ctx, int64(1)).Return(products, sql.ErrNoRows)

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp DefaultResponse
		wantErr  bool
	}{
		{
			name: "Test_service_GetProductByBrandID positif case",
			fields: fields{
				repo: &repository.Repository{
					ProductRepo: mockProduct,
				},
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantResp: DefaultResponse{
				Message: success,
				Data:    products,
			},
			wantErr: false,
		},

		{
			name: "Test_service_GetProductByBrandID error get product",
			fields: fields{
				repo: &repository.Repository{
					ProductRepo: mockProduct,
				},
			},
			args: args{
				ctx: ctx,
				id:  1,
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
			gotResp, err := s.GetProductByBrandID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetProductByBrandID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("service.GetProductByBrandID() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
