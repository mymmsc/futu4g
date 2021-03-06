package api

import (
	"fmt"

	"github.com/mymmsc/futu4g/protocol/Qot_GetBasicQot"
	"github.com/mymmsc/futu4g/protocol/Qot_GetBroker"
	"github.com/mymmsc/futu4g/protocol/Qot_GetCapitalDistribution"
	"github.com/mymmsc/futu4g/protocol/Qot_GetCapitalFlow"
	"github.com/mymmsc/futu4g/protocol/Qot_GetCodeChange"
	"github.com/mymmsc/futu4g/protocol/Qot_GetFutureInfo"
	"github.com/mymmsc/futu4g/protocol/Qot_GetHoldingChangeList"
	"github.com/mymmsc/futu4g/protocol/Qot_GetIpoList"
	"github.com/mymmsc/futu4g/protocol/Qot_GetKL"
	"github.com/mymmsc/futu4g/protocol/Qot_GetOptionChain"
	"github.com/mymmsc/futu4g/protocol/Qot_GetOrderBook"
	"github.com/mymmsc/futu4g/protocol/Qot_GetOwnerPlate"
	"github.com/mymmsc/futu4g/protocol/Qot_GetPlateSecurity"
	"github.com/mymmsc/futu4g/protocol/Qot_GetPlateSet"
	"github.com/mymmsc/futu4g/protocol/Qot_GetPriceReminder"
	"github.com/mymmsc/futu4g/protocol/Qot_GetRT"
	"github.com/mymmsc/futu4g/protocol/Qot_GetReference"
	"github.com/mymmsc/futu4g/protocol/Qot_GetSecuritySnapshot"
	"github.com/mymmsc/futu4g/protocol/Qot_GetStaticInfo"
	"github.com/mymmsc/futu4g/protocol/Qot_GetSubInfo"
	"github.com/mymmsc/futu4g/protocol/Qot_GetTicker"
	"github.com/mymmsc/futu4g/protocol/Qot_GetUserSecurity"
	"github.com/mymmsc/futu4g/protocol/Qot_GetWarrant"
	"github.com/mymmsc/futu4g/protocol/Qot_ModifyUserSecurity"
	"github.com/mymmsc/futu4g/protocol/Qot_RegQotPush"
	"github.com/mymmsc/futu4g/protocol/Qot_RequestHistoryKL"
	"github.com/mymmsc/futu4g/protocol/Qot_RequestHistoryKLQuota"
	"github.com/mymmsc/futu4g/protocol/Qot_RequestRehab"
	"github.com/mymmsc/futu4g/protocol/Qot_RequestTradeDate"
	"github.com/mymmsc/futu4g/protocol/Qot_SetPriceReminder"
	"github.com/mymmsc/futu4g/protocol/Qot_StockFilter"
	"github.com/mymmsc/futu4g/protocol/Qot_Sub"
	"github.com/mymmsc/futu4g/protocol/Qot_UpdateBasicQot"
	"github.com/mymmsc/futu4g/protocol/Qot_UpdateBroker"
	"github.com/mymmsc/futu4g/protocol/Qot_UpdateKL"
	"github.com/mymmsc/futu4g/protocol/Qot_UpdateOrderBook"
	"github.com/mymmsc/futu4g/protocol/Qot_UpdatePriceReminder"
	"github.com/mymmsc/futu4g/protocol/Qot_UpdateRT"
	"github.com/mymmsc/futu4g/protocol/Qot_UpdateTicker"
)

// QotSub ?????????????????????
func (api *FutuAPI) QotSub(req *Qot_Sub.Request) (<-chan *Qot_Sub.Response, error) {
	out := make(chan *Qot_Sub.Response)
	if err := api.send(ProtoIDQotSub, req, out); err != nil {
		return nil, fmt.Errorf("QotSub error: %w", err)
	}
	return out, nil
}

// QotRegQotPush ??????????????????
func (api *FutuAPI) QotRegQotPush(req *Qot_RegQotPush.Request) (<-chan *Qot_RegQotPush.Response, error) {
	out := make(chan *Qot_RegQotPush.Response)
	if err := api.send(ProtoIDQotRegQotPush, req, out); err != nil {
		return nil, fmt.Errorf("QotRegQotPush error: %w", err)
	}
	return out, nil
}

