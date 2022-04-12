package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	s "strings"
)

type Person struct {
	last_name  string
	first_name string
	movies     string
}

func (p Person) String() string {

	return fmt.Sprintf("First name: %s, Last name: %s, Series: %s", p.first_name, p.last_name, p.movies)
}

type Sort interface {
	sort_the_slice(p int64) list
}

type list []Person

func (people list) sort_the_slice(p int64) {

	if p == 0 {
		sort.Slice(people, func(i, j int) bool { return people[i].last_name < people[j].last_name })

	} else if p == 1 {
		sort.Slice(people, func(i, j int) bool { return people[i].first_name < people[j].first_name })

	} else {
		sort.Slice(people, func(i, j int) bool { return people[i].movies < people[j].movies })

	}

}

func main() {

	file_name := os.Args[1]
	sort_on, err := strconv.ParseInt(os.Args[2], 10, 0)

	println(file_name, sort_on)

	file_content, err := os.Open(file_name)
	defer file_content.Close()
	check(err)

	scanner := bufio.NewScanner(file_content)
	each_person_line := make([]string, 0)

	for scanner.Scan() {
		each_person_line = append(each_person_line, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	people := []Person{}
	for i := range each_person_line {
		splice := s.Split(each_person_line[i], ",")
		last_name := splice[0]
		first_name := splice[1]
		movies := splice[2]
		persons := Person{last_name: last_name, first_name: first_name, movies: movies}
		people = append(people, persons)
		i++

	}

	names := list(people)
	names.sort_the_slice(sort_on)

	for j := range people {
		fmt.Println(people[j])
		j++
	}

}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
