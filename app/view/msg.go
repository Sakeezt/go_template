package view

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type language struct {
	TH string
	EN string
}

var MsgData = map[string]language{
	"attempt_error":       {TH: "พบข้อผิดพลาดในการทำงาน", EN: "Attempt error"},
	"unauthorized":        {TH: "ไม่มีสิทธิ์เข้าใช้งาน", EN: "Unauthorized"},
	"not_found":           {TH: "ไม่พบข้อมูล", EN: "Data not found"},
	"validate_fail":       {TH: "ตรวจสอบพบข้อมูลไม่ถูกต้อง", EN: "Validate failed"},
	"get_data_success":    {TH: "เรียกดูข้อมูลสำเร็จ", EN: "Get data success"},
	"create_data_success": {TH: "สร้างข้อมูลสำเร็จ", EN: "Create data success"},
	"update_data_success": {TH: "อัพเดทข้อมูลสำเร็จ", EN: "Update data success"},
	"delete_data_success": {TH: "ลบข้อมูลสำเร็จ", EN: "Delete data success"},
	"login_success":       {TH: "เข้าสู่ระบบสำเร็จ", EN: "Login success"},
	"login_fail":          {TH: "ชื่อผู้ใช้งานหรือรหัสผ่านไม่ถูกต้อง", EN: "Username or password is incorrect"},
	"logout_success":      {TH: "ออกจากระบบสำเร็จ", EN: "Logout success"},
}

func getHeaderLanguage(c *gin.Context) (lang string) {
	if r := c.Request; r != nil {
		val := r.Header.Get("Accept-Language")
		if val == "" {
			return "en"
		}
		return strings.ToLower(val)
	}
	return "en"
}

func get(c *gin.Context, key string) (msg string) {
	var text string
	if val, ok := MsgData[key]; ok {
		switch getHeaderLanguage(c) {
		case "th":
			return val.TH
		case "en":
			return val.EN
		default:
			return val.TH
		}
	}
	return text
}

func MsgGetDataSuccess(c *gin.Context) string {
	return get(c, "get_data_success")
}

func MsgCreateDataSuccess(c *gin.Context) string {
	return get(c, "create_data_success")
}

func MsgUpdateDataSuccess(c *gin.Context) string {
	return get(c, "update_data_success")
}

func MsgDeleteDataSuccess(c *gin.Context) string {
	return get(c, "delete_data_success")
}

func MsgAttemptError(c *gin.Context) string {
	return get(c, "attempt_error")
}

func MsgNotFoundData(c *gin.Context) string {
	return get(c, "not_found")
}
