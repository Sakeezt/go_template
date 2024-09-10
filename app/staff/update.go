package staff

import (
	"go-template/app/view"
	"go-template/service/staff/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// Update godoc
// @Tags         Staffs
// @Summary      Update contents of a staff
// @Description  Update staff with a given staff ID according to a given data
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string                  false  "Language"  default(en)
// @Param        id               path      string                  true   "Staff ID"  default(123456789012345678)
// @Param        Request          body      inout.StaffUpdateInput  true   "Input"
// @Success      200              {object}  view.SuccessUpdateResp{data=object}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrorUpdateItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrorUpdateItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrorUpdateItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrorUpdateItem}
// @Router       /staffs/{id} [put]
func (ctrl *Controller) Update(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.staff.Update")
	defer sp.Finish()

	inp := &inout.StaffUpdateInput{
		ID: c.Param("id"),
	}
	if err := c.ShouldBindJSON(inp); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	if err := ctrl.service.Update(ctx, inp); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeUpdateSuccessResp(c)
}
