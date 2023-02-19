package service

import (
	"encoding/json"
	"errors"
	"eyesStars/app/common"
	"eyesStars/app/constant"
	"eyesStars/app/model/dto"
	"eyesStars/app/model/entity"
	"eyesStars/app/model/po"
	"eyesStars/app/model/receiver"
	"eyesStars/app/model/returnee"
	"eyesStars/app/rpc/generate/userThrift"
	"eyesStars/app/rpc/rpc"
	"eyesStars/app/utils"
	"eyesStars/global"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

/**
 * @author eyesYeager
 * @date 2023/1/28 20:51
 */

type starService struct {
	cropBuffer uint8 // 星星记忆数的裁剪缓冲值
}

var StarService = starService{5}

// AddStar 添加星星
func (starsService *starService) AddStar(data receiver.StarAdd, uid uint32) (err error, result returnee.StarAdd) {
	result = returnee.StarAdd{}
	star := entity.Star{
		Uid:       uid,
		Content:   data.Content,
		Name:      data.Name,
		Anonymous: data.Anonymous,
	}

	// 插入数据库
	if err = global.DB.Create(&star).Error; err != nil {
		return err, result
	}

	// 结构体赋值
	if err = copier.Copy(&result, &star); err != nil {
		return err, result
	}

	// 处理时间格式
	result.CreateTime = star.CreateTime.Format("2006.01.02 15:04:05")

	return err, result
}

// GetStarById 根据id查询星星
func (starsService *starService) GetStarById(id uint32) (err error, result returnee.StarGet) {
	result = returnee.StarGet{}

	// 查询数据库
	var star entity.Star
	err = global.DB.First(&star, id).Error

	// 结果判断
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return common.CustomError{}.SetErrorMsg("星星不存在"), result
		}
		return err, result
	}
	if star.Status != constant.StarStatus.Normal {
		// 星星被冻结或被删除
		return common.CustomError{}.SetErrorMsg("该定星星不存在"), result
	}

	// 结构体赋值
	if err = copier.Copy(&result, &star); err != nil {
		return err, result
	}

	// 处理时间格式
	result.CreateTime = star.CreateTime.Format("2006.01.02 15:04:05")

	return err, result
}

// GetStars 批量获取星星
func (starsService *starService) GetStars(c *gin.Context, ids string) (err error, result returnee.StarsGet) {
	// 拼接执行sql，随机获取记录
	sql := func(status uint8, notIn string, num uint8) string {
		return `SELECT t1.id, t1.uid, t1.content, t1.name, t1.anonymous, t1.create_time FROM star t1 
				JOIN (
					SELECT floor(
						rand()*((SELECT max(id) FROM star)-(SELECT min(id) FROM star)+1)+(SELECT min(id) FROM star)
					) id 
				) t2 
				WHERE t1.id >= t2.id AND t1.status=` + strconv.Itoa(int(status)) + ` AND t1.id NOT IN(` + notIn + `)
				ORDER BY t1.id
				LIMIT ` + strconv.Itoa(int(num)) + `;`
	}

	// 星星容器
	var stars []po.StarsGet

	// 循环获取
	missNum := global.Config.Program.StarsNum
	for uint8(len(stars)) != global.Config.Program.StarsNum {
		var container []po.StarsGet
		// 执行随机sql查询
		err := global.DB.Raw(sql(constant.StarStatus.Normal, ids, missNum)).Scan(&container).Error
		if err != nil {
			global.Log.Error("随机获取星星失败！" + err.Error())
			return common.CustomError{}.SetErrorMsg("程序出错，清缓存再刷新看看？"), result
		}
		// 随机获取的数量小于等于传进去的missNum
		stars = append(stars, container...)
		for _, v := range container {
			ids += "," + strconv.Itoa(int(v.Id))
		}
		if uint8(len(container)) < missNum { // 查询的值小于指定的数量时，需要重新计算missNum再做查询
			missNum -= uint8(len(container))
		} else { // 查询的值等于指定的数量时，退出循环
			break
		}
	}

	// 查看ids是否超出记忆数量，超出则裁剪
	if idNum := strings.Count(ids, ",") + 1; uint8(idNum) > global.Config.Program.StarMemoryLen {
		// 查找截取位置
		substrNum := uint8(idNum) - global.Config.Program.StarMemoryLen + starsService.cropBuffer
		count := uint8(0)
		for i, v := range ids {
			if string(v) == "," {
				count++
			}
			if count == substrNum {
				ids = ids[i+1:]
				break
			}
		}
	}

	// 组装非匿名uid
	var uidList []uint64
	for _, v := range stars {
		if !v.Anonymous && !utils.Contains(uidList, v.Uid) {
			uidList = append(uidList, v.Uid)
		}
	}

	// 获取非匿名用户信息
	var uid2Info = map[uint64]*userThrift.UserInfoReturnee{}
	if len(uidList) != 0 {
		// 连接耶瞳用户中心
		err, client, transport := rpc.User()
		if err != nil {
			return common.CustomError{}.SetErrorMsg("用户中心服务似乎宕机了"), result
		}
		defer func(transport thrift.TTransport) {
			closeErr := transport.Close()
			if closeErr != nil {
				global.Log.Error("UserService:thrift连接关闭异常！" + closeErr.Error())
			}
		}(transport)

		// 获取数据
		_, tUidList := utils.Uint64ToInt64Slice(uidList)
		userInfoList, err := client.GetBatchUserInfo(c, tUidList)
		for _, v := range userInfoList {
			uid2Info[uint64(v.ID)] = v
		}
	}

	// 构建数据
	var completeStarList []dto.CompleteStar
	for _, v := range stars {
		completeStar := dto.CompleteStar{}
		_ = copier.Copy(&completeStar, &v)
		if v.Anonymous {
			completeStar.Avatar = global.Config.Program.DefaultAvatar
			completeStar.Name = global.Config.Program.DefaultName
		} else {
			completeStar.Avatar = uid2Info[v.Uid].Avatar
			completeStar.Name = uid2Info[v.Uid].Username
		}
		if v.Name != "" {
			completeStar.Name = v.Name
		}
		completeStar.CreateTime = v.CreateTime.Format("2006-01-02 15:04:05")
		completeStarList = append(completeStarList, completeStar)
	}

	// 返回结果
	jsonStars, _ := json.Marshal(completeStarList)
	return err, returnee.StarsGet{
		Ids:   utils.AesEncrypt(ids),
		Stars: utils.AesEncrypt(string(jsonStars)),
	}
}
