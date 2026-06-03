package middlewares

import (
	"movie-reservation-system/src/application/usecases/auth/token"
	"movie-reservation-system/src/domain/settings"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middlewares struct {
	settings *settings.Settings
}

func NewMiddlewares(settings *settings.Settings) *Middlewares {
	return &Middlewares{settings}
}

func (middlewares *Middlewares) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerAuth := ctx.GetHeader("Authorization")

		if headerAuth == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "header Authorization requerido"})
			return
		}

		if !strings.HasPrefix(headerAuth, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "autentication de tipo Bearer requerida"})
			return
		}

		parts := strings.Split(headerAuth, " ")

		if len(parts) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "formato invalido para autentication Bearer <token>"})
			return
		}

		tokenString := parts[1]

		claims, err := token.VerifyToken(middlewares.settings.Key, tokenString)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("userID", claims.Subject)
		ctx.Set("userRole", claims.Role)

		ctx.Next()
	}
}

func (middlewares *Middlewares) Authorization(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole, exists := ctx.Get("userRole")

		if !exists {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "no se pudo determinar el rol"})
			return
		}

		for _, role := range roles {
			if userRole == role {
				ctx.Next()
				return
			}
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "permisos insuficientes"})
	}
}
