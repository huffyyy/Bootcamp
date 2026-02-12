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

	// Find project root by looking for go.mod
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Navigate up to find go.mod
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

	outPath := filepath.Join(projectRoot, "internal", "domain")

	g := gen.NewGenerator(gen.Config{
		OutPath:           outPath,
		ModelPkgPath:      "",
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	g.UseDB(db)

	motor := g.GenerateModelAs("dealer.motors", "Motor")
	motorType := g.GenerateModelAs("dealer.motor_types", "MotorType")

	g.ApplyBasic(
		motor,
		motorType,
	)

	g.Execute()
}
