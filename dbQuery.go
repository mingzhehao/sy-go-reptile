package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	UserAge  int    `json:"user_age"`
}

type LaGouUser struct {
	positionType string `json:"positionType"`
	positionName string `json:"positionName"`
	workYear     string `json:"workYear"`
	salary       string `json:"salary"`
	city         string `json:"city"`
}

type BaseJsonBean struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	UserAge  int    `json:"user_age"`
}

func NewBaseJsonBean() *BaseJsonBean {
	return &BaseJsonBean{}
}

/**
 * select query
 * @params user_id int
 * @return User json
 */
func GetUserByIdMethodOne(user_id int) *User {
	db, _ := GetDbConnection()
	defer db.Close()
	row := db.QueryRow("SELECT `user_id`, `user_name`,`user_age` FROM `user` WHERE user_id=?", user_id)
	user := new(User)
	row.Scan(&user.UserId, &user.UserName, &user.UserAge)
	return user
}

func GetUserByIdMethodTwo(user_id int) []byte {
	db, _ := GetDbConnection()
	defer db.Close()
	stmtOut, err := db.Prepare("SELECT `user_id`, `user_name`,`user_age` FROM `user` WHERE user_id=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	var (
		UserId   int    // we "scan" the result in here
		UserName string // we "scan" the result in here
		UserAge  int    // we "scan" the result in here
	)

	// Query the square-number of 1
	err = stmtOut.QueryRow(user_id).Scan(&UserId, &UserName, &UserAge) // WHERE number = 1
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	result := NewBaseJsonBean()
	result.UserId = UserId
	result.UserName = UserName
	result.UserAge = UserAge
	bytes, _ := json.Marshal(result)
	return bytes

}

func UpdateUserInfoByUserId(user_id, user_age, user_sex int) (e error) {
	db, err := GetDbConnection()
	defer db.Close()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`UPDATE user SET user_age=?,user_sex=? WHERE user_id=?`)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(user_age, user_sex, user_id)
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println(row)
	return nil
}

//删除数据
func DeleteUserInfoByUserId(user_id int) (e error) {
	db, err := GetDbConnection()
	defer db.Close()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`DELETE FROM user WHERE user_id=?`)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(user_id)
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println(row)
	return nil
}

/**
 * 数据插入
	positionType string `json:"positionType"`
	positionName string `json:"positionName"`
	workYear     string `json:"workYear"`
	salary       string `json:"salary"`
	city         string `json:"city"`
*/
func InsertLagouUser(positionType, positionName, workYear, salary, city, language string) (e error) {
	db, err := GetDbConnection()
	defer db.Close()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`insert into lagou_user values(null,?,?,?,?,?,? )`)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(positionType, positionName, workYear, salary, city, language)
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println(row)
	return nil
}
