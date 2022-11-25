package types

type JMdict struct {
	Entries []JMdictEntry `xml:"entry"`
}

type JMdictEntry struct {
	Sequence   int                `xml:"ent_seq"`
	Expression []JMdictExpression `xml:"k_ele"`
	Reading    []JMdictReading    `xml:"r_ele"`
	Sense      []JMdictSense      `xml:"sense"`
}

type JMdictExpression struct {
	Expression  string   `xml:"keb"`
	Information []string `xml:"ke_inf"`
	Priority    []string `xml:"ke_pri"`
}

type JMdictReading struct {
	Reading     string   `xml:"reb"`
	NoKanji     string   `xml:"re_nokanji"`
	Restriction []string `xml:"re_restr"`
	Information []string `xml:"re_inf"`
	Priority    []string `xml:"re_pri"`
}

type JMdictSense struct {
	RestrictedKanji   []string                `xml:"stagk"`
	RestrictedReading []string                `xml:"stagr"`
	Reference         []string                `xml:"xref"`
	Antonym           []string                `xml:"ant"`
	PartOfSpeech      []string                `xml:"pos"`
	Field             []string                `xml:"field"`
	Miscellaneous     []string                `xml:"misc"`
	LanguageSource    []JMdictSenseLangSource `xml:"lsource"`
	Dialect           []string                `xml:"dial"`
	Gloss             []JMdictSenseGloss      `xml:"gloss"`
	Priority          []string                `xml:"pri"`
	Information       []string                `xml:"s_inf"`
	Example           []JMdictSenseExample    `xml:"example"`
}

type JMdictSenseLangSource struct {
	Value    string `xml:",chardata"`
	Language string `xml:"lang,attr"`
	Type     string `xml:"ls_type,attr"`
	Wasei    string `xml:"ls_wasei,attr"`
}

type JMdictSenseGloss struct {
	Value    string `xml:",chardata"`
	Language string `xml:"lang,attr"`
	Gender   string `xml:"g_gend,attr"`
	Type     string `xml:"g_type,attr"`
}

type JMdictSenseExample struct {
	Source    JMdictExampleSource     `xml:"ex_srce"`
	Text      string                  `xml:"ex_text"`
	Sentences []JMdictExampleSentence `xml:"ex_sent"`
}

type JMdictExampleSource struct {
	ID   string `xml:",chardata"`
	Type string `xml:"exsrc_type,attr"`
}

type JMdictExampleSentence struct {
	Language string `xml:"lang,attr"`
	Text     string `xml:",chardata"`
}

type KanjiDic struct {
	Header  KanjiDicHeader  `xml:"header"`
	Entries []KanjiDicEntry `xml:"character"`
}

type KanjiDicHeader struct {
	FileVersion     string `xml:"file_version"`
	DatabaseVersion string `xml:"database_version"`
	DateOfCreation  string `xml:"date_of_creation"`
}

type KanjiDicEntry struct {
	Literal        string                      `xml:"literal"`
	Codepoint      []KanjiDicEntryCodepoint    `xml:"codepoint>cp_value"`
	Radical        []KanjiDicEntryRadical      `xml:"radical>rad_value"`
	Misc           KanjiDicEntryMisc           `xml:"misc"`
	DicNumber      []KanjiDicEntryDicNumber    `xml:"dic_number>dic_ref"`
	QueryCode      []KanjiDicEntryQueryCode    `xml:"query_code>q_code"`
	ReadingMeaning KanjiDicEntryReadingMeaning `xml:"reading_meaning"`
}

type KanjiDicEntryCodepoint struct {
	Value string `xml:",chardata"`
	Type  string `xml:"cp_type,attr"`
}

type KanjiDicEntryRadical struct {
	Value string `xml:",chardata"`
	Type  string `xml:"rad_type,attr"`
}

type KanjiDicEntryMisc struct {
	Grade       string                 `xml:"grade"`
	StrokeCount []string               `xml:"stroke_count"`
	Variant     []KanjiDicEntryVariant `xml:"variant"`
	Freq        string                 `xml:"freq"`
	RadName     string                 `xml:"rad_name"`
	Jlpt        string                 `xml:"jlpt"`
}

type KanjiDicEntryVariant struct {
	Value string `xml:",chardata"`
	Type  string `xml:"var_type"`
}

type KanjiDicEntryDicNumber struct {
	Value  string `xml:",chardata"`
	Type   string `xml:"dr_type,attr"`
	Volume string `xml:"m_vol,attr"`
	Page   string `xml:"m_page,attr"`
}

type KanjiDicEntryQueryCode struct {
	Value             string `xml:",chardata"`
	Type              string `xml:"qc_type,attr"`
	Misclassification string `xml:"skip_misclass,attr"`
}

type KanjiDicEntryReadingMeaning struct {
	Reading []KanjiDicEntryReading `xml:"rmgroup>reading"`
	Meaning []KanjiDicEntryMeaning `xml:"rmgroup>meaning"`
	Nanori  []string               `xml:"nanori"`
}

type KanjiDicEntryReading struct {
	Value  string `xml:",chardata"`
	Type   string `xml:"r_type,attr"`
	OnType string `xml:"on_type"`
	Status string `xml:"r_status"`
}

type KanjiDicEntryMeaning struct {
	Value string `xml:",chardata"`
	Lang  string `xml:"m_lang,attr"`
}
