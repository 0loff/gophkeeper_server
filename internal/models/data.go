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

type CardsdataEntry struct {
	ID       int
	Pan      string
	Expiry   string
	Holder   string
	Metainfo string
}

type BindataEntry struct {
	ID       int
	Binary   []byte
	Metainfo string
}
