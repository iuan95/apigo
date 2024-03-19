package handler

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/iuan95/apigo/db"
	"github.com/iuan95/apigo/model"
)


func GetItems(c *fiber.Ctx) error{
	var items []model.Items
	conn, err:=db.DB.Query(context.Background(), "SELECT * FROM items")
	if err != nil{
		return c.Status(403).JSON(&fiber.Map{
			"message": "error can not get items",
		})
	}
	defer conn.Close()
	for conn.Next() {
		var item model.Items
		conn.Scan(&item.Id, &item.Name, &item.Description)
		items= append(items, item)
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": "welcome !",
		"data": items,
	})
	
}

func GetItemByID(c *fiber.Ctx) error{
	id:=c.Params("id")
	parceId,err := strconv.Atoi(id)
	if err !=nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "no correct id",
			"id": id,
		})
	}
	var item model.Items
	err=db.DB.QueryRow(context.Background(),"SELECT * FROM items WHERE id = $1", parceId).Scan(&item.Id, &item.Name, &item.Description)
	if err!=nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "no item with id="+id,
			"id": id,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": "ok",
		"data": item,
	})
}