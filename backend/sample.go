package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type JobListing struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Company     string   `json:"company"`
	Location    string   `json:"location"`
	Salary      string   `json:"salary"`
	Skills      []string `json:"skills"`
}

func main() {
	// JSON ファイルから求人情報を読み込む
	jobListingData, err := ioutil.ReadFile("job_listing.json")
	if err != nil {
		panic(err)
	}

	// JSON データを JobListing 構造体にアンマーシャルする
	var jobListing JobListing
	err = json.Unmarshal(jobListingData, &jobListing)
	if err != nil {
		panic(err)
	}

	// 求人情報を出力する
	fmt.Println("求人情報:")
	fmt.Println("ID:", jobListing.ID)
	fmt.Println("タイトル:", jobListing.Title)
	fmt.Println("説明:", jobListing.Description)
	fmt.Println("会社:", jobListing.Company)
	fmt.Println("勤務地:", jobListing.Location)
	fmt.Println("給与:", jobListing.Salary)
	fmt.Println("スキル:", jobListing.Skills)
}
