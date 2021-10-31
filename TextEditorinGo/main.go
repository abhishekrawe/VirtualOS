package main

import(

   "fyne.io/fyne/v2"
   "fyne.io/fyne/v2/app"

)



func main() {

  a:= app.New()

  w:=a.NewWindow("Text Editor")

  w.Resize(fyne.NewSize(600, 600))

  w.ShowAndRun()

}