package handlers

import (
	"ginbasic/helper"
	"ginbasic/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProtectedHomePageHandler(c *gin.Context) {
	sessionUsername := c.GetString("usrname")
	var b1 []model.Book
	err := helper.Database.Preload("Author").Find(&b1).Error
	if err != nil {
		helper.ErrorPanic(err)
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"name":  sessionUsername,
		"entry": b1,
	})

}

func GetCreateNewHandler(c *gin.Context) {
	var authors []model.Author
	helper.Database.Find(&authors)
	//fmt.Println(authors)

	c.HTML(http.StatusOK, "create.html", gin.H{
		"authors": authors,
	})
}

func CreateAuth(c *gin.Context) {
	authName := c.PostForm("authorName")

	// Check if an author with the given name already exists
	var existingAuthor model.Author
	if helper.Database.Where("Name = ?", authName).First(&existingAuthor).Error == nil {
		c.HTML(http.StatusOK, "create.html", gin.H{
			"message": "Author already exists with the given name",
		})
		return
	}

	// Create a new author with the given name
	newAuthor := model.Author{
		Name: authName,
	}
	err := helper.Database.Create(&newAuthor).Error
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"message": "Error creating database entry",
		})
		return
	}

	// Redirect to a different page after successful form submission
	c.Redirect(http.StatusSeeOther, "/home")
}

func CreateBook(c *gin.Context) {
	bookTitle := c.PostForm("bookTitle")
	authorID := c.PostForm("authorID")

	// Check if the selected author exists
	var existingAuthor model.Author
	err := helper.Database.First(&existingAuthor, authorID).Error
	if err != nil {
		c.HTML(http.StatusOK, "create.html", gin.H{
			"message": "Selected author does not exist",
		})
		return
	}

	// Create a new book with the given title and author ID
	newBook := model.Book{
		Title:    bookTitle,
		AuthorID: existingAuthor.ID,
	}
	err = helper.Database.Create(&newBook).Error
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"message": "Error creating database entry",
		})
		return
	}

	// Redirect to a different page after successful form submission
	c.Redirect(http.StatusSeeOther, "/home")
}
