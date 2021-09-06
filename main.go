// Recipes API
//
// This is a sample recipes API. You can find out more about
//the API at https://github.com/PacktPublishing/Building-Distributed-Applications-in-Gin.
//
// Schemes: http
// Host: localhost:8080
// BasePath: /
// Version: 1.0.0
// Contact: Mohamed Labouardy <mohamed@labouardy.com> https://labouardy.com
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/rs/xid"
)

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
	file, _ := ioutil.ReadFile("recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)
	router.PUT("/recipes/:id", UpdateReceiptHandler)
	router.DELETE("/recipes/:id", DeleteRecipeHandler)
	router.GET("/recipes/search", SearchRecipeHandler)
	router.Run()
}

// Recipe represents the recipe for this application
//
// A recipe is the main entity for this application.
// It's also used as one of main axes for reporting.
//
//
//
// swagger:model
type Recipe struct {
	// the id for this recipe automatic calculated
	//
	// required: true
	ID string `json:"id"`
	// the name for this recipe
	//
	// required: true
	Name string `json:"name"`
	// tags to easily make searches on the recipes
	//
	// required: false
	Tags []string `json:"tags"`
	// the ingredients for this recipe
	//
	// required: true
	Ingredients []string `json:"ingredients"`
	// the instructions to build this recipe
	//
	// required: true
	Instructions []string `json:"instructions"`
	// the date when this recipe was published (automatic defined when creating a new recipe)
	//
	// required: false
	PublishedAt time.Time `json:"publisehdAt"`
}

// swagger:operation POST /recipes recipes newRecipe
// Creates new Recipe
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: Successful operation
func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)

}

// swagger:operation GET /recipes recipes listRecipes
// Returns list of recipes
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: Successful operation
func ListRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

// swagger:operation PUT /recipes/{id} recipes updateRecipe
//
// Update an existing recipe
//
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: ID of the recipe
//   required: true
//   type: string
// responses:
//   '200':
//     description: Succesful operation
//   '400':
//     description: invalid output
//   '404':
//     description: invalid recipe ID
func UpdateReceiptHandler(c *gin.Context) {
	id := c.Param("id")
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	index := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found",
		})
		return
	}

	recipe.ID = id
	recipes[index] = recipe
	c.JSON(http.StatusOK, recipe)
}

func DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	index := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i
		}
	}

	recipes = append(recipes[:index], recipes[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe has been deleted",
	})
}

func SearchRecipeHandler(c *gin.Context) {
	tag := c.Query("tag")
	listOfRecipes := make([]Recipe, 0)

	for i := 0; i < len(recipes); i++ {
		found := false
		for _, t := range recipes[i].Tags {
			if strings.EqualFold(t, tag) {
				found = true
			}
		}
		if found {
			listOfRecipes = append(listOfRecipes, recipes[i])
		}
	}

	c.JSON(http.StatusOK, listOfRecipes)
}
