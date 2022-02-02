package main

import (
	"os"
	"strings"

	"github.com/Hwisaek/Go/nomadcoder_go-for-beginners/scrapper"
	"github.com/labstack/echo"
)

<<<<<<< HEAD
func main() {
	dictionary := mydict.Dictioinary{"first": "First word"}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Update(baseWord, "Second")
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
=======
const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)

	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
>>>>>>> 2d21ccec975b5446a5bd706dc872c3403dea5dbb
}
