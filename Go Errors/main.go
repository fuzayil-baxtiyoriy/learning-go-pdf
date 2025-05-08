package main

import (
	"errors"
	"fmt"
	"os"
)

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d isn't an even number", i)
	}
	return i * 2, nil
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

// Wrapping Errors with defer

func doThing1(val int) (int, error) {
	panic("doThing1")

}

func doThing2(val string) (string, error) {
	panic("doThing2")
}

func doThing3(val1 int, val2 string) (string, error) {
	panic("doThing3")
}

func DoSomeThings(val1 int, val2 string) (string, error) {
	val3, err := doThing1(val1)
	if err != nil {
		return "", fmt.Errorf("in DoSomeThings: %w", err)
	}

	val4, err := doThing2(val2)
	if err != nil {
		return "", fmt.Errorf("in DoSomeThings: %w", err)
	}

	result, err := doThing3(val3, val4)
	if err != nil {
		return "", fmt.Errorf("in DoSomeThings: %w", err)
	}
	return result, nil
}

// You can simplify this code by using defer:

func DoSomeThingsWithDefer(val1 int, val2 string) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in DoSomeThings: %w", err)
		}
	}()
	val3, err := doThing1(val1)
	if err != nil {
		return "", err
	}
	val4, err := doThing2(val2)
	if err != nil {
		return "", err
	}
	return doThing3(val3, val4)
}

// panic and recover

func doPanic(msg string) {
	panic(msg)
}

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}

func main() {
	// numerator := 20
	// denominator := 0
	// remiander, mod, err := calcRemainderAndMod(numerator, denominator)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(remiander, mod)

	// Sentinel Errors
	// data := []byte("This is not a zip file")
	// notAzipFile := bytes.NewReader(data)
	// _, err := zip.NewReader(notAzipFile, int64(len(data)))
	// if err == zip.ErrFormat {
	// 	fmt.Println("Told you so")
	// }

	// Wrapping Errors
	// err := fileChecker("not_here.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
	// 		fmt.Println(wrappedErr)
	// 	}
	// }

	// Is and As
	// err := fileChecker("not_here.txt")
	// if err != nil {
	// 	if errors.Is(err, os.ErrNotExist) {
	// 		fmt.Println("That file doesn't exist")
	// 	}
	// }

	// doPanic(os.Args[0])

	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}

}
