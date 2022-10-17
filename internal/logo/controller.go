package logo

import (
	"crypto/md5"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (controller Controller) Get(c echo.Context) error {
	url := "./assets/membership.png" // 파일

	file, err := os.Stat(url) //파일을 가져온다
	if err != nil {
		return echo.ErrInternalServerError
	}

	modifiedTime := file.ModTime() //파일 수정시간을 가지고 온다. 수정시간이나 파일 사이즈를 이용한다.
	etag := fmt.Sprintf("%x", md5.Sum([]byte(modifiedTime.String())))

	if c.Request().Header.Get("If-Not-Modified") == etag {
		return c.NoContent(http.StatusNotModified)
	}

	c.Response().Header().Set("Etag", etag) //Etag를 가져와서 response 헤더에 셋팅을 해준다.
	return c.File(url)
}
