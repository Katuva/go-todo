package structs

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
)

// Todo Holds our structure reference
var Todo STodos

// STodo Todo item
type STodo struct {
	Context     []string `json:"context"`
	Tags        []string `json:"tags"`
	Complete    bool     `json:"complete"`
	Content     string   `json:"content"`
	Due         string   `json:"due"`
	Created     string   `json:"created"`
	RepeatEvery string   `json:"repeat_every"`
}

// STodos Todo items
type STodos struct {
	Todos []STodo `json:"todos"`
}

// Parse Parse json into struct
func (t *STodos) Parse(file string) {
	bTodo, err := ioutil.ReadFile(file)
	err = json.Unmarshal(bTodo, t)
	if err != nil {
		log.WithField("err", err).Fatal("Couldn't decode todo file")
	}
}

// Add Add a todo item to the struct
func (t *STodos) Add(todo STodo) *STodos {
	t.Todos = append(t.Todos, todo)
	return t
}
