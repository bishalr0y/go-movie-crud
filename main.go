package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Movie struct {
	Title    string    `json:"title" validate:"required"`
	Year     int       `json:"year" validate:"required"`
	Genre    string    `json:"genre" validate:"required"`
	Director *Director `json:"director" validate:"required"`
	// pointer is used here so the it holds the reference rather than a copy of the value
	// if pointer is not used here, then the changes made will not reflect
}

type Director struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
}

var movies []Movie

func main() {
	fmt.Println("Hello world")
	router := gin.Default()

	validate := validator.New(validator.WithRequiredStructEnabled())

	// insert some values inside the movies array
	movies = append(movies, Movie{
		Title: "Atlas",
		Year:  2024,
		Genre: "Action",
		Director: &Director{
			Firstname: "Brad",
			Lastname:  "Payton",
		},
	})

	movies = append(movies, Movie{
		Title: "Mad Max: Fury Road",
		Year:  2015,
		Genre: "Action",
		Director: &Director{
			Firstname: "Goerge",
			Lastname:  "Miller",
		},
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// * @GET /movies
	router.GET("/movies", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"movies": movies,
		})
	})

	// * @GET /movie?id=
	router.GET("movie/", func(c *gin.Context) {
		id := c.Query("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "id query is missing",
			})
			return
		}
		movieId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"movie": movies[movieId],
		})
	})

	// * @POST /movie
	router.POST("/movie", func(c *gin.Context) {
		var movie Movie
		c.Bind(&movie)
		if err := validate.Struct(movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		movies = append(movies, movie)
		c.JSON(http.StatusOK, gin.H{
			"movies": movies,
		})

	})

	// * @PUT /movie?id=
	router.PUT("/movie", func(c *gin.Context) {
		var movie Movie
		id := c.Query("id")
		c.Bind(&movie)

		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "id query is missing",
			})
			return
		}

		if err := validate.Struct(movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		movieId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// if the id is not present
		if movieId > len(movies) || movieId < 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "id is not present",
			})
			return
		}

		// update the movie
		movies[movieId] = movie
		c.JSON(http.StatusOK, gin.H{
			"movies": movies,
		})
	})

	// * @ DELETE /movie?id=
	router.DELETE("/movie", func(c *gin.Context) {
		id := c.Query("id")

		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "id query is missing",
			})
			return
		}

		movieId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		if movieId < 0 || movieId > len(movies) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "id is not present",
			})
		}
		movies = append(movies[:movieId], movies[movieId+1:]...)
		c.JSON(http.StatusOK, gin.H{
			"movies": movies,
		})
	})

	router.Run(":8080")
}
