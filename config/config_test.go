package config

import (
  "testing"
  "errors"
)

var testConfigJson=[]byte(`{"Port":"1234","PasswdFile":"/test/passwd","GroupFile":"/test/group"}`)

func mockReadFile(file string) ([]byte, error) {
  return testConfigJson, nil
}

func TestReadConfig_Success(t *testing.T) {
  ReadFile = mockReadFile
  config := Read("/fake/path")
  assertEqual(t, config.Port, "1234")
}

func TestReadConfig_Fail(t *testing.T) {
  ReadFile = func(file string) ([]byte, error) {
    return nil, errors.New("")
  }
  config := Read("/fake/path")
  assertEqual(t, config.Port, "8080")
}
