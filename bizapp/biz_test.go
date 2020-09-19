package bizapp

import (
	"encoding/json"
	"testing"
)

func TestNew(t *testing.T) {
	biz := NewBiz("wsc", "微商城")

	opts := MenuOptions()

	biz.AddMenu("dashboard", "概览")
	biz.AddMenu("goods", "商品")
	biz.AddMenu("goods:list", "商品列表", opts.ParentId("goods"))
	biz.AddMenu("goods:list:add", "添加商品", opts.ParentId("goods:list"))
	biz.AddMenu("shop", "店铺")
	biz.AddMenu("shop:design", "店铺装修", opts.ParentId("shop"))

	m := NewMenu("order", "订单")
	m.AddChild("order:list", "订单列表")

	biz.AddMenus(m)

	biz.Build()

	biz.EachMenu(func(menu Menu) bool {
		t.Logf("%+v\n", menu)
		return true
	})

	data, _ := json.Marshal(biz.GetMenus())

	t.Log(string(data))

}
