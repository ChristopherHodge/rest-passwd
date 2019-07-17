package restapi

import (
	"net/http"
  "strings"
  "rest-passwd/passwd"
)

type Groups struct {}

// allow us to mock this in our API tests
var GetGroup = passwd.GetGroup

func (groups *Groups) AllHandler(resp http.ResponseWriter, req *http.Request) {
  entries, err := GetGroup() ; if err != nil {
    log.Warn(err)
    http.Error(resp, ServerErrorStr, http.StatusInternalServerError)
    return
  }
  for field, vals := range req.URL.Query() {
    for _, val := range vals {
      entries, err = entries.Find(field, val) ; if err != nil {
        log.Warn(err)
        http.Error(resp, BadRequestStr, http.StatusBadRequest)
        return
      }
    }
  }
  sendResponse(resp, entries.Entries)
}

func (groups *Groups) ByIdHandler(resp http.ResponseWriter, req *http.Request) {
  group, err := GetGroup() ; if err != nil {
    log.Warn(err)
    http.Error(resp, ServerErrorStr, http.StatusInternalServerError)
    return
  }
  id := strings.Split(req.URL.Path, "/")[2]
  group, err = group.Find("gid", id) ; if err != nil {
    log.Warn(err)
    http.Error(resp, NotFoundStr, http.StatusNotFound)
    return
  }
  if len(group.Entries) == 0 {
    http.Error(resp, NotFoundStr, http.StatusNotFound)
    return
  }
  sendResponse(resp, group.Entries[0])
}
