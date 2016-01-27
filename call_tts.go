package alidayu

import "time"

func CallTTS(called_num, called_show_num, tts_code, tts_param string) (success bool, response string) {
	if called_num == "" || called_show_num == "" || tts_code == "" {
		return false, "Parameter not complete"
	}

	params := make(map[string]string)
	params["app_key"] = AppKey
	params["format"] = "json"
	params["method"] = Method_CallTTS
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["v"] = "2.0"
	params["called_show_num"] = called_show_num
	params["called_num"] = called_num
	params["tts_code"] = tts_code
	params["tts_param"] = tts_param

	return DoPost(params)
}
