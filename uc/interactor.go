package uc

import (
	"github.com/err0r500/go-realworld-clean/domain"
)

// NewHandler : the interactor constructor, use this in order to avoid null pointers at runtime
//func NewHandler(Logger Logger, uRW UserRW, arw ArticleRW, validator UserValidator, handler AuthHandler) Handler {
//	return interactor{
//		Logger:        Logger,
//		UserRW:        uRW,
//		ArticleRW:     arw,
//		UserValidator: validator,
//		AuthHandler:   handler,
//	}
//}

// interactor : the struct that will have as properties all the IMPLEMENTED interfaces
// in order to provide them to its methods : the use cases
type interactor struct {
	logger        Logger
	userRW        UserRW
	articleRW     ArticleRW
	userValidator UserValidator
	authHandler   AuthHandler
}

// Logger : only used to log stuff
type Logger interface {
	Log(...interface{})
}

type AuthHandler interface {
	GenUserToken(userName string) (token string, err error)
	GetUserName(token string) (userName string, err error)
}

type UserRW interface {
	Create(username, email, password string) (*domain.User, error)
	GetByName(userName string) (*domain.User, error)
	GetByEmailAndPassword(email, password string) (*domain.User, error)
	Save(user domain.User) error
}

type ArticleRW interface {
	GetByAuthorsNameOrderedByMostRecentAsc(usernames []string) ([]domain.Article, error)
	GetRecentFiltered(filters Filters) ([]domain.Article, error)
}

type UserValidator interface {
	CheckUser(user domain.User) error
}
