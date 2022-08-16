package service

import (
	"github/com/sulkifli27/golang_api/dto"
	"github/com/sulkifli27/golang_api/model"
	"github/com/sulkifli27/golang_api/repository"
	"log"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) model.User
    FindByEmail(email string) model.User
    IsDuplicateEmail(email string) bool
}

type authService struct {
    userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
    return & authService{
        userRepository: userRep,
    }
}

func (service *authService) VerifyCredential(email string, password string) interface{}{
    res := service.userRepository.VerifyCredential(email, password)
    if value, ok := res.(model.User) ; ok {
        comparedPassword := ComparePassword(value.Password, []byte(password))
        if value.Email == email && comparedPassword {
            return res
        }
        return false
    }
    return false
}

func (service authService) CreateUser(user dto.RegisterDTO) model.User {
    userToCreate := model.User{}
    err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
    if err != nil {
        log.Fatalf("Failed map %v", err)
    }
    res := service.userRepository.InsertUser(userToCreate)
    return res
}

func (service *authService) FindByEmail(email string) model.User {
    return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
    res := service.userRepository.IsDuplicateEmail(email)
    return !(res.Error == nil)
}

func ComparePassword(hashedPwd string, plainPassword []byte) bool {
    byteHash := []byte(hashedPwd) 
    err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
    if err != nil {
        log.Println(err)
        return false
    }
    return true
}
