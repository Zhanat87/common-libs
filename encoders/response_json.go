package encoders

import (
	"context"
	"encoding/json"
	"net/http"
)

func EncodeResponseJSON(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		EncodeErrorJSON(ctx, e, w)

		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(response)
}
