package UploadToBlob

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
)

func Testku (c echo.Context) error {
	c.FormFile("upload")
	return nil
}

func ProcessImage(c echo.Context) error {
	fmt.Println("masuk sini")
	file, err := c.FormFile("upload")
	email := c.FormValue("email")
	//files := c.FormValue("upload")
	//defer file.Close()
	if err != nil {
		fmt.Println("Error 1 : ", err.Error())
	}
	src, err := file.Open()
	if err != nil {
		fmt.Println("Error open  ",err.Error() )
	}

	dst, err := os.Create(file.Filename)
	fmt.Println("file : ", file.Filename)
	if err != nil {
		fmt.Println("Err dst : ", err.Error())
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		fmt.Println("Error 2 : ", err.Error())
	}
	fmt.Println(buf.Bytes())
	fmt.Println("berhasil",dst)
	UplodBytesToBlob(buf.Bytes(), email)
	//fmt.Println("isi out : ", file)
	//fmt.Println("isi in : ", buf)
	//do other stuff
	return c.JSON(http.StatusOK, "berhasil di upload")
}
