package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

type Todo struct {
	ID   int
	Task string
}

var (
	todos    []Todo
	nextID   int
	todoLock sync.Mutex
)

func main() {
	http.HandleFunc("/fodase", indexH)
	http.HandleFunc("/addessabosta", addEssaBosta)
	http.HandleFunc("/deleteessabosta", deleteEssaBosta)
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("_\\|/_:42069")
	http.ListenAndServe(":42069", nil)
}

func indexH(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.ExecuteTemplate(w, "index", todos)
}

func addEssaBosta(w http.ResponseWriter, r *http.Request) {
	task := r.FormValue("task")
	if task != "" {
		todoLock.Lock()
		newTodo := Todo{ID: nextID, Task: task}
		todos = append(todos, newTodo)
		nextID++
		todoLock.Unlock()
		html := fmt.Sprintf(`
			<div id=%d style="">
				<span>%s</span>
				<button hx-get="/deleteessabosta?id=%d" hx-target="closest div" hx-swap="outerHTML">apag</button>
			</div>
		`, newTodo.ID, newTodo.Task, newTodo.ID)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))
	}
}

func deleteEssaBosta(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	todoLock.Lock()
	defer todoLock.Unlock()
	for i, todo := range todos {
		if fmt.Sprintf("%d", todo.ID) == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusOK)
}
