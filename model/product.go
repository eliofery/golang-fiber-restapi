package model

import "time"

type Product struct {
	Id               int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Sku              string    `json:"sku"`
	Name             string    `json:"name"`
	Stock            int       `json:"stock"`
	Price            int       `json:"price"`
	Image            string    `json:"image"`
	TotalFinalPrice  int       `json:"totalFinalPrice"`
	TotalNormalPrice int       `json:"totalNormalPrice"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	CategoryId       int       `json:"categoryId"`
	DiscountId       int       `json:"discountId"`
}
