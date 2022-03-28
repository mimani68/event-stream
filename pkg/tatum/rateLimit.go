package tatum

func reachRateLimitOfTatum(response string) bool {
	if response == "<html>\n<head><title>429 Too Many Requests</title></head>\n<body>\n<center><h1>429 Too Many Requests</h1></center>\n<hr><center>nginx</center>\n</body>\n</html>\n" {
		return true
	} else {
		return false
	}
}
