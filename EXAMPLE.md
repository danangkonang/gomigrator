# example gomigrator seeder runner with fake data

```go
package seeder

import (
	"fmt"
	"os"
	"time"

	"github.com/danangkonang/gomigrator/db/migration"
	"syreclabs.com/go/faker"
)

func (s *Seeder) Users() {
	start := time.Now()
	query := `
		INSERT INTO
			users (name, created_at, updated_at)
		VALUES
			(?, ?, ?)
	`
	stmt, _ := migration.Connection().Db.Prepare(query)
	for i := 0; i < 1000000; i++ {
		_, err := stmt.Exec(
			faker.Internet().UserName(), time.Now(), time.Now(),
		)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
	}
	duration := time.Since(start)
	fmt.Println("insert table users", string(migration.Green), "success", string(migration.Reset), "in", fmt.Sprintf("%.2f", duration.Seconds()), "second")
}

```
