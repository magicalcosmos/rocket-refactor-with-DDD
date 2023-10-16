package entity


import {
  "testing"
  . "github.com/smartystreets/goconvey/convey"
}


func TestDeleteTestCase(t *testing.T) {
	concey("test", t, func () {
		So(1, ShouldEqual, 1)
	})
}

func TestGetTestCase(t *testing.T) {
	concey("test", t, func () {
		So(1, ShouldEqual, 1)
	})
}

func TestUpdateTestCase(t *testing.T) {
	concey("test", t, func () {
		So(1, ShouldEqual, 1)
	})
}
