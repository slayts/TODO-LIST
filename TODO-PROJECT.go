package main

import "fmt"

type element struct {
	id  int
	txt string
}

//view elements in our list
func viewlist(l map[int]string) {
	fmt.Println("to do ", l)
}

//add element to our list
func addtodo() {

}

//update element in our list
func updatetodo() {

}

//delete from list (takes id and delete from map by id)
func deletetodo() {

}

//marks if done returns boolean
func markasdone() {

}

//filters results based on date (to be done or when updated/added?)
func datefilter() {

}

func main() {

	todo := element{
		id:  0,
		txt: "bla bla",
	}
	list := map[int]string{
		todo.id: todo.txt,
	}

	//send list to all
	viewlist(list)
	addtodo()
	updatetodo()
	deletetodo()
	markasdone()
	datefilter()
}
