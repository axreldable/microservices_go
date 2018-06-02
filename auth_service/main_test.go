package main

import (
	"context"
	"golang-2018-1/7/99_hw/models"
	"reflect"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"sync"
)

func TestAuth(t *testing.T) {

	// Поменял логику логирования
	// Т.к. у нас нет авторизации, при первом логировании будем запоминать данные (email, пароль),
	// а при последующих подключениях с таким же email проверять пароль.

	cases := []*struct {
		Req *models.UserWithPassword
		Res *models.Session
		Err error
	}{
		{
			Req: &models.UserWithPassword{
				Email:    "test@mail.ru",
				Password: "1234",
			},
			Res: &models.Session{
				SessionId: "XVlBzgbaiC",
			},
			Err: nil,
		},
		// подключение с неправильным паролем
		{
			Req: &models.UserWithPassword{
				Email:    "test@mail.ru",
				Password: "12345",
			},
			Res: &models.Session{},
			Err: grpc.Errorf(codes.PermissionDenied, "invalid credentials"),
		},
		// второе подключение с правильным паролем
		{
			Req: &models.UserWithPassword{
				Email:    "test@mail.ru",
				Password: "1234",
			},
			Res: &models.Session{
				SessionId: "MRAjWwhTHc", // уже другая сессия
			},
			Err: nil,
		},
	}

	s := &server{
		Mutex: &sync.Mutex{},
		Users: map[string]*models.UserWithPassword{
			"test@mail.ru": &models.UserWithPassword{
				Email:    "test@mail.ru",
				Password: "1234",
			},
		},
		Sessions: make(map[string]*models.User),
	}

	for _, c := range cases {
		resp, err := s.Login(context.Background(), c.Req)

		if !reflect.DeepEqual(resp, c.Res) {
			t.Errorf("Expected: %#v\nGot: %#v", c.Res, resp)
		}

		if !reflect.DeepEqual(err, c.Err) {
			t.Errorf("Expected error: %#v\nGot: %#v", c.Err, err)
		}
	}
}

func TestCheck(t *testing.T) {
	cases := []*struct {
		Req *models.Session
		Res *models.User
		Err error
	}{
		{
			Req: &models.Session{
				SessionId: "XVlBzgbaiC",
			},
			Res: &models.User{
				Email: "test@mail.ru",
			},
			Err: nil,
		},
		{
			Req: &models.Session{
				SessionId: "1234",
			},
			Res: &models.User{},
			Err: grpc.Errorf(codes.PermissionDenied, "invalid credentials"),
		},
	}

	s := &server{
		Mutex: &sync.Mutex{},
		Users: make(map[string]*models.UserWithPassword),
		Sessions: map[string]*models.User{
			"XVlBzgbaiC": &models.User{
				Email: "test@mail.ru",
			},
		},
	}

	for _, c := range cases {
		resp, err := s.Check(context.Background(), c.Req)

		if !reflect.DeepEqual(resp, c.Res) {
			t.Errorf("Expected: %#v\nGot: %#v", c.Res, resp)
		}

		if !reflect.DeepEqual(err, c.Err) {
			t.Errorf("Expected error: %#v\nGot: %#v", c.Err, err)
		}

	}

}
