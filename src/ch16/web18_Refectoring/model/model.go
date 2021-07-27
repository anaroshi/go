package model

import (
	"time"
)

type Todo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var todoMap map[int]*Todo

// 처음에 한번만 읽어들인다. 초기화시 사용
func init() {
	todoMap = make(map[int]*Todo)
}

func GetTodoMap() []*Todo {
	// JSON을 반환
	list := []*Todo{}	
	for _, v := range todoMap {
		list = append(list, v)
	}
	return list
}

func AddTodoMap(name string) *Todo {
	id := len(todoMap)+1
	todo := &Todo{id, name, false, time.Now()}
	todoMap[id] = todo
	return todo
}

func RemoveTodoMap(id int)  bool {
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		return true
	}
	return false
}

func CompleteTodoMap(id int, complete bool) bool {
	if todo, ok := todoMap[id]; ok {
		todo.Completed = complete
		return true
	}
	return false
}
