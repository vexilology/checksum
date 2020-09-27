package main

import (
  "flag"
  "./src"
)

func main() {
  for_string := true

  stringFound := flag.Bool("m", false, "found message or string")
  fileFound := flag.String("f", "file", "found file")
  flag.Parse()

  if *stringFound {
    src.CheckMessage()
  } else if for_string {
    src.CheckFile(*fileFound)
  }
}
