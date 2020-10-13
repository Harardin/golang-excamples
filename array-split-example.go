/*
Excample of spliting Array in few minor arrays

Might be useful when transfering big data via socket

*/

package main

import (
  "fmt"
)

func main() {
  arr := []int{1, 2, 3, 5, 6, 6, 7, 7, 7, 8, 8, 8, 8, 9}
  limit := 5

  for i := 0; i < len(arr); i += limit {
    batch := arr[i:min(i+limit, len(arr))]
    fmt.Println(batch)
  }

}

func min(a, b int) int {
  if a <= b {
    return a
  }
  return b
}
