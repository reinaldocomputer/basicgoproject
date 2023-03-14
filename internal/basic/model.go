package basic

type Request struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Age  int    `json:"age" db:"age"`
}

type Basic struct {
	Id   int
	Name string
	Age  int
}
