package storage

var Storage map[string][][][]string

func Initialize() {
	Storage = make(map[string][][][]string)
}
