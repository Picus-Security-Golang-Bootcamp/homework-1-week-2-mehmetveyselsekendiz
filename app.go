package main

import "fmt"

type Movie struct {
	name string
	year int8
	gerne string
	leading_actor string
}

func add_movie(movie_catalog map[string]Movie, m Movie){
	movie_catalog[m.name] = m
}

func main() {
	
	var movie_catalog map[string]Movie
	movie_catalog = make(map[string]Movie)


}