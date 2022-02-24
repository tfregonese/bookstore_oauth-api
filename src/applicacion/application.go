package applicacion

import (
	"github.com/gin-gonic/gin"
	"github.com/tfregonese/bookstore_oauth-api/src/domain/access_token"
	"github.com/tfregonese/bookstore_oauth-api/src/http"
	"github.com/tfregonese/bookstore_oauth-api/src/repository/db"
)

var router = gin.Default()

func Start() {
	/*
		session, dbErr := cassandra.GetSession()
		if dbErr != nil {
			panic(dbErr)
		}
		fmt.Println("Cassandra OK!")
		session.Close()
	*/
	/* Clean Arch
	Los casos de uno no pueden depender de las capas externas,
	por lo que se les pasa como parametro los servicios que necesita para funcionar.
	*/
	dbRepository := db.New()
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
