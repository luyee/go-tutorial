package main

type NewString string

func main() {
	var my string = "a"
	var you NewString = my //cannot use my (type string) as type NewString in assignment
}
