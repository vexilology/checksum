package main

import (
  "testing"
)

func TestMap(t *testing.T) {
  for i := 0; i < len(algorithm_list); i++ {
    b := algorithm_list[i]
    if _, ok := h[b]; ok {
      t.Logf("found: %v", b)
    }
  }
}
