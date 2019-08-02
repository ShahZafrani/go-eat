package main

import (
  "fmt"
  "math/rand"
  "time"
)

var num chan int
var dist chan map[int]float32
var curr_num int

func main() {
  num = make(chan int)
  dist = make(chan map[int]float32)
  go roulette(9, 27)
  go testRand(100000)
  fmt.Println("distribution between 9 and 26 (in %) \n", <-dist)
  fmt.Println("\a \a \a final num: ", <-num)
}

func newRand(max int, min int) int {
  rand.Seed(time.Now().UnixNano())
  return rand.Intn(max - min) + min
}

func testRand(test_size int) {
  dist_map := make(map[int]float32)
  for i :=9; i<27; i++ {
    dist_map[i] = 0.0
  }
  for i :=0; i< test_size; i++ {
    num:= newRand(27, 9)
    dist_map[num] += 1.0
  }
  // range on a map doesn't happen in order, but we don't care about that
  for k := range dist_map {
    // get a percent value
    dist_map[k] /= float32(test_size / 100)
  }
  dist <- dist_map
}

func roulette(min int, max int) {
  for i := 0; i < 100; i++ {
    fmt.Println("roulette: ", curr_num)
    time.Sleep(time.Millisecond * time.Duration(int(i^2)))
    curr_num = newRand(max, min)
  }
  num <- curr_num
}
