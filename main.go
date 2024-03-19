package main

import ( 
  "fmt"
)

func main(){

  temp := []float64{
    21.0, 23.00, 24.95, 21.05, 21.0, 21, 23,
  }
  freq := make(map[float64]int)
  
  for _, t := range temp {
      freq[t]++   
  }
  for t, f := range freq {
    fmt.Printf("%+.2f appears %d times\n", t, f)
  }
}