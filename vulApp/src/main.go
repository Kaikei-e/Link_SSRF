package main

import (
	"context"
	"database/sql"
	"html/template"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"rssf/repository"
	"strconv"

	"github.com/labstack/echo/v4"
)

var ctx = context.Background()

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type User struct {
	ID          int
	Username    string
	ProfileLink string
}

type AuthedUserListResponse struct {
	Users []User
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	conn, err := repository.InitDBConn()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	e := echo.New()
	port := ":9000"
	t := &Template{
		templates: template.Must(template.ParseGlob("public/*.html")),
	}
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		slog.Info("GET /", "Remote Addr", c.Request().RemoteAddr)
		slog.Info("GET /", "", c.Request().RequestURI)

		return getHomePage(c)
	})

	// authed := e.Group("/authed")
	// {
	// authed.GET("/admin", func(c echo.Context) error {
	e.GET("/admin", func(c echo.Context) error {
		return getAdminPage(c, conn)
	})
	// }

	e.POST("/delete/:id", func(c echo.Context) error {
		maybeID := c.Param("id")
		slog.Info("POST /delete/:id", "Remote Addr", c.Request().RemoteAddr)
		slog.Info("POST /delete/:id", "", c.Request().RequestURI)

		id, err := strconv.Atoi(maybeID)
		if err != nil {
			slog.Error("Failed to convert string to int", "Error", err)
			c.Redirect(http.StatusBadRequest, "/")
		}

		return deleteUser(id, c, conn)
	})

	e.Logger.Fatal(e.Start(port))
}

func getHomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "hoge")
}

func getAdminPage(c echo.Context, conn *sql.DB) error {
	slog.Info("GET /admin", "Remote Addr", c.Request().RemoteAddr)
	slog.Info("GET /admin", "", c.Request().RequestURI)

	rows, err := conn.Query("SELECT * FROM vul_schema.users;")
	if err != nil {
		slog.Error("Failed to bind the results", "Error", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Username, &u.ProfileLink)
		if err != nil {
			slog.Error("Bindng result", "Error", err)
		}
		users = append(users, u)
	}

	res := &AuthedUserListResponse{
		Users: users,
	}

	return c.Render(http.StatusOK, "admin", res)
}

func deleteUser(id int, c echo.Context, conn *sql.DB) error {
	slog.Info("POST /delete/:id", "Remote Addr", c.Request().RemoteAddr)

	tx, err := conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		slog.Error("Failed to begin transaction", "Error", err)
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM vul_schema.users WHERE id = $1;", id)
	if err != nil {
		slog.Error("Failed to delete user", "Error", err)
		tx.Rollback()
		return c.Redirect(http.StatusInternalServerError, "/")
	}

	err = tx.Commit()
	if err != nil {
		slog.Error("Failed to commit transaction", "Error", err)
		return c.Redirect(http.StatusInternalServerError, "/")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})

}
