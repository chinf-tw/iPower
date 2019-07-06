package main

import (
	// "math/rand"
	// "time"
	"context"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

type item struct {
	Title    []string `json:"title"`
	Quantity []int    `json:"quantity"`
	Price    []int    `json:"price"`
	Item_id  []string `json:"item_id"`
}

func (i *item) getItem(db *sql.DB, id string, c *gin.Context) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		errHandle(err, c)
		return
	}

	tsql := fmt.Sprintf("Select Item.item_name,a.[value] FROM Item,Item_record a Where Item.item_id = a.item_id and a.team_id = '%s' and a.Item_date = (Select Max(b.Item_date) From Item_record b Where team_id = '%s' and a.item_id = b.item_id)  ORDER BY a.item_id;", id, id)

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	defer rows.Close()
	var itemNameList []string
	var valueList []int
	var item_name string
	var value int
	for rows.Next() {

		err := rows.Scan(&item_name, &value)
		if err != nil {
			errHandle(err, c)
			return
		}
		itemNameList = append(itemNameList, item_name)
		valueList = append(valueList, value)
	}
	i.Title = itemNameList
	i.Quantity = valueList
}
func (i *item) getPrice(db *sql.DB, c *gin.Context) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		errHandle(err, c)
		return
	}

	tsql := fmt.Sprintf("SELECT item_name,item_id,[value] FROM Item;")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	defer rows.Close()
	var titleList []string
	var valueList []int
	var idList []string
	var title string
	var id string
	var value int

	for rows.Next() {

		err := rows.Scan(&title, &id, &value)
		if err != nil {
			errHandle(err, c)
			return
		}
		titleList = append(titleList, title)
		valueList = append(valueList, value)
		idList = append(idList, id)
	}
	i.Title = titleList
	i.Price = valueList
	i.Item_id = idList
}
