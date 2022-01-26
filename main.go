package main

import (
      "fmt"
)

func Union(a, b []int) []int {
      m := make(map[int]bool)

      for _, item := range a {
              m[item] = true
      }

      for _, item := range b {
              if _, ok := m[item]; !ok {
                      a = append(a, item)
              }
      }
      return a
}

func main() {
	// 2 slices of int to Union them
      var a = []int{1, 2, 3, 4, 5}
      var b = []int{2, 3, 5, 7, 11}
      fmt.Println(Union(a, b))
}