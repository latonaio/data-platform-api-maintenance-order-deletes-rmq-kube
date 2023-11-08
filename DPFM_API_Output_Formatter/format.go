package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*Header, error) {
	defer rows.Close()
	header := Header{}
	i := 0

	for rows.Next() {
		i++
		err := rows.Scan(
			&header.MaintenanceOrder,
			&header.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToItem(rows *sql.Rows) (*[]ObjectListItem, error) {
	defer rows.Close()
	items := make([]ObjectListItem, 0)
	i := 0

	for rows.Next() {
		i++
		item := ObjectListItem{}
		err := rows.Scan(
			&item.MaintenanceOrder,
			&item.MaintenanceObjectListItem,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &items, err
		}

		items = append(items, item)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &items, nil
	}

	return &items, nil
}

