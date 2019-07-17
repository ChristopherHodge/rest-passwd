package passwd

import (
  "testing"
  "errors"
  "os"
  "rest-passwd/config"
)

var passwdBadEntry = "this:*:almost:looks:real"
var passwdBadUid = "baduid:*:1x3:321:Test User:/home/user:/path/to/shell"
var passwdValidEntry = "user:*:123:321:Test User:/home/user:/path/to/shell"
var passwdValidEntry2 = "user2:*:124:321:Test User 2:/home/user2:/path/to/shell"
var passwdValidEntry3 = "user3:*:125:320:Test User 3:/home/user3:/path/to/shell"

func testPasswdEntry() PasswdEntry {
  var entry PasswdEntry
  entry = make(map[string]interface{})
  entry["name"] = "user"
  entry["passwd"] = "*"
  entry["uid"] = 123
  entry["gid"] = 321
  entry["gecos"] = "Test User"
  entry["home"] = "/home/user"
  entry["shell"] = "/path/to/shell"
  return entry
}

func testPasswd() *Passwd {
  passwd := &Passwd{}
  passwd.AddLine(passwdValidEntry)
  passwd.AddLine(passwdValidEntry2)
  passwd.AddLine(passwdValidEntry3)
  return passwd
}

func TestPasswdAddLine_BadFormat(t *testing.T) {
  passwd := &Passwd{}
  res := passwd.AddLine(passwdBadEntry)

  assertTypeEqual(t, res, errors.New(""))
  assertEqual(t, res.Error(), InvalidFormat.Error())
}

func TestPasswdAddLine_BadUid(t *testing.T) {
  passwd := &Passwd{}
  res := passwd.AddLine(passwdBadUid)

  assertTypeEqual(t, res, errors.New(""))
  assertEqual(t, res.Error(), FieldNotInt.Error())
}

func TestPasswdAddLine_Valid(t *testing.T) {
  passwd := &Passwd{}
  passwd.AddLine(passwdValidEntry)

  assertTypeEqual(t, passwd, &Passwd{})
  assertEqual(t, len(passwd.Entries), 1)
  assertDeepEqual(t, passwd.Entries[0], testPasswdEntry())
}

func TestGetPasswd_InvalidFile(t *testing.T) {
  Config := &config.Config{}
  Config.PasswdFile = "invlaid-file"
  passwd, err := GetPasswd()

  assertTypeEqual(t, passwd, &Passwd{})
  assertTypeEqual(t, err, new(os.PathError))
}

func TestPasswdFind_IntOneExists(t *testing.T) {
  passwd := testPasswd()
  res, _ := passwd.Find("uid", "123")

  assertEqual(t, len(res.Entries), 1)
  assertDeepEqual(t, res.Entries[0], testPasswdEntry())
}

func TestPasswdFind_StringOneExists(t *testing.T) {
  passwd := testPasswd()
  res, _ := passwd.Find("name", "user")

  assertEqual(t, len(res.Entries), 1)
  assertDeepEqual(t, res.Entries[0], testPasswdEntry())
}

func TestPasswdFind_TwoExist(t *testing.T) {
  passwd := testPasswd()
  res, _ := passwd.Find("gid", "321")

  assertEqual(t, len(res.Entries), 2)
}

func TestPasswdFind_NotExists(t *testing.T) {
  passwd := testPasswd()
  res, _ := passwd.Find("name", "noexist")

  assertEqual(t, len(res.Entries), 0)
}

func TestPasswdFind_BadField(t *testing.T) {
  passwd := testPasswd()
  res, _ := passwd.Find("badfield", "value")

  assertEqual(t, len(res.Entries), 0)
}
