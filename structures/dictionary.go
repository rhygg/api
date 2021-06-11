package structures

type Phonetics struct {
	Text  string
	Audio string
}
type Definitions struct {
	Definition string
	Example    string
	Synonyms   []string
}
type Meanings struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  *Definitions `json:"definitions"`
}
type DictionarySearch struct {
	Phonetics *Phonetics `json:"phonetics"`
	Meanings  *Meanings  `json:"meanings"`
}
