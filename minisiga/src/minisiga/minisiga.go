package minisiga

type Student struct {
	Name string `json:"name"`
	Age  int `json:"age"`
}

type Class struct {
	Students []Student `json:"students"`
}

type Course struct {
	PK 	    int `json:"pk"`
	Name    string `json:"name"`
	Classes []Class `json:"classes"`
}
