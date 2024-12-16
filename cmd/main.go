package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error

	// 환경 변수로부터 DB 연결 정보 가져오기
	// 환경 변수는 deployment.yaml에 작성됨
	dbHost := os.Getenv("DB_HOST")         // MySQL 호스트 주소
	dbUser := os.Getenv("DB_USER")         // MySQL 사용자 이름
	dbPassword := os.Getenv("DB_PASSWORD") // MySQL 비밀번호
	dbName := os.Getenv("DB_NAME")         // MySQL 데이터베이스 이름

	// DB 연결
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("🚨 DB 연결 실패: %v", err)
	}
	defer db.Close()

	// 핸들러 등록
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/items", crudHandler)

	// 서버 시작
	port := ":8080"
	fmt.Printf("✅ API Server 시작: 포트 %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// Health check 핸들러
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// CRUD 핸들러
func crudHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getItems(w, r)
	case http.MethodPost:
		createItem(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("🚨 허용되지 않은 메서드"))
	}
}

func getItems(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name FROM items")
	if err != nil {
		http.Error(w, "🚨 데이터 조회 실패", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []map[string]interface{}
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			http.Error(w, "Failed to scan item", http.StatusInternalServerError)
			return
		}
		items = append(items, map[string]interface{}{"id": id, "name": name})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var item struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "🚨 유효하지 않은 요청", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO items (name) VALUES (?)", item.Name)
	if err != nil {
		http.Error(w, "🚨 데이터 삽입 실패", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("✅ 데이터 삽입"))
}
