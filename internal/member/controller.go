package member

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	customErrors "golang_web_programming/errors"
	"golang_web_programming/internal/user"
	"net/http"
	"strings"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}


// Create godoc
// @Summary      멤버십 생성
// @Description  멤버십을 생성합니다.
// @Accept       json
// @Tags         Memberships
// @Produce      json
// @Param        requestBody  body      CreateRequest  true  "user_name:사용자의 이름, membership_type:naver,toss,pacyco 중 하나"
// @Success      201          {object}  CreateResponse
// @Router       /v1/memberships [post]
func (controller *Controller) Create(c echo.Context) error {
	var request CreateRequest

	if err := c.Bind(&request); err != nil {
		return customErrors.ErrParamsBinding
	}

	if err := ValidateCreateRequest(request); err != nil {
		return err
	}

	createResponse, err := controller.service.Create(&request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, createResponse)
}

func (controller *Controller) Update(c echo.Context) error {
	var request UpdateRequest

	id := c.Param("id")

	if err := c.Bind(&request); err != nil {
		return customErrors.ErrParamsBinding
	}

	request.ID = id

	if err := ValidateUpdateRequest(request); err != nil {
		return err
	}

	updateResponse, err := controller.service.Update(&request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, updateResponse)
}


//GetByID godoc
//@Summary      멤버십 정보 단건 조회
//@Description  멤버십 정보를 조회합니다. (상세 설명)
//@Accept       json
//@Tags         Memberships
//@Produce      json
//@param        Authorization  header    string  true  "Authorization"  default(Bearer <Add access token here>)
//@Param        id             path      string  true  "Membership uuid"
//@Success      200            {object}  GetResponse
//@Failure      400            {object}  Fail400GetResponse
//@Router       /v1/memberships/{id} [get]
func (controller *Controller) GetByID(c echo.Context) error {
	id := c.Param("id")

	res, err := controller.service.GetByID(id)
	if err != nil {
		return echo.ErrInternalServerError
	}

	userInfo := c.Get("user").(*jwt.Token).Claims.(*user.Claims)
	if userInfo.Name != res.UserName && !userInfo.IsAdmin {
		return customErrors.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, res)
}

func (controller *Controller) Delete(c echo.Context) error {
	id := c.Param("id")

	if strings.Trim(id, " ") == "" {
		return customErrors.ErrInvalidUserID
	}

	if err := controller.service.Delete(id); err != nil {
		return err
	}

	return nil
}

func (controller *Controller) GetList(c echo.Context) error {

	res, err := controller.service.GetList()
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}
