package service

import (
	"errors"
	"hr/config"
	"hr/repository"
	"hr/util"
)

type LoginService interface {
	Login(account, password string) (string, error)
	Logout(token string) bool
}

// 管理后台登录服务
type adminLoginService struct {
	adminRepo repository.AdminRepository
}

// 企业后台登录服务
type companyLoginService struct {
	companyRepo repository.CompanyRepository
}

// 用户端登录服务
type userLoginService struct {
}

func NewAdminLoginService(adminRepo repository.AdminRepository) LoginService {
	return &adminLoginService{
		adminRepo: adminRepo,
	}
}

func NewCompanyLoginService(loginRepo repository.CompanyRepository) LoginService {
	return &companyLoginService{
		companyRepo: loginRepo,
	}
}

// 管理员登录
func (als *adminLoginService) Login(account, passwd string) (token string, err error) {
	admin, err := als.adminRepo.FindByAcount(account)
	if err != nil {
		return token, errors.New(config.ACCOUNT_PASSWORD_ERROR)
	}
	if !util.VerifyPwd(passwd, admin.Password) {
		return token, errors.New(config.ACCOUNT_PASSWORD_ERROR)
	}
	token, err = util.CreateToken(map[string]interface{}{
		"id":      admin.Id,
		"account": admin.Account,
		"uuid":    admin.Uuid,
	}, config.GetConfig().Get("jwt.secret").ToString())
	if err != nil {
		return token, err
	}
	return token, nil
}

// 管理员退出登录
func (als *adminLoginService) Logout(token string) bool {
	return true
}

// 企业账户登录
func (ls *companyLoginService) Login(account, password string) (token string,err error) {
	company := ls.companyRepo.FindByAccount(account)
	if company == nil {
		return "", errors.New(config.LOGIN_FAIL)
	}
	if !util.VerifyPwd(password, company.Password) {
		return "", errors.New(config.LOGIN_FAIL)
	}
	token, err = util.CreateToken(map[string]interface{}{
		"id":      company.Id,
		"account": company.Account,
		"uuid":    company.Uuid,
	}, config.GetConfig().Get("jwt.secret").ToString())
	if err != nil {
		return token, err
	}
	return token, nil
}

// 企业账号退出
func (ls *companyLoginService) Logout(token string) bool {
	return true
}
