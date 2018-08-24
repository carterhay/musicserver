
package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
  //"html/template"
  "net/http"
  "strings"
  tag "github.com/wtolson/go-taglib"
  "os"
  "path/filepath"
)

type Song struct{
  ID uint
  Title string
  Path string
  Artist string
  Album string
}

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


func get_files(path string) []Song {

  var list []Song

  err := filepath.Walk(path, visit)
  if err != nil {
    fmt.Println(err)
  }

  for index, path := range files {
    file, err := tag.Read(path)
    if err != nil {
      panic(err)
    }

    list = append(list, Song{uint(index), file.Title(), path, file.Artist(), file.Album()})

    file.Close()
  }

  return list

}


func main() {
  fmt.Println("Starting...")

  //Config
  port := ":8080"
  path := "/home/carter/Music"

  router := gin.Default()

  // I need to find a better way to do this!
  //router.StaticFS("/music", http.Dir(path))
  router.StaticFS("/assets", http.Dir("assets"))

  router.LoadHTMLFiles("templates/index.tmpl")

  list := get_files(path)

  //for i := 0; i < len(list); i++ {
  //  split := strings.Split(list[i].Path, "/")
  //  list[i].Path = split[len(split) - 1]
  //  fmt.Println("New Path: " + list[i].Path)
  //}

  for _,song := range list {
    fmt.Println(song.Path)
    fmt.Println(song.ID)
    fmt.Println("/music/" + fmt.Sprint(song.ID))
    router.StaticFile("/music/" + fmt.Sprint(song.ID), song.Path)
  }

  var test = make(map[string][]Song)
  test["songs"] = list

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", test)
  })
  router.Run(port)
}
