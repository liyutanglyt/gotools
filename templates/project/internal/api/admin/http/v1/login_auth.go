/**
 * Created by Wangwei on 2019-06-04 17:10.
 */

package v1

import (
	"goadmin/internal/common"
	"goadmin/internal/common/middleware/jwt"
	"goadmin/internal/model/sys"
	service "goadmin/internal/service/sys"
	"goadmin/pkg/security"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	myJwt           = &jwt.JWT{}
	employeeService = &service.EmployeeService{}
)

type AuthController struct{}

func (AuthController) Login(c *gin.Context) {
	clientId := c.Request.Header.Get("X-Client-Id")
	clientId = common.FormatClientId(clientId)

	var loginInfoReq *sys.LoginInfoReq
	if BindJSON(c, &loginInfoReq) != nil {
		ResponseError(c, "无效的请求数据")
		return
	}

	password := security.MD5Password(loginInfoReq.Password)
	employee, err := employeeService.GetByAccount(loginInfoReq.Account, password)
	if err != nil {
		ResponseError(c, err)
		return
	}

	routeLinks, err := roleMenuService.FindRouteLinksByRole(employee.RoleId)
	if err != nil {
		ResponseError(c, err)
		return
	}

	tokenStr, exp, err := myJwt.GenEmployeeToken(employee.Id, employee.Account,
		employee.RoleId, employee.OrgTypeId, employee.OrgId, employee.OrgName, clientId)
	if err != nil {
		ResponseError(c, "生成token失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":        0,
		"user":        employee,
		"route_links": routeLinks,
		"jwt": gin.H{
			"token":      tokenStr,
			"expires_in": exp,
		},
	})
}
