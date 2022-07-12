package handler

import (
	"github.com/agiledragon/gomonkey"
	"github.com/gin-gonic/gin"
	"net/http"
	"tdd/dao"
	"tdd/model"
	"tdd/web/response"
	"testing"
)

func TestGetUsers(t *testing.T) {
	const testURL = "/api/v1/users"

	tests := []struct {
		name         string
		request      *TestRequest
		wantResponse *WantResponse
	}{
		{
			name: "get users",
			request: &TestRequest{
				Method: http.MethodGet,
				URL:    testURL,
				Query: map[string]string{
					"page":      "1",
					"page_size": "10",
				},
			},
			wantResponse: &WantResponse{
				Code: http.StatusOK,
				Body: &response.Body{
					Code: response.ErrorCodeSuccess,
					Msg:  "",
					Data: gin.H{
						"users": []*model.User{
							{
								ID:       1,
								Username: "test",
								Password: "test",
							},
						},
					},
				},
			},
		},
		{
			name: "get users with invalid params",
			request: &TestRequest{
				Method: http.MethodGet,
				URL:    testURL,
				Query: map[string]string{
					"page":      "0",
					"page_size": "10",
				},
			},
			wantResponse: &WantResponse{
				Code: http.StatusOK,
				Body: &response.Body{
					Code: response.ErrorCodeInvalidParams,
					Msg:  "invalid params",
				},
			},
		},
		{
			name: "get users with invalid params",
			request: &TestRequest{
				Method: http.MethodGet,
				URL:    testURL,
				Query: map[string]string{
					"page": "a",
				},
			},
			wantResponse: &WantResponse{
				Code: http.StatusOK,
				Body: &response.Body{
					Code: response.ErrorCodeInvalidParams,
					Msg:  "invalid params",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			patches := gomonkey.ApplyFunc(
				dao.GetUsers,
				func(db *dao.DBManager, page, pageSize uint32) (users []*model.User, err error) {
					if tt.wantResponse.Body.Data == nil {
						return []*model.User{}, nil
					}

					return tt.wantResponse.Body.Data.(gin.H)["users"].([]*model.User), nil
				},
			)
			defer patches.Reset()

			RunTestAPI(t, tt.request, tt.wantResponse)
		})
	}
}
