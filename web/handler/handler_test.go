package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"tdd/web/response"
	"testing"
)

var testEngine *gin.Engine

func TestMain(m *testing.M) {
	testEngine = gin.New()
	Register(testEngine)

	os.Exit(m.Run())
}

type TestRequest struct {
	Method   string
	URL      string
	Headers  map[string]string
	Query    map[string]string
	RawQuery string
	Json     interface{}
}

func (r *TestRequest) Do() (*httptest.ResponseRecorder, error) {
	var body io.Reader
	if r.Json != nil {
		data, err := json.Marshal(r.Json)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(data)
	}

	request, err := http.NewRequest(r.Method, r.URL, body)
	if err != nil {
		return nil, err
	}

	if r.Query != nil {
		query := request.URL.Query()
		for key, value := range r.Query {
			query.Add(key, value)
		}
		request.URL.RawQuery = query.Encode()
	}
	if r.RawQuery != "" {
		request.URL.RawQuery = r.RawQuery
	}
	if r.Headers != nil {
		for key, value := range r.Headers {
			request.Header.Set(key, value)
		}
	}
	if r.Json != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	responseRecorder := httptest.NewRecorder()
	testEngine.ServeHTTP(responseRecorder, request)

	return responseRecorder, nil
}

type WantResponse struct {
	Code int
	Body *response.Body
}

func (r *WantResponse) Equal(gotResponseRecorder *httptest.ResponseRecorder) bool {
	if r.Code != gotResponseRecorder.Code {
		return false
	}

	var wantResponseBodyBytes []byte
	if r.Body != nil {
		wantResponseBodyBytes, _ = json.Marshal(r.Body)
	}
	if !reflect.DeepEqual(wantResponseBodyBytes, gotResponseRecorder.Body.Bytes()) {
		return false
	}

	return true
}

func RunTestAPI(t *testing.T, request *TestRequest, wantResponse *WantResponse) {
	gotResponseRecorder, err := request.Do()
	if err != nil {
		t.Errorf("%s %s error = %v", request.Method, request.URL, err)
		return
	}
	if !wantResponse.Equal(gotResponseRecorder) {
		t.Errorf(
			"%s %s gotResponse code = %d, body = %s, want code = %d, body = %s",
			request.Method,
			request.URL,
			gotResponseRecorder.Code,
			gotResponseRecorder.Body.String(),
			wantResponse.Code,
			wantResponse.Body.String(),
		)
		return
	}
}
