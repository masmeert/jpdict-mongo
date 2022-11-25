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

	fmt.Println("Connecting to MongoDB...")
	client, err := database.NewClient(os.Getenv("DB_URI"), os.Getenv("DB_NAME"))
	utils.RaiseError(err)
	fmt.Println("Connected to MongoDB!")

	fmt.Println("Parsing dictionary files...")
	jmdict, err := parsers.ParseJMdict()
	utils.RaiseError(err)
	kanjidic, err := parsers.ParseKanjiDic()
	utils.RaiseError(err)
	fmt.Println("Parsed dictionary files!")

	fmt.Println("Inserting kanjidic entries...")
	_, err = client.KanjiDic.InsertMany(context.TODO(), utils.KanjidicToInterface(kanjidic))
	utils.RaiseError(err)
	fmt.Println("Inserted kanjidic entries!")

	fmt.Println("Inserting jmdict entries...")
	_, err = client.JMdict.InsertMany(context.TODO(), utils.JMdictToInterface(jmdict))
	utils.RaiseError(err)
	fmt.Println("Inserted jmdict entries!")

	fmt.Println("Closing database connection...")
	utils.RaiseError(client.Close())
	fmt.Println("Closed database connection!")
}
