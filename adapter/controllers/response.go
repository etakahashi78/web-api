package controllers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

// ResponseJSON 200系で使用する共通関数
func ResponseJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.WriteHeader(statusCode)
	setJSONContentType(w)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		slog.Error("ResponseJSON failed.", "err", err)
		ResponseError(w, http.StatusInternalServerError, err)
	}
}

// ResponseError 400, 500エラーで使用する共通関数
func ResponseError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	setJSONContentType(w)

	res := make(map[string]string)
	res["Code"] = fmt.Sprintf("%d", statusCode)
	res["Message"] = fmt.Sprintf("%d %s: %v", statusCode, http.StatusText(statusCode), err)
	b, er := json.Marshal(res)
	if er != nil {
		slog.Error("json.Marshal in ResponseError failed.", "er", er)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	_, er = w.Write(b)
	if er != nil {
		slog.Error("w.Write in ResponseError failed.", "er", er)
	}
}

func setJSONContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}
