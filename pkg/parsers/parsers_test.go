package parsers

import "testing"

func TestGetJMdict(t *testing.T) {
	jmdict, err := GetJMdict()
	if err != nil {
		t.Error(err)
	}
	t.Log(len(jmdict.Entries))
}

func TestGetKanjiDic(t *testing.T) {
	kanjidic, err := GetKanjiDic()
	if err != nil {
		t.Error(err)
	}
	t.Log(len(kanjidic.Entries))
}

func TestParseKanjidic(t *testing.T) {
	kanjis, err := ParseKanjidic()
	if err != nil {
		t.Error(err)
	}
	t.Log(kanjis[0])
}
