package main

import (
	"log"
)

func main() {
	collections := ReadFile()
	postList := make([]*Post, 0)
	for i, c := range collections {
		log.Println("trying", i)
		id := Parse(c)
		post := GetPost(id)
		postList = append(postList, post)
		if i >= 1 {
			break
		}
	}
	err := WriteFile(postList)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("saved all data")
}
