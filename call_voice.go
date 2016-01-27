package alidayu

import "time"

func CallVoice(called_num, called_show_num, voice_code string) (success bool, response string) {
	if called_num == "" || called_show_num == "" || voice_code == "" {
		return false, "Parameter not complete"
	}

	params := make(map[string]string)
	params["app_key"] = AppKey
	params["format"] = "json"
	params["method"] = Method_CallVoice
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["v"] = "2.0"
	params["called_show_num"] = called_show_num
	params["called_num"] = called_num
	params["voice_code"] = voice_code

	return DoPost(params)
}
