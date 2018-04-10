package config

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/navisot/movierama/app/database"
	"github.com/navisot/movierama/app/helpers"
	"github.com/navisot/movierama/app/models"
)

func StartMigration() {

	var allModels = make(map[string]interface{})
	allModels["User"] = &models.User{}
	allModels["Rating"] = &models.Rating{}
	allModels["Movie"] = &models.Movie{}

	for _, table := range allModels {
		database.DB.DropTableIfExists(table)
		database.DB.AutoMigrate(table)
	}

	u1 := models.User{}
	u1.Username = "nikavr77"
	u1.Email = "nikavr77@gmail.com"
	u1.Password = helpers.HashAndSalt([]byte("12345678"))
	database.DB.Model(&u1).Save(&u1)

	u2 := models.User{}
	u2.Username = "demopap"
	u2.Email = "demopap@demo.com"
	u2.Password = helpers.HashAndSalt([]byte("12345678"))
	database.DB.Model(&u2).Save(&u2)

	u3 := models.User{}
	u3.Username = "bestgol"
	u3.Email = "bestgol@gmail.com"
	u3.Password = helpers.HashAndSalt([]byte("12345678"))
	database.DB.Model(&u3).Save(&u3)

	m1 := models.Movie{}
	m1.UserID = 1
	m1.Title = "Lady Bird"
	m1.Description = "Lady Bird begins applying to East-coast colleges despite her mother's insistence that the family cannot afford it. She gets accepted into a local state university and is upset because it is too close to home. She is elated to discover that she has been placed on the wait list for a New York college but does not share the news with her mother, fearing her response. She sets out for her high school prom with Kyle, Jenna, and Jenna’s boyfriend, but the four decide to go to a party instead. Lady Bird changes her mind during the car ride and asks them to drop her off at Julie's apartment, where the two rekindle their friendship and go to prom together."
	m1.PublicationDate = time.Now()
	database.DB.Model(&m1).Save(&m1)

	m2 := models.Movie{}
	m2.UserID = 2
	m2.Title = "Prison Break"
	m2.Description = "Prison Break follows a man, Michael Scofield, on his mission to break his brother out of prison. Michael is convinced that his brother, Lincoln Burrows, is innocent of the crime for which he was committed, which was the murder of the vice president’s brother. Burrows is on death row, and scheduled for execution within the month. Scofield must make connections on the prison block that will aid him in his escape, but enemies are sure to make slow work of Michael’s plans."
	m2.PublicationDate = time.Now()
	database.DB.Model(&m2).Save(&m2)

	m3 := models.Movie{}
	m3.UserID = 2
	m3.Title = "Prison Break New Episode"
	m3.Description = "Prison Break follows a man, Michael Scofield, on his mission to break his brother out of prison. Michael is convinced that his brother, Lincoln Burrows, is innocent of the crime for which he was committed, which was the murder of the vice president’s brother. Burrows is on death row, and scheduled for execution within the month. Scofield must make connections on the prison block that will aid him in his escape, but enemies are sure to make slow work of Michael’s plans."
	m3.PublicationDate = time.Now()
	database.DB.Model(&m3).Save(&m3)

	m4 := models.Movie{}
	m4.UserID = 3
	m4.Title = "Go Series"
	m4.Description = "Learn GO."
	m4.PublicationDate = time.Now()
	database.DB.Model(&m4).Save(&m4)

	r1 := models.Rating{}
	r1.UserID = 3
	r1.MovieID = 1
	r1.Vote = 1
	database.DB.Model(&r1).Save(&r1)

	r2 := models.Rating{}
	r2.UserID = 3
	r2.MovieID = 2
	r2.Vote = 1
	database.DB.Model(&r2).Save(&r2)

	r3 := models.Rating{}
	r3.UserID = 2
	r3.MovieID = 1
	r3.Vote = 1
	database.DB.Model(&r3).Save(&r3)

	r4 := models.Rating{}
	r4.UserID = 1
	r4.MovieID = 2
	r4.Vote = 2
	database.DB.Model(&r4).Save(&r4)

	r5 := models.Rating{}
	r5.UserID = 1
	r5.MovieID = 4
	r5.Vote = 2
	database.DB.Model(&r5).Save(&r5)

}
