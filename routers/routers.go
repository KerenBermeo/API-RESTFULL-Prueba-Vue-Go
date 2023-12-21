package routers

import (
	c_proyect "trucode2/apitaskmanager/controllers/project"
	c_task "trucode2/apitaskmanager/controllers/task"
	c_user "trucode2/apitaskmanager/controllers/user"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and configures the Gin router for the application.
// It includes the setup of routes for users, projects, and tasks.
func SetupRouter() *gin.Engine {
	// Create a new Gin router with default middleware.
	router := gin.Default()

	// Set up routes for users, projects, and tasks.
	UserRoutes(router)
	ProjectRoutes(router)
	TaskRoutes(router)

	// Return the configured router.
	return router
}

// UserRoutes configures the user-related routes on the provided Gin router.
func UserRoutes(router *gin.Engine) {
	// Create a route group for user-related endpoints.
	userRoutes := router.Group("/users")
	{
		// Define routes for creating, retrieving, updating, and deleting users.
		userRoutes.POST("", c_user.CreateUser)
		userRoutes.GET("", c_user.GetAllUsers)
		userRoutes.GET("/:user_id", c_user.GetUser)
		userRoutes.PUT("/:user_id", c_user.UpdateUser)
		userRoutes.DELETE("/:user_id", c_user.DeleteUser)
	}
}

// ProjectRoutes configures the project-related routes on the provided Gin router.
func ProjectRoutes(router *gin.Engine) {
	// Create a route group for project-related endpoints with user-specific path parameter.
	projectRoutes := router.Group("/users/:user_id/project")
	{
		// Define routes for creating, retrieving, updating, and deleting projects.
		projectRoutes.POST("", c_proyect.CreateProject)
		projectRoutes.GET("", c_proyect.GetAllProjects)
		projectRoutes.GET("/:project_id", c_proyect.GetProject)
		projectRoutes.PUT("/:project_id", c_proyect.UpdateProject)
		projectRoutes.DELETE("/:project_id", c_proyect.DeleteProject)
	}
}

// TaskRoutes configures the task-related routes on the provided Gin router.
func TaskRoutes(router *gin.Engine) {
	// Create a route group for task-related endpoints with user and project-specific path parameters.
	taskRoutes := router.Group("/users/:user_id/project/:project_id/tasks")
	{
		// Define routes for creating, retrieving, updating, and deleting tasks.
		taskRoutes.POST("", c_task.CreateTask)
		taskRoutes.GET("", c_task.GetAllTasks)
		taskRoutes.GET("/:task_id", c_task.GetTask)
		taskRoutes.PUT("/:task_id", c_task.UpdateTask)
		taskRoutes.DELETE("/:task_id", c_task.DeleteTask)
	}
}
