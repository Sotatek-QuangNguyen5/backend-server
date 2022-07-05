package database

type People struct {
	
	Id   int
	Name string
	Age  int
}

var Db = []People{

	{1, "quang", 21},
	{2, "hoan", 21},
	{3, "ly", 21},
}
