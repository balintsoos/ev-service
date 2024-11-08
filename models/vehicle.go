package models

type ElectricVehicle struct {
	ID              int     `json:"id"`
	Make            string  `json:"make"`
	Model           string  `json:"model"`
	Year            int     `json:"year"`
	BatteryCapacity int     `json:"battery_capacity"`
	RangeKm         int     `json:"range_km"`
	Price           float64 `json:"price"`
}
