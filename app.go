package main

import "fmt"

type movie struct {
	name string
	year int8
	gerne string
	leading_actor string
}

func (m movie) get_movie_name() string{
	return m.name
}

func main() {
	fmt.Println("hello world")
}