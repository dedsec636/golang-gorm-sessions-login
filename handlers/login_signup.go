package handlers

import (
	"fmt"
	"ginbasic/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("secretkeynga"))

func GetSignupHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func PostSignupHandler(c *gin.Context) {
	email1 := c.PostForm("email")
	username1 := c.PostForm("name")
	password1 := c.PostForm("password")
	fmt.Println(email1)
	user, err := model.FindUserByEmail(email1)
	fmt.Println(user)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"message": "Internal server error",
		})
		return
	}

	if user != nil {
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"error": "Account exists with given email",
		})
		return
	}

	usr := model.User{
		Email:    email1,
		Uname:    username1,
		PassHash: password1,
	}
	savedUser, err := usr.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, "error")
	}
	c.JSON(http.StatusCreated, gin.H{"user": savedUser})

}

func GetLoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func PostLoginHandler(c *gin.Context) {

	//to handle get request on refresh
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	}

	//post

	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := model.FindUserByEmail(email)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"message": "Internal server error",
		})
		return
	}

	if user == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"error": "No such account exists",
		})
		return
	}

	err = user.Validate(password)
	if err != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"error": "Wrong password",
		})
		return
	}
	Store.Options.MaxAge = 60 * 15

	session, _ := Store.Get(c.Request, "test-session")
	session.Values["email"] = user.Email
	session.Values["usrname"] = user.Uname
	session.Save(c.Request, c.Writer)
	c.Redirect(http.StatusFound, "/home")
}

func LogoutHandler(c *gin.Context) {
	session, _ := Store.Get(c.Request, "test-session")
	session.Options.MaxAge = -1
	err := session.Save(c.Request, c.Writer)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"message": "Internal server error",
		})
		return
	}
	c.Redirect(http.StatusSeeOther, "/login")
}
