package metadata

import (
	"gochat/lib/util"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
)

func GetModelTablename(md interface{}) string {
	return strings.ToLower(reflect.TypeOf(md).Name()) + "s"
}

func MakeModelMetadata(md interface{}) ModelMetadata {
	mtname := GetModelTablename(md)
	mMetadata := ModelMetadata{Tablename: mtname}

	tof := reflect.TypeOf(md)
	vof := reflect.ValueOf(md)
	flen := tof.NumField()

	fMetadata := []FieldMetadata{}

	for i := 0; i < flen; i++ {
		if vof.Field(i).CanInterface() {
			fMetadata = append(fMetadata, FieldMetadata{
				Name:  strings.ToLower(tof.Field(i).Name),
				Value: vof.Field(i).Interface(),
				Tags:  strings.Split(strings.ToLower(tof.Field(i).Tag.Get("sqldb")), ","),
			})
		}
	}

	mMetadata.Fields = fMetadata
	return mMetadata
}

type ModelMetadata struct {
	Tablename string
	Fields    []FieldMetadata
}

func RemovePrimary(m ModelMetadata) ModelMetadata {
	newFields := make([]FieldMetadata, 0, len(m.Fields))
	for _, field := range m.Fields {
		if !util.ArrStringContains(field.Tags, TAG_PRIMARY_KEY) {
			newFields = append(newFields, field)
		}
	}
	m.Fields = newFields
	return m
}

func (m *ModelMetadata) GetFieldnames() []string {
	flen := len(m.Fields)
	fieldnames := make([]string, flen)
	for i, field := range m.Fields {
		fieldnames[i] = field.Name
	}
	return fieldnames
}

func (m *ModelMetadata) GetValues() []interface{} {
	flen := len(m.Fields)
	values := make([]interface{}, flen)
	for i, field := range m.Fields {
		values[i] = field.Value
	}
	return values
}

type FieldMetadata struct {
	Name  string
	Tags  []string
	Value interface{}
}

func (f *FieldMetadata) GetSQLType() string {
	for _, tag := range f.Tags {
		tval := strings.Split(tag, "-")
		if tval[0] == MOD_TAG_VARCHAR {
			return "varchar(" + tval[1] + ")"
		} else if tval[0] == MOD_TAG_CHAR {
			return "char(" + tval[1] + ")"
		}
	}
	return f.GetValueType()
}

func (f *FieldMetadata) GetValueType() string {
	switch f.Value.(type) {
	case int, int64:
		return "INT"
	case int16:
		return "MEDIUMINT"
	case string:
		return "varchar(255)"
	default:
		logrus.Panic("Type Not Supported ", reflect.TypeOf(f.Value))
		return ""
	}
}

func (f *FieldMetadata) GetTagConstraints() string {
	tags := make([]string, len(f.Tags))
	for i := range f.Tags {
		tval := strings.Split(f.Tags[i], "-")
		tg := tagDef[tval[0]]

		switch tval[0] {
		case TAG_PRIMARY_KEY:
			tg = "AUTO_INCREMENT " + tg
		case TAG_FOREIGN_KEY:
			tg += " " + tval[1] + " ON DELETE CASCADE ON UPDATE CASCADE"
		case TAG_DEFAULT:
			tg += " " + tval[1]
		}

		tags[i] = tg
	}
	return strings.Join(tags, " ")
}
