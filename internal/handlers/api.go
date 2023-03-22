package handlers

import (
	"server/internal/db"
)

type HandlerService struct {
	DBService *db.DBService
}

func NewHandlerService(dbService *db.DBService) *HandlerService {
	return &HandlerService{dbService}
}
