package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"sort"
	"strconv"
)

var rd *render.Render

type Todo struct {
	ID        int    `json:"id,omitempty"` // omitempty는 "ID" -> "id"로 변환되고 생략 가능함을 표시
	Name      string `json:"name"`
	Completed bool   `json:"completed,omitempty"`
}

type Todos []Todo

func (t Todos) Len() int {
	return len(t)
}
func (t Todos) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t Todos) Less(i, j int) bool {
	return t[i].ID > t[j].ID
}

type Success struct {
	Success bool `json:"success"`
}

var todoMap map[int]Todo
var lastId int = 0

func main() {
	rd = render.New()
	m := MakeWebHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}

func MakeWebHandler() http.Handler {
	todoMap = make(map[int]Todo)
	mux := mux.NewRouter()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/todos", GetTodoListHandler).Methods(http.MethodGet)
	mux.HandleFunc("/todos", PostTodoListHandler).Methods(http.MethodPost)
	mux.HandleFunc("/todos/{id:[0-9]+}", DeleteTodoListHandler).Methods(http.MethodDelete)
	mux.HandleFunc("/todos/{id:[0-9]+}", PutTodoListHandler).Methods(http.MethodPut)

	return mux
}

func PutTodoListHandler(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if todo, ok := todoMap[id]; ok {
		todo.Name = newTodo.Name
		todo.Completed = newTodo.Completed
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusBadRequest, Success{false})
	}
}

func DeleteTodoListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusNotFound, Success{false})
	}
}

func PostTodoListHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastId++
	todo.ID = lastId
	todoMap[lastId] = todo
	rd.JSON(w, http.StatusCreated, todo)
}

func GetTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Todos, 0)
	for _, todo := range todoMap {
		list = append(list, todo)
	}
	sort.Sort(list)
	rd.JSON(w, http.StatusOK, list)
}
