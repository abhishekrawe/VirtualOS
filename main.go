package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")



	historyBtn:=widget.NewButton("history", func() {

	})
	backBtn:=widget.NewButton("back", func() {


	})
	clearBtn:= widget.NewButton("clear" , func() {})

	openBtn:= widget.NewButton("clear" , func() {})

	closeBtn:= widget.NewButton("clear" , func() {})

	divideBtn:= widget.NewButton("clear" , func() {})

	sevenBtn:= widget.NewButton("clear" , func() {})


	eightBtn:= widget.NewButton("clear" , func() {})

	nineBtn:= widget.NewButton("clear" , func() {})

	multiplyBtn:= widget.NewButton("clear" , func() {})
    

	fourBtn:= widget.NewButton("clear" , func() {})

	fiveBtn:= widget.NewButton("clear" , func() {})

	sixBtn:= widget.NewButton("clear" , func() {})

	minusBtn:= widget.NewButton("clear" , func() {})



	oneBtn:= widget.NewButton("clear" , func() {})

	twoBtn:= widget.NewButton("clear" , func() {})

	threeBtn:= widget.NewButton("clear" , func() {})

	plusBtn:= widget.NewButton("clear" , func() {})



	zeroBtn:= widget.NewButton("clear" , func() {})

	dotBtn:= widget.NewButton("clear" , func() {})

	equalBtn:= widget.NewButton("clear" , func() {})

	

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hello World!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}