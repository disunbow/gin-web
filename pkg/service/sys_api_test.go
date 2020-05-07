package service

import (
	"fmt"
	"go-shipment-api/pkg/global"
	"go-shipment-api/tests"
	"testing"
)

func TestGetRoleCategoryApisByRoleId(t *testing.T) {
	tests.InitTestEnv()

	fmt.Println(GetAllApiGroupByCategoryByRoleId(3))
	fmt.Println(GetAllApiGroupByCategoryByRoleId(1))

	defer global.Mysql.Close()
}