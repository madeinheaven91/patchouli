package router

import (
	"io"
	"net/http"
	"testing"
)

func TestRouting(t *testing.T) {
	r := NewRouter(http.NewServeMux())
	r.Route = NewRoute("/root").
		Handler(testHandler{}).
		Subroute(
			NewRoute("/unspecified"),
		).
		Subroute(
			NewRoute("/test1").
				Handler(testHandler{}).
				Subroute(
					NewRoute("/{wildcard}").
						Handler(wildcardHandler{}),
				),
		).
		Subroute(
			NewRoute("/test2").
				Handler(wildcardHandler{}).
				Subroute(
					NewRoute("/post").
						Method("POST"),
				),
		)
	go http.ListenAndServe(":8080", r.BuildMux())

	tests := []struct {
		name           string
		url            string
		fn             func(string) (resp *http.Response, err error)
		expectedStatus int
		expectedBody   string
	}{
		{
			"root",
			"http://localhost:8080/root",
			http.Get,
			200,
			"Test handler",
		},
		{
			"unspecified",
			"http://localhost:8080/root/unspecified",
			http.Get,
			500,
			"none",
		},
		{
			"not found",
			"http://localhost:8080/unknown",
			http.Get,
			404,
			"none",
		},
		{
			"test1",
			"http://localhost:8080/root/test1",
			http.Get,
			200,
			"Test handler",
		},
		{
			"wildcard",
			"http://localhost:8080/root/test1/wildcard",
			http.Get,
			200,
			"wildcard",
		},
		{
			"wildcard2",
			"http://localhost:8080/root/test1/wildcard2",
			http.Get,
			200,
			"wildcard2",
		},
		{
			"wrong method",
			"http://localhost:8080/root/test2/post",
			http.Get,
			405,
			"none",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := tt.fn(tt.url)
			if err != nil {
				panic(err)
			}
			if tt.expectedStatus != resp.StatusCode {
				t.Fatalf("status codes not maching. expected %d, got %d", tt.expectedStatus, resp.StatusCode)
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			if tt.expectedBody != "none" && string(body) != tt.expectedBody {
				t.Fatalf("bodies not maching. expected %s, got %s", tt.expectedBody, string(body))
			}
			resp.Body.Close()
		})
	}
}

type testHandler struct{}

func (d testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test handler"))
}

type wildcardHandler struct{}

func (d wildcardHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.PathValue("wildcard")))
}
