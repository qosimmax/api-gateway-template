package handler

import (
	"api-gateway-template/mocks"
	"net/http"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

type FakeAPIMock struct {
	Fetcher *mocks.DataFetcher
}

var (
	muxRouter *mux.Router
	fakeAPI   FakeAPIMock
)

func TestMain(m *testing.M) {
	muxRouter = mux.NewRouter()
	fakeAPI = FakeAPIMock{Fetcher: new(mocks.DataFetcher)}
	muxRouter.HandleFunc("/example", Example(fakeAPI.Fetcher)).Methods(http.MethodGet).Name("Example")
	exitVal := m.Run()
	os.Exit(exitVal)
}
