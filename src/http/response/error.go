package response

func GenerateErrorMsg(errmsgMap map[string]string) string {
	msg := "codef_error : " +
		errmsgMap["code"] + ", " +
		errmsgMap["message"]

	return msg
}
