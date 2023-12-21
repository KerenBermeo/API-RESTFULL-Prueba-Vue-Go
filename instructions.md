# RESTful API Task Manager

## Objective:

Create a web application using Go, Gorm, and Gin to manage data for Users, Projects, and Tasks. This exercise focuses on setting up routes, defining multiple models, and establishing relationships between them.

## Instructions:

### Setup:

Install Go, Gorm, and Gin on your machine.

Create a new Go project directory.

### Models:

Define three models: User, Project, and Task, with appropriate attributes for each.

#### Establish relationships between the models:

A User can be associated with multiple Projects.

A Project can have multiple Tasks.

Use Gorm to create the corresponding database tables and define associations.

### Router Setup:

Implement routes for the following actions using Gin:

- Create User: POST request to /users

- Get All Users: GET request to /users

- Get User by ID: GET request to /users/:id

- Create Project: POST request to /projects

- Get All Projects: GET request to /projects

- Get Project by ID: GET request to /projects/:id

- Create Task: POST request to /tasks

- Get All Tasks: GET request to /tasks

- Get Task by ID: GET request to /tasks/:id

### Controller:

Create controller functions for each route to handle the corresponding CRUD operation.

Use Gorm to perform database operations (create, read, update, delete) for each model.

### Documentation:

In a README file:

Provide simple documentation for your API, specifying the available endpoints and the expected request/response format.

Include any necessary setup instructions for users who want to run your application.

### Extra:

- Implement cascading deletion. When deleting a User, ensure that associated Projects and Tasks are also deleted.

- Enhance the API to allow updating a Task and associating it with a different Project.
