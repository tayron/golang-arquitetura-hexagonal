package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/tayron/golang-arquitetura-hexagonal/adapters/cli"
	mock_application "github.com/tayron/golang-arquitetura-hexagonal/application/mocks"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ProductId := "abc"
	productName := "Product Test"
	ProductPrice := 25.99
	ProductStatus := "enabled"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(ProductId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(ProductPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(ProductStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, ProductPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(ProductId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID # %s, with the name %s has been created with the price R$ %.2f and status %s",
		ProductId, productName, ProductPrice, ProductStatus)

	result, err := cli.Run(service, "create", "", productName, ProductPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been enabled", productName)
	result, err = cli.Run(service, "enable", ProductId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been disabled", productName)
	result, err = cli.Run(service, "disable", ProductId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: R$ %.2f\nStatus:: %s\n ",
		ProductId, productName, ProductPrice, ProductStatus)
	result, err = cli.Run(service, "", ProductId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

}
