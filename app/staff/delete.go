package staff

import (
	"go-template/app/view"
	"go-template/service/staff/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// Delete godoc
// @Tags         Staffs
// @Summary      Delete contents of a staff
// @Description  Delete staff with a given staff ID
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string  false  "Language"  default(en)
// @Param        id               path      string  true   "Staff ID"  default(123456789012345678)
// @Success      200              {object}  view.SuccessDeleteResp{data=object}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrorDeleteItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrorDeleteItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrorDeleteItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrorDeleteItem}
// @Router       /staffs/{id} [delete]
func (ctrl *Controller) Delete(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.staff.Delete")
	defer sp.Finish()

	inp := &inout.StaffDeleteInput{
		ID: c.Param("id"),
	}

	if err := ctrl.service.Delete(ctx, inp); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeDeleteSuccessResp(c)
}
