/**
* @Author:wjp
* @Date:2020/12/2 1:13 下午
 */
package main

import "database/sql"

type User struct {
	UserName string `json:"user_name"`
}

func UserInfo(id int) (User,error)  {

	return User{}, sql.ErrNoRows
}

