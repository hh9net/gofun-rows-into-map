package db

import (
	"strconv"

	"github.com/jmoiron/sqlx"
)

//M ... shorthand notation for our map
type M map[string]interface{}

// ScanIntoMaps ... load contents of all columns into map. Return array of maps (one per each row)
func ScanIntoMaps(rows *sqlx.Rows) ([]M, error) {

	cols, _ := rows.Columns()
	colTypes, _ := rows.ColumnTypes()

	var result []M

	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		var m = M{}
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})

			var realVal interface{}
			var asStr string = string((*val).([]byte))

			switch colType := colTypes[i].DatabaseTypeName(); colType {
			case "INT":
				fallthrough
			case "BIGINT":
				realVal, _ = strconv.ParseInt(asStr, 10, 64)
			case "DECIMAL":
				realVal, _ = strconv.ParseFloat(asStr, 64)
			case "BOOL":
				realVal, _ = strconv.ParseBool(asStr)
			default:
				realVal = asStr
			}
			m[colName] = realVal
		}

		result = append(result, m)
	}
	return result, nil
}
