package utils

import "strconv"

var (
	KeyType = map[string]string{
		"USER_REGISTER_KEY":                "user_register:",
		"USER_REGISTER_TIME_KEY":           "user_register:",
		"USER_LOGIN_KEY":                   "login_user_",
		"USER_MODIFY_PAYPASSWORD_KEY":      "user_modify_paypassword:",
		"USER_MODIFY_PAYPASSWORD_TIME_KEY": "user_modify_time_paypassword:",
		"USER_CHANGE_ACCOUNT_KEY":          "user_change_account_auth_code:",
		"USER_CHANGE_ACCOUNT_TIME_KEY":     "user_change_account_auth_time_code:",
		"USER_FOGOT_PASSWORD_KEY":          "user_fogot_password_auth_code:",
		"USER_FOGET_PASSWORD_TIME_KEY":     "user_fogot_password_auth_time_code:",
		"Resource_TreeGrid4_Menu":          "ResourceTreeGrid4Menu",
		"Resource_TreeGrid4_Action":        "ResourceTreeGrid4Action",
		"SERVER_LOGIN":                     "server_login_",
		"CEC_TRANS_PUSH_NOMAL_KEY":         "",
		"VIP_SEARCH_KEY":                   "vip_parent_id:",
		"CEC_TRANS_NOT_PARESE_KEY":         "-notparse",
	}
)

type KeyData struct {
	PairName string
	Mobile   string
	UserId   int
	Num      int
}

func (k *KeyData) GetKey(KeyName string) string {
	switch KeyName {
	case "USER_REGISTER_KEY":
		return KeyType[KeyName] + k.Mobile
	case "USER_REGISTER_TIME_KEY":
		return KeyType[KeyName] + k.Mobile
	case "USER_FOGOT_PASSWORD_KEY":
		return KeyType[KeyName] + k.Mobile
	case "USER_FOGET_PASSWORD_TIME_KEY":
		return KeyType[KeyName] + k.Mobile
	case "USER_LOGIN_KEY":
		return KeyType[KeyName] + strconv.Itoa(k.UserId)
	case "USER_FOGET_PWD_KEY":
		return KeyType[KeyName] + k.Mobile
	case "Resource_TreeGrid4_Menu":
		return KeyType[KeyName] + "_" + strconv.Itoa(k.UserId) + "_" + strconv.Itoa(k.Num)
	case "Resource_TreeGrid4_Action":
		return KeyType[KeyName] + "_" + strconv.Itoa(k.UserId) + "_" + strconv.Itoa(k.Num)
	case "SERVER_LOGIN":
		return KeyType[KeyName] + strconv.Itoa(k.UserId)
	case "USER_CHANGE_ACCOUNT_KEY":
		return KeyType[KeyName] + k.Mobile
	case "USER_CHANGE_ACCOUNT_TIME_KEY":
		return KeyType[KeyName] + k.Mobile
	case "USER_MODIFY_PAYPASSWORD_KEY":
		return KeyType[KeyName] + k.Mobile
	case "USER_MODIFY_PAYPASSWORD_TIME_KEY":
		return KeyType[KeyName] + k.Mobile
	case "CEC_TRANS_PUSH_NOMAL_KEY":
		return k.PairName
	case "VIP_SEARCH_KEY":
		return KeyType[KeyName] + strconv.Itoa(k.UserId)
	case "CEC_TRANS_NOT_PARESE_KEY":
		return k.PairName + KeyType[KeyName]
	}
	return ""
}
