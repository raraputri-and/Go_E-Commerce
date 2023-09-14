package user

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Create(signupRequest SignupRequest) (User, error)
	Login(LoginRequest LoginRequest) (string, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(signupRequest SignupRequest) (User, error) {
	// HASH the Password
	hash, err := bcrypt.GenerateFromPassword([]byte(signupRequest.Password), 10)

	if err != nil {
		return User{}, err
	}

	// SAVE USER
	user := User{
		Email:    signupRequest.Email,
		Password: string(hash),
	}

	newUser, err := s.repository.Create(user)
	return newUser, err
}

func (s *service) Login(loginRequest LoginRequest) (string, error) {
	// GET USER
	user, err := s.repository.FindByEmail(loginRequest.Email)
	if err != nil {
		return "", err
	} else if user.ID == 0 {
		return "", errors.New("Invalid email or password")
	}

	// COMPARE PASSWORD
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))

	if err != nil {
		return "", err
	}

	// SIGN TOKEN
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
