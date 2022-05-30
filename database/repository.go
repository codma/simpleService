package database

import (
	"fmt"
	"simpleService/model"

	"github.com/gin-gonic/gin"
)

//스토어 목록 조회
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

//스토어 추가
func (d *Database) AddStore(c *gin.Context) (bool, error) {

	storeInfo := &model.AddOndStore{}
	err := c.BindJSON(storeInfo)
	if err != nil {
		return false, err
	}

	addOne, err := d.db.Prepare("INSERT INTO stores (storeName, planCode, domain) VALUES (?, ?, ?);")
	if err != nil {
		return false, err
	}

	res, err := addOne.Exec(storeInfo.StoreName, storeInfo.PlanCode, storeInfo.Domain)
	if err != nil {
		return false, err
	}

	rowCount, err := res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

	return true, err
}

//스토어 도메인 변경
func (d *Database) UpdateDomain(c *gin.Context) (bool, error) {
	//storeName에 해당하는 도메인주소를 업데이트 한다.
	storeInfo := &model.AddOndStore{}
	err := c.BindJSON(storeInfo)
	if err != nil {
		return false, err
	}

	// Modify some data in table.
	rows, err := d.db.Exec("UPDATE stores SET domain = ? WHERE storeName = ?", storeInfo.Domain, storeInfo.StoreName)
	if err != nil {
		return false, err
	}

	rowCount, err := rows.RowsAffected()
	fmt.Printf("Updated %d row(s) of data.\n", rowCount)
	fmt.Println("Done.")

	return true, err
}

//스토어 삭제
func (d *Database) DeleteStore(c *gin.Context) (bool, error) {
	//storeName에 해당하는 열을 삭제한다.
	storeInfo := &model.AddOndStore{}
	err := c.BindJSON(storeInfo)
	if err != nil {
		return false, err
	}

	// Modify some data in table.
	rows, err := d.db.Exec("DELETE FROM stores WHERE StoreName = ?", storeInfo.StoreName)
	if err != nil {
		return false, err
	}

	rowCount, err := rows.RowsAffected()
	fmt.Printf("Deleted %d row(s) of data.\n", rowCount)
	fmt.Println("Done.")

	return true, err
}
