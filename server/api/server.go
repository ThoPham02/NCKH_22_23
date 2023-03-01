package api

import (
	"database/sql"
	"fmt"
	"github/ThoPham02/research_management/api/token"
	"github/ThoPham02/research_management/config"
	"github/ThoPham02/research_management/db/store"
	"github/ThoPham02/research_management/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Config     *config.Config
	Store      store.Store
	TokenMaker token.Maker
	Router     *gin.Engine
}

func NewServer(store store.Store) *Server {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load config::", err)
	}
	tokenMaker, err := token.NewPasetoMaker(config.TokenSemmetricKey)
	if err != nil {
		log.Fatal("Can't create a new token::", err)
	}
	router := RegisterRouter(store, gin.Default())

	server := &Server{
		Config:     config,
		Store:      store,
		TokenMaker: tokenMaker,
		Router:     router,
	}
	server.Router.POST("/user/login", server.UserRegister())

	return server
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type NewUserResponse struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Permission int64  `json:"permission"`
}

type LoginUserResponse struct {
	AccessToken string          `json:"access_token"`
	User        NewUserResponse `json:"user"`
}

func (server *Server) UserRegister() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req LoginUserRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ErrRequest(err))
			return
		}
		user, err := server.Store.GetUserByName(ctx, req.Username)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, utils.ErrResposive(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, utils.ErrResposive(err))
			return
		}

		isUser := utils.ComparePassword(req.Password, user.Password)
		if !isUser {
			ctx.JSON(http.StatusBadRequest, utils.ErrResposive(fmt.Errorf("%s is wrong password", req.Password)))
			return
		}

		userResponse := NewUserResponse{
			Username:   user.Username,
			Email:      user.Email,
			Permission: user.Permission,
		}

		res := LoginUserResponse{
			User: userResponse,
		}
		ctx.JSON(http.StatusOK, res)
	}
}
