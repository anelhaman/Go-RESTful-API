package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Owner : export all
type Owner struct {
	ID   string
	Name string
}

// Comment : export all
type Comment struct {
	Description string
	Owner       Owner
	TimeCreate  string
}

// Post : export all
type Post struct {
	PostName    string
	ImgURL      []string
	TimeCreate  string
	Description []string  //description
	Tag         []string  // tags of post
	Owner       Owner     // whom posted this online
	Comments    []Comment // who's  join
}

func main() {
	r := gin.Default()
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println(err.Error())
	}

	r.GET("/find", func(c *gin.Context) {
		var posts []Post
		collection := session.DB("pantip").C("posts")
		err = collection.Find(bson.M{}).All(&posts)
		if err != nil {
			fmt.Println(err.Error())
		}
		c.JSON(200, posts)

	})

	defer session.Close()

	r.Run(":4000")
}
