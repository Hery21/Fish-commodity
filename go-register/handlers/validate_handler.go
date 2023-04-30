package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) ValidateHandler(c *gin.Context) {
	// tokenString := c.GetHeader("Authorization")[7:]
	// token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	// validate the signing method
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	// 	}
	// 	// return the secret key used to sign the token
	// 	return []byte("SECRET_KEY"), nil
	// })

	// if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"name":  claims.Name,
	// 		"phone": claims.Phone,
	// 		"role":  claims.Role,
	// 	})
	// } else {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	// }
}
