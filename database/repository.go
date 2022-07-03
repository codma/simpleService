package database

import (
	"fmt"
	"simpleService/model"
)

//스토어 목록 조회
func FindStoreList() (any, error) {
	var results []model.Store
	var result model.Store

	//쿼리작성
	rows, err := Db.Query("SELECT storeId, storeName, planCode, domain, activate from store;")
	if err != nil {
		return "쿼리 실행 실패", err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result.StoreId, &result.StoreName, &result.PlanCode, &result.Domain, &result.Activate)
		if err != nil {
			return "데이터 읽기 실패", err
		}
		results = append(results, model.Store(result))
	}

	defer Db.Close()

	return results, nil
}

//스토어 추가
func AddStore(storeInfo model.StoreRequest) (bool, error) {

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
func UpdateDomain(storeInfo model.StoreRequest) (string, error) {
	//storeName에 해당하는 도메인주소를 업데이트 한다.

	_, err := Db.Exec("UPDATE store SET domain = ? WHERE storeName = ?", storeInfo.Domain, storeInfo.StoreName)
	if err != nil {
		return "업데이트 실패", err
	}

	return "업데이트 성공", err
}

//스토어 삭제
func DeleteStore(storeInfo model.StoreRequest) (string, error) {
	//storeName에 해당하는 열을 삭제한다.

	_, err := Db.Exec("DELETE FROM store WHERE StoreName = ?", storeInfo.StoreName)
	if err != nil {
		return "삭제 실패", err
	}

	return "삭제 성공", err
}
