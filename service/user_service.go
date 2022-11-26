package service

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/heqingbao/ginchat/models"
	"github.com/heqingbao/ginchat/utils"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// GetUserList
// @Summary 获取用户列表
// @Tags 用户模块
// @Success 200 {string} helloworld
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	list := models.GetUserList()
	c.JSON(http.StatusOK, list)
}

// FindUserByNameAndPwd
// @Summary 获取用户列表
// @Tags 用户模块
// @Param name formData string false "name"
// @Param password formData string false "password"
// @Success 200 {string} helloworld
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户名不存在",
		})
		return
	}
	if !utils.ValidPassword(password, user.Salt, user.Password) {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "密码错误",
		})
		return
	}

	// 生成token
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.Md5Encode(str)
	utils.DB.Model(user).Where("id = ?", user.ID).Update("identity", temp)

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    user,
	})
}

// CreateUser
// @Summary 创建用户
// @Tags 用户模块
// @Param name query string false "用户名"
// @Param password query string false "密码"
// @Param repassword query string false "再次确认密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "两次密码不一致",
		})
		return
	}

	if models.FindUserByName(user.Name).Name != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "名称已存在",
		})
		return
	}

	user.Salt = fmt.Sprintf("%06d", rand.Int31())
	user.Password = utils.MakePassword(password, user.Salt)
	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @Param id query string false "ID"
// @Success 200 {string} json{"code", "message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

// UpdateUser
// @Summary 更新用户
// @Tags 用户模块
// @Param id formData string false "ID"
// @Param name formData string false "name"
// @Param password formData string false "password"
// @Param phone formData string false "phone"
// @Param email formData string false "email"
// @Success 200 {string} json{"code", "message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "参数格式错误",
		})
		return
	}
	models.UpdateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}

// cors
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	MsgHandler(ws, c)
}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	for {
		msg, err := utils.Subscribe(c, utils.PublishKey)
		if err != nil {
			fmt.Println(err)
			return
		}
		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			fmt.Println(err)
		}
	}
}
