package controllers

import (
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
		c.RenderJSON(err)
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

	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}

	return c.RenderJSON(users)
}

func (c User) Create() revel.Result {
	sql := "INSERT INTO users (name) VALUES ($1);"

	result, err := app.DB.Exec(sql, c.Params.Form.Get("name"))
	if err != nil {
		c.RenderJSON(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		c.RenderJSON(err)
	}
	if rows != 1 {
		c.RenderJSON(fmt.Sprintf("expected to affect 1 row, affected %d", rows))
	}
	return c.Redirect(routes.User.Index())
}
