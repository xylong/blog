package dao

import (
	"blog/internal/model"
	"blog/tools"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTag_Create(t *testing.T) {
	Convey("创建标签", t, func() {
		dao := NewTagDao()
		id, err := dao.Create(&model.Tag{
			Name: tools.RandomString(4),
		})
		So(err, ShouldBeNil)
		So(id, ShouldBeGreaterThan, 0)
	})
}

func TestTag_Delete(t *testing.T) {
	Convey("删除标签", t, func() {
		dao := NewTagDao()
		id, err := dao.Create(&model.Tag{
			Name: "delete" + tools.RandomString(4),
		})
		So(err, ShouldBeNil)
		So(id, ShouldBeGreaterThan, 0)
		ok := dao.Delete(id)
		So(ok, ShouldBeTrue)
	})
}

func TestTag_Update(t *testing.T) {
	Convey("更新标签", t, func() {
		dao := NewTagDao()
		tag := &model.Tag{
			Name: tools.RandomString(4),
		}
		id, err := dao.Create(tag)
		So(err, ShouldBeNil)
		So(id, ShouldBeGreaterThan, 0)
		tag.Name = tag.Name + "update"
		err = dao.Update(tag)
		So(err, ShouldBeNil)
	})
}

func TestTag_Select(t *testing.T) {
	Convey("查询标签", t, func() {
		dao := NewTagDao()
		_, total, err := dao.Select(0, 10, "")
		So(err, ShouldBeNil)
		So(total, ShouldBeGreaterThanOrEqualTo, 0)
		t.Logf("count:%d", total)
	})
}

func TestTag_IsExist(t *testing.T) {
	Convey("判断标签存在", t, func() {
		dao := NewTagDao()
		tag := &model.Tag{
			Name: tools.RandomString(4),
		}
		_, err := dao.Create(tag)
		So(err, ShouldBeNil)
		ok := dao.IsExist(tag.Name)
		So(ok, ShouldBeTrue)
	})
}
