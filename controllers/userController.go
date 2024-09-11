package controllers

import (
	"context"
	"fmt"
	"io/ioutil"

	"net/http"

	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kalyanKumarPokkula/Go-jwt/helpers"
	"github.com/kalyanKumarPokkula/Go-jwt/initializers"
	"github.com/kalyanKumarPokkula/Go-jwt/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	// getting username/email/password form a user request body
	var body helpers.Signup_Body

	fmt.Println("inside the signup function")

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to read body",
			"success": "false",
		})

		return
	}

	println(body.Email)
	//Hash the password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to hash password",
			"success": "false",
		})

		return
	}
	user := models.User{UserName: body.UserName, Email: body.Email, Password: string(hashPassword)}
	result := initializers.DB.Create(&user)
	fmt.Println(*result)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to create a user",
			"success": "false",
		})

		return
	}
	// Response
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully created a user",
		"success": "true",
	})

}

func Login(c *gin.Context) {

	var body helpers.Login

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to read body",
			"success": "false",
		})

		return
	}
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invaild email",
			"success": "false",
		})

		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invaild password",
			"success": "false",
		})

		return
	}
	fmt.Println(user)
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	// secretKey, err := generateSecretKey(32)
	// fmt.Println(secretKey)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRECT")))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to create a token",
			"success": "false",
		})

		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully login a user",
		"token":   tokenString,
		"user":    user.UserName,
		"user_id" : user.ID,
		"userEmail" : user.Email,
	})
}

// func generateSecretKey(length int) (string, error) {
//     // Generate random bytes
//     secret := make([]byte, length)
//     if _, err := rand.Read(secret); err != nil {
//         return "", err
//     }

//     // Encode bytes as hexadecimal string
//     secretKey := hex.EncodeToString(secret)

//     return secretKey, nil
// }

func Vaildate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func GoogleLogin(c *gin.Context) {
	url := initializers.AppConfig.GoogleLoginConfig.AuthCodeURL("random")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != "random" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid oauth state"})
		return
	}
	googlecon := initializers.GoogleConfig()
	code := c.Query("code")
	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Code-Token Exchange Failed"})
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User Data Fetch Failed"})
		return
	}

	UserData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User Data Fetch Failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": UserData})

}

// func VerifyEmail(c *gin.Context){

// }