// QotGetSubInfo ??????????????????
func (api *FutuAPI) QotGetSubInfo(req *Qot_GetSubInfo.Request) (<-chan *Qot_GetSubInfo.Response, error) {
	out := make(chan *Qot_GetSubInfo.Response)
	if err := api.send(ProtoIDQotGetSubInfo, req, out); err != nil {
		return nil, fmt.Errorf("QotGetSubInfo error: %w", err)
	}
	return out, nil
}

// QotGetBasicQot ????????????????????????
func (api *FutuAPI) QotGetBasicQot(req *Qot_GetBasicQot.Request) (<-chan *Qot_GetBasicQot.Response, error) {
	out := make(chan *Qot_GetBasicQot.Response)
	if err := api.send(ProtoIDQotGetBasicQot, req, out); err != nil {
		return nil, fmt.Errorf("QotGetBasicQot error: %w", err)
	}
	return out, nil
}

// QotUpdateBasicQot ????????????????????????
func (api *FutuAPI) QotUpdateBasicQot() (<-chan *Qot_UpdateBasicQot.Response, error) {
	out := make(chan *Qot_UpdateBasicQot.Response)
	if err := api.subscribe(ProtoIDQotUpdateBasicQot, out); err != nil {
		return nil, fmt.Errorf("QotUpdateBasicQot error: %w", err)
	}
	return out, nil
}

// QotGetKL ????????????????????????
func (api *FutuAPI) QotGetKL(req *Qot_GetKL.Request) (<-chan *Qot_GetKL.Response, error) {
	out := make(chan *Qot_GetKL.Response)
	if err := api.send(ProtoIDQotGetKL, req, out); err != nil {
		return nil, fmt.Errorf("QotGetKL error: %w", err)
	}
	return out, nil
}

// QotUpdateKL ????????????????????????
func (api *FutuAPI) QotUpdateKL() (<-chan *Qot_UpdateKL.Response, error) {
	out := make(chan *Qot_UpdateKL.Response)
	if err := api.subscribe(ProtoIDQotUpdateKL, out); err != nil {
		return nil, fmt.Errorf("QotUpdateKL error: %w", err)
	}
	return out, nil
}

// QotGetRT ????????????
func (api *FutuAPI) QotGetRT(req *Qot_GetRT.Request) (<-chan *Qot_GetRT.Response, error) {
	out := make(chan *Qot_GetRT.Response)
	if err := api.send(ProtoIDQotGetRT, req, out); err != nil {
		return nil, fmt.Errorf("QotGetRT error: %w", err)
	}
	return out, nil
}

// QotUpdateRT ????????????
func (api *FutuAPI) QotUpdateRT() (<-chan *Qot_UpdateRT.Response, error) {
	out := make(chan *Qot_UpdateRT.Response)
	if err := api.subscribe(ProtoIDQotUpdateRT, out); err != nil {
		return nil, fmt.Errorf("QotUpdateRT error: %w", err)
	}
	return out, nil
}

// QotGetTicker ????????????
func (api *FutuAPI) QotGetTicker(req *Qot_GetTicker.Request) (<-chan *Qot_GetTicker.Response, error) {
	out := make(chan *Qot_GetTicker.Response)
	if err := api.send(ProtoIDQotGetTicker, req, out); err != nil {
		return nil, fmt.Errorf("QotGetTicker error: %w", err)
	}
	return out, nil
}

// QotUpdateTicker ????????????
func (api *FutuAPI) QotUpdateTicker() (<-chan *Qot_UpdateTicker.Response, error) {
	out := make(chan *Qot_UpdateTicker.Response)
	if err := api.subscribe(ProtoIDQotUpdateTicker, out); err != nil {
		return nil, fmt.Errorf("QotUpdateTicker error: %w", err)
	}
	return out, nil
}

// QotGetOrderBook ???????????????
func (api *FutuAPI) QotGetOrderBook(req *Qot_GetOrderBook.Request) (<-chan *Qot_GetOrderBook.Response, error) {
	out := make(chan *Qot_GetOrderBook.Response)
	if err := api.send(ProtoIDQotGetOrderBook, req, out); err != nil {
		return nil, fmt.Errorf("QotGetOrderBook error: %w", err)
	}
	return out, nil
}

