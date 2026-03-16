package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"musicproject.com/config"
	"musicproject.com/pkg/model"
)

func TestHandleOathGoogleLogin(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		wantStatus int
		wantRes    *model.GoogleUserInfo
	}{
		{
			name:       "login google oauth",
			method:     "GET",
			wantStatus: http.StatusOK,
			wantRes:    &model.GoogleUserInfo{},
		},
	}
	t.Skip()
	cfg := config.ReadConfigFile()
	oauthCfg := cfg.GoogleOathConfig()
	//jwtKey := cfg.JWTAccessKey()
	for _, tt := range tests {
		//t.Skip()
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(HandleOauthGoogleLogin(oauthCfg))
			defer ts.Close()
			//client := oauthCfg.Client(context.Background(), auth.GenerateToken(jwtKey))

			//re := httptest.NewRequest("GET", "/", nil)
			//w := httptest.NewRecorder()
			//handleOathGoogleLogin(oauthCfg).ServeHTTP(w, r)

			resp, err := http.Get(ts.URL)
			if err != nil {
				t.Error(err)
			}
			//r = httptest.NewRequest("GET", "/", nil)

			//handleOathGoogleRedirect([]byte("test"), oauthCfg).ServeHTTP(w, r)
			// body, err := io.ReadAll(resp.Body)
			// if err != nil {
			// 	t.Error(err)
			// }

			if tt.wantStatus != resp.StatusCode {
				t.Errorf("wrong status code want: %v got: %v", tt.wantStatus, resp.StatusCode)
			}
			userInfo, err := ReadJSON[model.GoogleUserInfo](resp.Body)
			if err != nil {

				t.Errorf("read error: %v", err)
			}

			assert.Equal(t, tt.wantRes, userInfo, tt.name)
		})
	}

}
