package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ArticleController struct {
	DB *gorm.DB
}

type Article struct {
	gorm.Model // ID, created_at, updated_at, deleted_at
	Title      string
	Author     string
	Content    string
}
type ArticlesResponse struct {
	Articles []Article
}
type ArticleResponse struct {
	Article Article
}

func DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	) + "?parseTime=true&collation=utf8mb4_bin"
}
func insertTestData(db *gorm.DB) {
	article1 := Article{
		Title:   "Hello HiCoder",
		Author:  "Bob",
		Content: "Super cool content",
	}
	article2 := Article{
		Title:   "GoodBye HiCoder",
		Author:  "Nancy",
		Content: "Super Hyper content",
	}
	article3 := Article{
		Title:   "Hello Hiroshima Univ",
		Author:  "Sato",
		Content: "Super extra content",
	}

	db.Create(&article1)
	db.Create(&article2)
	db.Create(&article3)
}
func main() {
	// TODO DBが立ち上がってないときに先に落ちてしまう。起動を遅らせるかリトライ処理を挟む
	db, err := gorm.Open(mysql.Open(DSN()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// TODO remove prod, this is dev for init
	db.Migrator().DropTable(&Article{})

	db.AutoMigrate(&Article{})
	// TODO remove this is data for dev
	insertTestData(db)

	ct := &ArticleController{
		DB: db,
	}
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// TODO remove production, avoid CORS in dev
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Routes
	e.GET("/", healthCheck)
	e.GET("/article", ct.GetAllArticles) // TODO change func
	e.GET("/article/:id", ct.GetArticle) // TODO change func

	// Start server
	e.Logger.Fatal(e.Start(":3001"))
}

// Handler
func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World! This server is alive.")
}

func (a *ArticleController) GetAllArticles(c echo.Context) error {
	var articles []Article
	res := a.DB.Find(&articles)
	if res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Error.")
	}
	var response ArticlesResponse
	response.Articles = articles
	return c.JSON(http.StatusOK, response)
}

func (a *ArticleController) GetArticle(c echo.Context) error {
	id := c.Param("id")
	var article Article
	res := a.DB.First(&article, "id = ?", id)
	if res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Error.")
	}
	response := ArticleResponse{
		Article: article,
	}
	return c.JSON(http.StatusOK, response)
}
