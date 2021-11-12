package models

import (
	"fmt"
	"log"
	"simplelogic/domain/entity"

	"simplelogic/config"

	_ "github.com/lib/pq" // postgres driver
)

//create struct user
//create Register User

func CreatTeam(user entity.User) string {
	db := config.CreateConnection()
	defer db.Close()
	inserting := `INSERT INTO teams (name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id`

	var id int64
	err := db.QueryRow(inserting, user.Name, user.Email, user.Password, user.Role).Scan(&id)
	if err != nil {
		log.Fatal("belum bisa daftar team yah!",err)
	}

	fmt.Printf("id %d", id)

	return user.Name

}
