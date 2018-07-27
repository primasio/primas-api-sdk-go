package account

import (
	"testing"
	"time"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestCreateAccount(t *testing.T) {
	name := "kevin"
	abstract := "kevin test"
	avatar := "06354a3ee794cf210e0db0d6729710ee047227d679668de697431fdbd1232ffc"
	account_id := ""
	sub_account_id := ""
	created := time.Now().Unix()
	extra_hash := "test_value"

	createAccount, err := CreateAccount(name, abstract, avatar, account_id, sub_account_id, int(created), extra_hash)
	if err != nil {
		t.Errorf("CreateAccount error:%v", err.Error())
		return
	}

	if createAccount != nil {
		t.Logf("CreateAccount response value:%v", createAccount)
		if createAccount.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("CreateAccount response error:%v", createAccount.ResultMsg)
			return
		}
		if createAccount.Data != nil {
			t.Logf("CreateAccount response value:%v", createAccount.Data)
		} else {
			t.Logf("CreateAccount response value don't find ")
		}
	}
}

func TestCreateAccount_sub(t *testing.T) {
	name := "kevin"
	abstract := "kevin test"
	avatar := "06354a3ee794cf210e0db0d6729710ee047227d679668de697431fdbd1232ffc"
	account_id := "0b9e00eecaea5e3aec8ebfa3218ced7dbd7c23fd988f959e42302f13647ae37b"
	sub_account_id := "a_0001"
	created := time.Now().Unix()
	extra_hash := "test_value"

	createAccount, err := CreateAccount(name, abstract, avatar, account_id, sub_account_id, int(created), extra_hash)
	if err != nil {
		t.Errorf("CreateAccount error:%v", err.Error())
		return
	}

	if createAccount != nil {
		t.Logf("CreateAccount response value:%v", createAccount)
		if createAccount.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("CreateAccount response error:%v", createAccount.ResultMsg)
			return
		}
		if createAccount.Data != nil {
			t.Logf("CreateAccount response value:%v", createAccount.Data)
		} else {
			t.Logf("CreateAccount response value don't find ")
		}
	}
}
