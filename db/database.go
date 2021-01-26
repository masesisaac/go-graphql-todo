package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Todo type
type Todo struct {
	gorm.Model
	Text string
	Done bool
}

var db *gorm.DB

// InitDB initialize the db
func InitDB() {
	_db, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	_db.AutoMigrate(&Todo{})

	db = _db
}

// GetTodos gets all todos from the db
func GetTodos() (*[]Todo, error) {
	var todos []Todo
	result := db.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todos, nil
}

// GetTodo gets 1 todo from db
func GetTodo(id string) (*Todo, error) {
	todo := Todo{}
	result := db.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

// AddTodo inserts todo in db
func AddTodo(text string) (*Todo, error) {
	todo := Todo{Text: text, Done: false}
	result := db.Create(&todo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

// UpdateTodoDone updates done of a todo
func UpdateTodoDone(id string, done bool) (*Todo, error) {
	todo := Todo{}
	result := db.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}
	todo.Done = done
	result2 := db.Save(&todo)
	if result2.Error != nil {
		return nil, result2.Error
	}

	return &todo, nil
}

// DeleteTodo deletes a todo from db
func DeleteTodo(id string) (bool, error) {
	result := db.Delete(&Todo{}, id)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
