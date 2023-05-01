package handlers

import (
	"AuthApp/config"
	"AuthApp/httperror"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func validateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, httperror.Unauthorized()
		}

		return config.Config.JWTSecret, nil
	})
}

func (h *Handler) ValidateHandler(c *gin.Context) {
	// 	fmt.Println("pass-----------------")
	// 	tokenString := c.GetHeader("Authorization")[7:]
	// 	fmt.Println(tokenString)
	// 	token, err := jwt.ParseWithClaims(tokenString, &dto.IdTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 		// validate the signing method
	// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	// 		}
	// 		// return the secret key used to sign the token
	// 		return []byte("SECRET_KEY"), nil
	// 	})

	//	if claims, ok := token.Claims.(*dto.IdTokenClaims); ok && token.Valid {
	// c.JSON(http.StatusOK, gin.H{
	// 	"iss":  claims.Issuer,
	// 	"exp":  claims.ExpiresAt,
	// 	"iat":  claims.IssuedAt,
	// 	"user": claims.User,
	// })
	//	} else {
	//
	//		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	//	}
	authHeader := c.GetHeader("Authorization")

	s := strings.Split(authHeader, "Bearer ")
	unauthorizedErr := httperror.Unauthorized()
	if len(s) < 2 {
		c.AbortWithStatusJSON(unauthorizedErr.StatusCode, unauthorizedErr)
		return
	}
	encodeToken := s[1]

	token, err := validateToken(encodeToken)
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(unauthorizedErr.StatusCode, unauthorizedErr)
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(unauthorizedErr.StatusCode, unauthorizedErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"claims": claims})
}
