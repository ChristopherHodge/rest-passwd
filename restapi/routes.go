package restapi

import (
	"net/http"
  "regexp"
)

type Handler func(http.ResponseWriter, *http.Request)

type Route struct {
  Pattern *regexp.Regexp
  Handler Handler
}

// Register route handler by regex
func (service *Service) AddRoute(pattern string, handler Handler) {
  route := Route{ Pattern: regexp.MustCompile(pattern), Handler: handler }
  service.Routes = append(service.Routes, route)
}

// Find matching handler, or return 404
func (service *Service) Handler(resp http.ResponseWriter, req *http.Request) {
  for _, route := range service.Routes {
    if route.Pattern.MatchString(req.URL.Path) {
      route.Handler(resp, req)
      return
    }
  }
  resp.WriteHeader(404)
}

