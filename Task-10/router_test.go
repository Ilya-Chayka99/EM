package Task_10

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	server := httptest.NewServer(router())
	defer server.Close()

	resp, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if string(body) != "Hello World" {
		t.Fatalf("expected Hello World, got %s", string(body))
	}
}
