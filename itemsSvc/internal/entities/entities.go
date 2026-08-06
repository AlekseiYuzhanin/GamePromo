package entities

import "time"

type ShopSlot struct {
	Shop      Shop      `json:"shop"`
	Slot      Slot      `json:"slot"`
	Price     float64   `json:"price"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Shop struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Slot struct {
	Id        string     `json:"id"`
	Title     string     `json:"title"`
	Rarity    SlotRarity `json:"rarity"`
	Type      SlotType   `json:"type"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type SlotType struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type SlotRarity struct {
	Id         string    `json:"id"`
	Title      string    `json:"title"`
	Weight     float64   `json:"weight"`
	Multiplier float64   `json:"multiplier"`
	Color      string    `json:"color"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
