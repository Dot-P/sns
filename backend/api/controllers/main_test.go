package controllers_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sns/backend/controllers"
	"github.com/sns/backend/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {

	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
