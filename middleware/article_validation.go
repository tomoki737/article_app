package middleware

import (
	"strings"
	"errors"
	"fmt"
)

func ValidateComment(comment string) error {
	if len(comment) == 0 {
			return errors.New("コメントが空です")
	}

	maxCommentLength := 1000
	if len(comment) > maxCommentLength {
			return fmt.Errorf("コメントの長さが最大値 %d を超えています", maxCommentLength)
	}

	invalidChars := []string{"<script>", "</script>", "<", ">"}
	for _, char := range invalidChars {
			if strings.Contains(comment, char) {
					return errors.New("コメントに不適切な文字が含まれています")
			}
	}

	return nil
}
