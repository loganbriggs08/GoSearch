package search

import (
	"fmt"
	"strings"
)

var (
	questionKeywords = []string{"what", "when", "where", "which", "who", "whom", "whose", "why", "how", "is", "are", "can", "do"}
)

func ContainsQuestionKeyword(query string) bool {
	fmt.Println("Checking query")
	reformedQuery := strings.ToLower(query)

	for _, keyword := range questionKeywords {

		if strings.HasPrefix(reformedQuery, keyword) {
			return true
		}
	}

	return false
}
