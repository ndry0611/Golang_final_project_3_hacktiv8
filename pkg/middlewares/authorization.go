package middlewares

import (
	"final_project_3/repository/task_repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthorizationService interface {
	TaskUpdateAuthorization() gin.HandlerFunc
}

type authorizationService struct {
	taskRepo task_repository.Repository
}

func NewAuthorizationService(taskRepo task_repository.Repository) AuthorizationService {
	return &authorizationService{
		taskRepo: taskRepo,
	}
}

func AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(jwt.MapClaims)
		userRole := user["role"].(string)
		if userRole != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "permission to access endpoint denied",
				"error":   "NOT_AUTHORIZED",
			})
			return
		}
		c.Next()
	}
}

func (authorization *authorizationService) TaskUpdateAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(jwt.MapClaims)
		userId := uint(user["id"].(float64))
		userRole := user["role"].(string)
		if userRole == "admin" {
			c.Next()
		} else {
			param := c.Param("taskId")
			taskId, errConv := strconv.Atoi(param)
			if errConv != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": errConv.Error(),
				})
				return
			}
			task, err := authorization.taskRepo.FindTaskById(taskId)
			if err != nil {
				c.AbortWithStatusJSON(err.Status(), err)
				return
			}
			if task.UserID != userId {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"error": "you are forbidden to modify this task",
				})
				return
			}
			c.Next()
		}
	}
}