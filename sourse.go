package main

func sumString(str1, str2 string, sep byte) string {
	return str1 + string(sep) + str2
}

func doubleString(str string) string {
	return str + str
}
