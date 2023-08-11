package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	h := http.HandlerFunc(handler)
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "", string(body))
}

func TestDump(t *testing.T) {
	testCases := []struct {
		method     string
		path       string
		wantInfo   string
		wantHeader http.Header
		wantForm   url.Values
		wantBody   string
	}{
		{
			"GET", "/",
			"HTTP/1.1 GET example.com/", http.Header{}, url.Values(nil), "",
		},
	}
	for _, tc := range testCases {
		req := httptest.NewRequest(tc.method, tc.path, nil)
		info, header, form, body, err := dump(req)
		require.NoError(t, err)
		require.Equal(t, tc.wantInfo, info)
		require.Equal(t, tc.wantHeader, header)
		require.Equal(t, tc.wantForm, form)
		require.Equal(t, "", body)
	}
}
