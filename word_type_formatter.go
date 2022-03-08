package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT.
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error reading stdin: %s", err)
	}
	inputWords := strings.Split(strings.TrimSpace(string(bytes)), " ")

	wordTypeFinders := map[WordType]Finder{
		Fruit:     NewFruitFinder(),
		Vegetable: NewVegetableFinder(),
		Animal:    NewAnimalFinder(),
	}

	result := WordTypeFormatter(inputWords, wordTypeFinders)

	fmt.Println(result)
}

func WordTypeFormatter(inputWords []string, wordTypeFinders map[WordType]Finder) string {
	var result []string
	var found bool

	for _, word := range inputWords {
		found = false
		for wordType, finder := range wordTypeFinders {
			if finder.Contains(word) {
				formatter := FormatterFactory(wordType)
				if formatter != nil {
					result = append(result, formatter.Format(word))
					found = true
					break
				}
			}
		}

		if !found {
			result = append(result, fmt.Sprintf("Unknown word: %s", word))
		}
	}

	return fmt.Sprintf(strings.Join(result, " "))
}

// Finder provides a generic way for determining if a word is in a list of valid words
type Finder interface {
	Contains(s string) bool
}

type FinderBase struct {
	names []string
}

func (f FinderBase) Contains(s string) bool {
	for _, name := range f.names {
		if name == s {
			return true
		}
	}

	return false
}

type FruitFinder struct {
	FinderBase
}

func NewFruitFinder() Finder {
	return FruitFinder{
		FinderBase: FinderBase{
			names: []string{"apple", "banana", "mango"},
		},
	}
}

type VegetableFinder struct {
	FinderBase
}

func NewVegetableFinder() Finder {
	return VegetableFinder{
		FinderBase: FinderBase{
			names: []string{"carrot", "zucchini", "broccoli"},
		},
	}
}

type AnimalFinder struct {
	FinderBase
}

func NewAnimalFinder() Finder {
	return AnimalFinder{
		FinderBase: FinderBase{
			names: []string{"horse", "giraffe", "mouse", "pigeon"},
		},
	}
}

// Formatter is an interface for formatting strings
type Formatter interface {
	Format(name string) string
}

type FruitFormatter struct{}

func (f FruitFormatter) Format(name string) string {
	return strings.ToUpper(name)
}

type VegetableFormatter struct{}

func (v VegetableFormatter) Format(name string) string {
	return fmt.Sprintf("[%s]", name)
}

type AnimalFormatter struct{}

func (a AnimalFormatter) Format(name string) string {
	var formattedName string
	for i, char := range name {
		if i < len(name)-1 {
			formattedName += string(char) + "*"
		} else {
			formattedName += string(char)
		}
	}

	return formattedName
}

type WordType int

const (
	Fruit WordType = iota
	Vegetable
	Animal
)

func FormatterFactory(wordType WordType) Formatter {
	switch wordType {
	case Fruit:
		return FruitFormatter{}
	case Vegetable:
		return VegetableFormatter{}
	case Animal:
		return AnimalFormatter{}
	}

	return nil
}
