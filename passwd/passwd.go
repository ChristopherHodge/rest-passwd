package passwd

import (
  "strings"
  "strconv"
)

type Passwd struct {
  Entries PasswdEntries
}

func (passwd *Passwd) AddEntry(entry PasswdEntry) {
  passwd.Entries = append(passwd.Entries, entry)
}

// Parse a passwd file line and return pointer to PasswdEntry struct
func (passwd *Passwd) ParseEntry(line string) (PasswdEntry, error) {
  var passwdEntry = PasswdEntry{}
  var err error
  var fields []string = strings.Split(line, ":")
  if(len(fields) != 7) {
    return passwdEntry, InvalidFormat
  }
  passwdEntry["name"] = fields[0]
  passwdEntry["passwd"] = fields[1]
  passwdEntry["uid"], err = strconv.Atoi(fields[2]) ; if(err != nil) {
    return passwdEntry, FieldNotInt
  }
  passwdEntry["gid"], err = strconv.Atoi(fields[3]) ; if(err != nil) {
    return passwdEntry, FieldNotInt
  }
  passwdEntry["gecos"] = fields[4]
  passwdEntry["home"] = fields[5]
  passwdEntry["shell"] = fields[6]
  return passwdEntry, nil
}

// Return pointer to new copy of Passwd data
func GetPasswd() (*Passwd, error) {
  passwd := &Passwd{}
  err := passwd.Read(Config.PasswdFile) ; if err != nil {
    return passwd, err
  }
  return passwd, nil
}

// use findEntries() to filter our data
func (passwd *Passwd) Find(field string, value string) (*Passwd, error) {
  entries, err := findEntries(passwd.Entries, field, value) ; if err != nil {
    return passwd, err
  }
  return &Passwd{entries}, nil
}

func (passwd *Passwd) AddLine(line string) error {
  return AddLine(passwd, line)
}

func (passwd *Passwd) Read(file string) error {
  return Read(passwd, file)
}
