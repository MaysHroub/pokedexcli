package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)

  for {
    fmt.Print("Pokedex > ")
    
    if !scanner.Scan() {
      break
    }  
  
    input := strings.TrimSpace(scanner.Text())
    words := cleanInput(input)
    fmt.Println("Your command was: " + words[0])
  }  
}


func cleanInput(text string) []string {
  text = strings.ToLower(text)
  return strings.Fields(text)  
}

