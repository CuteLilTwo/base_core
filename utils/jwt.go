/*
 * @Description:
 * @Author: Young (hao_youngg@163.com)
 * @LastEditors: Young (hao_youngg@163.com)
 * @Date: 2019-04-30 16:44:03
 * @LastEditTime: 2019-04-30 16:45:32
 */
package utils

import (
	"base_core/config"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func UnixToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0).In(config.SYS_TIME_LOCATION)
}

func AuthToken(tokenString string, Secret []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return Secret, nil
	})

	// TODO: TOKEN失效时间延长
	tokencheck := token.Claims.(jwt.MapClaims)
	authtime := UnixToTime(int64(tokencheck["Timeout"].(float64)))
	out, _ := time.ParseDuration("2h")
	authtime = authtime.Add(out)
	fmt.Println(authtime.After(time.Now()))
	if authtime.After(time.Now()) == false {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), err
}
