package main

import (
  "flag"
  "./mapping"
)

func main() {
  xx := true

  stringFound := flag.Bool("m", false, "found message or string")
  fileFound := flag.String("f", "file", "found file")
  flag.Parse()

  if *stringFound {
    mapping.CheckMessage()
  } else if xx {
    mapping.CheckFile(*fileFound)
  }
}
