package main
import (
  "fmt"
)

type Person struct {
FirstName, LastName string
Age       int
}
type Couple struct {
  personA, personB Person
  relationship string
}
func main(){
  //creates an instance of the struct with zeros initialized for the values
  var empty Person
  fmt.Println(empty)
  var p = Person{"Jordan", "Rust", 64}
  fmt.Println(p)
  introduce(p)
  var c = Couple{p, Person{"Jerry", "Seinfeld", 58}, "platonic"}
  gossip(c)
}

func introduce(p Person) {
  fmt.Printf("Hello, I am %s %s and I am %d years old\n", p.FirstName, p.LastName, p.Age)
}

func fullName(p Person) string {
  return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}
func gossip(c Couple) {
  fmt.Printf("Did you hear %s is in an %s relationship with %s?\n",
              fullName(c.personA), c.relationship, fullName(c.personB))
}
