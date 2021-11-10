package query

import (
	"strings"
)

func MakeInsertQueryChain(o interface{}) InsertQueryChain {
	return InsertQueryChain{
		tablename: getModelTablename(o),
	}
}

type InsertQueryChain struct {
	tablename string
	keys      []string
	values    [][]interface{}
}

func (iqc InsertQueryChain) ToString() string {
	b := strings.Builder{}
	b.WriteString("INSERT INTO ")
	b.WriteString(iqc.tablename)
	b.WriteString(" (")
	b.WriteString(strings.Join(iqc.keys, ", "))
	b.WriteString(")")
	b.WriteString(" VALUES ")
	b.WriteString(iqc.valuesToString())
	return b.String()
}

func (iqc *InsertQueryChain) GetValues() []interface{} {
	vals := make([]interface{}, 0, len(iqc.values)*len(iqc.values[0]))
	for _, pval := range iqc.values {
		vals = append(vals, pval...)
	}
	return vals
}

func (iqc *InsertQueryChain) valuesToString() string {
	result := make([]string, len(iqc.values))
	for i, values := range iqc.values {
		vstrArr := make([]string, len(values))
		for j := range values {
			vstrArr[j] = "?"
		}
		result[i] = "(" + strings.Join(vstrArr, ", ") + ")"
	}
	return strings.Join(result, ", ")
}

func (iqc InsertQueryChain) InsertModel(o interface{}) InsertQueryChain {
	mmeta := MakeModelMetadata(o).removePrimary()
	iqc.keys = mmeta.getFieldnames()
	iqc.values = [][]interface{}{mmeta.getValues()}
	return iqc
}

func (iqc InsertQueryChain) InsertManyModel(arr ...interface{}) InsertQueryChain {
	iqc.values = [][]interface{}{}
	for i, o := range arr {
		mmeta := MakeModelMetadata(o).removePrimary()
		if i == 1 {
			iqc.keys = mmeta.getFieldnames()
		}
		iqc.values = append(iqc.values, mmeta.getValues())
	}
	return iqc
}
