package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Verifieddanny/go-social/internal/db/auth"
	"github.com/Verifieddanny/go-social/internal/store"
	"github.com/Verifieddanny/go-social/internal/store/cache"

	"go.uber.org/zap"
)

func newTestApplication(t *testing.T, cfg config) *application {
	t.Helper()
	logger := zap.Must(zap.NewProduction()).Sugar()
	mockStore := store.NewMockStore()
	mockCacheStore := cache.NewMockStore()

	testAuth := &auth.TestAuthenticator{}

	return &application{
		logger:        logger,
		store:         mockStore,
		cacheStorage:  mockCacheStore,
		authenticator: testAuth,
		config:        cfg,
	}
}

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Exepected the response code to be %d and we got %d", expected, actual)
	}
}
