package api

import (
	db "back-end/db/sqlc"
	"back-end/token"
	"back-end/util"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err = v.RegisterValidation("currency", validCurrency)
		if err != nil {
			return nil, err
		}
	}
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)
	authRoutes.POST("/transfers", server.createTransfer)
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	server.router = router
}
