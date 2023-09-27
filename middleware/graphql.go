package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/huantt/go-graphql-sample/graphql/loader"
)

// GraphQLHandler is a graphql HTTP handler
type graphQLHandler struct {
	handler *relay.Handler
	loaders loader.Map
}

// NewHandler returns a new http.Handler
func NewHandler(schema *graphql.Schema, loaders loader.Map) gin.HandlerFunc {
	h := &graphQLHandler{
		handler: &relay.Handler{Schema: schema},
		loaders: loaders,
	}
	return h.ServeHTTP
}

// ServeHTTP serves http requests
func (h *graphQLHandler) ServeHTTP(c *gin.Context) {
	// Attach loaders to the request's context
	attachedLoaderCtx := h.loaders.Attach(c)
	request := c.Request.WithContext(attachedLoaderCtx)
	h.handler.ServeHTTP(c.Writer, request)
}
