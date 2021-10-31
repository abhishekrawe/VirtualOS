package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator with Go and Fyne")

    output:=""
	input := widget.NewLabel(output)


	historyBtn:=widget.NewButton("history", func() {

	})
	backBtn:=widget.NewButton("back", func() {


	})
	clearBtn:= widget.NewButton("clear" , func() {})

	openBtn:= widget.NewButton("(" , func() {})

	closeBtn:= widget.NewButton(")" , func() {})

	divideBtn:= widget.NewButton("/" , func() {})

	sevenBtn:= widget.NewButton("7" , func() {})


	eightBtn:= widget.NewButton("8" , func() {})

	nineBtn:= widget.NewButton("9" , func() {})

	multiplyBtn:= widget.NewButton("*" , func() {})
    

	fourBtn:= widget.NewButton("4" , func() {})

	fiveBtn:= widget.NewButton("5" , func() {})

	sixBtn:= widget.NewButton("6" , func() {})

	minusBtn:= widget.NewButton("-" , func() {})



	oneBtn:= widget.NewButton("1" , func() {})

	twoBtn:= widget.NewButton("2" , func() {})

	threeBtn:= widget.NewButton("3" , func() {})

	plusBtn:= widget.NewButton("+" , func() {})



	zeroBtn:= widget.NewButton("0" , func() {})

	dotBtn:= widget.NewButton("." , func() {})

	equalBtn:= widget.NewButton("=" , func() {})

	

	
	w.SetContent(container.NewVBox(
		hello,
		container.NewGridWithColumns(2,
		historyBtn,
	    backBtn,
	    ),
		container.NewGridWithColumns(4,
		clearBtn,
		openBtn,
		closeBtn,
		divideBtn,),

		container.NewGridWithColumns(4,
		nineBtn,
	    eightBtn,
        sevenBtn,
        multiplyBtn,),
        container.NewGridWithColumns(4,
         fourBtn,
		 fiveBtn,
	     sixBtn,
         minusBtn,),
		 container.NewGridWithColumns(4,
		oneBtn,
		twoBtn,
		threeBtn,
		plusBtn,),
		container.NewGridWithColumns(2,
		container.NewGridWithColumns(2,
		zeroBtn,
		dotBtn,),
		equalBtn,),	
		),	
	)
		

	w.ShowAndRun()
}