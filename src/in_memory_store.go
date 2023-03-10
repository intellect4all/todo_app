package main

import "time"

type InMemoryStore struct {
	storage map[int]Todo
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		storage: make(map[int]Todo),
	}
}

func (s *InMemoryStore) Set(value *Todo) (id int) {
	if count := s.Count(); count == 0 {
		id = 1
	} else {
		id = count + 1
	}
	value.CreatedAt = getNow()
	value.ID = id
	s.storage[id] = value.copyTodo()
	return
}

func (s *InMemoryStore) Get(key int) (Todo, bool) {
	value, ok := s.storage[key]
	return value, ok
}

func (s *InMemoryStore) Delete(key int) {
	delete(s.storage, key)
	return
}

// GetAll return all todos as an array
func (s *InMemoryStore) GetAll() (todos []Todo) {
	for _, value := range s.storage {
		todos = append(todos, value)
	}
	return
}

func (s *InMemoryStore) Count() int {
	return len(s.storage)
}

func (s *InMemoryStore) Clear() {
	s.storage = make(map[int]Todo)
	return
}

func (s *InMemoryStore) UpdateTodo(
	key int,
	title string,
	description string,
	status string,
) (updatedTodo Todo, ok bool) {
	updatedTodo, ok = s.Get(key)
	if !ok {
		return
	}
	updatedTodo.Title = title
	updatedTodo.Description = description
	updatedTodo.Status = status
	updatedTodo.UpdatedAt = getNow()
	s.storage[key] = updatedTodo
	return
}

func getNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
