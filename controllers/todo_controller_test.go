package controllers_test

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetTodoHandler(t *testing.T) {
	var tests = []struct {
		name       string
		resultCode int
	}{
		{
			name:       "正常系",
			resultCode: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/todo", nil)
			res := httptest.NewRecorder()
			con.GetTodoHandler(res, req)
			if res.Code != tt.resultCode {
				t.Errorf("invalid status code: %d", res.Code)
			}
		})
	}
}

func TestPostTodoHandler(t *testing.T) {
	var tests = []struct {
		name       string
		resultCode int
		reqBody    string
	}{
		{
			name:       "正常系",
			resultCode: 200,
			reqBody:    `{"task":"test","due_date":"2021-01-01T00:00:00Z"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/todo", strings.NewReader(tt.reqBody))
			res := httptest.NewRecorder()
			r := mux.NewRouter()
			r.HandleFunc("/todo", con.PostTodoHandler).Methods(http.MethodPost)
			r.ServeHTTP(res, req)
			if res.Code != tt.resultCode {
				t.Errorf("invalid status code: %d", res.Code)
			}
		})
	}
}

func TestPutTodoHandler(t *testing.T) {
	var tests = []struct {
		name       string
		query      string
		resultCode int
		body       string
	}{
		{
			name:       "正常なパラメータ",
			query:      "1",
			resultCode: 200,
			body:       `{"task":"test","due_date":"2021-01-01T00:00:00Z"}`,
		},
		{
			name:       "無効なパラメータ",
			query:      "a",
			resultCode: 400,
			body:       `{"task":"test","due_date":"2021-01-01T00:00:00Z"}`,
		},
		{
			name:       "無効なリクエストボディ",
			query:      "1",
			resultCode: 400,
			body:       `{"due_date":"２０２１年０１月０１日００時００分００秒"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("/todo/%s", tt.query)
			req := httptest.NewRequest(http.MethodPut, url, strings.NewReader(tt.body))
			res := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/todo/{id}", con.PutTodoHandler).Methods(http.MethodPut)
			r.ServeHTTP(res, req)
			if res.Code != tt.resultCode {
				t.Errorf("invalid status code: %d", res.Code)
			}
		})
	}
}

func TestDeleteTodoHandler(t *testing.T) {
	var tests = []struct {
		name       string
		query      string
		resultCode int
	}{
		{
			name:       "正常なパラメータ",
			query:      "1",
			resultCode: 200,
		},
		{
			name:       "無効なパラメータ",
			query:      "a",
			resultCode: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("/todo/%s", tt.query)
			req := httptest.NewRequest(http.MethodDelete, url, nil)
			res := httptest.NewRecorder()
			r := mux.NewRouter()
			r.HandleFunc("/todo/{id}", con.DeleteTodoHandler).Methods(http.MethodDelete)
			r.ServeHTTP(res, req)

			if res.Code != tt.resultCode {
				t.Errorf("invalid status code: %d", res.Code)
			}
		})
	}
}
