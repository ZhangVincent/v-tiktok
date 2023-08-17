// @description 违禁词校验
// @author zkp15
// @date 2023/8/12 9:42
// @version 1.0

package forbiddenWordCheck

import (
	"strings"
	"v-tiktok/model/constants"
	"v-tiktok/pkg/strs"
)

func Check(content *string) (hitWords []string) {
	if strs.IsBlank(*content) {
		return
	}
	words := constants.ForbiddenWords
	if len(words) == 0 {
		return
	}
	for _, word := range words {
		if strings.Contains(*content, word) {
			hitWords = append(hitWords, word)
			*content = strings.ReplaceAll(*content, word, "*")
		}
	}
	return
}
