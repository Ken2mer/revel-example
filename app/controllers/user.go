package controllers

import (
	"encoding/json"
	"fmt"
	"revel-example/app"
	"revel-example/app/models"
	"revel-example/app/routes"

	"github.com/revel/revel"
)

type User struct {
	*revel.Controller
}

func (c User) Index() revel.Result {
	rows, err := app.DB.Query("select * from users;")
	if err != nil {
		return c.RenderJSON(err)
	}
	defer rows.Close()

	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserId, &user.Name); err != nil {
			fmt.Println(err)
		}
		users = append(users, user)
	}

	return c.RenderJSON(users)
}

func (c User) Create() revel.Result {
	var user *models.User
	if err := json.Unmarshal(c.Params.JSON, &user); err != nil {
		return c.RenderJSON(err)
	}

	sql := "INSERT INTO users (name) VALUES ($1);"

	result, err := app.DB.Exec(sql, user.Name)
	if err != nil {
		return c.RenderJSON(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return c.RenderJSON(err)
	}
	if rows != 1 {
		return c.RenderJSON(fmt.Sprintf("expected to affect 1 row, affected %d", rows))
	}
	return c.Redirect(routes.User.Index())
}
