package restapi

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestAddRoute(t *testing.T) {
  service := &Service{}
  var someHandler Handler
  someHandler = func(http.ResponseWriter, *http.Request) {}
  service.AddRoute("/abc/[0-9]+", someHandler)
  assertEqual(t, len(service.Routes), 1)
}

func TestRouteHandler(t *testing.T) {
  service := &Service{}
  var someHandler Handler
  someHandler = func(resp http.ResponseWriter, req *http.Request) {
    resp.Write([]byte("success"))
  }
  service.AddRoute("/abc/[0-9]+", someHandler)
  req, err := http.NewRequest("GET", "/abc/123", nil) ; if err != nil {
    t.Fatal(err)
  }
  rr := httptest.NewRecorder()
  service.Handler(rr, req)
  assertEqual(t, string(rr.Body.Bytes()), "success")
}
