package user

import (
	"errors"
	"github.com/UritMedical/qf"
	"github.com/UritMedical/qf/user/model"
	"github.com/UritMedical/qf/util"
	"strings"
)

//defPassword 默认密码
const defPassword = "123456"

func (b *Bll) regUserApi(api qf.ApiMap) {
	//登录
	api.Reg(qf.EApiKindSave, "login", b.login)

	//用户增删改查
	api.Reg(qf.EApiKindSave, "user", b.saveUser)
	api.Reg(qf.EApiKindDelete, "user", b.deleteUser)
	api.Reg(qf.EApiKindGetModel, "user", b.getUserModel)
	api.Reg(qf.EApiKindGetList, "users", b.getAllUsers)

	//密码重置、修改
	api.Reg(qf.EApiKindSave, "user/pwd/reset", b.resetPassword)
	api.Reg(qf.EApiKindSave, "user/pwd", b.changePassword)
}

//
// login
//  @Description: 用户登录
//  @param ctx
//  @return interface{}
//  @return error
//
func (b *Bll) login(ctx *qf.Context) (interface{}, error) {
	var params = struct {
		LoginId  string
		Password string //md5
	}{}

	if err := ctx.Bind(&params); err != nil {
		return nil, err
	}
	params.LoginId = strings.Replace(params.LoginId, " ", "", -1)
	if user, ok := b.userDal.CheckLogin(params.LoginId, params.Password); ok {
		role, _ := b.userRoleDal.GetUsersByRoleId(user.Id)
		token, _ := util.GenerateToken(user.Id, role)

		//获取用户所在部门
		departs, _ := b.getDepartsByUserId(user.Id)

		//获取用户所拥有的角色
		roles, _ := b.getRolesByUserId(user.Id)

		return map[string]interface{}{
			"Token":    token,
			"Departs":  util.ToMaps(departs),
			"Roles":    util.ToMaps(roles),
			"UserInfo": util.ToMap(user),
		}, nil
	} else if params.LoginId == devUser.LoginId && params.Password == devUser.Password {
		//开发者账号
		token, _ := util.GenerateToken(devUser.Id, []uint64{})
		return map[string]interface{}{
			"Token":    token,
			"UserInfo": util.ToMap(devUser),
		}, nil
	} else {
		return nil, errors.New("loginId not exist or password error")
	}
}

func (b *Bll) saveUser(ctx *qf.Context) (interface{}, error) {
	user := &model.User{}
	if err := ctx.Bind(user); err != nil {
		return nil, err
	}
	if !b.userDal.CheckExists(user.Id) {
		user.Password = util.ConvertToMD5([]byte(defPassword))
	}
	//创建用户
	return nil, b.userDal.Save(user)
}

func (b *Bll) deleteUser(ctx *qf.Context) (interface{}, error) {
	uId := ctx.GetId()
	return nil, b.userDal.Delete(uId)
}

func (b *Bll) getUserModel(ctx *qf.Context) (interface{}, error) {
	var user model.User
	userId := ctx.LoginUser().UserId

	//获取用户所在部门
	departs, _ := b.getDepartsByUserId(userId)

	//获取用户所拥有的角色
	roles, _ := b.getRolesByUserId(userId)

	err := b.userDal.GetModel(userId, &user)
	ret := map[string]interface{}{
		"Info":        util.ToMap(user),
		"Roles":       util.ToMaps(roles),
		"Departments": util.ToMaps(departs),
	}

	return ret, err
}

//
// getAllUsers
//  @Description: 获取所有用户
//  @param ctx
//  @return interface{}
//  @return error
//
func (b *Bll) getAllUsers(ctx *qf.Context) (interface{}, error) {
	list, err := b.userDal.GetAllUsers()
	result := make([]map[string]interface{}, 0)
	for _, user := range list {
		//获取用户所在部门
		departs, _ := b.getDepartsByUserId(user.Id)

		//获取用户所拥有的角色
		roles, _ := b.getRolesByUserId(user.Id)

		ret := map[string]interface{}{
			"UserInfo":    util.ToMap(user),
			"Roles":       util.ToMaps(roles),
			"Departments": util.ToMaps(departs),
		}
		result = append(result, ret)
	}
	return result, err
}

//
// resetPassword
//  @Description: 重置密码
//  @param ctx
//  @return interface{}
//  @return error
//
func (b *Bll) resetPassword(ctx *qf.Context) (interface{}, error) {
	uId := ctx.GetId()
	return nil, b.userDal.SetPassword(uId, util.ConvertToMD5([]byte(defPassword)))
}

//
// changePassword
//  @Description: 修改密码
//  @param ctx
//  @return interface{}
//  @return error
//
func (b *Bll) changePassword(ctx *qf.Context) (interface{}, error) {
	var params = struct {
		OldPassword string
		NewPassword string
	}{}
	if err := ctx.Bind(&params); err != nil {
		return nil, err
	}
	if !b.userDal.CheckOldPassword(ctx.LoginUser().UserId, params.OldPassword) {
		return nil, errors.New("old password is incorrect")
	}
	return nil, b.userDal.SetPassword(ctx.LoginUser().UserId, params.NewPassword)
}
