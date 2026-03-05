package handlers

/*
  Gin 常用 API（白話）：
  - c *gin.Context：這一筆請求的資訊包，拿參數、body、回傳都透過 c
  - c.Param("id")：URL 路徑參數，例如 /api/partners/123 的 123
  - c.ShouldBindJSON(&input)：把 request body 的 JSON 轉成 struct，並依 binding 驗證
  - c.JSON(狀態碼, 資料)：回傳狀態碼與 JSON 給客戶端
  - gin.H：快速建 JSON 物件，如 gin.H{"error": "not found"}
*/

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"database/sql"
	"go-api-practice-2/database"
	"go-api-practice-2/models"

	"github.com/gin-gonic/gin"
)

var ErrNotFound = errors.New("not found")

func respondError(c *gin.Context, err error) {
	if err == nil {
		return
	}
	if errors.Is(err, ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func parseID(c *gin.Context, param string) (int, bool) {
	id, err := strconv.Atoi(c.Param(param))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return 0, false
	}
	return id, true
}

// GetPartners GET /api/partners - 取得人魚漢頓的所有夥伴
func GetPartners(c *gin.Context) {
	// 請實作：查 partners 表，回傳列表（使用 database.DB）
	query := "SELECT id, name, intro, created_at, updated_at FROM partners"

	rows, err := database.DB.Query(query)
	if err != nil {
		respondError(c, fmt.Errorf("查詢 partner 失敗: %w", err))
		return
	}

	defer rows.Close()

	var partners []models.Partner
	for rows.Next() {
		var p models.Partner
		if err := rows.Scan(&p.ID, &p.Name, &p.Intro, &p.CreatedAt, &p.UpdatedAt); err != nil {
			respondError(c, fmt.Errorf("讀取資料失敗: %w", err))
			return
		}
		partners = append(partners, p)
	}

	c.JSON(http.StatusOK, partners)
}

// GetPartnerByID GET /api/partners/:id - 依 ID 取得單一夥伴
func GetPartnerByID(c *gin.Context) {
	id, ok := parseID(c, "id")
	//先確認id存在 才傳進去sql查詢
	if !ok {
		return
	}
	//SELECT 欄位名稱 FROM 表格名稱;
	query := "SELECT id, name, intro, created_at, updated_at FROM partners WHERE id = $1"

	var p models.Partner
	err := database.DB.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Intro, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondError(c, ErrNotFound)
			return
		}
		respondError(c, fmt.Errorf("查詢單一夥伴失敗: %w", err))
		return
	}

	// 請實作：依 id 查單筆，找不到回傳 ErrNotFound
	c.JSON(http.StatusOK, p)
}

// CreatePartner POST /api/partners - 新增漢頓的夥伴
func CreatePartner(c *gin.Context) {
	var input models.Partner
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//INSERT INTO 表格名稱 (欄位1, 欄位2, 欄位3) VALUES (資料1, 資料2, 資料3);
	query := "INSERT INTO partners (name, intro) VALUES ($1, $2) RETURNING id, name, intro, created_at, updated_at"
	err := database.DB.QueryRow(query, input.Name, input.Intro).Scan(&input.ID, &input.Name, &input.Intro, &input.CreatedAt, &input.UpdatedAt)
	if err != nil {
		respondError(c, fmt.Errorf("新增夥伴失敗: %w", err))
		return
	}

	// 請實作：INSERT 進 partners，回傳新增的資料
	c.JSON(http.StatusOK, input)
}

// UpdatePartner PUT /api/partners/:id - 更新夥伴資料
func UpdatePartner(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	var input models.Partner
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//UPDATE 表格名稱 SET 欄位名稱 = 新的資料 WHERE 條件;
	query := "UPDATE partners SET name = $1, intro = $2 WHERE id = $3 RETURNING id, name, intro, created_at, updated_at"
	err := database.DB.QueryRow(query, input.Name, input.Intro, id).Scan(&input.ID, &input.Name, &input.Intro, &input.CreatedAt, &input.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondError(c, ErrNotFound)
			return
		}
		respondError(c, fmt.Errorf("更新單一夥伴失敗: %w", err))
		return
	}

	// 請實作：UPDATE 該 id，回傳更新後資料；不存在回傳 ErrNotFound
	c.JSON(http.StatusOK, input)
}

// DeletePartner DELETE /api/partners/:id - 刪除夥伴
func DeletePartner(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}

	//DELETE FROM 表格名稱 WHERE 條件;
	query := "DELETE FROM partners WHERE id = $1"
	result, err := database.DB.Exec(query, id)
	if err != nil {
		respondError(c, fmt.Errorf("刪除夥伴失敗: %w", err))
		return
	}

	rowsAffecte, err := result.RowsAffected()
	if err != nil {
		respondError(c, fmt.Errorf("取得受影響行數失敗: %w", err))
		return
	}

	// 請實作：DELETE 該 id，不存在回傳 ErrNotFound
	if rowsAffecte == 0 {
		respondError(c, ErrNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "刪除成功"})
}
