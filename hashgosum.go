package main

import (
  "fmt"
  "flag"
  "./src"
)

var (
  stringFound = flag.String("m", "", "found string")
  fileFound = flag.String("f", "", "found file")
  isHelp = flag.Bool("h", false, "help")
)

func parseFlags() {
  flag.Parse()

  if *isHelp {
    fmt.Println("Usage:")
    flag.PrintDefaults()
    return
  }

  if *stringFound != "" {
    src.CheckMessage(*stringFound)
    return
  }

  if *fileFound != "" {
    src.CheckFile(*fileFound)
    return
  } else {
    fmt.Println("Empty message, try again.")
    return
  }
}

func main() {
  parseFlags()
}
