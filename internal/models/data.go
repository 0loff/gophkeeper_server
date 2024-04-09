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
	Username []byte
	Password []byte
	Metainfo string
}

type CardsdataEntry struct {
	ID       int
	Pan      []byte
	Expiry   []byte
	Holder   []byte
	Metainfo string
}

type BindataEntry struct {
	ID       int
	Binary   []byte
	Metainfo string
}
