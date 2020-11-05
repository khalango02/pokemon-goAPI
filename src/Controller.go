package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	Id             string `json:"id,omitempty"`
	NameEnglish    string `json:"nameEnglish,omitempty"`
	Type0          string `json:"type0,omitempty"`
	Type1          string `json:"type1,omitempty"`
	HP             int    `json:"HP,omitempty"`
	BaseAttack     int    `json:"baseAttack,omitempty"`
	BaseDefense    int    `json:"baseDefense,omitempty"`
	SpecialAttack  int    `json:"specialAttack,omitempty"`
	SpecialDefense int    `json:"specialDefense,omitempty"`
	BaseSpeed      int    `json:"baseSpeed,omitempty"`
}

func GetPokemons(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.1:3307)/Poke")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT Id, nameEnglish, type0, type1, HP, baseAttack, baseDefense, specialAttack, specialDefense, baseSpeed type0 FROM `Pokemons`")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag Tag

		err = results.Scan(&tag.Id, &tag.NameEnglish, &tag.Type0, &tag.Type1, &tag.HP, &tag.BaseAttack, &tag.BaseDefense, &tag.SpecialAttack, &tag.SpecialDefense, &tag.BaseSpeed)
		if err != nil {
			panic(err.Error())
		}
		log.Print(results)
		json.NewEncoder(w).Encode(tag)
	}
}
