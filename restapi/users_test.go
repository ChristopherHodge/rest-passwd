package restapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
  "encoding/json"
  "rest-passwd/passwd"
)

var passwdValidEntry = "user:*:123:321:Test User:/home/user:/path/to/shell"
var passwdValidEntry2 = "user2:*:124:321:Test User 2:/home/user2:/path/to/shell"
var passwdValidEntry3 = "user3:*:125:320:Test User 3:/home/user3:/path/to/shell"

func mockGetPasswd() (*passwd.Passwd, error) {
  var mockPasswd = &passwd.Passwd{}
  mockPasswd.AddLine(passwdValidEntry)
  mockPasswd.AddLine(passwdValidEntry2)
  mockPasswd.AddLine(passwdValidEntry3)
  return mockPasswd, nil
}

func TestGetUsers_All(t *testing.T) {
  GetPasswd = mockGetPasswd
	req, err := http.NewRequest("GET", "/users", nil) ; if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
  handler := http.HandlerFunc(new(Users).AllHandler)
	handler.ServeHTTP(rr, req)
  assertEqual(t, rr.Code, http.StatusOK)
  body := &passwd.Passwd{}
  json.Unmarshal(rr.Body.Bytes(), &body.Entries)
  assertEqual(t, len(body.Entries), 3)
}

func TestGetUsers_QueryInt(t *testing.T) {
  GetPasswd = mockGetPasswd
	req, err := http.NewRequest("GET", "/users", nil) ; if err != nil {
		t.Fatal(err)
	}
  q := req.URL.Query()
  q.Add("uid", "123")
  req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
  handler := http.HandlerFunc(new(Users).AllHandler)
	handler.ServeHTTP(rr, req)
  assertEqual(t, rr.Code, http.StatusOK)
  body := &passwd.Passwd{}
  json.Unmarshal(rr.Body.Bytes(), &body.Entries)
  assertEqual(t, len(body.Entries), 1)
}

func TestGetUsers_QueryString(t *testing.T) {
  GetPasswd = mockGetPasswd
	req, err := http.NewRequest("GET", "/users", nil) ; if err != nil {
		t.Fatal(err)
	}
  q := req.URL.Query()
  q.Add("name", "user")
  req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
  handler := http.HandlerFunc(new(Users).AllHandler)
	handler.ServeHTTP(rr, req)
  assertEqual(t, rr.Code, http.StatusOK)
  body := &passwd.Passwd{}
  json.Unmarshal(rr.Body.Bytes(), &body.Entries)
  assertEqual(t, len(body.Entries), 1)
}

func TestGetUsers_QueryMultiParams(t *testing.T) {
  GetPasswd = mockGetPasswd
	req, err := http.NewRequest("GET", "/users", nil) ; if err != nil {
		t.Fatal(err)
	}
  q := req.URL.Query()
  q.Add("shell", "/path/to/shell")
  q.Add("gid", "320")
  req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
  handler := http.HandlerFunc(new(Users).AllHandler)
	handler.ServeHTTP(rr, req)
  assertEqual(t, rr.Code, http.StatusOK)
  body := &passwd.Passwd{}
  json.Unmarshal(rr.Body.Bytes(), &body.Entries)
  assertEqual(t, len(body.Entries), 1)
}

func TestGetUserById_Exists(t *testing.T) {
  GetPasswd = mockGetPasswd
	req, err := http.NewRequest("GET", "/users/123", nil) ; if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
  handler := http.HandlerFunc(new(Users).ByIdHandler)
	handler.ServeHTTP(rr, req)
  assertEqual(t, rr.Code, http.StatusOK)
  body := make(map[string]interface{})
  json.Unmarshal(rr.Body.Bytes(), &body)
  assertEqual(t, body["name"], "user")
}

func TestGetUserById_NotExists(t *testing.T) {
  GetPasswd = mockGetPasswd
	req, err := http.NewRequest("GET", "/users/333", nil) ; if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
  handler := http.HandlerFunc(new(Users).ByIdHandler)
	handler.ServeHTTP(rr, req)
  assertEqual(t, rr.Code, http.StatusNotFound)
}

func TestGetUserGroups(t *testing.T) {
  GetPasswd = mockGetPasswd
  GetGroup = mockGetGroup
	req, err := http.NewRequest("GET", "/users/124/groups", nil) ; if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
  handler := http.HandlerFunc(new(Users).GetGroupsHandler)
	handler.ServeHTTP(rr, req)
  assertEqual(t, rr.Code, http.StatusOK)
  body := &passwd.Group{}
  json.Unmarshal(rr.Body.Bytes(), &body.Entries)
  assertEqual(t, len(body.Entries), 2)
  assertEqual(t, body.Entries[0]["name"], "group")
}
