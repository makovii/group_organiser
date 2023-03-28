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

func (a AuthController) SignUp(c *gin.Context){
	var player database.Player

	if err := c.BindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dbplayer database.Player
	a.DB.Where("email = ?", player.Email).First(&dbplayer)

	if dbplayer.Email != "" {
		var err Error
		err = SetError(err, "Email already in use")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	} 
	var err error
	player.Password, err = GeneratehashPassword(player.Password)
	if err != nil {
		fmt.Println("Error in password hash", err)
	}

	a.DB.Create(&player)
	c.JSON(http.StatusOK, player)
}

func (a AuthController) SignIn(c *gin.Context){
	var authdetails database.Authentication
	
	if err := c.BindJSON(&authdetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var authPlayer database.Player
	a.DB.Where("email = ?", authdetails.Email).First(&authPlayer)
	if authPlayer.Email == "" {
		var err Error
		err = SetError(err, "Username or Password is incorrect")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	check := CheckPasswordHash(authdetails.Password, authPlayer.Password)

	if !check {
		var err Error
		err = SetError(err, "Username or Password is incorrect")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	validToken, err := GenerateJWT(authPlayer.Email, authPlayer.Name, authPlayer.Id, a.CFG.Secrets.Secret)
	if err != nil {
		var err Error
		err = SetError(err, "Failed to generate token")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var token database.Token
	token.Email = authPlayer.Email
	token.TokenString = validToken
	c.JSON(http.StatusOK, token)
}