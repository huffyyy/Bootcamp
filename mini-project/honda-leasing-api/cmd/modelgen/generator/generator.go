package main

import (
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=admin123 dbname=leasing_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	projectRoot := wd
	for {
		if _, err := os.Stat(filepath.Join(projectRoot, "go.mod")); err == nil {
			break
		}
		parent := filepath.Dir(projectRoot)
		if parent == projectRoot {
			log.Fatal("Could not find project root (go.mod)")
		}
		projectRoot = parent
	}

	queryPath := filepath.Join(projectRoot, "internal/domain/query")
	modelPath := filepath.Join(projectRoot, "internal/domain/models")

	g := gen.NewGenerator(gen.Config{
		OutPath:           queryPath,
		ModelPkgPath:      modelPath,
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	g.UseDB(db)

	motor := g.GenerateModelAs("dealer.motors", "Motor")
	motorType := g.GenerateModelAs("dealer.motor_types", "MotorType")

	customer := g.GenerateModelAs("dealer.customers", "Customer")

	paymentSchedule := g.GenerateModelAs("finance.payment_schedule", "PaymentSchedule")

	leasingProduct := g.GenerateModelAs("leasing.leasing_product", "LeasingProduct")
	leasingContract := g.GenerateModelAs("leasing.leasing_contract", "LeasingContract")
	leasingTask := g.GenerateModelAs("leasing.leasing_tasks", "LeasingTask")
	leasingTaskAtr := g.GenerateModelAs("leasing.leasing_tasks_attributes", "LeasingTaskAtr")

	g.ApplyBasic(
		motor,
		motorType,

		customer,

		paymentSchedule,

		leasingProduct,
		leasingContract,
		leasingTask,
		leasingTaskAtr,
	)

	g.Execute()
}
