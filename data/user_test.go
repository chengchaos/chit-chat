package data

import "testing"

func TestCreate(t *testing.T) {
	user := &User{
		Name:     "chengchao2",
		Email:    "chengchaos2@outlook.com",
		Password: "Password",
	}
	err := user.Create()

	if err != nil {
		t.Error("err -> ", err.Error())
	} else {
		t.Log("id => ", user.Id)
	}

}

func TestUser_CreateThread(t *testing.T) {
	user := &User{
		Id: 1,
	}

	conv, err := user.CreateThread("测试1")
	if err != nil {
		t.Fatalf("err : %s\n", err.Error())
	} else {
		t.Logf("conv : %v\n", conv)
	}

}
