package db

func Setup() {
	r := GetRedis()
	r.TESTCONN()
}
