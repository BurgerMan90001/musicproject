package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type handlerTest struct {
	name         string
	method       string
	url          string
	expect       string
	expectStatus int
}

func TestHandler(t *testing.T) {
	tests := []handlerTest{
		{
			name:         "put user",
			method:       "PUT",
			url:          "localhost:8081/user",
			expect:       `{"test":"test"}`,
			expectStatus: http.StatusOK,
		},
		{
			name:         "get user by id",
			method:       "GET",
			url:          "localhost:8081/user",
			expect:       `{"test":"test"}`,
			expectStatus: http.StatusOK,
		},
		{
			name:         "signup success",
			method:       "POST",
			url:          "localhost:8081/auth/signup",
			expect:       `{"test":"test"}`,
			expectStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Error(err)
			}

			rr := httptest.NewRecorder()

			//repo :=
			//authController := auth.New(repo, []byte(cfg.APIConfig.JWTKey))
			//userController := user.New(repo)

			//h := handler.New(authController, userController)
			//handler := http.HandlerFunc(h.handleUser)

			//handler.ServeHTTP(rr, req)
			//mux.ServeHTTP(rr, req)

			// if status := rr.Code; status != http.StatusOK {
			// 	t.Errorf("wrong status code got: %v wanted: %v", status, http.StatusOK)
			// }
			if status := rr.Code; status != tt.expectStatus {

			}

			fmt.Println(rr.Body.String())
		})
	}
}
