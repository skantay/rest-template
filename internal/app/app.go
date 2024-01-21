package app

import (
	"github.com/gin-gonic/gin"
	"github.com/skantay/web-1-clean/internal/controller"
	"github.com/skantay/web-1-clean/internal/core/server"
	"github.com/skantay/web-1-clean/internal/core/service"
	"github.com/skantay/web-1-clean/internal/infra/repository"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	instance := gin.New()
	instance.Use(gin.Recovery())

	db, err := repository.NewDB(
		config.DatabaseConfig{
			Driver:                  "mysql",
			Url:                     "user:password@tcp(127.0.0.1:3306)/your_database_name?charset=utf8mb4&parseTime=true&loc=UTC&tls=false&readTimeout=3s&writeTimeout=3s&timeout=3s&clientFoundRows=true",
			ConnMaxLifetimeInMinute: 3,
			MaxOpenConns:            10,
			MaxIdleConns:            1,
		},
	)
	if err != nil {
		log.Fatalf("failed to new database err=%s\n", err.Error())
	}

	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo)

	userController := controller.NewUserController(instance, userService)

	userController.InitRouter()

	httpServer := server.NewHttpServer(
		instance,
		config.HttpServerConfig{
			Port: 8000,
		},
	)

	log.Println("listening signals...")
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-c
	log.Println("graceful shutdown...")
}

}
