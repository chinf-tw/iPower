package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

type ability struct {
	Title    []string  `json:"title"`
	ID       []string  `json:"ability_id"`
	Quantity []float64 `json:"quantity"`
}

func (i *ability) getAbility(db *sql.DB, id string, c *gin.Context) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		errHandle(err, c)
		return
	}

	// tsql := fmt.Sprintf("SELECT Ability.ability_name,Ability_record.[value] FROM Ability, Ability_record WHERE team_id = '%s' and Ability.ability_id = Ability_record.ability_id;",id)
	tsql := fmt.Sprintf("Select Ability.ability_name,a.[value] FROM Ability,Ability_record a Where Ability.ability_id = a.ability_id and a.team_id = '%s' and a.Ability_date = (Select Max(b.Ability_date) From Ability_record b Where team_id = '%s' and a.ability_id = b.ability_id) ORDER BY a.ability_id;", id, id)
	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	defer rows.Close()
	var abilityNameList []string
	var valueList []float64
	var ability_name string
	var value float64
	for rows.Next() {

		err := rows.Scan(&ability_name, &value)
		if err != nil {
			errHandle(err, c)
			return
		}
		abilityNameList = append(abilityNameList, ability_name)
		valueList = append(valueList, value)
	}
	i.Title = abilityNameList
	i.Quantity = valueList
}

func (i *ability) getTitle(db *sql.DB, c *gin.Context) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		errHandle(err, c)
		return
	}

	tsql := `SELECT ability_name,ability_id FROM Ability;`
	rows, err := db.QueryContext(ctx, tsql)
	defer rows.Close()

	var abilityName, abilityID string
	var abilityNameList, abilityIDList []string

	for rows.Next() {
		err := rows.Scan(&abilityName, &abilityID)
		if err != nil {
			errHandle(err, c)
			return
		}
		abilityNameList = append(abilityNameList, abilityName)
		abilityIDList = append(abilityIDList, abilityID)
	}
	i.Title = abilityNameList
	i.ID = abilityIDList
}
