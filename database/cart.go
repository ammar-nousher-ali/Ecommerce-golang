package controllers

import "errors"

var (
	ErrCantFindProducts   = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("can't find the product")
	ErrUserIdIsNotValid   = errors.New("this user is not valid")
	ErroCantUpdateUser    = errors.New("cannot add this product to the cart")
	ErrCantRemoveItemCart = errors.New("cannot remove this item form cart")
	ErrCantGetItem        = errors.New("was unable to get item from the cart")
	ErrCanBuyCartItem     = errors.New("cannot update the purchase")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
