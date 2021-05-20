package output

import (
  "os"
  "log"
)

func CreateAndSave(input string) {
  file, err := os.Create("log.txt")
  if err != nil {
    log.Fatal(err)
  }

  _, err = file.WriteString(input)
  if err != nil {
    log.Fatal(err)
  }

  file.Close()
}
