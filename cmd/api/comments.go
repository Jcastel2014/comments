package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Jcastel2014/comments/internal/data"
	_ "github.com/Jcastel2014/comments/internal/data"
	"github.com/Jcastel2014/comments/internal/validator"
	_ "github.com/Jcastel2014/comments/internal/validator"
)

type CommentModel struct {
	DB *sql.DB
}

// func (c CommentModel) Insert(comment *Comment) error {
// 	query := `
// 		INSERT INTO comments(content,author)
// 		VALUES($1, $2)
// 		RETURNING id, created_at, version`

// 	args := []any{comment.Content, comment.Author}
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	return c.DB.QueryRowContext(ctx, query, args...).Scan(
// 		&comment.ID,
// 		&comment.CreatedAT,
// 		&comment.Version,
// 	)
// }

func (a *applicationDependencies) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	var incomingData struct {
		Content string `json:"content"`
		Author  string `json:"author"`
	}

	err := a.readJSON(w, r, &incomingData)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	comment := &data.Comment{
		Content: incomingData.Content,
		Author:  incomingData.Author,
	}

	v := validator.NEW()

	data.ValidateComment(v, comment)
	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	// err = a.commentModel.Insert(comment)

	fmt.Fprintf(w, "%+v\n", incomingData)
}
