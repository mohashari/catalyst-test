package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

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

	return r
}
