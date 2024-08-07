package model

import (
	"time"
)

type TongchengOrderHistory struct {
	Id                       uint      `gorm:"column:id;NOT NULL;AUTO_INCREMENT;primary_key"`
	AdminId                  uint      `gorm:"column:admin_id;NOT NULL"`
	ShopId                   uint      `gorm:"column:shop_id;default:0;NOT NULL"`    // 外卖店铺id
	TcShopId                 int64     `gorm:"column:tc_shop_id;default:0;NOT NULL"` // 同城店铺id
	LewaimaiCustomerId       uint      `gorm:"column:lewaimai_customer_id;default:0;NOT NULL"`
	OrderNo                  string    `gorm:"column:order_no;default:;NOT NULL"`               // 全局唯一的订单号，应该不允许重复，这个是内部查询用的
	TradeNo                  string    `gorm:"column:trade_no;default:"`                        // 订单号，这个是给商家看的，外部订单号
	District                 string    `gorm:"column:district;default:0"`                       // 顾客到店铺的距离(骑行距离)
	ZhixianDistrict          string    `gorm:"column:zhixian_district;default:0"`               // 顾客到店铺的距离(直线距离)
	FadanDistrictType        int64     `gorm:"column:fadan_district_type;default:1"`            // 商家发单费用的距离模式 0：直线 1：骑行
	CourierDistrictType      int64     `gorm:"column:courier_district_type;default:1"`          // 骑手计算收入的距离模式 0：直线 1：骑行
	Weight                   int64     `gorm:"column:weight;default:0;NOT NULL"`                // 重量
	Phone                    string    `gorm:"column:phone;NOT NULL"`                           // 取货店铺电话
	PhoneExt                 string    `gorm:"column:phone_ext;default:"`                       // 分机号
	Address                  string    `gorm:"column:address;NOT NULL"`                         // 取货店铺地址
	Nickname                 string    `gorm:"column:nickname"`                                 // 收货人
	ReceiverLat              float64   `gorm:"column:receiver_lat;default:0.000000"`            // 收货人纬度
	ReceiverLng              float64   `gorm:"column:receiver_lng;default:0.000000"`            // 收货人经度
	ShopPhone                string    `gorm:"column:shop_phone"`                               // 发货店铺手机号
	ShopAddress              string    `gorm:"column:shop_address"`                             // 发货店铺地址
	ShopNickname             string    `gorm:"column:shop_nickname"`                            // 发货店铺名称
	ShopReceiverLat          float64   `gorm:"column:shop_receiver_lat;default:0.000000"`       // 发货店铺纬度
	ShopReceiverLng          float64   `gorm:"column:shop_receiver_lng;default:0.000000"`       // 发货店铺经度
	Memo                     string    `gorm:"column:memo"`                                     // 订单备注
	Configmemo               string    `gorm:"column:configmemo"`                               // 商家备注
	CancelReason             uint      `gorm:"column:cancel_reason;default:0;NOT NULL"`         // 订单取消原因 默认0无 1配送时间太长 2商家联系我取消订单 3点错了，我要重新点 4临时有事，不想要了 5其他原因
	CancelDetail             string    `gorm:"column:cancel_detail;default:;NOT NULL"`          // 订单取消详细原因
	ChargeType               uint      `gorm:"column:charge_type;default:1;NOT NULL"`           // 1：货到付款 2：余额付款 3：在线支付
	PayType                  int64     `gorm:"column:pay_type;default:1;NOT NULL"`              // 支付类型 1：乐外卖自己通道的支付宝 2：乐外卖自己通道的微信支付 3：乐外卖自己通道的财付通 4：钱方通道的微信支付（需要加余额）5：乐刷通道微信支付（需要加余额） 6：乐刷通道支付宝（需要加余额） 7：商户自己的微信支付（非特约商户）8：商户自己的微信支付（特约商户） 9：乐刷T1结算微信支付 10：乐刷T1结算支付宝支付 11：商户自己的支付宝支付 12：天下支付T1微信 13：天下支付T1支付宝 14：爱财微信支付15：汇付天下微信支付16汇付天下支付宝支付
	DeliveryFee              float64   `gorm:"column:delivery_fee;default:0.00;NOT NULL"`       // 配送费
	TotalPrice               float64   `gorm:"column:total_price;default:0.00;NOT NULL"`        // 这里是指订单金额，不是同城订单的最后支付金额，如果是手动发单，这个值为0
	BuyingPriceTotal         float64   `gorm:"column:buying_price_total;default:0.00;NOT NULL"` // 成本总价
	TotalPoint               uint      `gorm:"column:total_point;NOT NULL"`
	InitDate                 time.Time `gorm:"column:init_date;NOT NULL"` // 创建订单的时间
	ConfirmeDate             time.Time `gorm:"column:confirme_date"`      // 确认订单的时间
	CompleteDate             time.Time `gorm:"column:complete_date"`      // 完成订单的时间
	OrderField               string    `gorm:"column:order_field"`
	DispatchStatus           string    `gorm:"column:dispatch_status;NOT NULL"` // 调度状态
	CourierId                int64     `gorm:"column:courier_id;default:0;NOT NULL"`
	CourierName              string    `gorm:"column:courier_name;default:;NOT NULL"`  // 配送员名称
	CourierPhone             string    `gorm:"column:courier_phone;default:;NOT NULL"` // 配送员手机号码
	AdminNum                 int64     `gorm:"column:admin_num"`
	ShopNum                  int64     `gorm:"column:shop_num"`
	IsPrinter                int64     `gorm:"column:is_printer;default:0;NOT NULL"`
	IsComment                int64     `gorm:"column:is_comment;default:0;NOT NULL"`  // 是否评论店铺
	IsSelftake               int64     `gorm:"column:is_selftake;default:0;NOT NULL"` // 0:商家配送，1：到店自取
	DeliveryDate             time.Time `gorm:"column:delivery_date"`
	Delivertime              string    `gorm:"column:delivertime;default:;NOT NULL"`
	Addservice               string    `gorm:"column:addservice"`
	Ip                       string    `gorm:"column:ip"`                                            // 下单顾客IP
	FailedReason             string    `gorm:"column:failed_reason"`                                 // 订单失败原因
	CourierTime              time.Time `gorm:"column:courier_time"`                                  // 设置配送员时间
	FoodPrice                float64   `gorm:"column:food_price;default:0.00;NOT NULL"`              // 商品原价
	IsCommentCourier         int64     `gorm:"column:is_comment_courier;default:0;NOT NULL"`         // 是否评论配送员
	PhoneCustomerId          uint      `gorm:"column:phone_customer_id;default:0;NOT NULL"`          // 电话顾客id
	IsRefund                 int64     `gorm:"column:is_refund;default:0;NOT NULL"`                  // 订单是否已退款
	RefundTime               time.Time `gorm:"column:refund_time"`                                   // 退款时间
	RefundUserType           int64     `gorm:"column:refund_user_type;default:0;NOT NULL"`           // 执行退款操作的账户的类型，0表示主帐号，1表示员工帐号,2表示智能机账号
	RefundUserId             int64     `gorm:"column:refund_user_id;default:0;NOT NULL"`             // 执行退款操作的员工帐号的ID
	RefundTradeNo            string    `gorm:"column:refund_trade_no;default:;NOT NULL"`             // 退款的退款号
	RefundResult             int64     `gorm:"column:refund_result;default:0;NOT NULL"`              // 退款的处理结果，默认为0，表示还在处理中，为1表示退款成功，为2表示退款失败（例如余额不足等），is_refund字段只表示商家提交了退款操作，具体退款接口是否调用成功以及退款结果如何，需要以这个字段为准，所有is_refund为1的订单，这个字段最终不能为0，也就是必须有个结果，要么成功要么失败
	RefundFailedReason       string    `gorm:"column:refund_failed_reason;default:;NOT NULL"`        // 订单退款失败的原因
	RefundStatus             int64     `gorm:"column:refund_status;default:0;NOT NULL"`              // 退款状态，主要用于饿了么、美团外卖等在线支付订单的处理流程 0：默认 1：顾客提交申请退款 2：退款已成功，但是设置为失败 3：餐厅拒绝退款 4：退款失败，订单变为已确认状态 5：退款仲裁中，平台客服处理
	FromType                 int64     `gorm:"column:from_type;default:0;NOT NULL"`                  // create_type指的是这个订单是在哪里创建的，from_type指的是订单最初是从哪里发来的。比如说智铺子接了美团饿了么订单，发送给乐外卖，那么create_type就是5-开放平台，from_type就是美团饿了么。 订单来源:0:乐外卖专送订单,1:美团外卖,2:饿了么,3:商家APP发单,4:开放平台（智铺子）,5:开放平台（有赞）,6:其他
	FromName                 string    `gorm:"column:from_name;default:"`                            // from_type 为6-其他时，开发者的名称
	AppId                    int64     `gorm:"column:app_id;default:0;NOT NULL"`                     // wx_admin_wxapp_setting数据表的自增id或者wx_admin_customerapp_setting的自增id
	PickupTime               time.Time `gorm:"column:pickup_time"`                                   // 配送员取货的时间
	ArrivedTime              time.Time `gorm:"column:arrived_time;NOT NULL"`                         // 配送员到店时间
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
	CourierType              int64     `gorm:"column:courier_type;default:1;NOT NULL"`               // 配送类型,1-商家配送,3-快服务,4-点我达,5-达达,6-乐外卖7智铺子
	ShipId                   string    `gorm:"column:ship_id;default:;NOT NULL"`                     // 快服务的物流id
	ThirdDeliveryFee         float64   `gorm:"column:third_delivery_fee;default:0.00;NOT NULL"`      // 第三方配送费
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
	AreaId                   int64     `gorm:"column:area_id;default:0;NOT NULL"`                   // 分区ID
	IsUnusual                uint      `gorm:"column:is_unusual;default:0;NOT NULL"`                // 是否为异常订单 默认0否 1是
	IsAlwaysUnusual          uint      `gorm:"column:is_always_unusual;default:0;NOT NULL"`         // 是否一直显示异常订单 默认0否 1是
	IsAbnormal               uint      `gorm:"column:is_abnormal;default:0;NOT NULL"`               // 是否异常订单 默认0否 1是
	IsOnceUnusual            uint      `gorm:"column:is_once_unusual;default:0;NOT NULL"`           // 是否曾经异常订单 默认0否 1是
	DeliveryStatus           uint      `gorm:"column:delivery_status;default:0;NOT NULL"`           // 配送状态 默认0待收到 1已收到
	TransferStatus           uint      `gorm:"column:transfer_status;default:0;NOT NULL"`           // 转单状态 默认0无需转单 1转单中 2转单成功 3转单失败 4确认转单失败
	TransferReason           string    `gorm:"column:transfer_reason;default:;NOT NULL"`            // 转单理由
	IsIgnore                 uint      `gorm:"column:is_ignore;default:0;NOT NULL"`                 // 是否忽略订单超过4小时未设置成功 默认0否 1是
	TransferId               uint      `gorm:"column:transfer_id;default:0;NOT NULL"`               // 转单者id
	ManzengId                uint      `gorm:"column:manzeng_id;default:0;NOT NULL"`                // 该订单参与的满赠活动id（wx_activity表id）0表示未参与满赠活动
	ManzengName              string    `gorm:"column:manzeng_name;default:;NOT NULL"`               // 参与的满赠活动的赠品名称
	DeliveryMode             uint      `gorm:"column:delivery_mode;default:1;NOT NULL"`             // 配送模式 默认1平台专送 2商家自配
	IsReserved               int64     `gorm:"column:is_reserved;default:0;NOT NULL"`               // 是否是预订单：0否，1是
	Delivertimerange         uint      `gorm:"column:delivertimerange;default:30;NOT NULL"`         // 预订单最少提前多少分钟下单
	XinkeDiscount            float64   `gorm:"column:xinke_discount;default:0.00;NOT NULL"`         // 新客立减活动的实际减去的金额
	IsOverTime               int64     `gorm:"column:is_over_time;default:0;NOT NULL"`              // 是否超时 0未超时 1超时
	OrderItem                string    `gorm:"column:order_item"`                                   // 订单详情
	WaimaiOrderId            int64     `gorm:"column:waimai_order_id;default:0"`                    // 平台专送订单id，对应lewaimai_order表的id
	Status                   int64     `gorm:"column:status;default:0"`                             // 订单状态:0未支付1、待发单（只针对美团饿了么订单） 2、已忽略（只针对美团饿了么订单） 3、待接单（商家已发单，还未分配骑手），4、取货中（已分配骑手，骑手未到店），5、已到店（骑手已到店，准备取货）6、送货中（骑手已取货，前往顾客途中），7、已送达，8、成功，9、失败 10、已取消（商家取消了待接单的订单）
	DistrictFee              float64   `gorm:"column:district_fee;default:0.00"`                    // 距离费用
	WeightFee                float64   `gorm:"column:weight_fee;default:0.00"`                      // 重量费用
	SpecialFee               float64   `gorm:"column:special_fee;default:0.00"`                     // 特殊场景费用
	TipFee                   float64   `gorm:"column:tip_fee;default:0.00"`                         // 小费
	TotalFee                 float64   `gorm:"column:total_fee;default:0.00"`                       // 总费用=距离费用+重量费用+特殊场景费用+小费
	SpecialRule              string    `gorm:"column:special_rule"`                                 // 特殊场景（本订单所有特殊场景费）
	CourierReward            float64   `gorm:"column:courier_reward;default:0.00"`                  // 特殊场景下配送员奖励金额(订单设置成功时才计算)
	DelayTime                int64     `gorm:"column:delay_time;default:0"`                         // 配送时间延长（分钟）
	IsArrived                int64     `gorm:"column:is_arrived;default:1"`                         // 是否到店 默认0否 1是
	DistributeType           int64     `gorm:"column:distribute_type;default:12"`                   // 配送来源 10自动派单 11人工派单 12抢单
	IsPickup                 int64     `gorm:"column:is_pickup;default:0"`                          // 是否取货，0：否，1：已取
	PrePickupTime            time.Time `gorm:"column:pre_pickup_time"`                              // 预计取货时间
	QucanStatus              int64     `gorm:"column:qucan_status;default:0"`                       // 取餐员状态 0-未取餐 1-已取餐 2-取餐完成 3-已刷卡 4-已交付 正常的流程是：0未取餐-》3已刷卡-》1已取餐-》4已交付-》2取餐完成
	FailedUserType           int64     `gorm:"column:failed_user_type;default:0;NOT NULL"`          // 执行失败操作的账户的类型，0表示主帐号，1表示员工帐号
	FailedUserId             int64     `gorm:"column:failed_user_id;default:0;NOT NULL"`            // 执行失败操作的员工帐号的ID
	CreateType               int64     `gorm:"column:create_type;default:0;NOT NULL"`               // 订单创建方式 1 乐外卖专送订单 2 商家app手动发单（包含各种来源） 3 美团外卖自动接单 4 饿了么自动接单 5 开放平台订单
	QucanId                  uint      `gorm:"column:qucan_id;default:0"`                           // 取餐员的employee id，取餐后会大于0
	QucanTime                time.Time `gorm:"column:qucan_time"`                                   // 取餐员的取餐时间
	ShopPreIncome            float64   `gorm:"column:shop_pre_income"`                              // 商家应得金额
	IsCheckCard              int64     `gorm:"column:is_check_card;default:1"`                      // 取餐员是否刷卡：0未刷卡1已刷卡
	TransferStationId        uint      `gorm:"column:transfer_station_id;default:0"`                // 取餐中转站id
	IsCabinet                int64     `gorm:"column:is_cabinet;default:0"`                         // 是否使用取餐柜
	CabinetType              int64     `gorm:"column:cabinet_type;default:0"`                       // 自提站点取餐方式, 1-取餐柜,2-固定站点,3-配送上楼
	IsTuancan                int64     `gorm:"column:is_tuancan;default:0"`                         // 订单是否是团餐
	TuancanAddressId         uint      `gorm:"column:tuancan_address_id;default:0"`                 // 团餐地址id
	CabinetStationId         int64     `gorm:"column:cabinet_station_id;default:0"`                 // 消费者设置的自提站点id
	FixedStationId           int64     `gorm:"column:fixed_station_id;default:0"`                   // 消费者设置的自提站点-固定站点id
	NotTransferOrderTime     int64     `gorm:"column:not_transfer_order_time;default:15;NOT NULL"`  // 骑手超过多少时间不能转单（分钟）
	CancelDistributionTime   int64     `gorm:"column:cancel_distribution_time;default:10;NOT NULL"` // 骑手接单超过多少分钟不能取消配送
	DeveloperId              int64     `gorm:"column:developer_id;default:0"`                       // 乐外卖作为开放配送平台接单时，开发者id
	UnusualType              int64     `gorm:"column:unusual_type;default:0;NOT NULL"`              // 异常类型 1：联系不上顾客  2顾客拒收  3顾客修改地址  4顾客定位不准   5商户位置不准   6封路/配送区域无法进入  7其他原因
	UnusualReason            string    `gorm:"column:unusual_reason;default:;NOT NULL"`             // 设置异常订单原因
	UnusualImages            string    `gorm:"column:unusual_images;default:;NOT NULL"`             // 异常订单图片展示
	ThirdOutTradeNo          string    `gorm:"column:third_out_trade_no"`                           // 第三方外部订单号（智铺子、美团、饿了么等第三方平台传过来的订单号）
	CheckCardTime            time.Time `gorm:"column:check_card_time"`                              // 刷卡时间
	CheckCardId              int64     `gorm:"column:check_card_id;default:0;NOT NULL"`             // 刷卡员id
	OrderOuttimeType         int64     `gorm:"column:order_outtime_type;default:0"`                 // 订单超时判断规则:0超出订单配送时间.1顾客下单选择的期望送达时间
	ExpectedDeliveryTime     time.Time `gorm:"column:expected_delivery_time"`                       // 订单的预计送达时间(对于骑手来说是配送超时时间)
	TuancanKeyword           string    `gorm:"column:tuancan_keyword;default:;NOT NULL"`            // 团餐地址关键字
	DeliveredDate            time.Time `gorm:"column:delivered_date"`                               // 已交付
	DeliveryUpstairsId       int64     `gorm:"column:delivery_upstairs_id;default:0;NOT NULL"`      // 上楼员id
	DeliveredQucanDate       time.Time `gorm:"column:delivered_qucan_date"`                         // 上楼员绑定订单时间
	DeliveryUpstairsStatus   int64     `gorm:"column:delivery_upstairs_status;default:0"`           // 上楼员状态 0-未上楼 1-上楼 2-已送达
	IsFoodDelivery           uint      `gorm:"column:is_food_delivery;default:0;NOT NULL"`          // 是否出餐 默认0否 1是
	QucanDeliverTime         time.Time `gorm:"column:qucan_deliver_time"`                           // 取餐员的交付时间
	ChucanDate               time.Time `gorm:"column:chucan_date"`                                  // 出餐时间
	IsQucanOvertime          int64     `gorm:"column:is_qucan_overtime;default:0"`                  // 取餐订单是否超时  0否  1是
	QucanOvertime            int64     `gorm:"column:qucan_overtime;default:0"`                     // 取餐超时时间（分钟）
	EstimatedDeliveryTime    time.Time `gorm:"column:estimated_delivery_time"`                      // 取餐员预计交付时间
	IsShowDelivery           int64     `gorm:"column:is_show_delivery;default:1"`                   // 是否让骑手在待抢列表看到 0否 1是
	DeliveryTimeStart        time.Time `gorm:"column:delivery_time_start"`                          // 顾客期望送达时间结束时间

}

func (m *TongchengOrderHistory) TableName() string {
	return "wx_tongcheng_order_history"
}
