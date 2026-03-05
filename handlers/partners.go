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

	"go-api-practice/models"

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
	if !ok {
		return
	}
	// 請實作：依 id 查單筆，找不到回傳 ErrNotFound
	respondError(c, fmt.Errorf("請實作：依 id 查詢單筆夥伴（id=%d）", id))
}

// CreatePartner POST /api/partners - 新增漢頓的夥伴
func CreatePartner(c *gin.Context) {
	var input models.Partner
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 請實作：INSERT 進 partners，回傳新增的資料
	respondError(c, fmt.Errorf("請實作：INSERT partners 並回傳（name=%s）", input.Name))
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
	// 請實作：UPDATE 該 id，回傳更新後資料；不存在回傳 ErrNotFound
	respondError(c, fmt.Errorf("請實作：UPDATE 夥伴並回傳（id=%d name=%s）", id, input.Name))
}

// DeletePartner DELETE /api/partners/:id - 刪除夥伴
func DeletePartner(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	// 請實作：DELETE 該 id，不存在回傳 ErrNotFound
	respondError(c, fmt.Errorf("請實作：DELETE 夥伴（id=%d）", id))
}
