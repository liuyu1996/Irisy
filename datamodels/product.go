package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"ID" ly:"ID"`
	ProductName  string `json:"ProductName" sql:"productName" ly:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum" ly:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"productImage" ly:"ProductImage"`
	ProductUrl   string `json:"ProductUrl" sql:"productUrl" ly:"ProductUrl"`
}
