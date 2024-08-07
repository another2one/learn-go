package model

import (
	"time"
)

type LewaimaiOrderHistory struct {
	Id                       uint      `gorm:"column:id;NOT NULL;primary_key"`
	AdminId                  uint      `gorm:"column:admin_id;NOT NULL"`
	ShopId                   uint      `gorm:"column:shop_id;NOT NULL"`
	LewaimaiCustomerId       uint      `gorm:"column:lewaimai_customer_id;NOT NULL"`
	OrderNo                  string    `gorm:"column:order_no;default:;NOT NULL"` // 全局唯一的订单号，应该不允许重复，这个是内部查询用的
	TradeNo                  string    `gorm:"column:trade_no;default:"`          // 订单号，这个是给商家看的
	Nickname                 string    `gorm:"column:nickname"`
	Phone                    string    `gorm:"column:phone;NOT NULL"`
	Address                  string    `gorm:"column:address;NOT NULL"`
	ReceiverLat              float64   `gorm:"column:receiver_lat;default:0.000000"` // 收货人纬度
	ReceiverLng              float64   `gorm:"column:receiver_lng;default:0.000000"` // 收货人经度
	Memo                     string    `gorm:"column:memo"`
	Configmemo               string    `gorm:"column:configmemo"`                        // 商家备注
	CancelReason             uint      `gorm:"column:cancel_reason;default:0;NOT NULL"`  // 订单取消原因 默认0无 1配送时间太长 2商家联系我取消订单 3点错了，我要重新点 4临时有事，不想要了 5其他原因
	CancelDetail             string    `gorm:"column:cancel_detail;default:;NOT NULL"`   // 订单取消详细原因
	ChargeType               uint      `gorm:"column:charge_type;default:1;NOT NULL"`    // 1：货到付款 2：余额付款 3：在线支付
	PayType                  int64     `gorm:"column:pay_type;default:1;NOT NULL"`       // 支付类型 1：乐外卖自己通道的支付宝 2：乐外卖自己通道的微信支付 3：乐外卖自己通道的财付通 4：钱方通道的微信支付（需要加余额）5：乐刷通道微信支付（需要加余额） 6：乐刷通道 支付宝（需要加余额） 7：商户自己的微信支付（非特约商户）8：商户自己的微信支付（特约商户） 9：乐刷T1结算微信支付 10：乐刷T1结算支付宝支付11、商户自己的支付宝支付
	IsDabao                  int64     `gorm:"column:is_dabao;default:0;NOT NULL"`       // 是否收取打包费
	DabaoMoney               float64   `gorm:"column:dabao_money;default:0.00;NOT NULL"` // 打包费金额
	Promotion                string    `gorm:"column:promotion;default:;NOT NULL"`
	IsMemberDelete           int64     `gorm:"column:is_member_delete;default:0;NOT NULL"`
	MemberDelete             float64   `gorm:"column:member_delete;default:0.00;NOT NULL"`
	IsDiscount               int64     `gorm:"column:is_discount;default:0;NOT NULL"`        // 是否打折
	DiscountValue            float64   `gorm:"column:discount_value;default:10.00;NOT NULL"` // 折扣值
	IsCoupon                 int64     `gorm:"column:is_coupon;default:0;NOT NULL"`          // 是否使用优惠券
	CouponValue              float64   `gorm:"column:coupon_value;default:0.00;NOT NULL"`    // 优惠券价值(这个表示使用优惠券实际优惠金额)
	CouponId                 int64     `gorm:"column:coupon_id"`                             // 优惠券ID
	DeliveryFee              float64   `gorm:"column:delivery_fee;default:0.00;NOT NULL"`    // 配送费
	TotalPrice               float64   `gorm:"column:total_price;NOT NULL"`
	BuyingPriceTotal         float64   `gorm:"column:buying_price_total;default:0.00;NOT NULL"` // 成本总价
	TotalPoint               uint      `gorm:"column:total_point;NOT NULL"`
	InitDate                 time.Time `gorm:"column:init_date;NOT NULL"` // 创建订单的时间
	ConfirmeDate             time.Time `gorm:"column:confirme_date"`      // 确认订单的时间
	CompleteDate             time.Time `gorm:"column:complete_date"`      // 完成订单的时间
	OrderField               string    `gorm:"column:order_field"`
	OrderStatus              string    `gorm:"column:order_status;NOT NULL"` // arrived表示已送达
	CourierId                int64     `gorm:"column:courier_id;default:0;NOT NULL"`
	AdminNum                 int64     `gorm:"column:admin_num"`
	ShopNum                  int64     `gorm:"column:shop_num"`
	IsPrinter                int64     `gorm:"column:is_printer;default:0;NOT NULL"`
	IsComment                int64     `gorm:"column:is_comment;default:0;NOT NULL"`  // 是否评论店铺
	IsSelftake               int64     `gorm:"column:is_selftake;default:0;NOT NULL"` // 0:商家配送，1：到店自取
	DeliveryDate             time.Time `gorm:"column:delivery_date"`
	Delivertime              string    `gorm:"column:delivertime;default:;NOT NULL"`
	Addservice               string    `gorm:"column:addservice"`
	IsFenxiao                int64     `gorm:"column:is_fenxiao;default:0;NOT NULL"`
	FenxiaoId                uint      `gorm:"column:fenxiao_id;default:0;NOT NULL"`                 // 配送来源 10自动派单 11人工派单 12抢单
	Ip                       string    `gorm:"column:ip"`                                            // 下单顾客IP
	FailedReason             string    `gorm:"column:failed_reason"`                                 // 订单失败原因
	IsSendcoupon             int64     `gorm:"column:is_sendcoupon;default:0;NOT NULL"`              // 是否发放满送优惠券
	SendcouponId             int64     `gorm:"column:sendcoupon_id;default:0;NOT NULL"`              // 发放的优惠券id
	CourierTime              time.Time `gorm:"column:courier_time"`                                  // 设置配送员时间
	IsFirstcut               int64     `gorm:"column:is_firstcut;default:0;NOT NULL"`                // 是否首单减免
	FirstcutValue            float64   `gorm:"column:firstcut_value;default:0.00;NOT NULL"`          // 首单减免金额（实际的首单减免金额）（防止统计出错）
	FirstorderValue          float64   `gorm:"column:firstorder_value;default:0.00;NOT NULL"`        // 首单减免金额（商家后台设置的减免金额）
	FoodPrice                float64   `gorm:"column:food_price;default:0.00;NOT NULL"`              // 商品原价
	DiscountPrice            float64   `gorm:"column:discount_price;default:0.00;NOT NULL"`          // 折扣优惠
	DiscountFoodDelete       float64   `gorm:"column:discount_food_delete;default:0.00;NOT NULL"`    // 特价优惠
	IsCommentCourier         int64     `gorm:"column:is_comment_courier;default:0;NOT NULL"`         // 是否评论配送员
	PhoneCustomerId          uint      `gorm:"column:phone_customer_id;default:0;NOT NULL"`          // 电话顾客id
	IsRefund                 int64     `gorm:"column:is_refund;default:0;NOT NULL"`                  // 订单是否已退款
	RefundTime               time.Time `gorm:"column:refund_time"`                                   // 退款时间
	RefundUserType           int64     `gorm:"column:refund_user_type;default:0;NOT NULL"`           // 执行退款操作的账户的类型，0表示主帐号，1表示员工帐号,2表示智能机账号
	RefundUserId             int64     `gorm:"column:refund_user_id;default:0;NOT NULL"`             // 执行退款操作的员工帐号的ID
	RefundTradeNo            string    `gorm:"column:refund_trade_no;default:;NOT NULL"`             // 退款的退款号
	RefundResult             int64     `gorm:"column:refund_result;default:0;NOT NULL"`              // 退款的处理结果，默认为0，表示还在处理中，为1表示退款成功，为2表示退款失败（例如余额不足等），is_refund字段只表示商家提交了退款操作，具体退款接口是否调用成功以 及退款结果如何，需要以这个字段为准，所有is_refund为1的订单，这个字段最终不能为0，也就是必须有个结果，要么成功要么失败
	RefundFailedReason       string    `gorm:"column:refund_failed_reason;default:;NOT NULL"`        // 订单退款失败的原因
	RefundStatus             int64     `gorm:"column:refund_status;default:0;NOT NULL"`              // 退款状态，主要用于饿了么、美团外卖等在线支付订单的处理流程 0：默认 1：顾客提交申请退款 2：退款已成功，但是设置为失败 3：餐厅拒绝退款 4：退款失败，订单变为已确 认状态 5：退款仲裁中，平台客服处理
	FromType                 int64     `gorm:"column:from_type;default:0;NOT NULL"`                  // 订单来源:0:外卖,1:后台新建,2:智能机,3:饿了么,4:美团外卖,5:百度外卖,6:小程序7:消费者APP
	AppId                    int64     `gorm:"column:app_id;NOT NULL"`                               // wx_admin_wxapp_setting数据表的自增id或者wx_admin_customerapp_setting的自增id
	IsPickup                 int64     `gorm:"column:is_pickup;default:0;NOT NULL"`                  // 此外卖订单是否取货，0：否，1：已取
	PickupTime               time.Time `gorm:"column:pickup_time"`                                   // 配送员取货的时间
	IsArrived                uint      `gorm:"column:is_arrived;default:0;NOT NULL"`                 // 是否到店 默认0否 1是
	ArrivedTime              time.Time `gorm:"column:arrived_time;NOT NULL"`                         // 配送员到店时间
	IsDelivered              int64     `gorm:"column:is_delivered;default:0"`                        // 订单是否送达0否1是
	DeliveredTime            time.Time `gorm:"column:delivered_time"`                                // 订单送达时间
	IsCount                  int64     `gorm:"column:is_count;default:0;NOT NULL"`                   // 临时字段，表示这个订单是否统计了总的订单数0未统计1已统计
	CompleteTime             int64     `gorm:"column:complete_time;default:0;NOT NULL"`              // 订单从下单到完成时间,以分为单位
	IsDelete                 int64     `gorm:"column:is_delete;default:0;NOT NULL"`                  // 顾客是否删除0显示1删除
	IncomeMoney              float64   `gorm:"column:income_money;default:0.00000;NOT NULL"`         // 每笔订单的收入
	AdminOrderNum            int64     `gorm:"column:admin_order_num;default:0;NOT NULL"`            // 顾客在平台的第几次下单
	ShopOrderNum             int64     `gorm:"column:shop_order_num;default:0;NOT NULL"`             // 顾客在店铺的第几次下单
	MeituanwaimaiOrderId     string    `gorm:"column:meituanwaimai_order_id;default:;NOT NULL"`      // 美团外卖的真实订单ID
	AdminMemo                string    `gorm:"column:admin_memo;default:;NOT NULL"`                  // 商家备注
	IsCheckRefund            int64     `gorm:"column:is_check_refund;default:0;NOT NULL"`            // 是否查询过订单退款
	RestaurantNumber         int64     `gorm:"column:restaurant_number;default:0;NOT NULL"`          // 订单流水号
	Operator                 string    `gorm:"column:operator"`                                      // Json存储操作者信息,每个字段类型1-主账号,2-员工PC,3-收银机,4-配送员
	VerifyTime               time.Time `gorm:"column:verify_time"`                                   // 核销时间
	VerifyUser               string    `gorm:"column:verify_user;default:;NOT NULL"`                 // 核销账号名称
	OrderFieldFee            float64   `gorm:"column:order_field_fee;default:0.00;NOT NULL"`         // 店铺设置的预设选项费用
	CourierType              int64     `gorm:"column:courier_type;default:1;NOT NULL"`               // 配送类型,1-商家配送,3-快服务,4-点我达,5-达达
	ShipId                   string    `gorm:"column:ship_id;default:;NOT NULL"`                     // 快服务的物流id
	SendMoney                float64   `gorm:"column:send_money;default:0.00"`                       // 第三方平台发单金额,发单金额=第三方平台配送费+溢价配送费
	YijiaMoney               float64   `gorm:"column:yijia_money;default:0.00"`                      // 第三方平台溢价配送费
	ThirdOrderPaytype        int64     `gorm:"column:third_order_paytype;default:0"`                 // 第三方平台发单付款方式(支付方式)0发单账户余额，1成功订单应得金额
	SendThirdTime            time.Time `gorm:"column:send_third_time"`                               // 发送给第三方平台配送的时间（发单时间）
	ThirdOrderNo             string    `gorm:"column:third_order_no"`                                // 第三方配送平台订单号
	ThirdOrderExt            string    `gorm:"column:third_order_ext"`                               // 发送给第三方配送的订单号扩展字段
	ThirdSendStatus          int64     `gorm:"column:third_send_status;default:0"`                   // 第三方配送发单状态：0为发单，1已发单，2发单成功，3发单失败，4发单取消
	ThirdCancelReason        string    `gorm:"column:third_cancel_reason"`                           // 第三方配送发单失败/取消原因
	JiajiaMoney              float64   `gorm:"column:jiajia_money;default:0.00"`                     // 商品加价总金额
	ThirdDeliveryFee         float64   `gorm:"column:third_delivery_fee;default:0.00;NOT NULL"`      // 第三方平台收取的配送费
	CourierName              string    `gorm:"column:courier_name;default:;NOT NULL"`                // 第三方配送平台配送员名称
	CourierPhone             string    `gorm:"column:courier_phone;default:;NOT NULL"`               // 第三方配送平台配送员手机号码
	ThirdCourierNumber       int64     `gorm:"column:third_courier_number;default:0;NOT NULL"`       // 第三方订单流水号
	ThirdCourierIsCancel     int64     `gorm:"column:third_courier_is_cancel;default:0;NOT NULL"`    // 第三方配送平台是否取消,0-正常,1-快服务取消,2-点我达取消,3-达达取消
	ThirdCourierCancelReason string    `gorm:"column:third_courier_cancel_reason;default:;NOT NULL"` // 第三方配送平台取消订单的原因
	IsVoiceNotice            int64     `gorm:"column:is_voice_notice;default:0;NOT NULL"`            // 是否已经语音提醒
	CustomerappId            int64     `gorm:"column:customerapp_id;default:0;NOT NULL"`             // 消费者appid 为wx_admin_customerapp_setting 表中的自增id
	CustomerappType          int64     `gorm:"column:customerapp_type;default:0;NOT NULL"`           // 消费者app类型，0：h5，1为安卓，2为IOS
	ReductionValue           float64   `gorm:"column:reduction_value;default:0.00;NOT NULL"`         // 优惠金额（目前只用于收银机）
	PricePlus                float64   `gorm:"column:price_plus;default:0.00;NOT NULL"`              // 订单加价金额
	PriceMoling              float64   `gorm:"column:price_moling;default:0.00;NOT NULL"`            // 订单抹零金额
	AddserviceFee            float64   `gorm:"column:addservice_fee;default:0.00;NOT NULL"`          // 增值服务费总金额
	SendToCourier            int64     `gorm:"column:send_to_courier;default:1;NOT NULL"`            // 是否发送到配送员app，1：是，0：否
	PromotionFee             float64   `gorm:"column:promotion_fee;default:0.00;NOT NULL"`           // 满减金额
	OrderFieldText           string    `gorm:"column:order_field_text"`
	AddserviceText           string    `gorm:"column:addservice_text"`
	AreaId                   int64     `gorm:"column:area_id;default:0;NOT NULL"`                        // 分区ID
	ActivityId               int64     `gorm:"column:activity_id;default:0;NOT NULL"`                    // 活动id（满减活动和首减活动，只能有其中一个）
	IsUnusual                uint      `gorm:"column:is_unusual;default:0;NOT NULL"`                     // 是否为异常订单 默认0否 1是
	IsAlwaysUnusual          uint      `gorm:"column:is_always_unusual;default:0;NOT NULL"`              // 是否一直显示异常订单 默认0否 1是
	IsAbnormal               uint      `gorm:"column:is_abnormal;default:0;NOT NULL"`                    // 是否异常订单 默认0否 1是
	IsApplyRefund            int64     `gorm:"column:is_apply_refund;default:0"`                         // 是否申请退款：0否1是
	IsRefunding              int64     `gorm:"column:is_refunding;default:0"`                            // 是否在退款中：0否1是
	RefundingDate            time.Time `gorm:"column:refunding_date"`                                    // 申请退款时间（每次申请都更新这个时间）
	IsOnceUnusual            uint      `gorm:"column:is_once_unusual;default:0;NOT NULL"`                // 是否曾经为异常订单 默认0否 1是
	DeliveryStatus           uint      `gorm:"column:delivery_status;default:0;NOT NULL"`                // 配送状态 默认0待收到 1已收到
	TransferStatus           uint      `gorm:"column:transfer_status;default:0;NOT NULL"`                // 转单状态 默认0无需转单 1转单中 2转单成功 3转单失败 4确认转单失败
	TransferReason           string    `gorm:"column:transfer_reason;default:;NOT NULL"`                 // 转单理由
	IsIgnore                 uint      `gorm:"column:is_ignore;default:0;NOT NULL"`                      // 是否忽略订单超过4小时未设置成功 默认0否 1是
	TransferId               uint      `gorm:"column:transfer_id;default:0;NOT NULL"`                    // 转单者id
	ManzengId                uint      `gorm:"column:manzeng_id;default:0;NOT NULL"`                     // 该订单参与的满赠活动id（wx_activity表id）0表示未参与满赠活动
	ManzengName              string    `gorm:"column:manzeng_name;default:;NOT NULL"`                    // 参与的满赠活动的赠品名称
	DeliveryMode             uint      `gorm:"column:delivery_mode;default:1;NOT NULL"`                  // 配送模式 默认1平台专送 2商家自配
	IsReserved               int64     `gorm:"column:is_reserved;default:0;NOT NULL"`                    // 是否是预订单：0否，1是
	Delivertimerange         uint      `gorm:"column:delivertimerange;default:30;NOT NULL"`              // 预订单最少提前多少分钟下单
	XinkeId                  uint      `gorm:"column:xinke_id;default:0;NOT NULL"`                       // 订单参与新客立减活动的活动id，wx_activity表id
	XinkeDiscount            float64   `gorm:"column:xinke_discount;default:0.00;NOT NULL"`              // 新客立减活动的实际减去的金额
	QucanStatus              int64     `gorm:"column:qucan_status;default:0"`                            // 取餐员状态 0-未取餐 1-已取餐 2-取餐完成
	Distance                 float64   `gorm:"column:distance;default:0.00"`                             // 店铺到顾客的骑行距离（前端传的）
	ZhixianDistance1         float64   `gorm:"column:zhixian_distance1;default:0.00"`                    // 店铺到顾客的直线距离（前端传的）废弃zhixian_distance  字段
	ZhixianDistance          float64   `gorm:"column:zhixian_distance"`                                  // 店铺到顾客的直线距离（前端传的）已废弃，统一用zhixian_distance1
	CourierDistanceType      int64     `gorm:"column:courier_distance_type;default:1"`                   // 骑手计算收入的距离模式 0：直线 1：骑行
	DistinctMode             int64     `gorm:"column:distinct_mode;default:1"`                           // 距离计算模式:0按直线距离计算；1按骑行距离计算
	VipPrice                 float64   `gorm:"column:vip_price;default:0.00"`                            // 购买超级会员价格
	OriginalDeliveryFee      float64   `gorm:"column:original_delivery_fee;default:0.00;NOT NULL"`       // 配送费原价
	ShopDeleteDeliveryFee    float64   `gorm:"column:shop_delete_delivery_fee;default:0.00;NOT NULL"`    // 商家满免配送费金额
	PingtaiDeleteDeliveryFee float64   `gorm:"column:pingtai_delete_delivery_fee;default:0.00;NOT NULL"` // 平台满免配送费金额
	ShopDeleteId             uint      `gorm:"column:shop_delete_id;default:0;NOT NULL"`                 // 订单参与商家配送费满免活动的活动id，wx_activity表id
	PingtaiDeleteId          uint      `gorm:"column:pingtai_delete_id;default:0;NOT NULL"`              // 订单参与平台配送费满免活动id，wx_activity表id
	ThirdType                int64     `gorm:"column:third_type;default:0"`                              // 第三方配送发单平台：1达达，2顺丰（专送），3顺丰（众包）
	NotTransferOrderTime     int64     `gorm:"column:not_transfer_order_time;default:15;NOT NULL"`       // 骑手超过多少时间不能转单（分钟）
	CancelDistributionTime   int64     `gorm:"column:cancel_distribution_time;default:10;NOT NULL"`      // 骑手接单超过多少分钟不能取消配送
	UnusualType              int64     `gorm:"column:unusual_type;default:0;NOT NULL"`                   // 异常类型 1：联系不上顾客  2顾客拒收  3顾客修改地址  4顾客定位不准   5商户位置不准   6封路/配送区域无法进入  7其他原因
	UnusualReason            string    `gorm:"column:unusual_reason;default:;NOT NULL"`                  // 设置异常订单原因
	UnusualImages            string    `gorm:"column:unusual_images;default:;NOT NULL"`                  // 异常订单图片展示
	TuancanKeyword           string    `gorm:"column:tuancan_keyword;default:;NOT NULL"`                 // 团餐地址关键字
	IsFoodDelivery           uint      `gorm:"column:is_food_delivery;default:0;NOT NULL"`               // 是否出餐 默认0否 1是
	ChucanDate               time.Time `gorm:"column:chucan_date"`                                       // 出餐时间
	CourierLng               float64   `gorm:"column:courier_lng;default:0.000000"`                      // 收货人经度
	CourierLat               float64   `gorm:"column:courier_lat;default:0.000000"`                      // 收货人纬度
	AddserviceNew            string    `gorm:"column:addservice_new;default:"`                           // 订单的预设信息
	OrderFieldNew            string    `gorm:"column:order_field_new;default:"`                          // 订单的预设信息
	ShopDeleteName           string    `gorm:"column:shop_delete_name;default:;NOT NULL"`                // 参与商家配送费满免活动名称
	PingtaiDeleteName        string    `gorm:"column:pingtai_delete_name;default:;NOT NULL"`             // 参与平台配送费满免活动名称
	RefundMoney              float64   `gorm:"column:refund_money;default:0.00"`                         // 部分退款的金额

}

func (m *LewaimaiOrderHistory) TableName() string {
	return "wx_lewaimai_order_history"
}
