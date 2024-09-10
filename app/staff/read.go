package staff

import (
	"go-template/app/view"
	"go-template/service/staff/inout"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// Read godoc
// @Tags         Staffs
// @Summary      Get a staff by staff ID
// @Description  Response a staff data with a given staff ID
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string  false  "Language"  default(en)
// @Param        id               path      string  true   "Staff ID"  default(123456789012345678)
// @Success      200              {object}  view.SuccessReadResp{data=inout.StaffView}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrorReadItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrorReadItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrorReadItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrorReadItem}
// @Router       /staffs/{id} [get]
func (ctrl *Controller) Read(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.staff.Read")
	defer sp.Finish()

	inp := &inout.StaffReadInput{
		ID: c.Param("id"),
	}

	staff, err := ctrl.service.Read(ctx, inp)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeReadSuccessResp(c, staff)
}
