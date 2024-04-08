package services

import (
	"database/sql"
	"errors"
	"sync"

	"github.com/sns/backend/apperrors"
	"github.com/sns/backend/models"
	"github.com/sns/backend/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	var amu sync.Mutex
	var cmu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	go func(db *sql.DB, articleID int) {
		defer wg.Done()
		amu.Lock()
		article, articleGetErr = repositories.SelectArticleDetail(db, articleID)
		amu.Unlock()
	}(s.db, articleID)

	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			articleGetErr = apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, articleGetErr
		}
		articleGetErr = apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
		return models.Article{}, articleGetErr
	}

	go func(db *sql.DB, articleID int) {
		defer wg.Done()
		cmu.Lock()
		commentList, commentGetErr = repositories.SelectCommentList(db, articleID)
		cmu.Unlock()
	}(s.db, articleID)

	wg.Wait()

	if commentGetErr != nil {
		commentGetErr = apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
		return models.Article{}, commentGetErr
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}
	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return articleList, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice count")
		return models.Article{}, err
	}

	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
