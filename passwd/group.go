package passwd

import (
  "strings"
  "strconv"
)

// Group struct
type Group struct {
  Entries PasswdEntries
}

func (group *Group) AddEntry(entry PasswdEntry) {
  group.Entries = append(group.Entries, entry)
}

// Parse a group file line and return pointer to GroupEntry struct
func (group *Group) ParseEntry(line string) (PasswdEntry, error) {
  groupEntry := PasswdEntry{}
  var err error
  var fields []string = strings.Split(line, ":")
  if(len(fields) != 4) {
    return groupEntry, InvalidFormat
  }
  groupEntry["name"] = fields[0]
  groupEntry["passwd"] = fields[1]
  groupEntry["gid"], err = strconv.Atoi(fields[2]) ; if(err != nil) {
    return groupEntry, FieldNotInt
  }
  groupEntry["member"] = strings.Split(fields[3], ",")
  return groupEntry, nil
}

// Return pointer to new copy of Group data
func GetGroup() (*Group, error)  {
  group := &Group{}
  err := group.Read(Config.GroupFile) ; if err != nil {
    return group, err
  }
  return group, nil
}

// use findEntries() to filer our data
func (group *Group) Find(field string, value string) (*Group, error) {
  entries, err := findEntries(group.Entries, field, value) ; if err != nil {
    return group, err
  }
  return &Group{entries}, nil
}

func (group *Group) AddLine(line string) error {
  return AddLine(group, line)
}

func (group *Group) Read(file string) (error) {
  return Read(group, file)
}
