package utility

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
		return
	}
}
