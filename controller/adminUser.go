package controller

import (
	"net/http"

	"github.com/Heath000/fzuSE2024/model"
	"github.com/gin-gonic/gin"
)

// AdminUserController is the controller for admin user operations
type AdminUserController struct{}

// GetUserList retrieves the list of all users
func (ctrl *AdminUserController) GetUserList(c *gin.Context) {
	var users []model.User

	// Retrieve all users from the database
	if err := model.GetAllUsers(&users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUser retrieves a single user by ID
func (ctrl *AdminUserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User

	// Get user by ID
	if err := user.GetFirstByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// CreateUser creates a new user by admin
func (ctrl *AdminUserController) CreateUser(c *gin.Context) {
	var form Signup

	// Bind the form data to the Signup struct
	if err := c.ShouldBind(&form); err == nil {
		// Check if passwords match
		if form.Password != form.Password2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password does not match with confirm password"})
			return
		}

		// Create a new user model
		var user model.User
		user.Name = form.Name
		user.Email = form.Email
		user.Password = form.Password

		// Call the AdminCreateUser function to create the user
		if err := model.AdminCreateUser(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// DeleteUser deletes a user by ID (admin only)
func (ctrl *AdminUserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// Call the DeleteUserByID function in the model to delete the user
	if err := model.DeleteUserByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// UpdateUser updates a user's information
func (ctrl *AdminUserController) UpdateUser(c *gin.Context) {
	var form Signup

	// Bind the form data to the Signup struct
	if err := c.ShouldBind(&form); err == nil {
		// Check if passwords match
		if form.Password != form.Password2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password does not match with confirm password"})
			return
		}

		// Update the user info
		var user model.User
		user.Name = form.Name
		user.Email = form.Email
		user.Password = form.Password

		// Call the AdminUpdateUser function to update the user's data
		if err := model.AdminUpdateUser(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
