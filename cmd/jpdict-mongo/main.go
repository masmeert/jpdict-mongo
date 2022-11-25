package main

import (
	"context"
	"fmt"
	"jpdict-mongo/pkg/database"
	"jpdict-mongo/pkg/parsers"
	"jpdict-mongo/pkg/utils"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	utils.RaiseError(godotenv.Load())

	fmt.Println("Parsing dictionary files...")
	kanjis, err := parsers.ParseKanjidic()
	utils.RaiseError(err)

	fmt.Println("Connecting to MongoDB...")
	client, err := database.NewClient(os.Getenv("DB_URI"), os.Getenv("DB_NAME"))
	utils.RaiseError(err)

	fmt.Println("Inserting kanjidic entries...")
	_, err = client.Kanjis.InsertMany(context.TODO(), utils.KanjisToInterface(kanjis))
	utils.RaiseError(err)

	fmt.Println("Closing database connection...")
	utils.RaiseError(client.Close())
}
