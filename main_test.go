package main

import (
	"net/http"
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Server starts and responds",
			want: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Start the server in a goroutine
			srv := &http.Server{Addr: ":9090"}
			go srv.ListenAndServe()

			// Give the server a moment to start
			time.Sleep(100 * time.Millisecond)

			// Test the server
			resp, err := http.Get("http://localhost:9090")
			if err != nil {
				t.Fatalf("Failed to send request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.want {
				t.Errorf("main() = %v, want %v", resp.StatusCode, tt.want)
			}

			// Shutdown the server after the test
			srv.Close()
		})
	}
}
