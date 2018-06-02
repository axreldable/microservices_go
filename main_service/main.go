package main

import (
	"net/http"

	"golang-2018-1/7/99_hw/models"

	"html/template"

	"github.com/gorilla/mux"
	"fmt"
	"context"
	"log"
	"google.golang.org/grpc"
	"strconv"
	"strings"
)

var id = 0

type TemplateData struct {
	Email string
	Tasks []*models.Task
}

type Handler struct {
	Auth  models.AuthClient
	Tasks models.TasksClient
	Tmpl  *template.Template
}

func (h *Handler) Root(w http.ResponseWriter, r *http.Request) {
	user, err := h.checkSession(w, r)
	if isError(err, w) {
		return
	}

	taskList, err := h.Tasks.List(
		context.Background(),
		user,
	)
	if isError(err, w) {
		return
	}

	h.Tmpl.ExecuteTemplate(w, "index", TemplateData{
		Email: user.Email,
		Tasks: taskList.Tasks,
	})
}

func (h *Handler) LoginForm(w http.ResponseWriter, r *http.Request) {
	h.Tmpl.ExecuteTemplate(w, "login", TemplateData{})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	sess, err := h.Auth.Login(
		context.Background(),
		&models.UserWithPassword{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		},
	)
	if err != nil && err.Error() == "rpc error: code = PermissionDenied desc = invalid credentials" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if isError(err, w) {
		return
	}

	cookie := http.Cookie{
		Name:  "token",
		Value: sess.SessionId,
	}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/", http.StatusFound)
}

func isError(err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return true
	}
	return false
}

func (h *Handler) NewTask(w http.ResponseWriter, r *http.Request) {
	user, err := h.checkSession(w, r)
	if isError(err, w) {
		return
	}

	id = id+1
	h.Tasks.Add(
		context.Background(),
		&models.Task{
			Id:                   int32(id),
			Email:                user.Email,
			Title:                r.FormValue("title"),
			Done:                 false,
		},
	)

	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) EditTask(w http.ResponseWriter, r *http.Request) {
	user, err := h.checkSession(w, r)
	if isError(err, w) {
		return
	}

	id, err := strconv.Atoi(strings.Split(r.URL.Path[1:], "/")[1])
	if isError(err, w) {
		return
	}
	_, err = h.Tasks.Update(
		context.Background(),
		&models.Task{
			Id:                   int32(id),
			Email:                user.Email,
			Done:                 true,
		},
	)
	if isError(err, w) {
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) checkSession(w http.ResponseWriter, r *http.Request) (*models.User, error) {
	cookieSessionID, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return nil, nil
	}

	user, err := h.Auth.Check(
		context.Background(),
		&models.Session{
			SessionId: cookieSessionID.Value,
		},
	)
	if err != nil && err.Error() == "rpc error: code = PermissionDenied desc = invalid credentials" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func main() {
	tmpl := template.Must(template.ParseFiles("main_service/index.html", "main_service/login.html"))
	h := &Handler{
		Tmpl: tmpl,
	}

	grcpConnAuth, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConnAuth.Close()
	h.Auth = models.NewAuthClient(grcpConnAuth)

	grcpConnTasks, err := grpc.Dial(
		"127.0.0.1:8082",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConnAuth.Close()
	h.Tasks = models.NewTasksClient(grcpConnTasks)

	r := mux.NewRouter()
	r.HandleFunc("/", h.Root).Methods("GET")
	r.HandleFunc("/login", h.LoginForm).Methods("GET")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/tasks", h.NewTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", h.EditTask).Methods("POST")

	http.Handle("/", r)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
