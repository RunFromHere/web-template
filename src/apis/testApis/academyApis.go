package testApis

import (
	"attendance/models"
	"attendance/util"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"koala/src/util/jsonUtil"
	"net/http"
	"strconv"
)

func TestApi(c *gin.Context) {
	var academy models.Academy
	//var err1 error
	//var err2 error
	//data, _ := c.GetRawData()
	//id, err1 := jsonparser.GetInt(data, "academyId")
	//academy.AcademyId = int(id)
	if c.BindJSON(&academy) != nil {
		jsonUtil.ErrJsonDefault(c, 400, "BINDJSON错误")
		return
	} else {
		fmt.Println("academyid: ", academy.AcademyId, "    name ", academy.Name)
		return
	}
}

func AddOrUpdateAcademyApi(c *gin.Context) {
	var academy models.Academy
	//var errForName error
	//var errForId error

	data, _ := c.GetRawData()
	name, err := jsonparser.GetString(data, "name")
	if err != nil || name == "" {
		jsonUtil.ErrJsonDefault(c, 400, "传递参数错误")
		return
	}
	academy.Name = name

	id, err := jsonparser.GetInt(data, "academyId")
	switch {
	case err != nil:
		//没有传academyId
		//添加
		AddAcademyApi(c, academy)
	case id <= 0:
		jsonUtil.ErrJsonDefault(c, 400, "传递参数错误")
		return
	case err == nil:
		academy.AcademyId = id
		//修改
		UpdateAcademyApi(c, academy)
	}
}

//添加学院
func AddAcademyApi(c *gin.Context, academy models.Academy) {
	err := academy.AddAcademy()
	if err != nil {
		jsonUtil.ErrJsonDefault(c, 500, "1数据库错误: "+err.Error())
		return
	}
	jsonUtil.OkJsonDefault(c, "1添加成功")
	return
}

//更新学院
func UpdateAcademyApi(c *gin.Context, academy models.Academy) {
	err := academy.UpdateAcademyById()
	if err != nil {
		jsonUtil.ErrJsonDefault(c, http.StatusInternalServerError, "2数据库错误: "+err.Error())
		return
	}
	jsonUtil.OkJsonDefault(c, "2修改成功")
	return
}

//删除一个学院
func DeleteAcademyApi(c *gin.Context) {
	var academy models.Academy
	//var err error
	academy.AcademyId, _ = strconv.ParseInt(c.Param("academyId"), 10, 64)

	if academy.AcademyId > 0 {
		err := academy.DeleteAcademyById()
		if err != nil {
			jsonUtil.ErrJsonDefault(c, 500, "数据库错误："+err.Error())
			return
		}
		jsonUtil.OkJsonDefault(c, "删除成功")
		return
	} else {
		jsonUtil.ErrJsonDefault(c, http.StatusBadRequest, "传递参数有误")
		return
	}
}

//查询一个学院
func FindOneAcademyApi(c *gin.Context) {
	var academy models.Academy
	academy.AcademyId, _ = strconv.ParseInt(c.Param("academyId"), 10, 64)

	if academy.AcademyId <= 0 {
		jsonUtil.ErrJsonDefault(c, http.StatusBadRequest, "传递参数有误")
		return
	}
	err := academy.SearchAcademyById()
	if err != nil {
		jsonUtil.ErrJsonDefault(c, 500, "数据库错误："+err.Error())
		return
	}
	jsonUtil.DataJsonDefault(c, "查询成功", academy)
	return
}

//分页查询学院
func SearchAcademyApi(c *gin.Context) {
	var academy models.Academy
	academies := make([]models.Academy, 0)
	page, _ := strconv.Atoi(c.Param("page"))
	limit, _ := strconv.Atoi(c.Param("limit"))
	offset, err := util.Page(page, limit)
	if err != nil {
		jsonUtil.ErrJsonDefault(c, http.StatusBadRequest, "分页参数错误："+err.Error())
		return
	}

	data, _ := c.GetRawData()
	academy.Name, err = jsonparser.GetString(data, "name")
	if err != nil || academy.Name == "" {
		academies, err = academy.SearchAllAcademy(offset, limit)
		if err != nil {
			jsonUtil.ErrJsonDefault(c, 500, "数据库错误： "+err.Error())
			return
		}
	} else {
		academies, err = academy.SearchAllAcademyLikeName(offset, limit)
		if err != nil {
			jsonUtil.ErrJsonDefault(c, 500, "数据库错误： "+err.Error())
			return
		}
	}

	jsonUtil.DataJsonDefault(c, "查询成功", academies)
	return
}
