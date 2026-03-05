# Go API Practice － 人魚漢頓的夥伴

Golang 後端 API 練習專案，主題為**三麗鷗「人魚漢頓」（ハンギョドン）**，用「漢頓的夥伴」做 CRUD 練習。  
使用 Gin + PostgreSQL + plain SQL，**只保留一組路由、一層 handler（不含 service）**，DB 邏輯在 `handlers/partners.go` 內實作。

## 主題簡介

- **人魚漢頓**：三麗鷗 1985 年推出的角色，藍色圓滾滾的半魚人，個性迷糊、浪漫，夢想當英雄。
- **夥伴**：漢頓的朋友，例如最好的朋友「小百合」（章魚女孩）。

## 環境需求

- Go 1.21+
- PostgreSQL

## 設定

1. 建立 PostgreSQL 資料庫（例如名稱 `practice`）。
2. 複製並編輯 `.env`：
   ```bash
   cp .env.example .env
   ```
   - `PORT`：API 監聽 port（預設 `8080`）
   - `DATABASE_URL`：PostgreSQL 連線字串

## 執行

```bash
go mod tidy
go run main.go
```

伺服器預設在 `http://localhost:8080`。

## 專案結構

| 目錄／檔案 | 說明 |
|------------|------|
| `handlers/` | 一層搞定：Gin（參數、JSON、回傳）+ DB 查詢／寫入，練習處在 `partners.go` |
| `models/` | 請求／回應 struct（如 `Partner`） |
| `database/` | 連線與建表（`partners` 表） |
| `routes/` | 只註冊一組 `/api/partners` 路由 |
| `postman/` | Postman collection 匯入後可測試 |

## API 路由（一組練習用）

| 方法 | 路徑 | 說明 |
|------|------|------|
| GET | `/api/partners` | 取得所有夥伴 |
| GET | `/api/partners/:id` | 依 ID 取得單一夥伴 |
| POST | `/api/partners` | 新增夥伴（body：`name` 必填、`intro` 選填） |
| PUT | `/api/partners/:id` | 更新夥伴 |
| DELETE | `/api/partners/:id` | 刪除夥伴 |

## 練習重點

在 `handlers/partners.go` 中，每個函式都有「請實作」註解，要你補上對 `partners` 表的 DB 操作（使用 `database.DB` 的 Query / QueryRow / Exec、Scan、RowsAffected）。找不到資料時回傳 `ErrNotFound`，handler 會回 404。
