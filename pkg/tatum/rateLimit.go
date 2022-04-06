package tatum

func reachRateLimitOfTatum(response string) bool {
	if response == "<html>\n<head><title>429 Too Many Requests</title></head>\n<body>\n<center><h1>429 Too Many Requests</h1></center>\n<hr><center>nginx</center>\n</body>\n</html>\n" {
		return true
	} else if response == "{\"statusCode\":429,\"errorCode\":\"subscription.suspended\",\"message\":\"Subscription for '1ab39066-9304-4f7a-adca-dd4ee10d165b_100' is suspended, more then 120% of credits were used. You can upgrade your subscription at https://dashboard.tatum.io\"}" {
		return true
	} else {
		return false
	}
}
