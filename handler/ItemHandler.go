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
	conn, err:=db.DB.Query(context.Background(), "SELECT * FROM items ORDER BY id")
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
type User struct {
    Name        string `json:"name"`
    Description string `json:"description"`
}
func CreateItem(c *fiber.Ctx) error{
	var u User
	err:= c.BodyParser(&u)
	if err != nil || u.Description == "" || u.Name == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "no correct data",
		})
	}
	_,err= db.DB.Exec(context.Background(), "INSERT INTO items (name, description) VALUES($1, $2)", &u.Name, &u.Description)
	if err!= nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "can not create item",
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "item was created",
	})

}
func UpdateItemById(c *fiber.Ctx) error{
	id:=c.Params("id")
	var u User
	err:= c.BodyParser(&u)
	_,errParse := strconv.Atoi(id)
	if err != nil || errParse!=nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "no correct data",
		})
	}
	_,err= db.DB.Exec(context.Background(), "UPDATE items SET name = $1, description = $2 WHERE id = $3", &u.Name, &u.Description, &id)
	if err!= nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "can not update item",
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "item was updated",
	})

}


func DeleteItemById(c *fiber.Ctx) error{
	id:=c.Params("id")
	_,err := strconv.Atoi(id)
	if err != nil || id =="" {
		return c.Status(405).JSON(fiber.Map{
			"message": "no valid id",
		})
	}
	_,err= db.DB.Exec(context.Background(), "DELETE FROM items WHERE id=$1", &id)
	if err!=nil {
		return c.Status(405).JSON(fiber.Map{
			"message": "can not delete item",
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "item was deleted",
	})
}