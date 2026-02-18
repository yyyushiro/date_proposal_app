package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "modernc.org/sqlite"
)

func main() {
	// Opens database
	db, err := sql.Open("sqlite", "datePlans.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	places := []string{"新宿", "渋谷", "横浜", "自宅", "公園", "水族館", "夜の海", "知らない駅"}
	actions := []string{"散歩する", "アイスを食べる", "夜景を見る", "ゲームをする", "深海魚を眺める", "お弁当を食べる"}

	fmt.Println("Started inserting data...")
	start := time.Now()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO datePlans(title, content) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= 1000000; i++ {
		title := fmt.Sprintf("%sで%s No.%d", places[rand.Intn(len(places))], actions[rand.Intn(len(actions))], i)
		body := "これはテスト用のデートプラン詳細テキストです。大量のデータの中でも爆速で動くか検証中。"

		_, err = stmt.Exec(title, body)

		if i%10000 == 0 {
			fmt.Printf("%d cases finished... \n", i)
		}
	}

	tx.Commit()

	fmt.Printf("Transaction completed; total time: %d \n", time.Since(start))

}
