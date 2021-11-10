package query

import (
	"gochat/lib/util"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
)

func getModelTablename(md interface{}) string {
	return strings.ToLower(reflect.TypeOf(md).Name()) + "s"
}

func MakeModelMetadata(md interface{}) ModelMetadata {
	mtname := getModelTablename(md)
	mMetadata := ModelMetadata{Tablename: mtname}

	tof := reflect.TypeOf(md)
	vof := reflect.ValueOf(md)
	flen := tof.NumField()

	fMetadata := []fieldMetadata{}

	for i := 0; i < flen; i++ {
		if vof.Field(i).CanInterface() {
			fMetadata = append(fMetadata, fieldMetadata{
				name:  strings.ToLower(tof.Field(i).Name),
				value: vof.Field(i).Interface(),
				tags:  strings.Split(strings.ToLower(tof.Field(i).Tag.Get("sqldb")), ","),
			})
		}
	}

	mMetadata.Fields = fMetadata
	return mMetadata
}

type ModelMetadata struct {
	Tablename string
	Fields    []fieldMetadata
}

func (m ModelMetadata) removePrimary() ModelMetadata {
	newFields := make([]fieldMetadata, 0, len(m.Fields))
	for _, field := range m.Fields {
		if !util.ArrStringContains(field.tags, TAG_PRIMARY_KEY) {
			newFields = append(newFields, field)
		}
	}
	m.Fields = newFields
	return m
}

func (m *ModelMetadata) getFieldnames() []string {
	flen := len(m.Fields)
	fieldnames := make([]string, flen)
	for i, field := range m.Fields {
		fieldnames[i] = field.name
	}
	return fieldnames
}

func (m *ModelMetadata) getValues() []interface{} {
	flen := len(m.Fields)
	values := make([]interface{}, flen)
	for i, field := range m.Fields {
		values[i] = field.value
	}
	return values
}

type fieldMetadata struct {
	name  string
	tags  []string
	value interface{}
}

func (f *fieldMetadata) getSQLType() string {
	for _, tag := range f.tags {
		tval := strings.Split(tag, "-")
		if tval[0] == MOD_TAG_VARCHAR {
			return "varchar(" + tval[1] + ")"
		} else if tval[0] == MOD_TAG_CHAR {
			return "char(" + tval[1] + ")"
		}
	}
	return f.getValueType()
}

func (f *fieldMetadata) getValueType() string {
	switch f.value.(type) {
	case int, int64:
		return "INT"
	case int16:
		return "MEDIUMINT"
	case string:
		return "varchar(255)"
	default:
		logrus.Panic("Type Not Supported ", reflect.TypeOf(f.value))
		return ""
	}
}

func (f *fieldMetadata) getTagConstraints() string {
	tags := make([]string, len(f.tags))
	for i := range f.tags {
		tval := strings.Split(f.tags[i], "-")
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
