package main

import (
	"fmt"
	"testing"
)

func TestGetGreeting(t *testing.T) {
	// 假定的某個人
	name := "Yu Yuan"

	// 呼叫取得問候語函式
	result := GetGreeting(name)

	// 預期結果
	exception := fmt.Sprintf("Hello, %s! Nice to meet you!", name)

	// 驗證結果是否一致
	if result != exception {
		t.Errorf("GetGreeting(%s) is Failed!\n Exception: %s", name, exception)
	}
}

func TestIsAdult(t *testing.T) {
	age := 20

	if !IsAdult(age) {
		t.Errorf("IsAdult(%d) is Failed!", age)
	}
}

func TestIsNotAdult(t *testing.T) {
	age := 16

	if IsAdult(age) {
		t.Errorf("IsAdult(%d) is Failed!", age)
	}
}

func TestIsAdultByTests(t *testing.T) {
	var tests = []struct {
		age       int
		exception bool
	}{
		{20, true},
		{16, false},
	}

	for _, test := range tests {
		if result := IsAdult(test.age); result != test.exception {
			t.Errorf("IsAdult(%d) is Failed!", test.age)
		}
	}
}
