package main

import (
  "fmt"
  "strings"
  "bufio"
  "io"
  "os"
  //"reflect"
)

func main(){
  var stdin *bufio.Reader
  var line []rune
  var c rune
  var err error
  var text string

  stdin = bufio.NewReader(os.Stdin)

  fmt.Printf("Type something: ")

  for {
          c, _, err = stdin.ReadRune()
          if err == io.EOF || c == '\n' { break }
          if err != nil {
                  fmt.Fprintf(os.Stderr,"Error reading standard input\n")
                  os.Exit(1)
          }
          line = append(line,c)
  }

  checkText := string(line[:len(line)])

  //remove the white space
  fix_text := strings.Replace(checkText, " ", "", -1)
  //fmt.Println(fix_text)

  //lower case the string
  text = strings.ToLower(fix_text)
  //fmt.Println(text)

  //fmt.Println(reflect.ValueOf(text).Kind())

  //trim null characters from the string
  text = strings.TrimSpace(text)

  //fmt.Println(strings.HasPrefix(text, "i"))
  //fmt.Println(strings.HasSuffix(text, "n"))
  //fmt.Println(strings.Contains(text, "a"))

  if(strings.HasPrefix(text, "i") == true && strings.HasSuffix(text, "n") == true && strings.Contains(text, "a") == true){
    fmt.Printf("\nFound!")
  }else{
    fmt.Printf("\nNot Found!")
  }
}
