package dbobj_test

import (
	"dbobj"
	"encoding/json"
	"fmt"
	"testing"
	)

type user struct {
	User_id             string `json:"user_id"`
	User_name           string `json:"user_name"`
	User_status_desc    string `json:"status_desc"`
	User_create_date    string `json:"create_date"`
	User_owner          string `json:"create_user"`
	User_email          string `json:"user_email"`
	User_phone          string `json:"user_phone"`
	Org_unit_id         string `json:"org_unit_id"`
	Org_unit_desc       string `json:"org_unit_desc"`
	Domain_id           string `json:"domain_id"`
	Domain_name         string `json:"domain_name"`
	User_maintance_date string `json:"modify_date"`
	User_maintance_user string `json:"modify_user"`
	User_status_id      string `json:"status_cd"`
}

func TestQuery(t *testing.T) {
	resInfo, _ := dbobj.MysqlEngin.Query("select * from sys_user_info user")
	rec, _ := json.Marshal(resInfo)
	var receiveInfos []user
	if err := json.Unmarshal(rec, &receiveInfos); err != nil {
		fmt.Println("---Unmarshal---receiveInfo----------->", err)
	}
	fmt.Println(len(receiveInfos))
	for _, val := range receiveInfos {
		fmt.Println(string(val.User_name))
	}
}