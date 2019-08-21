package error

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (argPtr *argError) Error() string {
	return fmt.Sprintf("%d - %s", argPtr.arg, argPtr.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it."}
	}
	return arg + 3, nil
}

func PrintValue() {
	for _, i := range []int{4, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 success: ", r)
		}
	}

	for _, i := range []int{4, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed: ", e)
		} else {
			fmt.Println("f2 success: ", r)
		}
	}

	_, ae := f2(42)
	if ae, ok := ae.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