// QotUpdateOrderBook ???????????????
func (api *FutuAPI) QotUpdateOrderBook() (<-chan *Qot_UpdateOrderBook.Response, error) {
	out := make(chan *Qot_UpdateOrderBook.Response)
	if err := api.subscribe(ProtoIDQotUpdateOrderBook, out); err != nil {
		return nil, fmt.Errorf("QotUpdateOrderBook error: %w", err)
	}
	return out, nil
}

// QotGetBroker ??????????????????
func (api *FutuAPI) QotGetBroker(req *Qot_GetBroker.Request) (<-chan *Qot_GetBroker.Response, error) {
	out := make(chan *Qot_GetBroker.Response)
	if err := api.send(ProtoIDQotGetBroker, req, out); err != nil {
		return nil, fmt.Errorf("QotGetBroker error: %w", err)
	}
	return out, nil
}

// QotUpdateBroker ??????????????????
func (api *FutuAPI) QotUpdateBroker() (<-chan *Qot_UpdateBroker.Response, error) {
	out := make(chan *Qot_UpdateBroker.Response)
	if err := api.subscribe(ProtoIDQotUpdateBroker, out); err != nil {
		return nil, fmt.Errorf("QotUpdateBroker error: %w", err)
	}
	return out, nil
}

// QotRequestRehab ????????????????????????????????????
func (api *FutuAPI) QotRequestRehab(req *Qot_RequestRehab.Request) (<-chan *Qot_RequestRehab.Response, error) {
	out := make(chan *Qot_RequestRehab.Response)
	if err := api.send(ProtoIDQotRequestRehab, req, out); err != nil {
		return nil, fmt.Errorf("QotRequestRehab error: %w", err)
	}
	return out, nil
}

// QotRequestHistoryKL ????????????????????????????????????K???
func (api *FutuAPI) QotRequestHistoryKL(req *Qot_RequestHistoryKL.Request) (<-chan *Qot_RequestHistoryKL.Response, error) {
	out := make(chan *Qot_RequestHistoryKL.Response)
	if err := api.send(ProtoIDQotRequestHistoryKL, req, out); err != nil {
		return nil, fmt.Errorf("QotRequestHistoryKL error: %w", err)
	}
	return out, nil
}

// QotGetStaticInfo ????????????????????????
func (api *FutuAPI) QotGetStaticInfo(req *Qot_GetStaticInfo.Request) (<-chan *Qot_GetStaticInfo.Response, error) {
	out := make(chan *Qot_GetStaticInfo.Response)
	if err := api.send(ProtoIDQotGetStaticInfo, req, out); err != nil {
		return nil, fmt.Errorf("QotGetStaticInfo error: %w", err)
	}
	return out, nil
}

// QotGetSecuritySnapshot ??????????????????
func (api *FutuAPI) QotGetSecuritySnapshot(req *Qot_GetSecuritySnapshot.Request) (<-chan *Qot_GetSecuritySnapshot.Response, error) {
	out := make(chan *Qot_GetSecuritySnapshot.Response)
	if err := api.subscribe(ProtoIDQotGetSecuritySnapshot, out); err != nil {
		return nil, fmt.Errorf("QotGetSecuritySnapshot error: %w", err)
	}
	return out, nil
}

// QotGetPlateSet ??????????????????????????????
func (api *FutuAPI) QotGetPlateSet(req *Qot_GetPlateSet.Request) (<-chan *Qot_GetPlateSet.Response, error) {
	out := make(chan *Qot_GetPlateSet.Response)
	if err := api.send(ProtoIDQotGetPlateSet, req, out); err != nil {
		return nil, fmt.Errorf("QotGetPlateSet error: %w", err)
	}
	return out, nil
}

// QotGetPlateSecurity ????????????????????????
func (api *FutuAPI) QotGetPlateSecurity(req *Qot_GetPlateSecurity.Request) (<-chan *Qot_GetPlateSecurity.Response, error) {
	out := make(chan *Qot_GetPlateSecurity.Response)
	if err := api.send(ProtoIDQotGetPlateSecurity, req, out); err != nil {
		return nil, fmt.Errorf("QotGetPlateSecurity error: %w", err)
	}
	return out, nil
}

