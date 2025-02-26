package routehandlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	mydb "github.com/klymenok/go-playground/db"
	"log"
	"net/http"
	"strconv"
)

// @BasePath /api/v1

// GetUser godoc
// @Summary  get users
// @Schemes
// @Description  get list of all users
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object}  mydb.User
// @Router       /users/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(chi.URLParam(r, "userId"))
	user, err := mydb.GetUserById(int64(userId))
	log.Println(err)
	if err != nil {
		w.WriteHeader(404)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

// @BasePath /api/v1

// GetUsers godoc
// @Summary  get users
// @Schemes
// @Description  get list of all users
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {array}  mydb.User
// @Router       /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("users"))
}

// @BasePath /api/v1

// CreateUser godoc
// @Summary  create user
// @Schemes
// @Description  create a new user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        first_name  body      string  true  "First name"
// @Param        last_name   body      string  true  "Last name"
// @Success      201         {object}  mydb.User
// @Router       /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// TODO add data validation
	var user mydb.User
	json.NewDecoder(r.Body).Decode(&user)
	user.Create()
	json.NewEncoder(w).Encode(user)
}

// @BasePath /api/v1

// UpdateUser godoc
// @Summary  update user
// @Schemes
// @Description  update an existing user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        first_name  body      string  false  "First name"
// @Param        last_name   body      string  false  "Last name"
// @Success      201         {object}  mydb.User
// @Router       /users [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// TODO add data validation
	userId, _ := strconv.Atoi(chi.URLParam(r, "userId"))
	user, err := mydb.GetUserById(int64(userId))
	if err != nil {
		w.WriteHeader(404)
	} else {
		json.NewDecoder(r.Body).Decode(&user)
		user.Update()
		json.NewEncoder(w).Encode(user)
	}

}

// @BasePath /api/v1

// DeleteUser godoc
// @Summary  delete user
// @Schemes
// @Description  delete an existing user
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      204  string  User  deleted
// @Router       /users [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(chi.URLParam(r, "userId"))
	mydb.DeleteUserById(int64(userId))
	w.Write([]byte("User deleted"))
}

// @BasePath /api/v1

// CreateTaskForUser godoc
// @Summary  create a task for user
// @Schemes
// @Description  create a new task and assign it for existing user
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      201  {object}  mydb.Task
// @Router       /users/{id}/create-task [get]
func CreateTaskForUser(w http.ResponseWriter, r *http.Request) {
	var task mydb.Task
	userId, _ := strconv.Atoi(chi.URLParam(r, "userId"))
	json.NewDecoder(r.Body).Decode(&task)
	task.CreatedBy = int64(userId)
	task.Create()
	json.NewEncoder(w).Encode(task)
}
