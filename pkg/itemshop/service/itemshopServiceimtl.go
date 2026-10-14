package service

import (
	_itemShopRepository "github.com/Doonminus2/go-api-workshop/pkg/itemshop/repository"
)
type itemshopServiceImpl struct { 
	itemShopRepository _itemShopRepository.ItemShopRepository	
}

func NewItemShopServiceImpl(
	itemShopRepository _itemShopRepository.ItemShopRepository,
) ItemShopService { 
	return &itemshopServiceImpl{itemShopRepository}

}