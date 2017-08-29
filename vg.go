package main

import (
    "path/filepath"
    "fmt"
    "os"
)


func printFile(path string, f os.FileInfo, err error) error {
  fmt.Printf("%s\n", path)
  return nil
} 

func main() {
    filepath.Walk("./", printFile)
}
