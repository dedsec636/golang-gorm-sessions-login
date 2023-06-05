package middleware

import (
	"ginbasic/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := handlers.Store.Get(c.Request, "test-session")
		email := session.Values["email"]
		if email != nil {
			// Set the email in the context
			c.Set("email", email)
		}
		if email == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		usrname := session.Values["usrname"]
		if usrname != nil {
			// Set the uname in the context
			c.Set("usrname", usrname)
		}
		c.Next()
	}
}
