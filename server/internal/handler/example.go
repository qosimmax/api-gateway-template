package handler

import (
	"api-gateway-template/app"
	"net/http"
)

// Example is handler that provides an example of how handlers should be written.
//
//	GET /api/v1/api
//	Responds: 200, 500
//
// The handler should accept an interface(s), and should contain only high level
// business logic. There should be no implementation details here (except I guess
// stuff specific to http, like writing the response).
func Example(f app.DataFetcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_, err := f.GetExampleData(ctx)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError, true)
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		//_, _ = w.Write(response)
	}
}
