/*
 * @Description: System startup and parameter configuration global variables for third party services
 * @Author: Young (hao_youngg@163.com)
 * @LastEditors: Young (hao_youngg@163.com)
 * @Date: 2019-04-30 11:41:33
 * @LastEditTime: 2019-04-30 16:32:30
 */
package config

import "time"

var (
	Conf *Config // System configuration config parameter

	SYS_TIME_LOCATION, _ = time.LoadLocation("Asia/Chongqing") // China's time zone
)

// Time format string
const (
	SYS_TIME_YMDHMS string = "2006-01-02 15:04:05" // Format time YYYY:MM:DD HH:mm:ss
	SYS_TIME_YMDHM  string = "2006-01-02 15:04"    // Format time YYYY:MM:DD HH:mm
	SYS_TIME_YMD    string = "2006-01-02"          // Format time YYYY:MM:DD
)

const (
	SUCC_CODE               = 200 + iota // Successful network request without any errors
	ERROR_CODE_CAN_NOT_SKIP              // Successful network request with some problems
	BACK_TO_LOGIN                        // A successful network request was made, but the user information is invalid and you need to login again
)
