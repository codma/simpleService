package database

import (
	"fmt"
	"log"
	"simpleService/model"

	"github.com/gin-gonic/gin"
)

//스토어 목록 조회
func FindStoreList() []model.Store {
	var results []model.Store
	var result model.Store

	//쿼리작성
	rows, err := Db.Query("SELECT storeId, storeName, planCode, domain, activate from store;")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	fmt.Println("Reading data:")
	for rows.Next() {
		err := rows.Scan(&result.StoreId, &result.StoreName, &result.PlanCode, &result.Domain, &result.Activate)
		if err != nil {
			log.Println(err)
		}
		results = append(results, model.Store(result))
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Done.")

	defer Db.Close()

	return results
}

//스토어 추가
func AddStore(storeInfo model.AddOneStore) (bool, error) {

	addOne, err := Db.Prepare("INSERT INTO store (storeName, planCode, domain) VALUES (?, ?, ?);")
	if err != nil {
		return false, err
	}

	res, err := addOne.Exec(storeInfo.StoreName, storeInfo.PlanCode, storeInfo.Domain)
	if err != nil {
		return false, err
	}

	//확인용
	rowCount, err := res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

	return true, err
}

//스토어 도메인 변경
func UpdateDomain(c *gin.Context) (bool, error) {
	//storeName에 해당하는 도메인주소를 업데이트 한다.
	storeInfo := &model.AddOneStore{}
	err := c.BindJSON(storeInfo)
	if err != nil {
		return false, err
	}

	// Modify some data in table.
	rows, err := Db.Exec("UPDATE store SET domain = ? WHERE storeName = ?", storeInfo.Domain, storeInfo.StoreName)
	if err != nil {
		return false, err
	}

	//확인용
	rowCount, err := rows.RowsAffected()
	fmt.Printf("Updated %d row(s) of data.\n", rowCount)
	fmt.Println("Done.")

	return true, err
}

//스토어 삭제
func DeleteStore(c *gin.Context) (bool, error) {
	//storeName에 해당하는 열을 삭제한다.
	storeInfo := &model.AddOneStore{}
	err := c.BindJSON(storeInfo)
	if err != nil {
		return false, err
	}

	// Modify some data in table.
	rows, err := Db.Exec("DELETE FROM store WHERE StoreName = ?", storeInfo.StoreName)
	if err != nil {
		return false, err
	}

	//확인용
	rowCount, err := rows.RowsAffected()
	fmt.Printf("Deleted %d row(s) of data.\n", rowCount)
	fmt.Println("Done.")

	return true, err
}
