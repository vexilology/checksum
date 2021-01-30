package main

import (
  "flag"
  "testing"
)

var (
  checkMessage = flag.Bool("message", false, "testing message.go")
  checkFile = flag.Bool("file", false, "testing file.go")
)

func TestFlagMessage(t *testing.T) {
  if *checkMessage {
    t.Skip()
  }
}

func TestFlagFile(t *testing.T) {
  if *checkFile {
    t.Skip()
  }
}
