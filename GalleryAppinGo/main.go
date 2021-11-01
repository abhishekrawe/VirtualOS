package main

import (
	"fmt"
    "io/ioutil"
    "log"
	"strings"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	//"fyne.io/fyne/v2/theme"
)

func main() {
	a := app.New()
	w := a.NewWindow("Image Viewer")
	w.Resize(fyne.NewSize(800,630));
	root_src:="C:\\Users\\HP\\Pictures";
	files, err := ioutil.ReadDir(root_src);
	if err != nil {
        log.Fatal(err)
    }

	tabs := container.NewAppTabs(
		
	)

	for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir(){
           extension:= strings.Split(file.Name(), ".")[1];
		   if extension == "png" || extension == "jpeg" {
              
			  image:=canvas.NewImageFromFile(root_src + "\\"+file.Name());
			  image.FillMode = canvas.ImageFillOriginal
			  tabs.Append(container.NewTabItem(file.Name(), image))
			
			
			
			  } 
		}
    }	
	tabs.SetTabLocation(container.TabLocationLeading)
    w.SetContent(tabs);
	w.ShowAndRun()
} 



	
