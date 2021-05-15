package api

const (
	ProtoIDInitConnect        = 1001 // 初始化连接
	ProtoIDGetGlobalState     = 1002 // 获取全局状态
	ProtoIDNotify             = 1003 // 通知推送
	ProtoIDKeepAlive          = 1004 // 通知推送
	ProtoIDGetUserInfo        = 1005 // 获取用户信息
	ProtoIDVerification       = 1006 // 请求或输入验证码
	ProtoIDGetDelayStatistics = 1007 // 获取延迟统计
	ProtoIDTestCmd            = 1008

	ProtoIDTrdGetAccList  = 2001 // 获取业务账户列表
	ProtoIDTrdUnlockTrade = 2005 // 解锁或锁定交易
	ProtoIDTrdSubAccPush  = 2008 // 订阅业务账户的交易推送数据

	ProtoIDTrdGetFunds        = 2101 // 获取账户资金
	ProtoIDTrdGetPositionList = 2102 // 获取账户持仓

	ProtoIDTrdGetOrderList = 2201 // 获取订单列表
	ProtoIDTrdPlaceOrder   = 2202 // 下单
	ProtoIDTrdModifyOrder  = 2205 // 修改订单
	ProtoIDTrdUpdateOrder  = 2208 // 订单状态变动通知(推送)

	ProtoIDTrdGetOrderFillList = 2211 // 获取成交列表
	ProtoIDTrdUpdateOrderFill  = 2218 // 成交通知(推送)

	ProtoIDTrdGetHistoryOrderList     = 2221 // 获取历史订单列表
	ProtoIDTrdGetHistoryOrderFillList = 2222 // 获取历史成交列表
	ProtoIDTrdGetMaxTrdQtys           = 2111 // 查询最大买卖数量

	// 订阅数据
	ProtoIDQotSub                 = 3001 // 订阅或者反订阅
	ProtoIDQotRegQotPush          = 3002 // 注册推送
	ProtoIDQotGetSubInfo          = 3003 // 获取订阅信息
	ProtoIDQotGetBasicQot         = 3004 // 获取股票基本行情
	ProtoIDQotUpdateBasicQot      = 3005 // 推送股票基本行情
	ProtoIDQotGetKL               = 3006 // 获取K线
	ProtoIDQotUpdateKL            = 3007 // 推送K线
	ProtoIDQotGetRT               = 3008 // 获取分时
	ProtoIDQotUpdateRT            = 3009 // 推送分时
	ProtoIDQotGetTicker           = 3010 // 获取逐笔
	ProtoIDQotUpdateTicker        = 3011 // 推送逐笔
	ProtoIDQotGetOrderBook        = 3012 // 获取买卖盘
	ProtoIDQotUpdateOrderBook     = 3013 // 推送买卖盘
	ProtoIDQotGetBroker           = 3014 // 获取经纪队列
	ProtoIDQotUpdateBroker        = 3015 // 推送经纪队列
	ProtoIDQotUpdatePriceReminder = 3019 //到价提醒通知

	// 历史数据
	ProtoIDQotRequestHistoryKL      = 3103 // 拉取历史K线
	ProtoIDQotRequestHistoryKLQuota = 3104 // 拉取历史K线已经用掉的额度
	ProtoIDQotRequestRehab          = 3105 // 获取除权信息

	// 其他行情数据
	ProtoIDQotGetTradeDate         = 3200 // 获取市场交易日
	ProtoIDQotGetSuspend           = 3201 // 获取股票停牌信息
	ProtoIDQotGetStaticInfo        = 3202 // 获取股票列表
	ProtoIDQotGetSecuritySnapshot  = 3203 // 获取股票快照
	ProtoIDQotGetPlateSet          = 3204 // 获取板块集合下的板块
	ProtoIDQotGetPlateSecurity     = 3205 // 获取板块下的股票
	ProtoIDQotGetReference         = 3206 // 获取正股相关股票，暂时只有窝轮
	ProtoIDQotGetOwnerPlate        = 3207 // 获取股票所属板块
	ProtoIDQotGetHoldingChangeList = 3208 // 获取高管持股变动
	ProtoIDQotGetOptionChain       = 3209 // 获取期权链

	ProtoIDQotGetOrderDetail    = 3016 // 获取委托明细
	ProtoIDQotUpdateOrderDetail = 3017 // 推送委托明细

	ProtoIDQotGetWarrant             = 3210 // 拉取窝轮信息
	ProtoIDQotGetCapitalFlow         = 3211 // 获取资金流向
	ProtoIDQotGetCapitalDistribution = 3212 // 获取资金分布

	ProtoIDQotGetUserSecurity    = 3213 // 获取自选股分组下的股票
	ProtoIDQotModifyUserSecurity = 3214 // 修改自选股分组下的股票
	ProtoIDQotStockFilter        = 3215 // 条件选股
	ProtoIDQotGetCodeChange      = 3216 // 代码变换
	ProtoIDQotGetIpoList         = 3217 // 获取新股Ipo
	ProtoIDQotGetFutureInfo      = 3218 // 获取期货资料
	ProtoIDQotRequestTradeDate   = 3219 // 在线拉取交易日
	ProtoIDQotSetPriceReminder   = 3220 // 设置到价提醒
	ProtoIDQotGetPriceReminder   = 3221 // 获取到价提醒

	ProtoIDQotGetUserSecurityGroup = 3222 // 获取自选股分组
	ProtoIDQotGetMarketState       = 3223 // 获取指定品种的市场状态
)
