package main

import (
	"flag"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/ez4bk/ezboot/tools/reflect"
	_ "github.com/go-sql-driver/mysql"

	"example.ezboot/myna/dbgen/mods"
	"example.ezboot/myna/dbgen/mods/myna"
)

var options struct {
	DSN         string
	withDefault bool
	mods        []mods.Mod
}

func init() {
	flag.StringVar(&options.DSN, "dsn", "", "database connection string")
	flag.BoolVar(&options.withDefault, "with-default", true,
		"generate DAO with the with package; false injects *xorm.Engine and *redis.Client")
	flag.Parse()

	if env := os.Getenv("DBGEN_DSN"); env != "" {
		options.DSN = env
	}

	options.mods = append(options.mods, &myna.Mod{})
}

func main() {
	for _, elem := range options.mods {
		logger := slog.With("name", elem.Name())
		if policy := elem.Policy(); policy != nil {
			ref, err := reflect.New(options.DSN, reflect.WithPolicy(policy))
			if err != nil {
				logger.Error("failed to generate models", "error", err, "name", elem.Name())
				os.Exit(1)
			}

			for table, tablePolicy := range policy.Tables {
				if err = generate(ref, table, elem, tablePolicy); err != nil {
					logger.Error("failed to generate models", "error", err, "table", table)
					os.Exit(1)
				}
			}
		}
	}
}

func generate(r *reflect.Reflection, table string, mod mods.Mod, policy *reflect.TablePolicy) error {
	modelFilename := filepath.Join(mod.Target(), table+".go")
	if err := os.MkdirAll(filepath.Dir(modelFilename), 0755); err != nil {
		return err
	}

	modelFile, err := os.Create(modelFilename)
	if err != nil {
		return err
	}
	defer func() { _ = modelFile.Close() }()

	renderOptions := []reflect.RenderOption{
		reflect.WithFieldTagMapper(reflect.JsonTagMapper()),
		reflect.WithFieldTagMapper(reflect.XORMTagMapper()),
	}

	if daoTarget := mod.DaoTarget(); daoTarget != "" && policy.DAOPolicy != nil {
		daoFilename := filepath.Join(daoTarget, table+".dao.go")
		if err = os.MkdirAll(filepath.Dir(daoFilename), 0755); err != nil {
			return err
		}

		daoFile, err := os.Create(daoFilename)
		if err != nil {
			return err
		}

		renderOptions = append(renderOptions,
			reflect.WithDAORender(daoFile),
			reflect.WithDAORenderOptions(reflect.WithDAOWithDefault(options.withDefault)),
		)
	}

	return r.Render(table, modelFile, renderOptions...)
}
