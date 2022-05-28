package database

import (
	"fmt"
	"simpleService/model"
)

func (d *Database) FindStoreList() []model.DbColumn {
	var results []model.DbColumn
	var result model.DbColumn

	//쿼리작성
	rows, err := d.db.Query("SELECT storeId, storeName, planCode, domain, activate from stores;")
	checkError(err)
	defer rows.Close()

	fmt.Println("Reading data:")
	for rows.Next() {
		err := rows.Scan(&result.StoreId, &result.StoreName, &result.PlanCode, &result.Domain, &result.Activate)
		checkError(err)
		results = append(results, model.DbColumn(result))
	}
	err = rows.Err()
	checkError(err)
	fmt.Println("Done.")

	defer d.db.Close()

	return results
}
