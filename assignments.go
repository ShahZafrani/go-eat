package main

import "fmt"

//global declaration
const company = "Slalom"

func main() {
  // scoped assignment
  profession := "Consultant"
  fmt.Println(formatBusinessCard("Shah", company, profession))
  // printBusinessCard("Shah Zafrani", company, profession)
}

// func _funcName_(param, param, param, type) return_type
func formatBusinessCard(n, c, p string) string {
  return c + "\n" + p + "\n" + n
}

// string formatting. Notice the Printf instead of Println
// func printBusinessCard(n, c, p string) {
//   fmt.Printf("%s\n%s\n%s\n", n, c, p)
// }
