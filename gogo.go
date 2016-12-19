package main


import (
  "fmt"
  "bufio"
  "strings"
  "strconv"
  "os"
)


const TOP = 0
const MIDDLE = 1
const BOTTOM = 2
const LEFT = 0
const RIGHT = 2

var ROW_DEF = [3][3]rune {
   {'\u250C','\u252C','\u2510'},
   {'\u251C','\u253C','\u2524'},
   {'\u2514','\u2534','\u2518'},
}

type Goban struct {
     size int
     board [][]rune
}

func NewGoban(size int)*Goban{
   goban := new(Goban)
   goban.size = size
   goban.board = make([][]rune, size)
   for i := range goban.board {
      goban.board[i] = make([]rune, size, 'G')
   }

   clear(goban)
   return goban
}

func draw_row(goban *Goban, row int){
  var row_def int
  if row == 0{
    row_def = TOP
  } else if row == goban.size - 1{
    row_def = BOTTOM
  } else {
    row_def = MIDDLE
  }
   x := row
   y := 0
   goban.board[x][y] = ROW_DEF[row_def][LEFT]
   for y := 1 ; y < goban.size - 1 ; y++{
      goban.board[x][y] = ROW_DEF[row_def][MIDDLE]
   }
   y = goban.size-1
   goban.board[x][y] = ROW_DEF[row_def][RIGHT]
}

func clear(goban *Goban){
   for x := 0 ; x < goban.size; x++{
      draw_row(goban, x)
   }
}

func display(goban *Goban){
   fmt.Println()
   fmt.Printf("%c", ' ')

   for x := 0 ; x < goban.size ; x++{
        fmt.Printf("%c", x + 'a')
   }
   fmt.Println()

   for x := 0 ; x < goban.size ; x++{
      fmt.Printf("%c", x + 'a')
      for y := 0 ; y < goban.size; y++{
          fmt.Printf("%c", goban.board[x][y])
      }
      fmt.Println()
   }
}


func main() {

   dat, err := os.Open("sample2.sgf")
   if err != nil {
     panic(err)
   }

   var goban *Goban
   var size int
   var color byte
   var col byte
   var row byte


   scanner := bufio.NewScanner(dat)
   for scanner.Scan() {
      text := scanner.Text()

      if strings.HasPrefix(text, "SZ"){
        pieces := strings.Split(text, "]")
        pieces = strings.Split(pieces[0], "[")

        sz, err := strconv.ParseInt(pieces[1],10, 32)
        if err != nil {
          panic(err)
        }
        size = int(sz)
        goban = NewGoban(size)
      } else if strings.HasPrefix(text, ";"){
        color = text[1]
        col_letter := text[3]
        row_letter := text[4]
        col = col_letter - 'a'
        row = row_letter - 'a'
        if color == 'B' {
           goban.board[row][col] = '\u25CB'
        } else if color == 'W'{
           goban.board[row][col] = '\u25CF'
        }
      }
   }
   //Tag the last stone placed
   if color == 'B' {
       goban.board[row][col] = '\u25CE'
   } else if color == 'W'{
       goban.board[row][col] = '\u25C9'
   }


   display(goban)
}