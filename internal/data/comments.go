package data

import (
	"time"

	"github.com/Jcastel2014/comments/internal/validator"
)

type Comment struct {
	ID        int64
	Content   string
	Author    string
	CreatedAt time.Time
	Version   int32
}

func ValidateComment(v *validator.Validator, comment *Comment) {
	v.Check(comment.Content != "", "content", "must be provided")
	v.Check(comment.Author != "", "author", "must be provided")
	v.Check(len(comment.Content) <= 100, "content", "must not be more than 100 bytes long")
	v.Check(len(comment.Author) <= 100, "author", "must not be more than 25 butes longs")
}