// QotGetReference ????????????????????????
func (api *FutuAPI) QotGetReference(req *Qot_GetReference.Request) (<-chan *Qot_GetReference.Response, error) {
	out := make(chan *Qot_GetReference.Response)
	if err := api.send(ProtoIDQotGetReference, req, out); err != nil {
		return nil, fmt.Errorf("QotGetReference error: %w", err)
	}
	return out, nil
}

// QotGetOwnerPlate ????????????????????????
func (api *FutuAPI) QotGetOwnerPlate(req *Qot_GetOwnerPlate.Request) (<-chan *Qot_GetOwnerPlate.Response, error) {
	out := make(chan *Qot_GetOwnerPlate.Response)
	if err := api.send(ProtoIDQotGetOwnerPlate, req, out); err != nil {
		return nil, fmt.Errorf("QotGetOwnerPlate error: %w", err)
	}
	return out, nil
}

// QotGetHoldingChangeList ????????????????????????
func (api *FutuAPI) QotGetHoldingChangeList(req *Qot_GetHoldingChangeList.Request) (<-chan *Qot_GetHoldingChangeList.Response, error) {
	out := make(chan *Qot_GetHoldingChangeList.Response)
	if err := api.send(ProtoIDQotGetHoldingChangeList, req, out); err != nil {
		return nil, fmt.Errorf("QotGetHoldingChangeList error: %w", err)
	}
	return out, nil
}

// QotGetOptionChain ???????????????
func (api *FutuAPI) QotGetOptionChain(req *Qot_GetOptionChain.Request) (<-chan *Qot_GetOptionChain.Response, error) {
	out := make(chan *Qot_GetOptionChain.Response)
	if err := api.send(ProtoIDQotGetOptionChain, req, out); err != nil {
		return nil, fmt.Errorf("QotGetOptionChain error: %w", err)
	}
	return out, nil
}

// QotGetWarrant ????????????
func (api *FutuAPI) QotGetWarrant(req *Qot_GetWarrant.Request) (<-chan *Qot_GetWarrant.Response, error) {
	out := make(chan *Qot_GetWarrant.Response)
	if err := api.send(ProtoIDQotGetWarrant, req, out); err != nil {
		return nil, fmt.Errorf("QotGetWarrant error: %w", err)
	}
	return out, nil
}

// QotRequestHistoryKLQuota ????????????K????????????????????????
func (api *FutuAPI) QotRequestHistoryKLQuota(req *Qot_RequestHistoryKLQuota.Request) (<-chan *Qot_RequestHistoryKLQuota.Response, error) {
	out := make(chan *Qot_RequestHistoryKLQuota.Response)
	if err := api.send(ProtoIDQotRequestHistoryKLQuota, req, out); err != nil {
		return nil, fmt.Errorf("QotRequestHistoryKLQuota error: %w", err)
	}
	return out, nil
}

// QotGetCapitalFlow ??????????????????
func (api *FutuAPI) QotGetCapitalFlow(req *Qot_GetCapitalFlow.Request) (<-chan *Qot_GetCapitalFlow.Response, error) {
	out := make(chan *Qot_GetCapitalFlow.Response)
	if err := api.send(ProtoIDQotGetCapitalFlow, req, out); err != nil {
		return nil, fmt.Errorf("QotGetCapitalFlow error: %w", err)
	}
	return out, nil
}

// QotGetCapitalDistribution ??????????????????
func (api *FutuAPI) QotGetCapitalDistribution(req *Qot_GetCapitalDistribution.Request) (<-chan *Qot_GetCapitalDistribution.Response, error) {
	out := make(chan *Qot_GetCapitalDistribution.Response)
	if err := api.send(ProtoIDQotGetCapitalDistribution, req, out); err != nil {
		return nil, fmt.Errorf("QotGetCapitalDistribution error: %w", err)
	}
	return out, nil
}

// QotGetUserSecurity ?????????????????????????????????
func (api *FutuAPI) QotGetUserSecurity(req *Qot_GetUserSecurity.Request) (<-chan *Qot_GetUserSecurity.Response, error) {
	out := make(chan *Qot_GetUserSecurity.Response)
	if err := api.send(ProtoIDQotGetUserSecurity, req, out); err != nil {
		return nil, fmt.Errorf("QotGetUserSecurity error: %w", err)
	}
	return out, nil
}

