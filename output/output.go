package output

import (
  "os"
  "log"
)

func GlobErr(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func CreateAndSave(input, create string) {
  if create != "" {
    file, err := os.Create(create)
    GlobErr(err)

    _, err = file.WriteString(input)
    GlobErr(err)

    file.Close()
  } else {
    file, err := os.Create("log.txt")
    GlobErr(err)

    _, err = file.WriteString(input)
    GlobErr(err)

    file.Close()
  }
}
