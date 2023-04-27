package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/makovii/group_organiser/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"github.com/makovii/group_organiser/config"
)


type AuthController struct {
	DB *gorm.DB
	CFG *config.Config
}

type CreateBody struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}

func NewAuthController(db *gorm.DB, cfg *config.Config) *AuthController {
	return &AuthController{DB: db, CFG: cfg}
}

func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateJWT(email, name string, id uint, secret string) (string, error) {
	var mySigningKey = []byte(secret)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["name"] = name
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour + 24).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokenString, nil
}

func CheckPasswordHash(password, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}

func CreateAuthManagerRequest(a AuthController, idManager uint) {
	request := database.Request{
		From: idManager,
		To: uint(a.CFG.Admin.Id),
		StatusId: uint(a.CFG.Status.WaitId),
		TypeId: uint(a.CFG.Type.RegistrationId),
	}

	// обработать ошибку на неудачное создание
	a.DB.Create(&request)
}

func (a AuthController) SignUp(c *gin.Context){
	var user database.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dbUser database.User
	a.DB.Where("email = ?", user.Email).First(&dbUser)

	if dbUser.Email != "" {
		var err Error
		err = SetError(err, "Email already in use")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	} 
	
	var err error
	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		fmt.Println("Error in password hash", err)
	}

	if user.Role == a.CFG.Role.ManagerId {
		user.Ban = true
	}

	a.DB.Create(&user)

	if user.Role == a.CFG.Role.ManagerId {
		CreateAuthManagerRequest(a, user.Id)
	}

	c.JSON(http.StatusOK, user)
}

func (a AuthController) SignIn(c *gin.Context){
	var authdetails database.Authentication
	
	if err := c.BindJSON(&authdetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var authUser database.User
	a.DB.Where("email = ?", authdetails.Email).First(&authUser)
	if authUser.Email == "" {
		var err Error
		err = SetError(err, "Username or Password is incorrect")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if authUser.Ban && authUser.Role == a.CFG.Role.ManagerId {
		var err Error
		err = SetError(err, "The administrator has not yet allowed you to log in")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	
	if authUser.Ban && authUser.Role == a.CFG.Role.PlayerId {
		var err Error
		err = SetError(err, "You was banned")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var check bool
	if authUser.Role != a.CFG.Role.AdminId {
		check = CheckPasswordHash(authdetails.Password, authUser.Password)		
	} else {
		check = authUser.Password == authdetails.Password
	}


	if !check {
		var err Error
		err = SetError(err, "Username or Password is incorrect")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	validToken, err := GenerateJWT(authUser.Email, authUser.Name, authUser.Id, a.CFG.Secrets.Secret)
	if err != nil {
		var err Error
		err = SetError(err, "Failed to generate token")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var token database.Token
	token.Email = authUser.Email
	token.Role = authUser.Role
	token.TokenString = validToken
	c.JSON(http.StatusOK, token)
}