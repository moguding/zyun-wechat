// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"zyun-wechat-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/mgt_resource",
			beego.NSInclude(
				&controllers.MgtResourceController{},
			),
		),

		beego.NSNamespace("/mgt_role",
			beego.NSInclude(
				&controllers.MgtRoleController{},
			),
		),

		beego.NSNamespace("/mgt_user",
			beego.NSInclude(
				&controllers.MgtUserController{},
			),
		),

		beego.NSNamespace("/zy_abroad_parcel",
			beego.NSInclude(
				&controllers.ZyAbroadParcelController{},
			),
		),

		beego.NSNamespace("/zy_address_template",
			beego.NSInclude(
				&controllers.ZyAddressTemplateController{},
			),
		),

		beego.NSNamespace("/zy_bill_order",
			beego.NSInclude(
				&controllers.ZyBillOrderController{},
			),
		),

		beego.NSNamespace("/zy_bill_record",
			beego.NSInclude(
				&controllers.ZyBillRecordController{},
			),
		),

		beego.NSNamespace("/zy_category",
			beego.NSInclude(
				&controllers.ZyCategoryController{},
			),
		),

		beego.NSNamespace("/zy_domestic_parcel",
			beego.NSInclude(
				&controllers.ZyDomesticParcelController{},
			),
		),

		beego.NSNamespace("/zy_link",
			beego.NSInclude(
				&controllers.ZyLinkController{},
			),
		),

		beego.NSNamespace("/zy_menu",
			beego.NSInclude(
				&controllers.ZyMenuController{},
			),
		),

		beego.NSNamespace("/zy_option",
			beego.NSInclude(
				&controllers.ZyOptionController{},
			),
		),

		beego.NSNamespace("/zy_parcel",
			beego.NSInclude(
				&controllers.ZyParcelController{},
			),
		),

		beego.NSNamespace("/zy_parcel_charge",
			beego.NSInclude(
				&controllers.ZyParcelChargeController{},
			),
		),

		beego.NSNamespace("/zy_parcel_ext",
			beego.NSInclude(
				&controllers.ZyParcelExtController{},
			),
		),

		beego.NSNamespace("/zy_parcel_relation",
			beego.NSInclude(
				&controllers.ZyParcelRelationController{},
			),
		),

		beego.NSNamespace("/zy_parcel_route",
			beego.NSInclude(
				&controllers.ZyParcelRouteController{},
			),
		),

		beego.NSNamespace("/zy_picture",
			beego.NSInclude(
				&controllers.ZyPictureController{},
			),
		),

		beego.NSNamespace("/zy_post",
			beego.NSInclude(
				&controllers.ZyPostController{},
			),
		),

		beego.NSNamespace("/zy_recommend",
			beego.NSInclude(
				&controllers.ZyRecommendController{},
			),
		),

		beego.NSNamespace("/zy_system_notice",
			beego.NSInclude(
				&controllers.ZySystemNoticeController{},
			),
		),

		beego.NSNamespace("/zy_unique_code",
			beego.NSInclude(
				&controllers.ZyUniqueCodeController{},
			),
		),

		beego.NSNamespace("/zy_user_account",
			beego.NSInclude(
				&controllers.ZyUserAccountController{},
			),
		),

		beego.NSNamespace("/zy_user_address",
			beego.NSInclude(
				&controllers.ZyUserAddressController{},
			),
		),

		beego.NSNamespace("/zy_user_advice",
			beego.NSInclude(
				&controllers.ZyUserAdviceController{},
			),
		),

		beego.NSNamespace("/zy_user_consume",
			beego.NSInclude(
				&controllers.ZyUserConsumeController{},
			),
		),

		beego.NSNamespace("/zy_user_option",
			beego.NSInclude(
				&controllers.ZyUserOptionController{},
			),
		),

		beego.NSNamespace("/zy_user_social",
			beego.NSInclude(
				&controllers.ZyUserSocialController{},
			),
		),

		beego.NSNamespace("/zy_user_wallet",
			beego.NSInclude(
				&controllers.ZyUserWalletController{},
			),
		),

		beego.NSNamespace("/zy_verification_record",
			beego.NSInclude(
				&controllers.ZyVerificationRecordController{},
			),
		),

		beego.NSNamespace("/zy_wechat",
			beego.NSInclude(
				&controllers.WechatController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
