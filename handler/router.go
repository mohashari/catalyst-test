package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/mohashari/catalyst-test/service"
)

//ErrorResp ...
type ErrorResp struct {
	Message string `json:"message"`
}

//Router ...
func Router(ctx context.Context, r *http.ServeMux, svc service.Service) *http.ServeMux {

	r.HandleFunc("/brand", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			reqBody, _ := ioutil.ReadAll(r.Body)
			var req service.BrandRequest
			if err := json.Unmarshal(reqBody, &req); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ErrorResp{Message: err.Error()})
				return
			}
			resp, err := svc.CreateBrand(ctx, req)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ErrorResp{Message: err.Error()})
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
			return
		}
	})

	r.HandleFunc("/product/brand", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if r.Method == http.MethodGet {
			query := r.URL.Query()
			brandID, present := query["id"]
			if !present || len(brandID) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ErrorResp{Message: "id required"})
				return
			}

			id, _ := strconv.Atoi(brandID[0])

			resp, err := svc.GetProductByBrandID(ctx, int64(id))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ErrorResp{Message: err.Error()})
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
			return
		}
	})

	r.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		switch r.Method {
		case http.MethodPost:
			reqBody, _ := ioutil.ReadAll(r.Body)
			var req service.ProductCreateReq
			if err := json.Unmarshal(reqBody, &req); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ErrorResp{Message: err.Error()})
				return
			}

			resp, err := svc.CreateProduct(ctx, req)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ErrorResp{Message: err.Error()})
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
			return

		case http.MethodGet:
			query := r.URL.Query()
			productID, present := query["id"]
			if !present || len(productID) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ErrorResp{Message: "id required"})
				return
			}

			id, _ := strconv.Atoi(productID[0])

			resp, err := svc.GetProductByID(ctx, int64(id))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ErrorResp{Message: err.Error()})
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
			return
		}
	})
	return r
}
