package dal

import (
	"database/sql"
	"fmt"
	"github.com/daniel/AnserBlock/common/utils"
	"github.com/daniel/AnserBlock/dao"
	"github.com/daniel/AnserBlock/dao/model"
)

// 注册操作
func Register(user *model.UserInfo) (err error) {
	// 1、查询数据库中该用户名的条目
	var count int64
	sqlstr := "select user_id from user where username=?"
	err = dao.DB.Get(&count, sqlstr, user.UserName)
	// 2、 判断当前的错误信息
	if err != nil && err != sql.ErrNoRows {
		return
	}
	// 3、若已经存在 则返回
	if count > 0 {
		err = utils.ErrUserAlreadyExisted
		return
	}

	// 4、成功 则插入
	password := user.Password + utils.PassWorldSolt
	dbPassword := utils.GetMD5([]byte(password))
	sqlstr = "insert into user (username,password,email,user_id,nickname,sex) values (?,?,?,?,?,?)"
	fmt.Printf("---->  %#v , dbPassword:  %s",user,dbPassword)

	_,err =  dao.DB.Exec(sqlstr, user.UserName, dbPassword, user.Email, user.UserId,user.NickName,user.Sex)
	if err != nil {

	}
	return

}

// 登陆操作数据库
func Login(user * model.UserInfo)(err error)  {

	// 1.	先保存传入过来的密码
	postPwd := user.Password
	// 2、 取出账号密码信息
	sqlstr := "select username,password,user_id from user where username=?"
	err = dao.DB.Get(user, sqlstr, user.UserName)


	// 3、若不存在存在 则返回不存在的错误码
	if err  == sql.ErrNoRows{
		err = utils.ErrUserNotExisted
		return
	}

	// 4、比较密码
	dbPwd := postPwd+utils.PassWorldSolt
	md5pwd:= utils.GetMD5([]byte(dbPwd))

	if md5pwd != user.Password {
		err = utils.ErrUserPwdWrong
       return
	}
	return
}

