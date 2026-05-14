package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/Verifieddanny/go-social/internal/store/cache"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func TestGetUser(t *testing.T) {
	withRedis := config{
		redisCfg: redisCfg{
			enabled: true,
		},
	}
	app := newTestApplication(t, withRedis)
	mux := app.mount()
	testToken, err := app.authenticator.GenerateToken(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("should not allow authenticated requests", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/users/12d2f706-3727-4101-ab4c-6710f787d8dc", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusUnauthorized, rr.Code)

	})

	t.Run("should allow authenticated requests", func(t *testing.T) {
		withRedis := config{
			redisCfg: redisCfg{
				enabled: false,
			},
		}

		app := newTestApplication(t, withRedis)
		mux := app.mount()
		mockCacheStore := app.cacheStorage.Users.(*cache.MockUserStore)

		req, err := http.NewRequest(http.MethodGet, "/v1/users/00000000-0000-0000-0000-000000000000", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+testToken)

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)
		log.Println(rr.Body)
		mockCacheStore.Calls = nil

	})

	t.Run("should hit the cache first and if not exists it sets the user on the cache", func(t *testing.T) {
		mockCacheStore := app.cacheStorage.Users.(*cache.MockUserStore)

		mockCacheStore.On("Get", uuid.MustParse("12d2f706-3727-4101-ab4c-6710f787d8dc")).Return(nil, nil)
		mockCacheStore.On("Get", uuid.MustParse("00000000-0000-0000-0000-000000000000")).Return(nil, nil)
		mockCacheStore.On("Set", mock.Anything, mock.Anything).Return(nil)

		req, err := http.NewRequest(http.MethodGet, "/v1/users/00000000-0000-0000-0000-000000000000", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+testToken)

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)

		mockCacheStore.AssertNumberOfCalls(t, "Get", 2)
		mockCacheStore.Calls = nil
	})

	t.Run("should NOT hit the cache if it is not enabled", func(t *testing.T) {
		withRedis := config{
			redisCfg: redisCfg{
				enabled: false,
			},
		}

		app := newTestApplication(t, withRedis)
		mux := app.mount()
		mockCacheStore := app.cacheStorage.Users.(*cache.MockUserStore)

		req, err := http.NewRequest(http.MethodGet, "/v1/users/00000000-0000-0000-0000-000000000000", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+testToken)

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)

		mockCacheStore.AssertNotCalled(t, "Get")
		mockCacheStore.Calls = nil
	})
}
