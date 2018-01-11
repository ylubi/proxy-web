package util

func CompressCommand(compress string) string {
	if compress == "1" {
		return " --c"
	}
	return ""
}
