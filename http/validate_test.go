package http

import (
	"fmt"
)

type Hh struct {
	Email string `json:"email" validate:"email"`
}

func ExampleValidateReq() {
	result := ValidateReq(&Hh{Email: "1234@qq.com"})
	fmt.Print(result)
	// Output: <nil>
}
