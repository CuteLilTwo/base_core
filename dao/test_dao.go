/*
 * @Description:
 * @Author: Young (hao_youngg@163.com)
 * @LastEditors: Young (hao_youngg@163.com)
 * @Date: 2019-04-30 14:31:35
 * @LastEditTime: 2019-04-30 15:42:58
 */
package dao

import (
	"base_core/data"

	"github.com/go-xorm/xorm"
)

type TestDao struct {
	engine *xorm.Engine
	sess   *xorm.Session
}

func NewTestDao(engine *xorm.Engine) *TestDao {
	return &TestDao{
		engine: engine,
		sess:   engine.NewSession(),
	}
}

/**
 * @Description: This Dao For Test
 * @Author: Young (hao_youngg@163.com)
 * @LastEditors: Young (hao_youngg@163.com)
 * @param {
	 data.TestDaoInData
 }
 * @return: {
	 data.TestDaoOutData
 }
 * @Date: 2019-04-30 14:44:08
*/
func (d *TestDao) TestDao(inputdata *data.TestDaoInputData) *data.TestDaoOutputData {
	// Data processing
	outdata := new(data.TestDaoOutputData)
	outdata.TestDaoOutputStr = inputdata.TestDaoInputStr + " Dao Data Output ->"
	return outdata
}
