package utils

import "strings"

// UrlIsEq 判断 restful 风格的api是否和请求的api相等
// 待匹配 请求的url
// 匹配的url
func UrlIsEq(source, target string) bool {

	sourceLen := len(strings.Split(target, "/"))
	targetLen := len(strings.Split(source, "/"))
	if sourceLen != targetLen {
		return false
	}
	// 优先比较固定规则
	// 冒号所在的 位置
	index := -1
	index1 := strings.Index(target, ":")
	index2 := strings.Index(target, "*")
	// 谁在前取谁
	if index1 != -1 {
		index = index1
	}

	if index == -1 || (index2 != -1 && index > index2) {
		index = index2
	}

	if index != -1 {

		publicTarget := target[:index]

		// 如果请求的url不包含 公共部分，则认为不是相同的url
		if !strings.HasPrefix(source, publicTarget) {
			return false
		}
		// 只取不同的部分进行比对
		target = target[index:]
		source = source[index:]
	}

	targetParts := strings.Split(target, "/")
	sourceParts := strings.Split(source, "/")

	for i := 0; i < len(targetParts); i++ {
		targetS := targetParts[i]

		// 如果 target 包含占位符 ":", 则认为匹配 后面的不做匹配
		if strings.HasPrefix(targetS, "*") {
			return true
		}

		// 如果 path1 包含占位符 ":"，则认为匹配 继续匹配下一个
		if strings.HasPrefix(targetS, ":") {
			continue
		}

		if len(sourceParts) <= i {
			// 没有了，认为不匹配
			return false
		}

		//全匹配
		sourceS := sourceParts[i]

		if targetS == sourceS {
			continue
		}

		// 如果以上条件都不满足，说明不匹配
		return false
	}

	// 如果以上所有条件都通过，则认为 URL 相等
	return true
}
