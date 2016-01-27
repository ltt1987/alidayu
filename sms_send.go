package alidayu

import (
	"time"
)

func SendSMS(rec_num, sms_free_sign_name, sms_template_code, sms_param string) (success bool, response string) {
	if rec_num == "" || sms_free_sign_name == "" || sms_template_code == "" {
		return false, "Parameter not complete"
	}

	params := make(map[string]string)
	params["app_key"] = AppKey
	params["format"] = "json"
	params["method"] = Method_SendSMS
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["v"] = "2.0"
	params["sms_type"] = "normal"
	params["sms_free_sign_name"] = sms_free_sign_name
	params["rec_num"] = rec_num
	params["sms_template_code"] = sms_template_code
	params["sms_param"] = sms_param

	return DoPost(params)

}
