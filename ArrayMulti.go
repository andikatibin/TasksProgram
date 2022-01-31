package main

import "fmt"

func main() {
  fmt.Println("Array Multi Dimensi")
  var angka = [10][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}, {7, 8}, {8, 9}, {9, 10}, {10, 11}}
  
  for i := 0; i < 10; i++ {
  for j := 0; j < 2; i++ {
    
  fmt.Printf("[%d][%d] = %d\n", i, i, j, angka [i][j])
  }
  }
}
