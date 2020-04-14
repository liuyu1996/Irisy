package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"product/common"
	"product/datamodels"
	"product/services"
)

type ProductController struct {
	Ctx iris.Context
	ProductService services.IProductService
}

func (p *ProductController) GetAll() mvc.View {
	productArray, _ := p.ProductService.GetAllProduct()
	return mvc.View{
		Name:"product/view.html",
		Data:iris.Map{
			"productArray":productArray,
		},
	}
}
//修改商品
func (p *ProductController) PostUpdate() {
	product := &datamodels.Product{}
	_ = p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName:"ly"})
	if err := dec.Decode(p.Ctx.Request().Form, product) ; err != nil{
		p.Ctx.Application().Logger().Debug(err)
	}
	err := p.ProductService.UpdateProduct(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

func (p *ProductController) GetAdd() mvc.View {
	return mvc.View{
		Name:"product/add.html",
	}
}

func (p *ProductController) PostAdd() {
	product := &datamodels.Product{}
	_ = p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName:"ly"})
	if err := dec.Decode(p.Ctx.Request().Form, product) ; err != nil{
		p.Ctx.Application().Logger().Debug(err)
	}
	_, err := p.ProductService.InsertProduct(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}