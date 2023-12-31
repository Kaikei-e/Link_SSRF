package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log/slog"
	"net/http"
	"os"

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
	ID          int    `json:"id,omitempty"`
	Username    string `json:"username" param:"username"`
	ProfileLink string `json:"profile_link" param:"profile_link"`
}

type registerRequest struct {
	Username    string `json:"username"`
	ProfileLink string `json:"profile_link"`
}

type registerResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type AuthedUserListResponse struct {
	Users []User
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// conn, err := repository.InitDBConn()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer conn.Close()

	e := echo.New()
	port := ":9000"

	apiV1 := e.Group("/api/v1")
	apiV1.GET("/profile_link", func(c echo.Context) error {
		slog.Info("GET /api/v1/myprofile", "Remote Addr", c.Request().RemoteAddr)
		slog.Info("GET /api/v1/myprofile", "", c.Request().RequestURI)

		return c.JSON(http.StatusOK, map[string]string{
			"message": "success",
		})
	})

	apiV1.GET("/users", func(c echo.Context) error {
		users := []User{
			{
				ID:          1,
				ProfileLink: "https://owasp.org/API-Security/editions/2023/en/0x11-t10/",
				Username:    "hoge",
			},
			{
				ID:          2,
				ProfileLink: "https://owasp.org/API-Security/editions/2023/en/0x11-t10/",
				Username:    "fuga",
			},
		}

		return c.JSON(http.StatusOK, users)
	})

	apiV1.POST("/register", func(c echo.Context) error {
		slog.Info("POST /api/v1/register", "", c.Request().RequestURI)

		var req registerRequest
		if err := c.Bind(&req); err != nil {
			slog.Error("Failed to bind request", "Error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "failed to bind request",
			})
		}

		// u, err := url.Parse(req.ProfileLink)
		// if err != nil {
		// 	slog.Error("Failed to parse url", "Error", err)
		// 	return c.JSON(http.StatusBadRequest, map[string]string{
		// 		"message": "failed to parse url",
		// 	})
		// }

		// tx, err := conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
		// if err != nil {
		// 	slog.Error("Failed to begin transaction", "Error", err)
		// }

		// _, err = tx.ExecContext(ctx, "INSERT INTO vul_schema.users (username, profile_link) VALUES ($1, $2);", req.Username, u.String())
		// if err != nil {
		// 	slog.Error("Failed to insert user", "Error", err)
		// 	tx.Rollback()
		// 	return c.JSON(http.StatusInternalServerError, map[string]string{
		// 		"message": "failed to insert user",
		// 	})
		// }

		// err = tx.Commit()
		// if err != nil {
		// 	slog.Error("Failed to commit transaction", "Error", err)
		// 	return c.JSON(http.StatusInternalServerError, map[string]string{
		// 		"message": "failed to commit transaction",
		// 	})
		// }

		rq, err := http.NewRequest(http.MethodGet, req.ProfileLink, nil)
		if err != nil {
			slog.Error("Failed to get profile", "Error", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "failed to get profile",
			})
		}
		cl := &http.Client{}
		re, err := cl.Do(rq)
		if err != nil {
			slog.Error("Failed to get profile", "Error", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "failed to get profile",
			})
		}
		defer re.Body.Close()

		// client := &http.Client{}
		// re, err := client.Do(&http.Request{
		// 	Method: http.MethodGet,
		// 	URL:    u,
		// })
		// if err != nil {
		// 	slog.Error("Failed to get profile", "Error", err)
		// 	return c.JSON(http.StatusInternalServerError, map[string]string{
		// 		"message": "failed to get profile",
		// 	})
		// }
		// defer re.Body.Close()

		// rawHTMLBobyBytes := re.Body.Read()
		rawHTMLBodyBytes := make([]byte, 300)
		re.Body.Read(rawHTMLBodyBytes)

		fmt.Println(string(rawHTMLBodyBytes))

		res := &registerResponse{
			Status:  http.StatusOK,
			Message: string(rawHTMLBodyBytes),
		}
		// rawHTMLBody := string(rawHTMLBobyBytes)
		return c.JSON(http.StatusOK, res)
	})

	// e.POST("/delete/:id", func(c echo.Context) error {
	// 	maybeID := c.Param("id")
	// 	slog.Info("POST /delete/:id", "Remote Addr", c.Request().RemoteAddr)
	// 	slog.Info("POST /delete/:id", "", c.Request().RequestURI)

	// 	id, err := strconv.Atoi(maybeID)
	// 	if err != nil {
	// 		slog.Error("Failed to convert string to int", "Error", err)
	// 		c.Redirect(http.StatusBadRequest, "/")
	// 	}

	// 	return deleteUser(id, c, conn)
	// })

	// authed := e.Group("/authed")
	// authed.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	c.Redirect(http.StatusMovedPermanently, "/")

	// 	return false, nil
	// }))

	// e.GET("/admin", func(c echo.Context) error {
	// 	return getAdminPage(c, conn)
	// })

	// authed.POST("/delete/:id", func(c echo.Context) error {

	// 	maybeID := c.Param("id")
	// 	slog.Info("POST /delete/:id", "Remote Addr", c.Request().RemoteAddr)
	// 	slog.Info("POST /delete/:id", "", c.Request().RequestURI)

	// 	id, err := strconv.Atoi(maybeID)
	// 	if err != nil {
	// 		slog.Error("Failed to convert string to int", "Error", err)
	// 		c.Redirect(http.StatusBadRequest, "/")
	// 	}

	// 	return deleteUser(id, c, conn)
	// })

	e.Logger.Fatal(e.Start(port))
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
