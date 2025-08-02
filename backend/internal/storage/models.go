package storage

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `db:"id"`
	Phone        string    `db:"phone"`
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type UserAddress struct {
	ID          uuid.UUID `db:"id"`
	UserID      uuid.UUID `db:"user_id"`
	AddressText string    `db:"address_text"`
	Coordinates string    `db:"coordinates"` // WKT / GeoJSON
	IsDefault   bool      `db:"is_default"`
}

type Category struct {
	ID       uuid.UUID  `db:"id"`
	Name     string     `db:"name"`
	ParentID *uuid.UUID `db:"parent_id"`
	ImageURL string     `db:"image_url"`
}

type Product struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	WeightGrams int       `db:"weight_grams"`
	Available   bool      `db:"available"`
}

type ProductCategory struct {
	ID         uuid.UUID `db:"id"`
	ProductID  uuid.UUID `db:"product_id"`
	CategoryID uuid.UUID `db:"category_id"`
}

type ProductImage struct {
	ID        uuid.UUID `db:"id"`
	ProductID uuid.UUID `db:"product_id"`
	URL       string    `db:"url"`
	Type      string    `db:"type"`
}

type Warehouse struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Address     string    `db:"address"`
	Coordinates string    `db:"coordinates"`
}

type WarehouseCategory struct {
	ID          uuid.UUID `db:"id"`
	WarehouseID uuid.UUID `db:"warehouse_id"`
	CategoryID  uuid.UUID `db:"category_id"`
}

type WarehouseStock struct {
	ID          uuid.UUID `db:"id"`
	WarehouseID uuid.UUID `db:"warehouse_id"`
	ProductID   uuid.UUID `db:"product_id"`
	Quantity    int       `db:"quantity"`
	Price       float64   `db:"price"`
}

type Order struct {
	ID         uuid.UUID  `db:"id"`
	UserID     *uuid.UUID `db:"user_id"`
	AddressID  *uuid.UUID `db:"address_id"`
	Status     string     `db:"status"`
	TotalPrice float64    `db:"total_price"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
}

type OrderItem struct {
	ID        uuid.UUID `db:"id"`
	OrderID   uuid.UUID `db:"order_id"`
	ProductID uuid.UUID `db:"product_id"`
	Quantity  int       `db:"quantity"`
	Price     float64   `db:"price"`
}

type Promotion struct {
	ID        uuid.UUID  `db:"id"`
	Title     string     `db:"title"`
	Subtitle  string     `db:"subtitle"`
	ImageURL  string     `db:"image_url"`
	Deeplink  string     `db:"deeplink"`
	SortOrder int        `db:"sort_order"`
	IsActive  bool       `db:"is_active"`
	StartsAt  *time.Time `db:"starts_at"`
	EndsAt    *time.Time `db:"ends_at"`
}

type Story struct {
	ID        uuid.UUID `db:"id"`
	Title     string    `db:"title"`
	ImageURL  string    `db:"image_url"`
	Deeplink  string    `db:"deeplink"`
	SortOrder int       `db:"sort_order"`
	IsActive  bool      `db:"is_active"`
}

type ProductCollection struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	SortOrder   int       `db:"sort_order"`
	IsActive    bool      `db:"is_active"`
}

type ProductCollectionItem struct {
	ID           uuid.UUID `db:"id"`
	CollectionID uuid.UUID `db:"collection_id"`
	ProductID    uuid.UUID `db:"product_id"`
	Position     int       `db:"position"`
}
