package passwd

import(
  "strconv"
  "os"
  "bufio"
  "path/filepath"
  "rest-passwd/config"
  "rest-passwd/logger"
)

type PasswdEntry map[string]interface{}

type PasswdEntries []PasswdEntry

type PasswdType interface {
  ParseEntry(string) (PasswdEntry, error)
  AddEntry(PasswdEntry)
}

var Config = config.Get()

var log = logger.Get()

// call necessary parser and add line to struct
func AddLine(passwd PasswdType, line string) (error) {
  entry, err := passwd.ParseEntry(line) ; if err != nil {
    return err
  }
  passwd.AddEntry(entry)
  return nil
}

// parse the file for the appropraite type
func Read(passwd PasswdType, file string) (error) {
  return read(file, func(line string) (error) {
    return AddLine(passwd, line)
  })
}

// read file with appropriate type adder function
func read(file string, adder func(entry string) error) error {
  fh, err := os.Open(file) ; if err != nil {
    return err
  }
  scanner := bufio.NewScanner(fh)
  ln := 0  // line number for log
  for scanner.Scan() {
    ln += 1
    // Ignore invalid entries, but continue parsing
    err := adder(scanner.Text()) ; if(err != nil) {
      log.Warn("failed to parse entry: ", filepath.Base(file), ":", ln, err)
    }
  }
  return nil
}

// search PasswdEntries key/value matches
func findEntries(entries PasswdEntries, field string, value string) (PasswdEntries, error) {
  var match PasswdEntries
  for _, entry := range entries {
    switch entry[field].(type) {
      case string:
        if entry[field] == value {
          match = append(match, entry)
        }
      case int:
        value_i, err := strconv.Atoi(value) ; if err != nil {
          return match, err
        }
        if entry[field] == value_i {
          match = append(match, entry)
        }
      case []string:
        for _, item := range entry[field].([]string) {
          if item == value {
            match = append(match, entry)
          }
        }
    }
  }
  return match, nil
}
