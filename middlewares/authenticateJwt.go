package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kalyanKumarPokkula/Go-jwt/initializers"
	"github.com/kalyanKumarPokkula/Go-jwt/models"
)

func AuthenticateJwt(c *gin.Context){
	//get the cookie off request

	tokenString, err := c.Cookie("Authorization")
	// tokenString:= c.GetHeader("Authorization")


	
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// println(tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
	
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRECT")), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		initializers.DB.First(&user , claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}


		c.Set("user" , user)

		c.Next()
		
	} else {
		fmt.Println(err)
	}
}