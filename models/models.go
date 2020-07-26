package models

func GetModels() []interface{} {
	return []interface{}{
		&Application{},
		&NameSpace{},
		&File{},
	}
}
