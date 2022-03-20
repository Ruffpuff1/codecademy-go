package main

import "fmt"

func main () {
	var publisher string = "DizzyBooks Publishing Inc."
	var writer string = "Tracy Hatchet"
	var artist string = "Jewel Tampson"
	var title string = "Mr. GoToSleep"

	var year int = 1997
	var pageNumber int = 14

	var grade float32 = 6.5

	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher)
	fmt.Println("Released in", year, "graded as", grade, "with", pageNumber, "pages")

	writer = "DizzyBooks Publishing Inc."
	artist = "Phoebe Paperclips"
	title = "Epic Vol. 1"

	year = 2013
	pageNumber = 160

	grade = 9.0

	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher)
	fmt.Println("Released in", year, "graded as", grade, "with", pageNumber, "pages")
}