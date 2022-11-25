package utils

import (
	"encoding/xml"
	"jpdict-mongo/pkg/types"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	rootpath   = filepath.Join(filepath.Dir(b), "../../.")
)

func RaiseError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetRootPath() string {
	return rootpath
}

func OpenXMLFile(filename string) (*os.File, error) {
	xmlFile, err := os.Open(filepath.Join(GetRootPath(), "assets", filename+".xml"))
	if err != nil {
		return nil, err
	}
	return xmlFile, nil
}

func ParseXMLEntities(directive xml.Directive) map[string]string {
	entity := regexp.MustCompile(`ENTITY\s+([^\s]+)\s+"([^"]+)"`)
	entities := make(map[string]string)
	for _, m := range entity.FindAllStringSubmatch(string(directive), -1) {
		entities[string(m[1])] = string(m[2])
	}
	return entities
}

func KanjidicToInterface(kanjidic types.KanjiDic) []interface{} {
	entries := make([]interface{}, len(kanjidic.Entries))
	for i, entry := range kanjidic.Entries {
		entries[i] = entry
	}
	return entries
}

func JMdictToInterface(jmdict types.JMdict) []interface{} {
	entries := make([]interface{}, len(jmdict.Entries))
	for i, entry := range jmdict.Entries {
		entries[i] = entry
	}
	return entries
}
