package search

import "strings"

var (
	questionKeywords = []string{"what", "when", "where", "which", "who", "whom", "whose", "why", "how", "is", "are", "can", "do"}
)

func ContainsQuestionKeyword(query string) bool {
	reformedQuery := strings.ToLower(query)

	for _, keyword := range questionKeywords {
		if strings.HasPrefix(reformedQuery, keyword) {
			return true
		} else {
			return false
		}
	}

	return false
}
