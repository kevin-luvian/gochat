package gorm

import (
	"gochat/env"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type tModel struct {
	gorm.Model
	State string
}

func init() {
	env.LoadDotEnvForTest()
}

func makeTestDB() *gorm.DB {
	db := GetDB()
	db.Migrator().DropTable(&tModel{})
	db.AutoMigrate(&tModel{})
	return db
}

func makeNewModel() tModel {
	return tModel{State: "new state"}
}

func isModelEqual(a, b tModel) bool {
	return a.State == b.State
}

func TestFindAll(t *testing.T) {
	db := makeTestDB()
	var models []tModel

	result := db.Find(&models)
	if result.Error != nil {
		t.Fatal(result.Error.Error())
	}

	if int(result.RowsAffected)+len(models) > 0 {
		t.Fatalf(`there shouldn't be a record in new db`)
	}
}

func TestCreate(t *testing.T) {
	db := makeTestDB()
	var model = makeNewModel()

	result := db.Create(&model)
	if result.Error != nil {
		t.Fatal(result.Error.Error())
	}

	raff := result.RowsAffected
	if raff == 0 || model.ID == 0 {
		t.Fatalf(`db is not updated`)
	}

	var models []tModel
	result = db.Find(&models)
	if result.Error != nil {
		t.Fatal(result.Error.Error())
	}

	raff = result.RowsAffected
	if int(raff)+len(models) == 0 {
		t.Fatalf(`there are no new record in db`)
	}

	if !isModelEqual(models[0], model) {
		t.Fatalf(`model in db is not equal with the given model`)
	}
}

func TestUpdate(t *testing.T) {
	db := makeTestDB()
	var model = makeNewModel()
	var updateTo = tModel{State: "updated state"}

	result := db.Create(&model)
	if result.Error != nil {
		t.Fatal(result.Error.Error())
	}

	var updated []tModel
	result = db.Model(&updated).Clauses(clause.Returning{}).Where("id = ?", model.ID).Updates(updateTo)
	if result.Error != nil {
		t.Fatal(result.Error.Error())
	}

	if result.RowsAffected != 1 || len(updated) != 1 {
		t.Fatalf(`db is not updated`)
	}

	if !isModelEqual(updated[0], updateTo) {
		t.Fatalf(`model in db is not equal with the given model`)
	}
}
