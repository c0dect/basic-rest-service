package dal

import (
	"github.com/c0dect/basic-rest-service/models"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"log"
	"strconv"
)

type ProductManager interface {
	AddProduct(prod models.Product) (newProduct models.Product, err error)
	GetProducts() ([]models.Product, error)
	GetProduct(id string) (models.Product, error)
	DeleteProduct(id string) (models.Product, error)
	UpdateProduct(id string, prod models.Product) (models.Product, error)
}

const entity_product = "Product"

type productDAL struct {
	context context.Context
}

func NewProductDAL(context context.Context) *productDAL {
	prodDAL := new(productDAL)
	prodDAL.context = context
	return prodDAL
}

func (prodDAL *productDAL) AddProduct(prod models.Product) (newProduct models.Product, err error) {

	productKey := datastore.NewIncompleteKey(prodDAL.context, entity_product, nil)
	productKey, err = datastore.Put(prodDAL.context, productKey, &prod)
	prod.ProductId = productKey.IntID()
	productKey, err = datastore.Put(prodDAL.context, productKey, &prod)

	return prod, err
}

func (prodDAL *productDAL) GetProducts() ([]models.Product, error) {

	query := datastore.NewQuery(entity_product)
	var products []models.Product
	_, err := query.GetAll(prodDAL.context, &products)

	return products, err
}

func (prodDAL *productDAL) GetProduct(productId string) (models.Product, error) {

	prodId, err := strconv.Atoi(productId)
	if err != nil {
		return models.Product{}, err
	}

	query := datastore.NewQuery(entity_product).Filter("ProductId =", prodId)

	var products []models.Product
	_, err = query.GetAll(prodDAL.context, &products)
	log.Println(products)
	if len(products) != 1 || err != nil {
		return models.Product{}, err
	}

	return products[0], nil
}

func (prodDAL *productDAL) DeleteProduct(productId string) (models.Product, error) {

	prodId, err := strconv.Atoi(productId)
	if err != nil {
		return models.Product{}, err
	}

	query := datastore.NewQuery(entity_product).Filter("ProductId =", prodId)

	var products []models.Product
	keys, err := query.GetAll(prodDAL.context, &products)
	log.Println(keys)
	if len(products) != 1 {
		return models.Product{}, err
	}
	err = datastore.Delete(prodDAL.context, keys[0])
	log.Println(err)
	if err != nil {
		return models.Product{}, err
	}
	return products[0], err
}

func (prodDAL *productDAL) UpdateProduct(productId string, requestProduct models.Product) (models.Product, error) {

	prodId, err := strconv.Atoi(productId)
	if err != nil {
		return models.Product{}, err
	}

	query := datastore.NewQuery(entity_product).Filter("ProductId =", prodId)

	var products []models.Product
	keys, err := query.GetAll(prodDAL.context, &products)
	log.Println(keys)
	if len(products) != 1 {
		return models.Product{}, err
	}
	log.Println(requestProduct)

	if requestProduct.CategoryId == "" {
		requestProduct.CategoryId = products[0].CategoryId
	}
	if requestProduct.Name == "" {
		requestProduct.Name = products[0].Name
	}
	if requestProduct.Price == 0 {
		requestProduct.Price = products[0].Price
	}
	_, err = datastore.Put(prodDAL.context, keys[0], &requestProduct)
	log.Println(err)
	if err != nil {
		return models.Product{}, err
	}
	return requestProduct, err
}
