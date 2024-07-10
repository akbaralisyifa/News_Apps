package helper

func ResponseFormat(code int, message string, data any) map[string]any {
	var result = make(map[string]any)
	result["code"] = code
	result["message"] = message
	if data != nil {
		result["data"] = data
	}
	return result
}
