package main

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

func main() {
	var stringer Stringer
	stringer.(*Student)
}
