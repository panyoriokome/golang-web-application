package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	// 設定した環境変数が適用されているか
	wantPort := 3333
	t.Setenv("PORT", fmt.Sprint(wantPort))

	got, err := New()
	if err != nil {
		t.Fatalf("cannnot create config: %v", err)
	}
	if got.Port != wantPort {
		t.Errorf("want %d, but got %d", wantPort, got.Port)
	}

	// デフォルト値が設定されているか
	wantEnv := "dev"
	if got.Env != wantEnv {
		t.Errorf("want %s, but got %s", wantEnv, got.Env)
	}
}
