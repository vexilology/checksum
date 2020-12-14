package main

import (
  "flag"
  "testing"
)

var checkMessage = flag.Bool("message", false, "testing src/message.go")
var checkFile = flag.Bool("file", false, "testing src/file.go")

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
