package test

func Prime(num uint) bool {
  switch num {
  case 0:
    return false
  case 1:
    return false
  }

  var i uint = 2

  for i < num {
    if num % i == 0 {
      return false
    }
    i++
  }
  return true
}
