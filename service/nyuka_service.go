package service

import (
	"context"
	"go-studying/api/request"
	"go-studying/api/response"
	"go-studying/models"
	"go-studying/repo"
)

type nyukaService struct {
	productRepo repo.IProductRepo
	nyukaRepo   repo.INyukaRepo
	venderRepo  repo.IVenderRepo
}

func NewNyukaService(prod repo.IProductRepo, nyuka repo.INyukaRepo, vender repo.IVenderRepo) INyukaService {
	return &nyukaService{
		productRepo: prod,
		nyukaRepo:   nyuka,
		venderRepo:  vender}

}

func (ns nyukaService) GetNyukaCount(ctx context.Context, nyukaInput *request.NyukaInputRequest) (res response.NyukaInputResponse, err error) {

	info := []models.NyukaInfo{}
	info, err = ns.nyukaRepo.GetNyukaCount(ctx, nyukaInput)

	if err != nil {
		return
	}

	// TODO TEST
	res = response.NyukaInputResponse{}

	res.COUNT = len(info)

	order_list := []string{}
	for i := 0; i < len(info); i++ {
		if !isValueInList(info[i].ORDNO, order_list) {
			order_list = append(order_list, info[i].ORDNO)
		}
	}

	res.ORDERNOLIST = order_list

	return
}

func (ns nyukaService) GetNyukaJANInfo(ctx context.Context, nyukaJAN *request.NyukaInputJANRequest) (res response.NyukaInputJANResponse, err error) {

	nyukaInfo := models.NyukaInfo{}
	nyukaInfo, err = ns.nyukaRepo.GetNyukaJANInfo(ctx, nyukaJAN)
	if err != nil {
		return
	}

	// JANによる商品名取得
	janInfo := models.Product{}
	janInfo, err = ns.productRepo.GetByProductCD(ctx, nyukaJAN.ITEMCD)
	if err != nil {
		return
	}

	// SIRENOによる仕入れ名取得
	venderInfo := models.Vender{}
	venderInfo, err = ns.venderRepo.GetByVenderNo(ctx, nyukaInfo.SIRENO)
	if err != nil {
		return
	}
	
	res = response.NyukaInputJANResponse{}
	res.ORDNO = nyukaInfo.ORDNO
	res.ENO = nyukaInfo.ENO
	res.ITEMCD = nyukaInfo.JAN
	res.ITEMNM = janInfo.ITEMNM
	res.SIRENO = nyukaInfo.SIRENO
	res.SIRENM = venderInfo.SIRENM
	res.ORDERSU = nyukaInfo.ORDERSU
	res.NYUKAYOSU = nyukaInfo.NKAYOSU
	res.NYUKAJISU = nyukaInfo.NKAJISU

	return
}

func (ns nyukaService) UpdateNyukaInfo(ctx context.Context, nyukaJANJISU *request.NyukaInputJANNKAJISURequest) (IsSuccess bool, err error) {
	IsSuccess = false
	err = ns.nyukaRepo.UpdateNyukaInfo(ctx, nyukaJANJISU)
	if err != nil {
		return
	}
	IsSuccess = true
	return
}

func isValueInList(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
