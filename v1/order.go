package v1

type Order struct {
	Owner string `json:"owner"`
	NamaBarang string `json:"namaBarang"`
	Quantity int32 `json:"quantity"`
	Harga int32 `json:"harga"`
	Ongkir int32 `json:"ongkir"`
	Customer Customer `json:"customer"`
}