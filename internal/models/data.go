package models

type Userdata struct {
	Textdata []TextdataEntry
}

type TextdataEntry struct {
	ID       int
	Text     string
	Metainfo string
}

type CredsdataEntry struct {
	ID       int
	Username string
	Password string
	Metainfo string
}
