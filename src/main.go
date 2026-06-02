package main

import (
	"log"
	"movie-reservation-system/src/application/usecases/auth"
	"movie-reservation-system/src/domain/entities"
	"movie-reservation-system/src/domain/settings"
	"movie-reservation-system/src/domain/validators"
	"movie-reservation-system/src/infrastructure/database"
	"movie-reservation-system/src/infrastructure/repositories"
	"movie-reservation-system/src/infrastructure/routes"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/sqlite"
)

func main() {
	settings, err := settings.NewSettings()

	if err != nil {
		log.Fatal(err)
	}

	dialector := sqlite.Open(settings.DbUri)
	connection, err := database.NewConnection(dialector)

	if err != nil {
		log.Fatal(err)
	}

	db := connection.GetDB()
	db.AutoMigrate(&entities.User{}, &entities.Role{})

	userRepository := repositories.NewUserRepositoryGorm(connection)

	validate := validator.New()
	validate.RegisterValidation("password", validators.ValidatePassword)

	registerUser := auth.NewRegisterUserUseCase(userRepository, validate)

	router := routes.NewRouter(registerUser)
	engine := gin.Default()
	router.RegisterRoutes(engine)

	if err := engine.Run(settings.AppPort); err != nil {
		log.Fatal(err)
	}
}
