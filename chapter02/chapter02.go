package chapter02

import (
	"fmt"
	"gin_project/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserInfo struct {
	Id   int
	Name string
	Age  int
	Addr string
}

type User struct {
	Name     string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Email    string `form:"email" json:"email"`
}

func Str(ctx *gin.Context) {
	str := "葛新"
	ctx.HTML(http.StatusOK, "chapter02/str.html", str)
}

func UserInfoStruct(ctx *gin.Context) {
	stuInfo := UserInfo{
		Id:   1,
		Name: "葛新",
		Age:  29,
		Addr: "北京市昌平区",
	}
	ctx.HTML(http.StatusOK, "chapter02/userInfoStruct.html", stuInfo)
}

func Arr(ctx *gin.Context) {
	arr := [3]int{1, 3, 5}
	ctx.HTML(http.StatusOK, "chapter02/arr.html", arr)
}

func ArrStruct(ctx *gin.Context) {
	arrStruct := [3]UserInfo{
		{Id: 1, Name: "葛新", Addr: "北京市", Age: 28},
		{Id: 2, Name: "小米", Addr: "北京市", Age: 15},
		{Id: 3, Name: "haha", Addr: "上海市", Age: 30},
	}
	ctx.HTML(http.StatusOK, "chapter02/structArr.html", arrStruct)
}

func MapController(ctx *gin.Context) {
	mapData1 := map[string]string{"name": "葛新", "addr": "北京市昌平区"}
	mapData2 := map[string]string{"名字": "葛新", "住址": "北京市昌平区龙泽园街道", "age": "28"}
	mapData := map[string]interface{}{"mapData1": mapData1, "mapData2": mapData2}
	ctx.HTML(http.StatusOK, "chapter02/map.html", mapData)
}

func MapArrController(ctx *gin.Context) {
	mapStruct := map[string]UserInfo{"user": {Id: 1, Name: "葛新", Addr: "回龙观", Age: 28},
		"user1": {Id: 2, Name: "王岩", Age: 27, Addr: "***"},
	}
	ctx.HTML(http.StatusOK, "chapter02/map1.html", mapStruct)
}

func SliceController(ctx *gin.Context) {
	sliceDemo := []int{1, 2, 3, 4, 5}
	ctx.HTML(http.StatusOK, "chapter02/slice.html", sliceDemo)
}

func Param1Controller(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Printf("id: %v\n", id)
	ctx.String(http.StatusOK, "Success!")
}

func Param2Controller(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Printf("id: %v\n", id)
	ctx.String(http.StatusOK, "Success!")
}

func GetQueryDataController(ctx *gin.Context) {
	id, _ := ctx.GetQuery("id")
	name, _ := ctx.GetQuery("name")
	age := ctx.DefaultQuery("age", strconv.Itoa(18))
	fmt.Printf("id: %v, name: %v, age: %v\n", id, name, age)
	ctx.String(http.StatusOK, "Success!")
}

func GetQueryArrayDataController(ctx *gin.Context) {
	ids := ctx.QueryArray("ids")
	fmt.Printf("id: %v\n", ids)
	ctx.String(http.StatusOK, "Success!")
}

func GetQueryMapDataController(ctx *gin.Context) {
	user := ctx.QueryMap("user")
	fmt.Printf("id: %v\n", user)
	ctx.String(http.StatusOK, "Success!")
}

func ToUserAddController(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter02/userAdd1.html", nil)
}

func DoUserAddController(ctx *gin.Context) {
	var user User
	errBindData := ctx.ShouldBind(&user)
	if errBindData != nil {
		return
	} else {
		fmt.Println(user)
	}
	hobbies := ctx.PostFormArray("hobby")
	fmt.Printf("hobbies: %v\n", hobbies)
	u := ctx.PostFormMap("user")
	fmt.Println(u["age"], u["address"])
	ctx.String(http.StatusOK, "Success")
}

func ToUserAddAJAXController(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter02/userAdd2.html", nil)
}

func DoUserAddAJAXController(ctx *gin.Context) {
	username := ctx.PostForm("username")
	fmt.Println(username)
	data := map[string]interface{}{
		"code": 200,
		"msg":  "success",
	}
	ctx.JSON(http.StatusOK, data)
}

func ToSingleFileUploadController(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter02/upload3.html", nil)
}

func DoSingleFileUploadController(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fileName := utils.FileNameFormat(file.Filename)
	fmt.Println(fileName)
	dst := "upload/" + fileName
	errSaveFile := ctx.SaveUploadedFile(file, dst)
	if errSaveFile != nil {
		return
	}
	ctx.String(http.StatusOK, "Success!")
}

func ToMultiFileUploadController(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter02/upload2.html", nil)
}

func DoMultiFileUploadController(ctx *gin.Context) {
	form, errGetPostData := ctx.MultipartForm()
	if errGetPostData != nil {
		fmt.Println(errGetPostData.Error())
		return
	}
	files := form.File["file"]
	for _, file := range files {
		fileName := utils.FileNameFormat(file.Filename)
		fmt.Println(fileName)
		dst := "upload/" + fileName
		errSaveFile := ctx.SaveUploadedFile(file, dst)
		if errSaveFile != nil {
			return
		}
	}

	ctx.String(http.StatusOK, "Success!")
}

func ToSingleFileUploadAJAXController(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter02/upload3.html", nil)
}

func DoSingleFileUploadAJAXController(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	name := ctx.PostForm("name")
	fmt.Printf("name: %v\n", name)
	fileName := utils.FileNameFormat(file.Filename)
	fmt.Println(fileName)
	dst := "upload/" + fileName
	errSaveFile := ctx.SaveUploadedFile(file, dst)
	if errSaveFile != nil {
		return
	}
	data := map[string]interface{}{
		"code": 200,
		"msg":  "Success!",
	}
	ctx.JSON(http.StatusOK, data)
}

func ToMultiFileUploadAJAXController(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter02/upload4.html", nil)
}

func DoMultiFileUploadAJAXController(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "获取数据失败",
		})
		return
	}
	files := form.File["files"]
	name := ctx.PostForm("name")
	fmt.Printf("name: %v\n", name)
	for _, file := range files {
		fileName := utils.FileNameFormat(file.Filename)
		fmt.Println(fileName)
		dst := "upload/" + fileName
		errSaveFile := ctx.SaveUploadedFile(file, dst)
		if errSaveFile != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "上传文件解析失败",
			})
			return
		}
	}
	//data := map[string]interface{}{
	//	"code": 200,
	//	"msg":  "Success!",
	//}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
	})
}
