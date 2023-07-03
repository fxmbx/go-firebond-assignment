package api

import (
	"go-firebond-assignment/config"
	db "go-firebond-assignment/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config config.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(config config.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}
	server.setUpRouter()
	return server, nil
}

func (server *Server) setUpRouter() {
	router := gin.Default()

	router.GET("/api", func(c *gin.Context) {
		c.String(200, "Welcome to firebomnd_assignment api")
	})
	router.GET("/api/rates/:cryptocurrency/:fiat", server.getCurrentCryptoToFiatRate)
	router.GET("/api/rates/:cryptocurrency/:fiat", server.getCurrentCryptoToFiatRate)
	router.GET("/api/rates/:cryptocurrency", server.getCryptoCurrencyRates)
	router.GET("/api/rates/history/:cryptocurrency/:fiat", server.getExchangeRateHistory)
	router.GET("/api/rates", server.getRates)

	router.GET("/api/balance/:address", server.getWalletBalance)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"name":    "Not Found",
			"message": "Page not found.",
			"code":    404,
			"status":  http.StatusNotFound,
		})
	})
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
