package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var rows, columns int //global variable of rows and columns of the canvas and the digit images
var size int          //global variable of area of each image in pixels (28*28 = 784)
var n int             //global variable of mnist image count
const outCount = 10   //output count of the digits from 0 to 9

func main() {

	//######################## MNIST data read #########################

	//reading MNIST data set
	//pixels := readGZ("data/train-images-idx3-ubyte.gz")

	//Printing one digit from data set
	//i := 59999 //digit index (min - 0 | max - 59999)
	//printMnist(pixels.RawMatrix().Data[i*size : (i+1)*size])

	//####################### Ebiten #########################

	//Window
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Digitopia")

	//App instance
	a := NewApp()

	//Running app
	if err := ebiten.RunGame(a); err != nil {
		log.Fatal(err)
	}

	//####################### Print #########################

	//print the stuff here
}
