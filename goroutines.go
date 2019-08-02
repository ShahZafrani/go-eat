package main

import "fmt"

func main() {
  var boxes = createBoxes(100)
  // fmt.Println(boxes)
  // moveBoxes(boxes)
  // moveBoxes(boxes[:50])
  // moveBoxes(boxes[50:])
  mover1 := make(chan string)
  mover2 := make(chan string)
  go moveBoxes(boxes[:70], mover1)
  go moveBoxes(boxes[70:], mover2)
  fmt.Println(<-mover1)
  fmt.Println(<-mover2)
}

func createBoxes(numboxes int) []int {
  boxes := make([]int, numboxes)
  for i := 0; i < numboxes; i ++ {
    boxes[i] = i
  }
  return boxes
}

func moveBoxes(boxes []int, mover chan string) {
  for _, box := range boxes {
    fmt.Printf("moving %d\n", box)
  }
  mover <- fmt.Sprintf("moved %d boxes", len(boxes))
}
