package util

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Printf("%+v", errors.WithStack(err))
		os.Exit(1)
	}
}
