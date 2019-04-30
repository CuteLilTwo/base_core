/*
 * @Description:
 * @Author: Young (hao_youngg@163.com)
 * @LastEditors: Young (hao_youngg@163.com)
 * @Date: 2019-04-30 14:57:49
 * @LastEditTime: 2019-04-30 15:50:45
 */
package services

import (
	"base_core/connect"
	"base_core/dao"
	"base_core/data"
)

type TestService interface {
	TestServices(inputSerdata *data.TestControllerInputData) *data.TestControllerOutputData
}

type testService struct {
	testDao *dao.TestDao
}

func NewTestService() TestService {
	return &testService{
		testDao: dao.NewTestDao(connect.PsqlEngine),
	}
}

/**
 * @Description:
 * @Author: Young (hao_youngg@163.com)
 * @LastEditors: Young (hao_youngg@163.com)
 * @param {
	serdata.TestServiceInputData
 }
 * @return: {
	serdata.TestServiceOutputData
 }
 * @Date: 2019-04-30 15:12:14
*/
func (s *testService) TestServices(inputdata *data.TestControllerInputData) *data.TestControllerOutputData {
	// Logical processing

	// dao data input
	inputDaodata := new(data.TestDaoInputData)

	// controller data output
	outputCondata := new(data.TestControllerOutputData)

	// controller data handle for dao data
	inputDaodata.TestDaoInputStr = inputdata.TestControllerInputstr + " Service Data Input ->"

	// dao processing
	outputDaodata := s.testDao.TestDao(inputDaodata)

	// service processing
	outputCondata.TestControllerOutputStr = outputDaodata.TestDaoOutputStr + " Service Data Output ->"

	return outputCondata
}
