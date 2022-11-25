package main

import (
	"context"
	"jpdict-mongo/pkg/database"
	"jpdict-mongo/pkg/parsers"
	"jpdict-mongo/pkg/utils"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	utils.RaiseError(godotenv.Load())

	client, err := database.NewClient(os.Getenv("DB_URI"), os.Getenv("DB_NAME"))
	utils.RaiseError(err)

	jmdict, err := parsers.ParseJMdict()
	if err != nil {
		panic(err)
	}

	kanjidic, err := parsers.ParseKanjiDic()
	utils.RaiseError(err)

	_, err = client.KanjiDic.InsertMany(context.TODO(), utils.KanjidicToInterface(kanjidic))
	utils.RaiseError(err)

	_, err = client.JMdict.InsertMany(context.TODO(), utils.JMdictToInterface(jmdict))
	utils.RaiseError(err)

	utils.RaiseError(client.Close())
}
