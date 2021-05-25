package main

import (
	"fmt"
	"time"

	"github.com/segmentio/ksuid"
)

type element struct {
	id          ksuid.KSUID
	title       string
	description string
	duedate     time.Time
}

//view elements in our list
func viewlist(l map[int]string) {
	fmt.Println("\nto do\n", l)
}

//add element to our list
func addtodo(e element, l map[int]string) {
	//assigning a UUID
	id := ksuid.New()
	e.id = id
	//inputing the title
	title := e.title
	fmt.Println("what are you going TO-DO?")
	fmt.Scanln(&title)
	//inputing description
	description := e.description
	fmt.Println("describe what do you want to do exaclt?")
	fmt.Scanln(&description)
	//inputing due date for the element
	//fmt.Println("when is tis task due (year, month, day) ?")

	fmt.Println(title)
	fmt.Println(description)

}

//update element in our list
func updatetodo(l map[int]string) {
	fmt.Println("\nthis is update")
}

//delete from list (takes id and delete from map by id)
func deletetodo(l map[int]string) {
	fmt.Println("\nthis is delete")
}

//marks if done returns boolean
func markasdone(l map[int]string) {
	fmt.Println("\nthis is marking as done / not done")
}

//filters results based on date (to be done or when updated/added?)
func datefilter(l map[int]string) {
	fmt.Println("\nthis is finding by date")
}

func main() {

	var s int
	fmt.Println("\nSellect from the following:-\n\n1.View List\n\n2.Add an element\n\n3.Update an element\n\n4.Delete an element\n\n5.Mark if done\n\n6.Find task by date")
	fmt.Scanln(&s)
	//initializing element
	todo := element{}
	//this is the TODO list
	list := map[int]string{}

	switch s {
	case 1:
		viewlist(list)
	case 2:
		addtodo(todo, list)
	case 3:
		updatetodo(list)
	case 4:
		deletetodo(list)
	case 5:
		markasdone(list)
	case 6:
		datefilter(list)

	}

	//send list to all

}
