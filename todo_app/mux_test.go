package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMux(t *testing.T) {
	// httptestを使ってServeHTTPに渡すためのモックを生成
	// NewRecorder関数を使うとResponseWriterインタフェースを満たす
	// *ResponseRecorder型の値を取得できる
	w := httptest.NewRecorder()
	// NewRequest関数を呼んで*http.Request型の値を生成する
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	sut := NewMux()
	sut.ServeHTTP(w, r)
	// Resultメソッドを呼ぶことでクライアントが受け取るレスポンス内容が含まれる
	// http.Response型の値を取得できる
	resp := w.Result()
	t.Cleanup(func() { _ = resp.Body.Close() })

	// ここまでできたらあとはレスポンスの内容を検証していく
	if resp.StatusCode != http.StatusOK {
		t.Error("expect status code 200, but", resp.StatusCode)
	}
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	expect := `{"status": "ok"}`
	if string(got) != expect {
		t.Errorf("expect %q, but got %q", expect, got)
	}
}
