package types

type Kanji struct {
	Character string
	Grade     int
	Stroke    int
	Frequency int
	JLPT      int
	Meaning   []string
	Onyomi    []string
	Kunyomi   []string
}
