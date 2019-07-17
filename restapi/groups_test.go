package restapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
  "encoding/json"
  "rest-passwd/passwd"
)

var groupValidEntry = "group:*:123:user1,user2,user3"
var groupValidEntry2 = "group2:*:124:user4,user5,user6"
var groupValidEntry3 = "group3:*:125:user2,user3,user4"

func mockGetGroup() (*passwd.Group, error) {
  var mockGroup = &passwd.Group{}
  mockGroup.AddLine(groupValidEntry)
  mockGroup.AddLine(groupValidEntry2)
  mockGroup.AddLine(groupValidEntry3)
  return mockGroup, nil
}

func TestGetGroups_All(t *testing.T) {
  GetGroup = mockGetGroup
	req, err := http.NewRequest("GET", "/groups", nil) ; if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
  handler := http.HandlerFunc(new(Groups).AllHandler)
	handler.ServeHTTP(rr, req)
  assertEqual(t, rr.Code, http.StatusOK)
  body := &passwd.Group{}
  json.Unmarshal(rr.Body.Bytes(), &body.Entries)
  assertEqual(t, len(body.Entries), 3)
}

func TestGetGroups_QueryInt(t *testing.T) {
  GetGroup = mockGetGroup
  req, err := http.NewRequest("GET", "/groups", nil) ; if err != nil {
    t.Fatal(err)
  }
  q := req.URL.Query()
  q.Add("gid", "123")
  req.URL.RawQuery = q.Encode()
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(new(Groups).AllHandler)
  handler.ServeHTTP(rr, req)
  assertEqual(t, rr.Code, http.StatusOK)
  body := &passwd.Group{}
  json.Unmarshal(rr.Body.Bytes(), &body.Entries)
  assertEqual(t, len(body.Entries), 1)
}

func TestGetGroups_QueryString(t *testing.T) {
  GetGroup = mockGetGroup
  req, err := http.NewRequest("GET", "/groups", nil) ; if err != nil {
    t.Fatal(err)
  }
  q := req.URL.Query()
  q.Add("name", "group")
  req.URL.RawQuery = q.Encode()
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(new(Groups).AllHandler)
  handler.ServeHTTP(rr, req)
  assertEqual(t, rr.Code, http.StatusOK)
  body := &passwd.Group{}
  json.Unmarshal(rr.Body.Bytes(), &body.Entries)
  assertEqual(t, len(body.Entries), 1)
}

func TestGetGroups_QueryMultiParams(t *testing.T) {
  GetGroup = mockGetGroup
  req, err := http.NewRequest("GET", "/groups", nil) ; if err != nil {
    t.Fatal(err)
  }
  q := req.URL.Query()
  q.Add("member", "user3")
  q.Add("member", "user1")
  req.URL.RawQuery = q.Encode()
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(new(Groups).AllHandler)
  handler.ServeHTTP(rr, req)
  assertEqual(t, rr.Code, http.StatusOK)
  body := &passwd.Group{}
  json.Unmarshal(rr.Body.Bytes(), &body.Entries)
  assertEqual(t, len(body.Entries), 1)
}

func TestGetGroupById_Exists(t *testing.T) {
  GetGroup = mockGetGroup
  req, err := http.NewRequest("GET", "/groups/123", nil) ; if err != nil {
    t.Fatal(err)
  }
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(new(Groups).ByIdHandler)
  handler.ServeHTTP(rr, req)
  assertEqual(t, rr.Code, http.StatusOK)
  body := make(map[string]interface{})
  json.Unmarshal(rr.Body.Bytes(), &body)
  assertEqual(t, body["name"], "group")
}

func TestGetGroupById_NotExists(t *testing.T) {
  GetGroup = mockGetGroup
  req, err := http.NewRequest("GET", "/groups/333", nil) ; if err != nil {
    t.Fatal(err)
  }
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(new(Groups).ByIdHandler)
  handler.ServeHTTP(rr, req)
  assertEqual(t, rr.Code, http.StatusNotFound)
}

