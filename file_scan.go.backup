package main

import (
  tag "github.com/wtolson/go-taglib"
  "fmt"
  "os"
  "path/filepath"
  "strings"
)

var files []string

func IsMusic(file string) bool {
  split := strings.Split(file, ".")
  filetype := split[len(split) - 1]
  if filetype == "flac" || filetype == "mp3" || filetype == "m4a" {
    return true
  }
  return false
}

func visit(path string, r os.FileInfo, err error) error {
  if !r.IsDir() {
    if IsMusic(r.Name()) {
      //fmt.Println(r.Name())
      files = append(files, path)
    }
  }
  return nil
}

func main() {
  fmt.Println("Starting...")

  err := filepath.Walk("/home/carter/Music", visit)
  if err != nil {
    fmt.Println(err)
  }

  for _, path := range files {
    file, err := tag.Read(path)
    if err != nil {
      panic(err)
    }

    fmt.Println("Title: " + file.Title())
    fmt.Println("\tArtist: " + file.Artist())
    fmt.Println("\tAlbum: " + file.Album())
    //fmt.Println("Year: " + file.Year())
    //fmt.Println("Genre: " + file.Genre())
    //fmt.Println("Length: " + file.Length())

    file.Close()
  }
}
