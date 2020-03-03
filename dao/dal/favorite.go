package dal

import (
	"github.com/NielJun/go-logger"
	"github.com/daniel/AnserBlock/common/utils"
	"github.com/daniel/AnserBlock/dao"
	"github.com/daniel/AnserBlock/dao/model"
)

// 创建一个收藏夹目录
// 通过数据库事务来查询和写入
// 先查询是否该用户已经创建了收藏夹
// 写入收藏夹
func CreateFavoriteDir(fd *model.FavoriteDir) (err error) {

	dbTx, err := dao.DB.Beginx()
	if err != nil {
		logger.Error("create data base Tx failed,err: %#v", err)
		return
	}

	var favoriteDirCount int
	sqlStr := "select count(dir_id) from favorite_dir where user_id = ? and dir_name = ?"
	err = dbTx.Get(&favoriteDirCount, sqlStr, fd.UserId, fd.DirName)
	if err != nil {
		logger.Error("get favoriteDir count failed,err: %#v", err)
		return
	}
	// 已经存在
	if favoriteDirCount > 0 {
		dbTx.Rollback()
		err = utils.ErrRecordExisted
		return
	}

	// 插入
	sqlStr = "insert into favorite_dir (user_id,dir_id,dir_name,count)values(?,?,?,?)"
	_, err = dbTx.Exec(sqlStr, fd.UserId, fd.DirId, fd.DirName, fd.Count)
	if err != nil {
		logger.Error("insert into favorite_dir failed,  favorite dir model is %#v  ,err: %#v", *fd, err)
		return
	}
	err = dbTx.Commit()
	if err != nil {
		logger.Error("insert into favorite_dir failed,  favorite dir model is %#v  ,err: %#v", *fd, err)
		return
	}

	return
}

// 收藏模块
// 1.判断是否收藏过了
// 2.插入收藏
func CreateFavorite(f *model.Favorite) (err error) {

	//1.判断
	var favoriteCount int
	sqlStr := "select count(answer_id) from favorite where dir_id = ? and answer_id = ?"
	err = dao.DB.Get(&favoriteCount, sqlStr, f.DirId, f.AnswerId)
	if err != nil {
		logger.Error("get count (answer_id) failed,sqlStr: %s,err: %#v", sqlStr, err)
		return
	}

	if favoriteCount > 0 {
		err = utils.ErrRecordExisted
		logger.Error("answer : %s have been recored!", f.AnswerId)
		return
	}

	//2.插入

	sqlStr = "insert into favorite (answer_id,user_id,dir_id)values (?,?,?)"
	_, err = dao.DB.Exec(sqlStr, f.AnswerId, f.UserId, f.DirId)
	if err != nil {
		logger.Error("insert into favorite table failed,sqlStr: %s,err: %#v", sqlStr, err)
		return
	}

	return
}

// 获取用户的所有的收藏夹列表
func GetFavoriteDirList(userId int64) (fevoriteDirList []*model.FavoriteDir, err error) {

	sqlStr := "select dir_id,dir_name,user_id,count from favorite_dir where user_id = ?"
	err = dao.DB.Select(&fevoriteDirList, sqlStr, userId)
	if err != nil {
		logger.Error("select fevorite dirs faild, userId is %v ,err: %#v", userId, err)
		return
	}
	return
}

func GetFavoriteList(userId, favoriteDirId int64, offset, limit int64) (favoriteList []*model.Favorite, err error) {
	sqlStr := "select dir_id,answer_id,user_id from favorite where user_id = ? and dir_id = ? limit ?,?"
	err = dao.DB.Select(&favoriteList, sqlStr, userId, favoriteDirId, offset, limit)
	if err != nil {
		logger.Error("select fevorite dirs faild, userId is %v ,err: %#v", userId, err)
		return
	}
	return
}
