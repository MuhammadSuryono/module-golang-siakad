package common

import (
	"github.com/kr/pretty"
	"testing"
)

func TestResponse(t *testing.T) {
	res := CommonResponse{
		Code: 200,
		IsSuccess: false,
		Data: nil,
		Message: "Gagal",
	}

	pretty.Println(res)
}