package main


import (

	"fyne.io/fyne"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

)


var myApp fyne.App = app.New()


var myWindow fyne.Window = myApp.NewWindow("My OS")


var btn1 fyne.Widget

var btn2 fyne.Widget

var btn3 fyne.Widget

var btn4 fyne.Widget

var img fyne.CanvasObject;

var DeskBtn fyne.Widget

var panelContent *fyne.Container

func main() {
	myApp.Settings().SetTheme(theme.LightTheme())
	img = canvas.NewImageFromFile("D:\\VirtualOS\\MY OS\\images.jpg")
   
	btn1 = widget.NewButtonWithIcon("Weather App", theme.InfoIcon(), func() {
		showWeatherApp(myWindow)
	})



}



