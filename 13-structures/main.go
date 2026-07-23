package main

import "fmt"

func main() {
	// Think about this as a constructor in Java
	type Movie struct {
		id      int
		title   string
		author  string
		genre   string
		rating  float64
		watched bool
	}

	// Struct (object) instance
	movie := Movie{
		id:      01,
		title:   "Soul",
		author:  "Peter Docter",
		genre:   "comedy-drama",
		rating:  9.3,
		watched: true,
	}

	fmt.Println("Movie id:", movie.id)
	fmt.Println("Movie title:", movie.title)
	fmt.Println("Movie author:", movie.author)
	fmt.Println("Movie genre:", movie.genre)
	fmt.Println("Movie rating:", movie.rating)
	fmt.Println("Did I watched it?:", movie.watched)
}
