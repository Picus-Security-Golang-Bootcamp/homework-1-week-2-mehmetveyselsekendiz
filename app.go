package main

import (
	"fmt" 
	"strings"
)

type Movie struct {
	name string
	year int8
	gerne string
	leading_actor string
}

func add_movie(movie_catalog map[string]Movie, m Movie){
	movie_catalog[strings.ToLower(m.name)] = m
}

func search_movie_catalog(movie_catalog map[string]Movie, filter string){

	var status = []bool{true}

	for movie := range movie_catalog{
		if(movie == strings.ToLower(filter)){
			fmt.Println(movie_catalog[movie])
			status[0] = false
			break
		}
	}

	if(status[0]){fmt.Println("Movie is not available.")}

}

func main() {
	
	var movie_catalog map[string]Movie
	movie_catalog = make(map[string]Movie)

	movie1 := Movie{"name1",1,"gerne1","actor1"}
	movie2 := Movie{"name2",2,"gerne2","actor2"}


	add_movie(movie_catalog, movie1)
	add_movie(movie_catalog, movie2)

	search_movie_catalog(movie_catalog, "name1")
}