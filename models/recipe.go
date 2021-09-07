package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//swagger:parameters recipes newRecipe
type Recipe struct {
	//swagger:ignore
	ID primitive.ObjectID `json:"id" bson:"_id"`
	// the name for this recipe
	//
	// required: true
	Name string `json:"name" bson:"name"`
	// tags to easily make searches on the recipes
	//
	// required: false
	Tags []string `json:"tags" bson:"tags"`
	// the ingredients for this recipe
	//
	// required: true
	Ingredients []string `json:"ingredients" bson:"ingredients"`
	// the instructions to build this recipe
	//
	// required: true
	Instructions []string `json:"instructions" bson:"instructions"`
	// the date when this recipe was published (automatic defined when creating a new recipe)
	//
	// required: false
	PublishedAt time.Time `json:"publisehdAt" bson:"publisehdAt"`
}
