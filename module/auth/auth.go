package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type USER struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	password string
	Role     string `json:"role"`
	Verified bool   `json:"verified"`
}

var userList []USER

func InitializeUser() {
	// add manager
	userList = append(userList, USER{"1", "manager", "manager", "manager", true})

	// add agent
	userList = append(userList, USER{"2", "agent", "agent", "agent", true})

	// add institution
	userList = append(userList, USER{"3", "institution", "institution", "institution", true})

	// add lender
	userList = append(userList, USER{"4", "lender", "lender", "lender", true})

	// add privilege
	userList = append(userList, USER{"5", "privilege", "privilege", "privilege", true})

	// add borrower
	userList = append(userList, USER{"6", "borrower", "borrower", "borrower", true})

}

func SignIn(ctx *gin.Context) {
	var cred = struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if error := ctx.BindJSON(&cred); error != nil {
		logrus.Fatal("error happend while handling the login")
		ctx.JSON(400, gin.H{"status": "failure", "msg": "invalid data provided"})
		return
	}

	for _, item := range userList {
		if item.Username == cred.Username {
			if item.password == cred.Password {
				ctx.JSON(200, gin.H{"status": "success", "msg": "login success", "data": item})
				return
			}
			ctx.JSON(200, gin.H{"status": "failure", "msg": "invalid username or password"})
			return
		}
	}

	ctx.JSON(200, gin.H{"status": "failure", "msg": "invalid username or password"})
	return
}

func SignUp(ctx *gin.Context) {

	var userInfo = struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
		Role     string `json:"role,omitempty"`
	}{}

	if err := ctx.BindJSON(&userInfo); err != nil {
		logrus.Fatal("error happened while signUp")
		ctx.JSON(400, gin.H{"status": "failure", "msg": "invalid data provided"})
		return
	}

	if userInfo.Password == "" || userInfo.Role == "" || userInfo.Username == "" {
		ctx.JSON(400, gin.H{"status": "failure", "msg": "payload missing"})
		return
	}

	availableRoles := map[string]bool{
		"manager":     true,
		"agent":       true,
		"institution": true,
		"lender":      true,
		"privilege":   true,
		"borrower":    true,
	}

	if !availableRoles[userInfo.Role] {
		ctx.JSON(200, gin.H{"status": "failure", "msg": "invalid role"})
		return
	}

	for _, item := range userList {

		if item.Username == userInfo.Username {
			ctx.JSON(200, gin.H{"status": "failure", "msg": "username already exist"})
			return
		}
	}

	var newUser USER

	newUser.Role = userInfo.Role
	newUser.Username = userInfo.Username
	newUser.password = userInfo.Password

	id, err := uuid.NewUUID()
	if err != nil {
		ctx.JSON(200, gin.H{"status": "failure", "msg": "error while creating new user"})
		return
	}

	newUser.Id = id.String()
	newUser.Verified = false

	userList = append(userList, newUser)

	ctx.JSON(200, gin.H{"status": "pending", "msg": "user created successfully", "data": newUser})

}
