package alidayu

import "time"

func CallDouble(caller_num, caller_show_num, called_num, called_show_num string) (success bool, response string) {
	if caller_num == "" || caller_show_num == "" || called_num == "" || called_show_num == "" {
		return false, "Parameter not complete"
	}

	params := make(map[string]string)
	params["app_key"] = AppKey
	params["format"] = "json"
	params["method"] = Method_CallVoice
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["v"] = "2.0"
	params["caller_num"] = caller_num
	params["caller_show_num"] = caller_show_num
	params["called_num"] = called_num
	params["called_show_num"] = called_show_num

	return DoPost(params)
}
