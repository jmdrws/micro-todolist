package core

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"errors"
	"user/model"
	"user/service"
)

//序列化
func BuildUser(item model.User) *service.UserModel {
	userModel := service.UserModel{
		ID: uint32(item.ID),
		UserName: item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}

//登录
func (*UserService) UserLogin(ctx context.Context,req *service.UserRequest,resp *service.UserDetailResponse) error {
	var user model.User
	resp.Code = 200
	if err := model.DB.Where("user_name=?",req.UserName).First(&user).Error;err!=nil{
		if gorm.IsRecordNotFoundError(err){
			resp.Code = 400
			return nil
		}
		resp.Code = 500
		return nil
	}
	if user.CheckPassword(req.Password)==false{
		resp.Code = 400
		return nil
	}
	resp.UserDetail = BuildUser(user)
	return nil
}

//注册
func (*UserService) UserRegister(ctx context.Context,req *service.UserRequest,resp *service.UserDetailResponse) error{
	if req.Password != req.PasswordConfirm{
		err := errors.New("两次密码输入不一致")
		return err
	}
	count := 0
	if err := model.DB.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
		return err
	}
	if count>0 {
		err := errors.New("用户名已存在")
		return err
	}
	user:=model.User{
		UserName: req.UserName,
	}
	if err := user.SetPassword(req.Password); err != nil {
		return err
	}
	if err := model.DB.Create(&user).Error;err!=nil{
		return err
	}
	resp.UserDetail = BuildUser(user)
	return nil
}
