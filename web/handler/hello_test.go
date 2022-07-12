package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tdd/web/response"
	"testing"
)

func TestHello(t *testing.T) {
	const testURL = "/api/v1/hello"

	tests := []struct {
		name         string
		request      *TestRequest
		wantResponse *WantResponse
	}{
		{
			name: "get hello",
			request: &TestRequest{
				Method: http.MethodGet,
				URL:    testURL,
			},
			wantResponse: &WantResponse{
				Code: http.StatusOK,
				Body: &response.Body{
					Code: response.ErrorCodeSuccess,
					Msg:  "",
					Data: gin.H{
						"hello": "world",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RunTestAPI(t, tt.request, tt.wantResponse)
		})
	}
}
