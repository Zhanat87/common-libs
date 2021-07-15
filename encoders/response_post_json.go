package encoders

import (
	"context"
	"encoding/json"
	"net/http"
)

func EncodeResponsePostJSON(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		EncodeErrorJSON(ctx, e, w)

		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, _ := json.Marshal(response)
	var dat map[string]interface{}
	_ = json.Unmarshal(b, &dat)
	if _, ok := dat["location"]; ok {
		w.Header().Set("Location", dat["location"].(string))
	}
	w.WriteHeader(http.StatusCreated)

	return json.NewEncoder(w).Encode(response)
}
