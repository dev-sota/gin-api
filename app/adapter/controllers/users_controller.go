package users

import (
	"fmt"

	"github.com/gin-gonic/gin"

	users "github.com/dev-sota/gin-api/app/adapter/gateways"
)

// Controller is users controlller
type Controller struct{}

// Index action: GET /users
func (pc Controller) Index(c *gin.Context) {
	var g users.Gateway
	p, err := g.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /users
func (pc Controller) Create(c *gin.Context) {
	var g users.Gateway
	p, err := g.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Show action: GET /users/:id
func (pc Controller) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var g users.Gateway
	p, err := g.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: PUT /users/:id
func (pc Controller) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var g users.Gateway
	p, err := g.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /users/:id
func (pc Controller) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var g users.Gateway

	if err := g.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
