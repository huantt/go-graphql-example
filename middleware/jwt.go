package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/huantt/go-graphql-sample/pkg/jwt"
	"net/http"
	"strings"
)

const (
	tokenHeader  = "Authorization"
	bearerPrefix = "Bearer "
	authContext  = "m_auth_user"
)

type handler struct {
	JwtService *jwt.Service
}

type AuthInfo struct {
	Token         string
	UserID        int
	Email         string
	WalletAddress string
}

func InitAuthMdw(jwtService *jwt.Service) gin.HandlerFunc {
	return (&handler{JwtService: jwtService}).Authenticate
}

func (h *handler) Authenticate(c *gin.Context) {
	reqToken := c.GetHeader(tokenHeader)
	splitReqToken := strings.Split(reqToken, bearerPrefix)
	if len(splitReqToken) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token := splitReqToken[1]
	claims, err := h.JwtService.Verify(token)
	if err != nil || claims == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	authInfo := &AuthInfo{
		UserID:        claims.UserID,
		WalletAddress: strings.ToLower(claims.Address),
	}
	c.Set(authContext, authInfo)
	c.Next()
}

func GetAuthInfo(c *gin.Context) *AuthInfo {
	userInfo, existed := c.Get(authContext)
	if !existed {
		return nil
	}
	user, ok := userInfo.(*AuthInfo)
	if !ok {
		return nil
	}
	return user
}
