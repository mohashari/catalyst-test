package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mohashari/catalyst-test/mock"
	"github.com/mohashari/catalyst-test/model"
	"github.com/mohashari/catalyst-test/repository"
)

func Test_service_CreateBrand(t *testing.T) {
	type fields struct {
		repo *repository.Repository
	}
	type args struct {
		ctx context.Context
		req BrandRequest
	}

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	brandRepo := mock.NewMockBrandRepo(ctrl)
	brandRepo.EXPECT().Insert(ctx, model.Brand{
		Name: "Cata",
	}).Return(int64(1), nil)

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp DefaultResponse
		wantErr  bool
	}{
		{
			name: "Test_service_CreateBrand positif case",
			fields: fields{
				repo: &repository.Repository{
					BrandRepo: brandRepo,
				},
			},
			args: args{
				ctx: ctx,
				req: BrandRequest{
					Name: "Cata",
				},
			},
			wantResp: DefaultResponse{
				Message: "success",
				Data:    int64(1),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			gotResp, err := s.CreateBrand(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateBrand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("service.CreateBrand() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
