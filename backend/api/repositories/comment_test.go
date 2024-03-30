package repositories_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sns/backend/models"
	"github.com/sns/backend/repositories"
)

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "CommentInsertTest",
	}
	expectedCommentID := 3
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}

	if newComment.CommentID != expectedCommentID {
		t.Errorf("new comment id is expected %d but got %d\n", expectedCommentID, newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
		delete from comments
		where message = ?
		`
		testDB.Exec(sqlStr, comment.Message)
	})
}
