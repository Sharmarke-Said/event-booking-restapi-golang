package user

import (
	"net/http"
	"strconv"

	"example.com/event-booking-restapi/internal/auth"
	"example.com/event-booking-restapi/internal/auth/roles"
	"github.com/gin-gonic/gin"
)

// Auth
func signup(context *gin.Context) {
	var user User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data"})
		return
	}

	// set default role if not provided
	// user.Role = roles.User
	if user.Role == "" {
        user.Role = roles.User
    }

	// Hash the password first
    // hashedPassword, err := auth.HashPassword(user.Password)
    // if err != nil {
    //     context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user. Try again later."})
    //     return
    // }
    
    // Update the user object's Password field with the hashed password
  // user.Password = hashedPassword

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create user. Try again later.",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}
func login(context *gin.Context) {
	var user User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
		return
	}

	token, err := auth.GenerateToken(user.Email, user.ID, user.Role)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}


func GetUsers(context *gin.Context) {
    users, err := GetAllUsers()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch users. Try again later.",
		})
        return
    }
    context.JSON(http.StatusOK, users)
}

func GetUser(context *gin.Context) {
    userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err != nil {
        	context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not bind JSON",
		})
        return
    }

    user, err := GetUserByID(userId)
    if err != nil {
        	context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch user. Try again",
		})
		return
    }
    context.JSON(http.StatusOK, user)
}

func UpdateUser(context *gin.Context) {
    userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{
            "message":"Invalid user ID.",
        })
        return
    }

    var updatedUser User
    if err := context.ShouldBindJSON(&updatedUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data."})
		return
    }

    updatedUser.ID = userId
    if err := updatedUser.Update(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update user."})
		return
    }
    context.JSON(http.StatusOK, gin.H{"message": "User updated successfully."})
}

func DeleteUser(context *gin.Context) {
    userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID."})
		return
    }

    if err := DeleteUserByID(userId); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete user."})
		return
    }
    context.JSON(http.StatusOK, gin.H{"message": "User deleted successfully."})
}