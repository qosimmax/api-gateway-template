package handler

import (
	"api-gateway-template/client"
	"encoding/json"
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	log "github.com/sirupsen/logrus"
)

func handleError(
	w http.ResponseWriter,
	err error,
	statusCode int,
	shouldLog bool,
) {
	if shouldLog {
		log.Error(err.Error())
	}

	errorBody, _ := json.Marshal(struct {
		Error string `json:"error"`
	}{
		Error: err.Error(), // Change to something generic if this API is publicly exposed
	})

	// grpc error check
	if code, ok := status.FromError(err); ok {
		errorBody, _ = json.Marshal(struct {
			Error string `json:"error"`
		}{
			Error: code.Message(),
		})

		statusCode = HTTPStatusFromCode(code.Code())

	}

	var errorCodeWrapper client.ErrorCodeWrapper
	if errors.As(err, &errorCodeWrapper) {
		w.Header().Add("X-preserve-error", "1")

		statusCode = errorCodeWrapper.StatusCode
		errorBody, err = errorCodeWrapper.GetResponseBody()
		if err != nil {
			log.Error(err.Error())
		}
	}

	w.WriteHeader(statusCode)
	_, _ = w.Write(errorBody)
}

func HTTPStatusFromCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		// Note, this deliberately doesn't translate to the similarly named '412 Precondition Failed' HTTP response status.
		return http.StatusBadRequest
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	}

	return http.StatusInternalServerError
}
