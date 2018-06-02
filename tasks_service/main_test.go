package main

import (
	"context"
	"golang-2018-1/7/99_hw/models"
	"reflect"
	"testing"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc"
	"sync"
)

func TestAddTasks(t *testing.T) {
	s := &server{
		Mutex: &sync.Mutex{},
		Tasks: make(map[string]map[int]*models.Task),
	}
	ctx := context.Background()

	task := &models.Task{
		Email: "test@mail.ru",
		Title: "test task",
	}

	addedTask, err := s.Add(ctx, task)

	if err != nil {
		t.Error("add task error")
	}

	expectedList := &models.TasksList{
		Tasks: []*models.Task{
			addedTask,
		},
	}

	list, err := s.List(ctx, &models.User{
		Email: "test@mail.ru",
	})

	if err != nil {
		t.Error("list task error")
	}

	if !reflect.DeepEqual(list, expectedList) {
		t.Errorf("Expected: %#v \nGot: %#v", expectedList.Tasks, list.Tasks)
	}

}

func TestUpdateTasks(t *testing.T) {
	s := &server{
		Mutex: &sync.Mutex{},
		Tasks: map[string]map[int]*models.Task{
			"test@mail.ru": map[int]*models.Task{
				1: &models.Task{
					Id:    1,
					Email: "test@mail.ru",
					Title: "updated task",
					Done:  true,
				},
			},
		},
	}
	ctx := context.Background()
	updatedTask := &models.Task{
		Id:    1,
		Email: "test@mail.ru",
		Title: "updated task",
		Done:  true,
	}

	_, err := s.Update(ctx, updatedTask)

	if err != nil {
		t.Error("update task error")
	}

	list, err := s.List(ctx, &models.User{
		Email: "test@mail.ru",
	})

	expectedList := &models.TasksList{
		Tasks: []*models.Task{
			updatedTask,
		},
	}

	if err != nil {
		t.Error("list task error")
	}

	if !reflect.DeepEqual(list, expectedList) {
		t.Errorf("Expected: %#v \nGot: %#v", expectedList, list)
	}
}

func TestUpdateTasksErr(t *testing.T) {
	s := &server{
		Mutex: &sync.Mutex{},
		Tasks: make(map[string]map[int]*models.Task),
	}
	ctx := context.Background()
	updatedTask := &models.Task{
		Id:    700,
		Email: "test@mail.ru",
		Title: "updated task",
		Done:  true,
	}

	expectedErr := grpc.Errorf(codes.NotFound, "no task found")

	_, err := s.Update(ctx, updatedTask)

	if !reflect.DeepEqual(err, expectedErr) {
		t.Errorf("Expected error: %#v\nGot: %#v", expectedErr, err)
	}

}
