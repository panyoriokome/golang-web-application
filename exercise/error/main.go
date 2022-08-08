package main

import (
	"errors"
	"fmt"
)

var ErrErrosNew = errors.New("errors.New")
var ErrFmtErrorf = fmt.Errorf("fmt.Errorf")

func main() {
	err := errorTestErrorNew()
	if err != nil {
		fmt.Println("Error errors.New")
	}

	err = errorTestFmtErrorf()
	if err != nil {
		fmt.Println("Error fmt.Errorf")
	}

	err = errorChain()
	if err != nil {
		fmt.Printf("errorChainの呼び出しでエラー: %v", err)
	}

	err = errorChainWrap()
	if err != nil {
		fmt.Printf("errorChainの呼び出しでエラー: %v", err)
	}

	// errors.Isを使ってラップしたエラーオブジェクトの型を判別
	err = errorChainWrap()
	if err != nil {
		if errors.Is(err, ErrErrosNew) {
			fmt.Printf("ErrErrosNewのエラーを識別 %v", err)
		}
	}
}

func errorTestErrorNew() error {
	return ErrErrosNew
}

func errorTestFmtErrorf() error {
	return ErrFmtErrorf
}

func errorChain() error {
	err := errorTestErrorNew()
	if err != nil {
		return fmt.Errorf("errorChain: %v", err)
	}
	return nil
}

func errorChainWrap() error {
	err := errorTestErrorNew()
	if err != nil {
		// %wを使うことで元のエラーのオブジェクト情報を保持したままにする
		return fmt.Errorf("errorChain: %w", err)
	}
	return nil
}
