package template

const (
	Mod = 3 // multiple DB with different which,  multiple table with same name
)

type Template struct {
	Field1 int    `bson:"field_1" json:"field_1"`
	Field2 string `bson:"field_2" json:"field_2"`
}
