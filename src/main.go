package main

import (
	"log"
	"movie-reservation-system/src/application/usecases/admin"
	"movie-reservation-system/src/application/usecases/auth"
	"movie-reservation-system/src/application/usecases/role"
	"movie-reservation-system/src/domain/entities"
	"movie-reservation-system/src/domain/settings"
	"movie-reservation-system/src/domain/validators"
	"movie-reservation-system/src/infrastructure/database"
	"movie-reservation-system/src/infrastructure/middlewares"
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
	roleRepository := repositories.NewRoleRepositoryGorm(connection)

	validate := validator.New()
	validate.RegisterValidation("password", validators.ValidatePassword)
	validate.RegisterValidation("role", validators.ValidateRole(settings.RegexRole))

	registerUser := auth.NewRegisterUserUseCase(userRepository, validate)
	loginUser := auth.NewLoginUserUseCase(userRepository, settings, validate)

	loginAdmin := admin.NewLoginAdminUseCase(userRepository, settings, validate)

	getRoles := role.NewGetRolesUseCase(roleRepository)
	getRole := role.NewGetRoleUseCase(roleRepository)
	saveRole := role.NewSaveRoleUseCase(roleRepository, validate)
	updateRole := role.NewUpdateRoleUseCase(roleRepository, validate)
	deleteRole := role.NewDeleteRoleUseCase(roleRepository)

	middlewares := middlewares.NewMiddlewares(settings)

	router := routes.NewRouter(registerUser, loginUser, loginAdmin,
		getRoles, getRole, saveRole, updateRole, deleteRole, middlewares)

	engine := gin.Default()
	router.RegisterRoutes(engine)

	if err := engine.Run(settings.AppPort); err != nil {
		log.Fatal(err)
	}
}
