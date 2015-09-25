package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtResourceController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtResourceController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtResourceController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtResourceController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtResourceController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtResourceController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtResourceController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtResourceController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtResourceController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtResourceController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtRoleController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtRoleController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtRoleController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtRoleController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtRoleController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtRoleController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtRoleController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtRoleController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtRoleController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtRoleController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtUserController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtUserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtUserController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtUserController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtUserController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtUserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtUserController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtUserController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtUserController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:MgtUserController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAbroadParcelController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAbroadParcelController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAbroadParcelController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAbroadParcelController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAbroadParcelController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAbroadParcelController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAbroadParcelController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAbroadParcelController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAbroadParcelController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAbroadParcelController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAddressTemplateController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAddressTemplateController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAddressTemplateController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAddressTemplateController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAddressTemplateController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAddressTemplateController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAddressTemplateController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAddressTemplateController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAddressTemplateController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyAddressTemplateController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillOrderController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillOrderController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillOrderController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillOrderController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillOrderController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillOrderController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillOrderController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillOrderController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillOrderController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillOrderController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillRecordController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillRecordController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillRecordController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillRecordController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillRecordController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillRecordController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillRecordController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillRecordController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillRecordController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyBillRecordController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyCategoryController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyCategoryController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyCategoryController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyCategoryController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyCategoryController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyCategoryController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyCategoryController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyCategoryController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyCategoryController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyCategoryController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyDomesticParcelController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyDomesticParcelController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyDomesticParcelController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyDomesticParcelController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyDomesticParcelController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyDomesticParcelController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyDomesticParcelController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyDomesticParcelController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyDomesticParcelController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyDomesticParcelController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyLinkController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyLinkController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyLinkController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyLinkController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyLinkController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyLinkController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyLinkController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyLinkController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyLinkController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyLinkController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyMenuController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyMenuController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyMenuController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyMenuController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyMenuController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyMenuController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyMenuController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyMenuController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyMenuController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyMenuController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyOptionController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyOptionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyOptionController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyOptionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyOptionController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyOptionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyOptionController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyOptionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyOptionController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyOptionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelChargeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelChargeController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelChargeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelChargeController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelChargeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelChargeController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelChargeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelChargeController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelChargeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelChargeController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelExtController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelExtController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelExtController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelExtController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelExtController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelExtController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelExtController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelExtController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelExtController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelExtController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRelationController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRelationController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRelationController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRelationController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRelationController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRelationController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRelationController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRelationController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRelationController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRelationController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRouteController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRouteController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRouteController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRouteController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRouteController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRouteController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRouteController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRouteController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRouteController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyParcelRouteController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPictureController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPictureController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPictureController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPictureController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPictureController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPictureController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPictureController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPictureController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPictureController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPictureController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPostController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPostController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPostController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPostController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPostController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPostController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPostController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPostController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPostController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyPostController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyRecommendController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyRecommendController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyRecommendController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyRecommendController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyRecommendController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyRecommendController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyRecommendController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyRecommendController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyRecommendController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyRecommendController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZySystemNoticeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZySystemNoticeController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZySystemNoticeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZySystemNoticeController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZySystemNoticeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZySystemNoticeController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZySystemNoticeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZySystemNoticeController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZySystemNoticeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZySystemNoticeController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUniqueCodeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUniqueCodeController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUniqueCodeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUniqueCodeController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUniqueCodeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUniqueCodeController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUniqueCodeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUniqueCodeController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUniqueCodeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUniqueCodeController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAccountController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAccountController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAccountController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAccountController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAccountController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAccountController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAccountController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAccountController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAccountController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAccountController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAddressController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAddressController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAddressController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAddressController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAddressController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAddressController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAddressController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAddressController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAddressController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAddressController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAdviceController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAdviceController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAdviceController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAdviceController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAdviceController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAdviceController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAdviceController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAdviceController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAdviceController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserAdviceController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserConsumeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserConsumeController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserConsumeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserConsumeController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserConsumeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserConsumeController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserConsumeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserConsumeController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserConsumeController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserConsumeController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserOptionController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserOptionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserOptionController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserOptionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserOptionController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserOptionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserOptionController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserOptionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserOptionController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserOptionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserSocialController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserSocialController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserSocialController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserSocialController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserSocialController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserSocialController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserSocialController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserSocialController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserSocialController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserSocialController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserWalletController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserWalletController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserWalletController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserWalletController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserWalletController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserWalletController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserWalletController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserWalletController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserWalletController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyUserWalletController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyVerificationRecordController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyVerificationRecordController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyVerificationRecordController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyVerificationRecordController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyVerificationRecordController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyVerificationRecordController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyVerificationRecordController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyVerificationRecordController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyVerificationRecordController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyVerificationRecordController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyWechatController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyWechatController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyWechatController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyWechatController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyWechatController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyWechatController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyWechatController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyWechatController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyWechatController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyWechatController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyZzController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyZzController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyZzController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyZzController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyZzController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyZzController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyZzController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyZzController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyZzController"] = append(beego.GlobalControllerRouter["zyun-wechat-api/controllers:ZyZzController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
