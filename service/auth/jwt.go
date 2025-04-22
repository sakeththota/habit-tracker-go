package auth

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sakeththota/habit-tracker-go/config"
	"github.com/sakeththota/habit-tracker-go/types"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateOwnership(handlerFunc gin.HandlerFunc, store types.HabitStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get userID and habitID from context / request
		userID := GetUserIDFromContext(c)
		habitID, err := strconv.Atoi(c.Param("id"))

		fmt.Printf("Validating User: %d with Habit: %d\n", userID, habitID)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("invalid habit id: %v", err)})
			return
		}

		// check if habit with habitID has a matching user ID for ownership
		habit, err := store.GetHabitById(habitID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("failed to get habit by id: %v", err)})
			return
		}
		if habit.UserID != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": fmt.Errorf("permission denied, failed to get progress of habit with id: %v", err)})
			return
		}

		handlerFunc(c)
	}
}

func WithJWTAuth(handlerFunc gin.HandlerFunc, store types.UserStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token from user request
		tokenString := getTokenFromRequest(c)

		// validate jwt
		token, err := validateToken(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			c.JSON(http.StatusForbidden, gin.H{"error": fmt.Errorf("permission denied")})
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			c.JSON(http.StatusForbidden, gin.H{"error": fmt.Errorf("permission denied")})
			return
		}

		// if is we need to fetch the userId from the DB (id from the token)
		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)

		userID, err := strconv.Atoi(str)
		if err != nil {
			log.Printf("failed to convert userID to int: %v", err)
			c.JSON(http.StatusForbidden, gin.H{"error": fmt.Errorf("permission denied")})
			return
		}

		u, err := store.GetUserById(userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			c.JSON(http.StatusForbidden, gin.H{"error": fmt.Errorf("permission denied")})
			return
		}

		fmt.Printf("logged in user with id: %d\n", u.ID)

		// set context "userId"
		c.Set("userID", u.ID)

		handlerFunc(c)
	}
}

func getTokenFromRequest(c *gin.Context) string {
	tokenAuth := c.GetHeader("Authorization")
	if tokenAuth != "" {
		return tokenAuth
	}
	return ""
}

func validateToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}

func GetUserIDFromContext(c *gin.Context) int {
	userID, ok := c.Get("userID")
	if !ok {
		return -1
	}

	return userID.(int)
}
