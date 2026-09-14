package main

import "fmt"

func main() {
	courses = append(courses, Course{
		Courseid:   "2",
		Coursename: "Go API Development",
		Price:      199,
		Author: &Author{
			Name:    "Hitesh Choudhary",
			Website: "learncodeonline.in",
		},
	})

	fmt.Println(courses[0].isEmpty())
}

var courses []Course

type Course struct {
	Courseid   string  `json:"courseid"`
	Coursename string  `json:"coursename"`
	Price      int     `json:"price"`
	Author     *Author `json:"author"` // Pointer to the Author stuct
}
type Author struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}

func (c *Course) isEmpty() bool {
	return c.Courseid == "" && c.Coursename == ""
}
