package staff

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"go-template/app/view"
	"go-template/service/staff/inout"
)

// Create godoc
// @Tags         Staffs
// @Summary      Create a new staff
// @Description  A newly created staff, ID will be given in a Content-Location response header
// @Accept       json
// @Produce      json
// @Param        Accept-Language  header    string                  false  "Language"  default(en)
// @Param        Request          body      inout.StaffCreateInput  true   "Input"
// @Success      201              {object}  view.SuccessCreateResp{data=string}
// @Failure      400              {object}  view.Error400Resp{errors=[]view.ErrorCreateItem}
// @Failure      404              {object}  view.Error404Resp{errors=[]view.ErrorCreateItem}
// @Failure      422              {object}  view.Error422Resp{errors=[]view.ErrorCreateItem}
// @Failure      500              {object}  view.Error500Resp{errors=[]view.ErrorCreateItem}
// @Router       /staffs [post]
func (ctrl *Controller) Create(c *gin.Context) {
	sp, ctx := opentracing.StartSpanFromContext(c, "handler.staff.Create")
	defer sp.Finish()

	inp := &inout.StaffCreateInput{}
	if err := c.ShouldBindJSON(inp); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	id, err := ctrl.service.Create(ctx, inp)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeCreateSuccessResp(c, id)
}
