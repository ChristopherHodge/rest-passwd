package passwd

import (
  "testing"
  "reflect"
  "errors"
  "os"
  "rest-passwd/config"
)

var groupBadEntry = "this:*:almost:looks:real"
var groupBadGid = "badgid:*:1x3:user1,user2,user3"
var groupValidEntry = "group:*:123:user1,user2,user3"
var groupValidEntry2 = "group2:*:124:user4,user5,user6"
var groupValidEntry3 = "group3:*:125:user2,user3,user4"

// the expected valid group PasswdEntry
func testGroupEntry() PasswdEntry {
  var entry PasswdEntry
  entry = make(map[string]interface{})
  entry["name"] = "group"
  entry["passwd"] = "*"
  entry["gid"] = 123
  member := make([]string, 0)
  member = append(member, "user1")
  member = append(member, "user2")
  member = append(member, "user3")
  entry["member"] = member
  return entry
}

func testGroup() *Group {
  group := &Group{}
  group.AddLine(groupValidEntry)
  group.AddLine(groupValidEntry2)
  group.AddLine(groupValidEntry3)
  return group
}

func TestGroupAddLine_BadFormat(t *testing.T) {
  group := &Group{}
  res := group.AddLine(groupBadEntry)

  assertTypeEqual(t, res, errors.New(""))
  assertEqual(t, res.Error(), InvalidFormat.Error())
}

func TestGroupAddLine_BadUid(t *testing.T) {
  group := &Group{}
  res := group.AddLine(groupBadGid)

  assertTypeEqual(t, res, errors.New(""))
  assertEqual(t, res.Error(), FieldNotInt.Error())
}

func TestGroupAddLine_Valid(t *testing.T) {
  group := &Group{}
  group.AddLine(groupValidEntry)

  assertEqual(t, len(group.Entries), 1)

  got := group.Entries[0]
  valid := testGroupEntry()

  assertTypeEqual(t, group, &Group{})
  if !reflect.DeepEqual(got, valid) {
    t.Fatalf("expecting: %+v\ngot: %+v\n", valid, got)
  }
}

func TestGetGroup_InvalidFile(t *testing.T) {
  Config := &config.Config{}
  Config.GroupFile = "invlaid-file"
  group, err := GetGroup()

  assertTypeEqual(t, group, &Group{})
  assertTypeEqual(t, err, new(os.PathError))
}

func TestGroupFind_IntOneExists(t *testing.T) {
  group := testGroup()
  res, _ := group.Find("gid", "123")

  assertEqual(t, len(res.Entries), 1)
  assertDeepEqual(t, res.Entries[0], testGroupEntry())
}

func TestGroupFind_StringOneExists(t *testing.T) {
  group := testGroup()
  res, _ := group.Find("name", "group")

  assertEqual(t, len(res.Entries), 1)
  assertDeepEqual(t, res.Entries[0], testGroupEntry())
}

func TestGroupFind_MemberOneExist(t *testing.T) {
  group := testGroup()
  res, _ := group.Find("member", "user1")

  assertEqual(t, len(res.Entries), 1)
}

func TestGroupFind_MemberTwoExist(t *testing.T) {
  group := testGroup()
  res, _ := group.Find("member", "user3")

  assertEqual(t, len(res.Entries), 2)
}

func TestGroupFind_NotExists(t *testing.T) {
  group := testGroup()
  res, _ := group.Find("name", "noexist")

  assertEqual(t, len(res.Entries), 0)
}

func TestGroupFind_BadField(t *testing.T) {
  group := testGroup()
  res, _ := group.Find("badfield", "value")

  assertEqual(t, len(res.Entries), 0)
}

