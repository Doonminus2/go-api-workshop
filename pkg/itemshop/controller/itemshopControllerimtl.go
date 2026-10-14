package controller

import (
	_itemShopService "github.com/Doonminus2/go-api-workshop/pkg/itemshop/service"
)

type itemshopControllerImpl struct { 
	itemshopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl( 
	itemshopService _itemShopService.ItemShopService,

) ItemShopController {
	// {itemshopService} คือ dp injection เข้ามา
	return &itemshopControllerImpl{itemshopService}
}