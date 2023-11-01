package template

const (
	Mod = 1 // one master with one slave
)

type Template struct {
	Column1 int    `gorm:"column:column_1;type:int(20) not null;primary_key;comment:column one" json:"column_1,string"`
	Column2 string `gorm:"column:column_2;type:varchar(64) not null;default:'value_of_column_2';comment:column two" json:"column_2"`
}
