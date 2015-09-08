package main

func main() {
  yo := make([]int, 3, 10)
  
  yo[0] = 124
  yo[1] = 335
  yo[2] = 35
  yo[3] = 0   // Throws index out of range error
}