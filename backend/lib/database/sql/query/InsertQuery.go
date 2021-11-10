package query

import (
	"gochat/lib/database/sql/query/metadata"
	"strings"
)

func MakeInsertQueryChain(o interface{}) InsertQueryChain {
	return InsertQueryChain{
		tablename: metadata.GetModelTablename(o),
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
	mmeta := metadata.MakeModelMetadata(o)
	mmeta = metadata.RemovePrimary(mmeta)

	iqc.keys = mmeta.GetFieldnames()
	iqc.values = [][]interface{}{mmeta.GetValues()}
	return iqc
}

func (iqc InsertQueryChain) InsertManyModel(arr ...interface{}) InsertQueryChain {
	iqc.values = [][]interface{}{}
	for i, o := range arr {
		mmeta := metadata.MakeModelMetadata(o)
		mmeta = metadata.RemovePrimary(mmeta)

		if i == 1 {
			iqc.keys = mmeta.GetFieldnames()
		}
		iqc.values = append(iqc.values, mmeta.GetValues())
	}
	return iqc
}