// QotModifyUserSecurity ?????????????????????????????????
func (api *FutuAPI) QotModifyUserSecurity(req *Qot_ModifyUserSecurity.Request) (<-chan *Qot_ModifyUserSecurity.Response, error) {
	out := make(chan *Qot_ModifyUserSecurity.Response)
	if err := api.send(ProtoIDQotModifyUserSecurity, req, out); err != nil {
		return nil, fmt.Errorf("QotModifyUserSecurity error: %w", err)
	}
	return out, nil
}

// QotStockFilter ??????????????????
func (api *FutuAPI) QotStockFilter(req *Qot_StockFilter.Request) (<-chan *Qot_StockFilter.Response, error) {
	out := make(chan *Qot_StockFilter.Response)
	if err := api.send(ProtoIDQotStockFilter, req, out); err != nil {
		return nil, fmt.Errorf("QotStockFilter error: %w", err)
	}
	return out, nil
}

// QotGetCodeChange ??????????????????????????????
func (api *FutuAPI) QotGetCodeChange(req *Qot_GetCodeChange.Request) (<-chan *Qot_GetCodeChange.Response, error) {
	out := make(chan *Qot_GetCodeChange.Response)
	if err := api.send(ProtoIDQotGetCodeChange, req, out); err != nil {
		return nil, fmt.Errorf("QotGetCodeChange error: %w", err)
	}
	return out, nil
}

// QotGetIpoList ??????IPO??????
func (api *FutuAPI) QotGetIpoList(req *Qot_GetIpoList.Request) (<-chan *Qot_GetIpoList.Response, error) {
	out := make(chan *Qot_GetIpoList.Response)
	if err := api.send(ProtoIDQotGetIpoList, req, out); err != nil {
		return nil, fmt.Errorf("QotGetIpoList error: %w", err)
	}
	return out, nil
}

// QotGetFutureInfo ????????????????????????
func (api *FutuAPI) QotGetFutureInfo(req *Qot_GetFutureInfo.Request) (<-chan *Qot_GetFutureInfo.Response, error) {
	out := make(chan *Qot_GetFutureInfo.Response)
	if err := api.send(ProtoIDQotGetFutureInfo, req, out); err != nil {
		return nil, fmt.Errorf("QotGetFutureInfo error: %w", err)
	}
	return out, nil
}

// QotRequestTradeDate ?????????????????????
func (api *FutuAPI) QotRequestTradeDate(req *Qot_RequestTradeDate.Request) (<-chan *Qot_RequestTradeDate.Response, error) {
	out := make(chan *Qot_RequestTradeDate.Response)
	if err := api.send(ProtoIDQotRequestTradeDate, req, out); err != nil {
		return nil, fmt.Errorf("QotRequestTradeDate error: %w", err)
	}
	return out, nil
}

// QotSetPriceReminder ??????????????????
func (api *FutuAPI) QotSetPriceReminder(req *Qot_SetPriceReminder.Request) (<-chan *Qot_SetPriceReminder.Response, error) {
	out := make(chan *Qot_SetPriceReminder.Response)
	if err := api.send(ProtoIDQotSetPriceReminder, req, out); err != nil {
		return nil, fmt.Errorf("QotSetPriceReminder error: %w", err)
	}
	return out, nil
}

// QotGetPriceReminder ??????????????????
func (api *FutuAPI) QotGetPriceReminder(req *Qot_GetPriceReminder.Request) (<-chan *Qot_GetPriceReminder.Response, error) {
	out := make(chan *Qot_GetPriceReminder.Response)
	if err := api.send(ProtoIDQotGetPriceReminder, req, out); err != nil {
		return nil, fmt.Errorf("QotGetPriceReminder error: %w", err)
	}
	return out, nil
}

// QotUpdatePriceReminder ??????????????????
func (api *FutuAPI) QotUpdatePriceReminder() (<-chan *Qot_UpdatePriceReminder.Response, error) {
	out := make(chan *Qot_UpdatePriceReminder.Response)
	if err := api.subscribe(ProtoIDQotUpdatePriceReminder, out); err != nil {
		return nil, fmt.Errorf("QotUpdatePriceReminder error: %w", err)
	}
	return out, nil
}
