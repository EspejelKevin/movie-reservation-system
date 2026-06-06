package main

import (
	"log"
	"movie-reservation-system/src/application/usecases/admin"
	"movie-reservation-system/src/application/usecases/auth"
	"movie-reservation-system/src/application/usecases/genre"
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
	db.AutoMigrate(&entities.User{}, &entities.Role{},
		&entities.Genre{}, &entities.Movie{}, &entities.Seat{},
		&entities.ShowTime{}, &entities.Reservation{})

	userRepository := repositories.NewUserRepositoryGorm(connection)
	roleRepository := repositories.NewRoleRepositoryGorm(connection)
	genreRepository := repositories.NewGenreRepositoryGorm(connection)

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

	getGenres := genre.NewGetGenresUseCase(genreRepository)
	getGenre := genre.NewGetGenreUseCase(genreRepository)
	saveGenre := genre.NewSaveGenreUseCase(genreRepository, validate)
	updateGenre := genre.NewUpdateGenreUseCase(genreRepository, validate)
	deleteGenre := genre.NewDeleteGenreUseCase(genreRepository)

	middlewares := middlewares.NewMiddlewares(settings)

	routerAdmin := routes.NewRouterAdmin(loginAdmin)
	routerUser := routes.NewRouterUser(registerUser, loginUser)
	routerRole := routes.NewRouterRole(getRoles, getRole, saveRole, updateRole, deleteRole, middlewares)
	routerGenre := routes.NewRouterGenre(getGenres, getGenre, saveGenre, updateGenre, deleteGenre, middlewares)

	engine := gin.Default()
	routerAdmin.RegisterRoutesAdmin(engine)
	routerUser.RegisterRoutesUser(engine)
	routerRole.RegisterRoutesRole(engine)
	routerGenre.RegisterRoutesGenre(engine)

	if err := engine.Run(settings.AppPort); err != nil {
		log.Fatal(err)
	}
}
