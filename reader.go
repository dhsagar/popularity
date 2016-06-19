package main

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type collection struct {
	URL string `csv:"url"`
}

func ReadFile() []*collection {
	file, err := os.OpenFile("OnlineNewsPopularity.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	list := []*collection{}
	if err := gocsv.UnmarshalFile(file, &list); err != nil {
		log.Fatal(err)
	}
	return list
}

func WriteFile(post []*Post) error {
	file, err := os.OpenFile("result.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return gocsv.MarshalFile(post, file)
}

