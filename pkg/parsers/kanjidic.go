package parsers

import (
	"encoding/xml"
	"io"
	"jpdict-mongo/pkg/types"
	"jpdict-mongo/pkg/utils"
	"strconv"
)

func GetKanjiDic() (types.KanjiDic, error) {
	var kanjidic types.KanjiDic
	file, err := utils.OpenXMLFile("kanjidic2")
	if err != nil {
		return kanjidic, err
	}

	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return kanjidic, err
	}

	err = xml.Unmarshal(bytes, &kanjidic)
	if err != nil {
		return kanjidic, err
	}

	return kanjidic, nil
}

func ParseKanjidic() ([]types.Kanji, error) {
	var kanjis []types.Kanji
	kanjidic, err := GetKanjiDic()
	if err != nil {
		return kanjis, err
	}

	for _, entry := range kanjidic.Entries {
		kanji := types.Kanji{
			Character: entry.Literal,
			Grade:     parseMiscString(entry.Misc.Grade),
			Stroke:    parseMiscString(entry.Misc.StrokeCount[0]),
			Frequency: parseMiscString(entry.Misc.Freq),
			JLPT:      parseMiscString(entry.Misc.Jlpt),
			Meaning:   parseMeaning(entry.ReadingMeaning.Meaning),
			Onyomi:    parseOnyomi(entry.ReadingMeaning.Reading),
			Kunyomi:   parseKunyomi(entry.ReadingMeaning.Reading),
		}
		kanjis = append(kanjis, kanji)
	}
	return kanjis, nil
}

func parseMeaning(meaning []types.KanjiDicEntryMeaning) []string {
	var meanings []string
	for _, m := range meaning {
		if m.Lang == "" {
			meanings = append(meanings, m.Value)
		}
	}
	return meanings
}

func parseOnyomi(reading []types.KanjiDicEntryReading) []string {
	var onyomi []string
	for _, r := range reading {
		if r.Type == "ja_on" {
			onyomi = append(onyomi, r.Value)
		}
	}
	return onyomi
}

func parseKunyomi(reading []types.KanjiDicEntryReading) []string {
	var kunyomi []string
	for _, r := range reading {
		if r.Type == "ja_kun" {
			kunyomi = append(kunyomi, r.Value)
		}
	}
	return kunyomi
}

func parseMiscString(miscStr string) int {
	if miscStr == "" {
		return 0
	}
	i, _ := strconv.Atoi(miscStr)
	return i
}
