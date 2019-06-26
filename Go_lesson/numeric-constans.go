package main

import "fmt"

const (
  // Create a huge number by shifting a 1 bit left 100 followed by 100 zeroes.
  Big = 1 << 100
  // In other words, the binary number that is 1 followed by 100 zeroes.
  Small = Big >> 99

)

func needInt(x int) int { return x * 10 + 1}
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
  fmt.Println(needFloat(Small))
  fmt.Println(needFloat(Big))
}
