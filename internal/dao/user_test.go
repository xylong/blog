package dao

import (
	"blog/internal/model"
	_ "blog/test/data"
	"blog/tools"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestUserDao_Create(t *testing.T) {
	Convey("创建用户", t, func() {
		hash, _ := bcrypt.GenerateFromPassword([]byte("132456"), bcrypt.DefaultCost)
		dao := NewUserDao()
		user := &model.User{
			Name:     tools.RandomString(6),
			Avatar:   "",
			Phone:    fmt.Sprintf("135%s", tools.RandomNumber(8)),
			Email:    fmt.Sprintf("%s@qq.com", tools.RandomString(8)),
			Password: string(hash),
		}
		id, err := dao.Create(user)
		So(err, ShouldBeNil)
		So(id, ShouldBeGreaterThan, 0)
	})
}

func TestUserDao_IsExist(t *testing.T) {
	Convey("判断用户是否存在", t, func() {
		hash, _ := bcrypt.GenerateFromPassword([]byte("132456"), bcrypt.DefaultCost)
		dao := NewUserDao()
		user := &model.User{
			Name:     tools.RandomString(6),
			Avatar:   "",
			Phone:    fmt.Sprintf("135%s", tools.RandomNumber(8)),
			Email:    fmt.Sprintf("%s@qq.com", tools.RandomNumber(8)),
			Password: string(hash),
		}
		id, err := dao.Create(user)
		So(err, ShouldBeNil)
		So(id, ShouldBeGreaterThan, 0)
		ok1 := dao.IsExist(user.Email, "")
		ok2 := dao.IsExist("", user.Phone)
		ok3 := dao.IsExist(user.Email, user.Phone)
		So(ok1, ShouldBeTrue)
		So(ok2, ShouldBeTrue)
		So(ok3, ShouldBeTrue)
	})
}
