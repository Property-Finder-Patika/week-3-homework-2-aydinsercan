package main

import (
	"fmt"
	"log"
	"os"
	errRect "rectangleStruct/errors"
)

const LOG_FILE = "/Users/SERCAN/Documents/GitHub/week-3-homework-2-aydinsercan/RectangleAbstraction/Golog.txt"

type Rectangle struct {
	length float32
	width  float32
}

func main() {

	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()
	// Set log out put and enjoy :)
	log.SetOutput(logFile)
	// optional: log date-time, filename, and line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("Logging to the file")

	var l, w float32
	fmt.Println("Enter Length value : ")
	fmt.Scanln(&l)
	fmt.Println("Enter Width value : ")
	fmt.Scanln(&w)

	rect, err := createRect(l, w)
	if err != nil {
		log.Println("Creation failed: ", err)
		fmt.Println("Creating a new rectangle encounters Error")
		fmt.Println(err)
		return
	}

	//printing rectangle's length, width, area and circumference
	fmt.Println("Length : ", rect.length)
	fmt.Println("Width : ", rect.width)

	fmt.Println("Rectangle Area : ", rect.area())
	fmt.Println("Rectangle Circumference : ", rect.circumference())

}

//Init rectangle function, takes length and width, returns rect,err
func createRect(length float32, width float32) (*Rectangle, error) {
	rekt := Rectangle{}

	if length <= 0 {
		err := errRect.CheckLength(length)
		return nil, err
	}

	if width <= 0 {
		err := errRect.CheckWidth(width)
		return nil, err
	}

	//struct initialization
	rekt.width = width
	rekt.length = length

	return &rekt, nil
}

func (r Rectangle) area() float32 {
	return r.length * r.width
}

func (r Rectangle) circumference() float32 {
	return 2*r.length + 2*r.width
}

/*
Enter Length value :
12
Enter Width value :
12
12
12
Rectangle Area :  144
Rectangle Circumference :  48
*/
