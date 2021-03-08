package sync

import "fmt"

type Book struct {
	Id            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	{
		Id:            1,
		Title:         "book 1",
		Author:        "Cortazar",
		YearPublished: 1992,
	},
	{
		Id:            2,
		Title:         "book 2",
		Author:        "Tolkien",
		YearPublished: 1992,
	},
	{
		Id:            3,
		Title:         "book 3",
		Author:        "Ende",
		YearPublished: 1992,
	},
	{
		Id:            4,
		Title:         "book 4",
		Author:        "Marquez",
		YearPublished: 1992,
	},
	{
		Id:            5,
		Title:         "book 5",
		Author:        "Galdos",
		YearPublished: 1992,
	},
	{
		Id:            6,
		Title:         "book 6",
		Author:        "Reverte",
		YearPublished: 1992,
	},
	{
		Id:            7,
		Title:         "book 7",
		Author:        "Hemingway",
		YearPublished: 1992,
	},
	{
		Id:            8,
		Title:         "book 8",
		Author:        "Clarín",
		YearPublished: 1992,
	},
	{
		Id:            9,
		Title:         "book 9",
		Author:        "Follet",
		YearPublished: 1992,
	},
	{
		Id:            10,
		Title:         "book 10",
		Author:        "Twain",
		YearPublished: 1992,
	},
}
