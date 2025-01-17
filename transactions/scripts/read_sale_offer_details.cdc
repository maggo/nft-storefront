import NFTStorefront from 0xNFTSTOREFRONT

// This script returns the details for a sale offer within a storefront

pub fun main(account: Address, saleOfferResourceID: UInt64): [UInt64] {
    let storefrontRef = account
        .getCapability<&NFTStorefront.Storefront{NFTStorefront.StorefrontPublic}>(
            NFTStorefront.StorefrontPublicPath
        )
        .borrow()
        ?? panic("Could not borrow public storefront from address")

    let saleOffer = storefrontRef.borrowSaleOffer(saleOfferResourceID: saleOfferResourceID)
        ?? panic("No item with that ID")
    
    return saleOffer.getDetails()
}
