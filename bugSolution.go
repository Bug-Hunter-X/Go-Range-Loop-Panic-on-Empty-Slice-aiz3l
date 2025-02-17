func myFunc(a []int) {
  if len(a) > 0 {
    for i := range a {
      fmt.Println(a[i])
    }
  }
}