package parsers

import (
	"encoding/xml"
	"jpdict-mongo/pkg/types"
	"jpdict-mongo/pkg/utils"
)

func GetJMdict() (types.JMdict, error) {
	// See https://groups.google.com/g/golang-nuts/c/yF9RM9rnkYc/m/9FbL0B6x7BYJ
	var jmdict types.JMdict
	file, err := utils.OpenXMLFile("JMdict")
	if err != nil {
		return jmdict, err
	}
	defer file.Close()

	d := xml.NewDecoder(file)
	for {
		token, err := d.Token()
		if err != nil || token == nil {
			break
		}
		switch se := token.(type) {
		case xml.StartElement:
			if err := d.DecodeElement(&jmdict, &se); err != nil {
				return jmdict, err
			}
		case xml.Directive:
			directive := token.(xml.Directive)
			entities := utils.ParseXMLEntities(directive)
			d.Entity = entities
		}
	}
	return jmdict, nil
}
