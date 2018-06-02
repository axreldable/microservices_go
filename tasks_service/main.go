package main

import (
	"context"
	"golang-2018-1/7/99_hw/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc"
	"net"
	"fmt"
	"log"
	"sync"
)

type server struct {
	*sync.Mutex
	Tasks map[string]map[int]*models.Task
}

func NewServer() *server {
	return &server{
		Mutex: &sync.Mutex{},
		Tasks: make(map[string]map[int]*models.Task),
	}
}

func (s *server) List(ctx context.Context, req *models.User) (*models.TasksList, error) {
	fmt.Println("List")
	s.Lock()
	tasks, isExist := s.Tasks[req.Email]
	s.Unlock()
	if isExist {
		var values []*models.Task
		for _, value := range tasks {
			values = append(values, value)
		}
		return &models.TasksList{
			Tasks: values,
		}, nil
	}
	return &models.TasksList{}, nil
}

func (s *server) Add(ctx context.Context, req *models.Task) (*models.Task, error) {
	fmt.Println("Add")
	s.Lock()
	userTasks, isExist := s.Tasks[req.Email]
	s.Unlock()
	if !isExist {
		taskMap := make(map[int]*models.Task)
		s.Lock()
		taskMap[int(req.Id)] = req
		s.Tasks[req.Email] = taskMap
		s.Unlock()
	} else {
		s.Lock()
		userTasks[int(req.Id)] = req
		s.Unlock()
	}
	return req, nil
}

func (s *server) Update(ctx context.Context, req *models.Task) (*models.Task, error) {
	fmt.Println("Update")
	s.Lock()
	userTasks, isExist := s.Tasks[req.Email]
	s.Unlock()
	if isExist {
		s.Lock()
		task, isExist := userTasks[int(req.Id)]
		s.Unlock()
		if isExist {
			task.Done = req.Done
			return task, nil
		}
	}
	return &models.Task{}, grpc.Errorf(codes.NotFound, "no task found")
}

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalln("cant listet port", err)
	}

	server := grpc.NewServer()

	models.RegisterTasksServer(server, NewServer())

	fmt.Println("starting server at :8082")
	server.Serve(lis)
}
