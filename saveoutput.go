package main

import (
  "os"
  "fmt"
)

var create_output = "output.txt"

func SaveOutput() {
  createFile()
  writeFile()
}

func isError(err error) bool {
  if err != nil {
    fmt.Println(err.Error())
  }
  return (err != nil)
}

func createFile() {
  var _, err = os.Stat(create_output)

  if os.IsNotExist(err) {
    var file, err = os.Create(create_output)
    if isError(err) {
      return
    }
    defer file.Close()
  }
}

func writeFile() {
  var file, err = os.OpenFile(create_output, os.O_RDWR, 0644)
  if isError(err) {
    return
  }
  defer file.Close()

  // FIXME: need idea for realization
  _, err = file.WriteString(*fileFound)
  if isError(err) {
    return
  }

  _, err = file.WriteString(*algorithmName)
  if isError(err) {
    return
  }
}

func DeleteFile() {
  var err = os.Remove(create_output)
  if isError(err) {
    return
  }
}
