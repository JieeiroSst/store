package menues

type Menues struct {
	Id int
	Text string
	Link string
	DisplayOrder int
	Target string
	Status int
}

type InputMenues struct {
	Text string
	Link string
	DisplayOrder int
	Target string
	Status int
}