package main

import "fmt"

func main() {
  fmt.Println("Array Multi Dimensi")
  var angka = [5][2]int{{0, 0}, {2, 2}, {3, 4}, {5, 6}, {7, 8}}
  
  for i := 0; i < 5; i++ {
  for j := 0; j < 2; i++ {
    
  fmt.Printf("Date Ke [%d][%d] = %d\n", i, j, angka [i][j])
  }
  }
}
