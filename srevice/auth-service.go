package srevice

import (
	"log"

	"github.com/Indra-riswan/go-rest-api/dto"
	"github.com/Indra-riswan/go-rest-api/entity"
	"github.com/Indra-riswan/go-rest-api/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredintial(email string, password string) interface{}
	CreateUser(user dto.RegisterDto) entity.User
	FindByEmail(email string) entity.User
	IsduplicatEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepo,
	}
}

func (service *authService) VerifyCredintial(email string, password string) interface{} {
	res := service.userRepository.VerifyCredintial(email, password)
	if v, ok := res.(entity.User); ok {
		comparepassword := ComparePassword(v.Password, []byte(password))
		if v.Email == email && comparepassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDto) entity.User {
	usercreate := entity.User{}
	err := smapping.FillStruct(&usercreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Erorr Create %v", err)

	}
	res := service.userRepository.InsertUser(usercreate)
	return res
}

func (service *authService) FindByEmail(email string) entity.User {
	return service.userRepository.FindByEmail(email)
}
func (service *authService) IsduplicatEmail(email string) bool {
	res := service.userRepository.IsduplicatEmail(email)

	return !(res.Error == nil)
}

func ComparePassword(hashedPwd string, plainPassword []byte) bool {
	bytehash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(bytehash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
