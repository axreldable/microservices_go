package main

import (
	"context"
	"golang-2018-1/7/99_hw/models"
	"net"
	"fmt"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"sync"
)

const sessKeyLen = 10

type server struct {
	*sync.Mutex
	Users    map[string]*models.UserWithPassword
	Sessions map[string]*models.User
}

func NewServer() *server {
	return &server{
		Mutex: &sync.Mutex{},
		Users:    make(map[string]*models.UserWithPassword),
		Sessions: make(map[string]*models.User),
	}
}

func (s *server) Check(ctx context.Context, req *models.Session) (*models.User, error) {
	fmt.Println("Check")
	s.Lock()
	inStoreUser, isExist := s.Sessions[req.SessionId]
	s.Unlock()
	if isExist {
		return inStoreUser, nil
	}
	return &models.User{}, grpc.Errorf(codes.PermissionDenied, "invalid credentials")
}

func (s *server) Login(ctx context.Context, req *models.UserWithPassword) (*models.Session, error) {
	fmt.Println("Login")
	s.Lock()
	inStoreUser, isExist := s.Users[req.Email]
	s.Unlock()
	if isExist && inStoreUser.Password == req.Password {
		sessId := RandStringRunes(sessKeyLen)
		s.Lock()
		s.Sessions[sessId] = &models.User{
			Email: req.Email,
		}
		s.Unlock()
		return &models.Session{
			SessionId: sessId,
		}, nil
	} else if isExist {
		return &models.Session{}, grpc.Errorf(codes.PermissionDenied, "invalid credentials")
	}
	s.Lock()
	s.Users[req.Email] = req
	s.Unlock()

	sessId := RandStringRunes(sessKeyLen)

	s.Lock()
	s.Sessions[sessId] = &models.User{
		Email: req.Email,
	}
	s.Unlock()

	return &models.Session{
		SessionId: sessId,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("cant listet port", err)
	}

	server := grpc.NewServer()

	models.RegisterAuthServer(server, NewServer())

	fmt.Println("starting server at :8081")
	server.Serve(lis)
}
