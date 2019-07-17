package restapi

import (
	"net/http"
  "strings"
  "rest-passwd/passwd"
)

type Users struct {}

// allow us to mock this in our API tests
var GetPasswd = passwd.GetPasswd

func (users *Users) AllHandler(resp http.ResponseWriter, req *http.Request) {
  entries, err := GetPasswd() ; if err != nil {
    http.Error(resp, ServerErrorStr, http.StatusInternalServerError)
    return
  }
  for field, vals := range req.URL.Query() {
    for _, val := range vals {
      entries, err = entries.Find(field, val) ; if err != nil {
        http.Error(resp, BadRequestStr, http.StatusBadRequest)
        return
      }
    }
  }
  sendResponse(resp, entries.Entries)
}

func (users *Users) ByIdHandler(resp http.ResponseWriter, req *http.Request) {
  id := strings.Split(req.URL.Path, "/")[2]
  user, err := GetPasswd() ; if err != nil {
    log.Warn(err)
    http.Error(resp, ServerErrorStr, http.StatusInternalServerError)
    return
  }
  user, err = user.Find("uid", id) ; if err != nil {
    log.Warn(err)
    http.Error(resp, ServerErrorStr, http.StatusInternalServerError)
    return
  }
  if len(user.Entries) == 0 {
    http.Error(resp, NotFoundStr, http.StatusNotFound)
    return
  }
  sendResponse(resp, user.Entries[0])
}

func (users *Users) GetGroupsHandler(resp http.ResponseWriter, req *http.Request) {
  id := strings.Split(req.URL.Path, "/")[2]
  user, err := GetPasswd() ; if err != nil {
    log.Warn(err)
    http.Error(resp, ServerErrorStr, http.StatusInternalServerError)
    return
  }
  user, err = user.Find("uid", id) ; if err != nil {
    log.Warn(err)
    http.Error(resp, ServerErrorStr, http.StatusInternalServerError)
    return
  }
  if len(user.Entries) == 0 {
    http.Error(resp, NotFoundStr, http.StatusNotFound)
    return
  }
  group, err := GetGroup() ; if err != nil {
    log.Warn(err)
    http.Error(resp, err.Error(), http.StatusInternalServerError)
    return
  }
  group, err = group.Find("member", user.Entries[0]["name"].(string)) ; if err != nil {
    log.Warn(err)
    http.Error(resp, ServerErrorStr, http.StatusInternalServerError)
    return
  }
  sendResponse(resp, group.Entries)
}
