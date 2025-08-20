package main
import (
	"learn_golang/models"
	"gorm.io/gen"
)

func main() {
  g := gen.NewGenerator(gen.Config{
    OutPath: "./database/myquery",
    Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface, // generate mode
  })

  // Generate basic type-safe DAO API for struct `model.User` following conventions
  g.ApplyBasic(models.User{})

  // Generate the code
  g.Execute()
}