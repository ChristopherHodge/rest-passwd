package restapi

import (
	"net/http"
  "encoding/json"
  "rest-passwd/logger"
)

type Server struct {}

type Service struct {
  Routes []Route
}

var log = logger.Get()

const ServerErrorStr = "error"
const NotFoundStr = "not found"
const BadRequestStr = "bad request"

// Respond to ServeHTTP to satisfy interface for http.ListenAndServe()
func (service *Service) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
  for _, route := range service.Routes {
    if route.Pattern.MatchString(req.URL.Path) {
      route.Handler(resp, req)
      return
    }
  }
  http.Error(resp, NotFoundStr, http.StatusNotFound)
}

// Start server and configure routes
func (server *Server) Listen(port string) {
  service := &Service{}
  users := &Users{}
  groups := &Groups{}
  service.AddRoute("/users$", users.AllHandler)
  service.AddRoute("/users/[0-9]+$", users.ByIdHandler)
  service.AddRoute("/users/[0-9]+/groups$", users.GetGroupsHandler)
  service.AddRoute("/groups$", groups.AllHandler)
  service.AddRoute("/groups/[0-9]+$", groups.ByIdHandler)
	log.Warn(http.ListenAndServe(":"+port, service))
}

// Marhsal and send JSON response
func sendResponse(resp http.ResponseWriter, body interface{}) error {
  js, err := json.Marshal(body) ; if err != nil {
    http.Error(resp, ServerErrorStr, http.StatusInternalServerError)
    return err
  }
  if string(js) == "null" {
    js = []byte("[]")
  }
  resp.Header().Set("Content-Type", "application/json")
  resp.Write(js)
  return nil
}
