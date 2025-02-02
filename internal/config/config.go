package config

type Task struct {
	Id      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment,omitempty"`
	Repeat  string `json:"repeat"`
}

type Err struct {
	Err string `json:"error,omitempty"`
}

type Id struct {
	Id int64 `json:"id,omitempty"`
}

type Tasks struct {
	ListOfTasks []Task `json:"tasks"`
}
