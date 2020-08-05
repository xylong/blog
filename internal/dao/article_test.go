package dao

import (
	"blog/internal/model"
	"blog/tools"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestArticle_Create(t *testing.T) {
	Convey("创建文章", t, func() {
		dao := NewArticleDao()
		article := model.Article{
			CategoryId: 1,
			Title:      "test create",
			Content:    "this is a test",
			Visits:     1,
		}
		id, err := dao.Create(&article)
		So(err, ShouldBeNil)
		So(id, ShouldBeGreaterThan, 0)
		t.Log(id)
	})
}

func TestArticle_Delete(t *testing.T) {
	Convey("删除文章", t, func() {
		dao := NewArticleDao()
		article := &model.Article{
			CategoryId: 1,
			Title:      "delete" + tools.RandomString(10),
			Content:    "testing for delete article",
			Visits:     1,
		}
		id, err := dao.Create(article)
		So(err, ShouldBeNil)
		So(id, ShouldBeGreaterThan, 0)
		ok := dao.Delete(id)
		So(ok, ShouldBeTrue)
	})
}

func TestArticle_Update(t *testing.T) {
	Convey("更新文章", t, func() {
		dao := NewArticleDao()
		article := &model.Article{
			CategoryId: 1,
			Title:      "update-" + tools.RandomString(10),
			Content:    "testing for update article",
			Visits:     1,
		}
		id, err := dao.Create(article)
		So(err, ShouldBeNil)
		So(id, ShouldBeGreaterThan, 0)
		article.Title = article.Title + tools.RandomNumber(4)
		article.Content = article.Content + " - " + tools.RandomString(50)
		ok := dao.Update(article)
		So(ok, ShouldBeTrue)
	})
}

func TestArticle_Select(t *testing.T) {
	Convey("查询文章", t, func() {
		dao := NewArticleDao()
		_, total, err := dao.Select(0, 10, "")
		So(err, ShouldBeNil)
		So(total, ShouldBeGreaterThanOrEqualTo, 0)
		t.Logf("count:%d", total)
	})
}

func TestArticle_Find(t *testing.T) {
	Convey("获取文章", t, func() {
		dao := NewArticleDao()
		article := model.Article{
			CategoryId: 1,
			Title:      "test show",
			Content:    "this is a test for get detail",
			Visits:     1,
		}
		id, err := dao.Create(&article)
		So(err, ShouldBeNil)
		So(id, ShouldBeGreaterThan, 0)
		detail, err := dao.Find(id)
		So(err, ShouldBeNil)
		t.Log(detail)
	})
}
