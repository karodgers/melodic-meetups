package database

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchData(t *testing.T) {
	type args struct {
		url    string
		target interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		handler http.HandlerFunc
	}{
		{
			name: "Successful Fetch and Decode",
			args: args{
				url:    "/success",
				target: &map[string]string{},
			},
			wantErr: false,
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"key": "value"}`))
			},
		},
		{
			name: "Invalid URL",
			args: args{
				url:    "http://%",
				target: &map[string]string{},
			},
			wantErr: true,
			handler: nil, // not needed for invalid URL
		},
		{
			name: "Non-OK Status Code",
			args: args{
				url:    "/non-ok",
				target: &map[string]string{},
			},
			wantErr: true,
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			},
		},
		{
			name: "Invalid JSON Format",
			args: args{
				url:    "/invalid-json",
				target: &map[string]string{},
			},
			wantErr: true,
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"key":`)) // invalid JSON
			},
		},
		{
			name: "Empty Response",
			args: args{
				url:    "/empty",
				target: &map[string]string{},
			},
			wantErr: true,
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(""))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test server if handler is provided
			var server *httptest.Server
			if tt.handler != nil {
				server = httptest.NewServer(tt.handler)
				defer server.Close()
				tt.args.url = server.URL + tt.args.url
			}

			err := FetchData(tt.args.url, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
