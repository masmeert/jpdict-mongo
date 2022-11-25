package parsers

import "testing"

func TestJMdictParser(t *testing.T) {
	jmdict, err := ParseJMdict()
	if err != nil {
		t.Error(err)
	}
	t.Log(len(jmdict.Entries))
}

func TestKanjiDicParser(t *testing.T) {
	kanjidic, err := ParseKanjiDic()
	if err != nil {
		t.Error(err)
	}
	t.Log(len(kanjidic.Entries))
}
