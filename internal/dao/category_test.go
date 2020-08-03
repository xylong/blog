package dao

import (
	"blog/internal/model"
	"blog/tools"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCategory_Create(t *testing.T) {
	Convey("创建分类", t, func() {
		dao := NewCategoryDao()
		id, err := dao.Create(&model.Category{
			Name: tools.RandomString(4),
		})
		So(err, ShouldBeNil)
		So(id, ShouldBeGreaterThan, 0)
		t.Log(id)
	})
}

func TestCategory_Delete(t *testing.T) {
	Convey("创建分类", t, func() {
		dao := NewCategoryDao()
		id, err := dao.Create(&model.Category{
			Name: tools.RandomString(4),
		})
		So(err, ShouldBeNil)
		So(id, ShouldBeGreaterThan, 0)
		ok := dao.Delete(id)
		So(ok, ShouldBeTrue)
	})
}

func TestCategory_Update(t *testing.T) {
	Convey("更新分类", t, func() {
		dao := NewCategoryDao()
		category := &model.Category{
			Name: tools.RandomString(4),
		}
		id, err := dao.Create(category)
		So(err, ShouldBeNil)
		So(id, ShouldBeGreaterThan, 0)
		category.Name = category.Name + "-update"
		err = dao.Update(category)
		So(err, ShouldBeNil)
	})
}

func TestCategory_Select(t *testing.T) {
	Convey("查询份分类", t, func() {
		dao := NewCategoryDao()
		_, total, err := dao.Select(0, 10, "")
		So(err, ShouldBeNil)
		So(total, ShouldBeGreaterThanOrEqualTo, 0)
		t.Logf("count:%d", total)
	})
}

func TestCategory_IsExist(t *testing.T) {
	Convey("判断分类是否存在", t, func() {
		dao := NewCategoryDao()
		category := &model.Category{
			Name: tools.RandomString(4),
		}
		_, err := dao.Create(category)
		So(err, ShouldBeNil)
		ok := dao.IsExist(category.Name)
		So(ok, ShouldBeTrue)
		ok = dao.IsExist(tools.RandomString(5))
		So(ok, ShouldBeFalse)
	})
}
