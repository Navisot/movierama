package controllers

import (
	"fmt"

	"github.com/navisot/movierama/app/database"
	"github.com/navisot/movierama/app/models"
	"github.com/pkg/errors"
)

func SaveMovie(mov *models.Movie) (*models.Movie, error) {

	saveError := database.DB.Debug().Save(&mov).Error

	if saveError != nil {
		return nil, saveError
	}

	return mov, nil
}

func GetAllMovies(isLoggedIn bool, email string) ([]models.MovieResult, error) {

	var result []models.MovieResult

	if isLoggedIn {

		userID := GetUserIDByEmail(email)

		err := database.DB.Raw("select u.user_id as user_id, mv.movie_id as movie_id,u.username as username,u.email as user_email,mv.title as movie_title,mv.description as movie_description,mv.publication_date as movie_date, (select count(r.vote) from ratings r where r.movie_id = mv.movie_id and r.vote = 1) as movie_likes,(select count(r.vote) from ratings r where r.movie_id = mv.movie_id and r.vote = 2) as movie_hates, (select r.vote from ratings r where r.user_id = ? and mv.movie_id = r.movie_id) as user_like from users u inner join movies mv on mv.user_id = u.user_id GROUP BY u.user_id, mv.movie_id", userID).Scan(&result).Error

		if err != nil {
			fmt.Println("ERROR")
		}

	} else {

		database.DB.Raw("select u.user_id as user_id, mv.movie_id as movie_id,u.username as username,u.email as user_email,mv.title as movie_title,mv.description as movie_description,mv.publication_date as movie_date, (select count(r.vote) from ratings r where r.movie_id = mv.movie_id and r.vote = 1) as movie_likes,(select count(r.vote) from ratings r where r.movie_id = mv.movie_id and r.vote = 2) as movie_hates from users u inner join movies mv on mv.user_id = u.user_id GROUP BY u.user_id, mv.movie_id").Scan(&result)

	}
	return result, nil

}

func GetAllMoviesSorted(item string, isLoggedIn bool, email string) ([]models.MovieResult, error) {

	var result []models.MovieResult

	if isLoggedIn {

		userID := GetUserIDByEmail(email)

		switch item {
		case "likes":
			err := database.DB.Raw("select u.user_id as user_id, mv.movie_id as movie_id,u.username as username,u.email as user_email,mv.title as movie_title,mv.description as movie_description,mv.publication_date as movie_date, (select count(r.vote) from ratings r where r.movie_id = mv.movie_id and r.vote = 1) as movie_likes,(select count(r.vote) from ratings r where r.movie_id = mv.movie_id and r.vote = 2) as movie_hates, (select r.vote from ratings r where r.user_id = ? and mv.movie_id = r.movie_id) as user_like from users u inner join movies mv on mv.user_id = u.user_id GROUP BY u.user_id, mv.movie_id order by movie_likes desc", userID).Scan(&result).Error
			if err != nil {
				fmt.Println("ERROR")
			}
		case "hates":
			err := database.DB.Raw("select u.user_id as user_id, mv.movie_id as movie_id,u.username as username,u.email as user_email,mv.title as movie_title,mv.description as movie_description,mv.publication_date as movie_date, (select count(r.vote) from ratings r where r.movie_id = mv.movie_id and r.vote = 1) as movie_likes,(select count(r.vote) from ratings r where r.movie_id = mv.movie_id and r.vote = 2) as movie_hates, (select r.vote from ratings r where r.user_id = ? and mv.movie_id = r.movie_id) as user_like from users u inner join movies mv on mv.user_id = u.user_id GROUP BY u.user_id, mv.movie_id order by movie_hates desc", userID).Scan(&result).Error
			if err != nil {
				fmt.Println("ERROR")
			}
		case "date":
			err := database.DB.Raw("select u.user_id as user_id, mv.movie_id as movie_id,u.username as username,u.email as user_email,mv.title as movie_title,mv.description as movie_description,mv.publication_date as movie_date, (select count(r.vote) from ratings r where r.movie_id = mv.movie_id and r.vote = 1) as movie_likes,(select count(r.vote) from ratings r where r.movie_id = mv.movie_id and r.vote = 2) as movie_hates, (select r.vote from ratings r where r.user_id = ? and mv.movie_id = r.movie_id) as user_like from users u inner join movies mv on mv.user_id = u.user_id GROUP BY u.user_id, mv.movie_id order by movie_date", userID).Scan(&result).Error
			if err != nil {
				fmt.Println("ERROR")
			}
		default:
			return nil, errors.New("Not Such Order Method")
		}
	} else {
		switch item {
		case "likes":
			database.DB.Raw("select u.user_id as user_id, mv.movie_id as movie_id,u.username as username,u.email as user_email,mv.title as movie_title,mv.description as movie_description,mv.publication_date as movie_date, (select count(r.vote) from ratings r where r.movie_id = mv.movie_id and r.vote = 1) as movie_likes,(select count(r.vote) from ratings r where r.movie_id = mv.movie_id and r.vote = 2) as movie_hates from users u inner join movies mv on mv.user_id = u.user_id GROUP BY u.user_id, mv.movie_id order by movie_likes desc").Scan(&result)
		case "hates":
			database.DB.Raw("select u.user_id as user_id, mv.movie_id as movie_id,u.username as username,u.email as user_email,mv.title as movie_title,mv.description as movie_description,mv.publication_date as movie_date, (select count(r.vote) from ratings r where r.movie_id = mv.movie_id and r.vote = 1) as movie_likes,(select count(r.vote) from ratings r where r.movie_id = mv.movie_id and r.vote = 2) as movie_hates from users u inner join movies mv on mv.user_id = u.user_id GROUP BY u.user_id, mv.movie_id order by movie_hates desc").Scan(&result)
		case "date":
			database.DB.Raw("select u.user_id as user_id, mv.movie_id as movie_id,u.username as username,u.email as user_email,mv.title as movie_title,mv.description as movie_description,mv.publication_date as movie_date, (select count(r.vote) from ratings r where r.movie_id = mv.movie_id and r.vote = 1) as movie_likes,(select count(r.vote) from ratings r where r.movie_id = mv.movie_id and r.vote = 2) as movie_hates from users u inner join movies mv on mv.user_id = u.user_id GROUP BY u.user_id, mv.movie_id order by movie_date desc").Scan(&result)
		default:
			return nil, errors.New("Not Such Order Method")
		}
	}

	return result, nil

}
