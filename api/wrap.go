package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Validatable interface {
	Validate() error
}

func WrapHandler[Req Validatable, Res any](handler func(ctx context.Context, req Req) (Res, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Req

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		res, err := handler(ctx, req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = req.Validate(); err != nil {
			http.Error(w, fmt.Sprintf("validation failed: %v", err), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
