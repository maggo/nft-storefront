package test

// go test -timeout 30s . -run ^TestNFTStorefront -v

import (
	"testing"
)

func TestNFTStorefrontDeployContracts(t *testing.T) {
	b := newEmulator()
	nftStorefrontDeployContracts(t, b)
}

func TestNFTStorefrontSetupAccount(t *testing.T) {
	b := newEmulator()

	contracts := nftStorefrontDeployContracts(t, b)

	userAddress, userSigner := createAccount(t, b)
	setupNFTStorefront(t, b, userAddress, userSigner, contracts)
}

func TestNFTStorefrontCreateSaleSell(t *testing.T) {
	b := newEmulator()

	contracts := nftStorefrontDeployContracts(t, b)

	t.Run("Should be able to list a sale offer", func(t *testing.T) {
		tokenToList := uint64(0)
		tokenPrice := "1.11"

		sellerAddress, sellerSigner := createAccount(t, b)
		setupAccount(t, b, sellerAddress, sellerSigner, contracts)

		// Contract mints item to seller account
		mintExampleNFT(
			t,
			b,
			sellerAddress,
			contracts.NFTAddress,
			contracts.ExampleNFTAddress,
			contracts.ExampleNFTSigner,
		)

		// Seller account lists the item
		sellItem(
			t,
			b,
			contracts,
			sellerAddress,
			sellerSigner,
			tokenToList,
			tokenPrice,
			false,
		)
	})

	t.Run("Should be able to accept a sale offer", func(t *testing.T) {
		tokenToList := uint64(1)
		tokenPrice := "1.11"

		sellerAddress, sellerSigner := createAccount(t, b)
		setupAccount(t, b, sellerAddress, sellerSigner, contracts)

		// Contract mints item to seller account
		mintExampleNFT(
			t,
			b,
			sellerAddress,
			contracts.NFTAddress,
			contracts.ExampleNFTAddress,
			contracts.ExampleNFTSigner,
		)

		// Seller account lists the item
		saleOfferResourceID := sellItem(
			t,
			b,
			contracts,
			sellerAddress,
			sellerSigner,
			tokenToList,
			tokenPrice,
			false,
		)

		buyerAddress, buyerSigner := createAccount(t, b)
		setupAccount(t, b, buyerAddress, buyerSigner, contracts)

		// Make the purchase
		buyItem(
			b,
			t,
			contracts,
			buyerAddress,
			buyerSigner,
			sellerAddress,
			saleOfferResourceID,
			false,
		)
	})

	t.Run("Should be able to remove a sale offer", func(t *testing.T) {
		tokenToList := uint64(2)
		tokenPrice := "1.11"

		sellerAddress, sellerSigner := createAccount(t, b)
		setupAccount(t, b, sellerAddress, sellerSigner, contracts)

		// Contract mints item to seller account
		mintExampleNFT(
			t,
			b,
			sellerAddress,
			contracts.NFTAddress,
			contracts.ExampleNFTAddress,
			contracts.ExampleNFTSigner,
		)

		// Seller account lists the item
		saleOfferResourceID := sellItem(
			t,
			b,
			contracts,
			sellerAddress,
			sellerSigner,
			tokenToList,
			tokenPrice,
			false,
		)

		// Cancel the sale
		removeItem(
			b,
			t,
			contracts,
			sellerAddress,
			sellerSigner,
			saleOfferResourceID,
			false,
		)
	})
}
