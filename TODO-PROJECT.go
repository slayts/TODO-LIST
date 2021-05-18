package main

import "fmt"

type element struct {
	id  int
	txt string
}

//view elements in our list
func viewlist(l map[int]string) {
	fmt.Println("\nto do ", l)
}

//add element to our list
func addtodo() {
	fmt.Println("\nthis is add")
}

//update element in our list
func updatetodo() {
	fmt.Println("\nthis is update")
}

//delete from list (takes id and delete from map by id)
func deletetodo() {
	fmt.Println("\nthis is delete")
}

//marks if done returns boolean
func markasdone() {
	fmt.Println("\nthis is marking as done / not done")
}

//filters results based on date (to be done or when updated/added?)
func datefilter() {
	fmt.Println("\nthis is finding by date")
}

func main() {
	var s int
	fmt.Println("\nSellect from the following:-\n\n1.View List\n\n2.Add an element\n\n3.Update an element\n\n4.Delete an element\n\n5.Mark if done\n\n6.Find task by date")
	fmt.Scanln(&s)

	todo := element{
		id:  0,
		txt: "bla bla",
	}
	list := map[int]string{
		todo.id: todo.txt,
	}

	switch s {
	case 1:
		viewlist(list)
	case 2:
		addtodo()
	case 3:
		updatetodo()
	case 4:
		deletetodo()
	case 5:
		markasdone()
	case 6:
		datefilter()

	}

	//send list to all

}
