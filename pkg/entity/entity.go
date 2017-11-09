package entity

// AutoMigrateTables stores the entity tables that we can auto migrate in gorm
var AutoMigrateTables = []interface{}{
	Constraint{},
	Distribution{},
	Flag{},
	Segment{},
	User{},
	Variant{},
}
