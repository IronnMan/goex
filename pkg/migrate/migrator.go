package migrate

import (
	"goex/pkg/console"
	"goex/pkg/database"
	"goex/pkg/file"
	"gorm.io/gorm"
	"os"
)

type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

func (m Migrator) createMigrationsTable() {
	migration := Migration{}

	if !m.Migrator.HasTable(&migration) {
		m.Migrator.CreateTable(&migration)
	}
}

type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

func NewMigrator() *Migrator {
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}

	migrator.createMigrationsTable()

	return migrator
}

func (m *Migrator) Up() {
	migrateFiles := m.readAllMigrationFiles()

	batch := m.getBatch()

	migrations := []Migration{}
	m.DB.Find(&migrations)

	runed := false

	for _, mfile := range migrateFiles {
		if mfile.isNotMigrated(migrations) {
			m.runUpMigration(mfile, batch)
			runed = true
		}
	}

	if !runed {
		console.Success("database is up to date.")
	}
}

func (m *Migrator) Rollback() {

	lastMigration := Migration{}
	m.DB.Order("id DESC").First(&lastMigration)
	migrations := []Migration{}
	m.DB.Where("batch = ?", lastMigration.Batch).Order("id DESC").Find(&migrations)

	if !m.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to rollback.")
	}
}

func (m Migrator) readAllMigrationFiles() []MigrationFile {
	files, err := os.ReadDir(m.Folder)
	console.ExitIf(err)

	var migrateFiles []MigrationFile
	for _, f := range files {
		fileName := file.FileNameWithoutExtension(f.Name())

		mfile := getMigrationFile(fileName)

		if len(mfile.FileName) > 0 {
			migrateFiles = append(migrateFiles, mfile)
		}
	}

	return migrateFiles
}

func (m *Migrator) getBatch() int {
	batch := 1

	lastMigration := Migration{}
	m.DB.Order("id DESC").First(&lastMigration)

	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}
	return batch
}

func (m Migrator) runUpMigration(mfile MigrationFile, batch int) {
	if mfile.Up != nil {
		console.Warning("migrating " + mfile.FileName)

		mfile.Up(database.DB.Migrator(), database.SQLDB)

		console.Success("migrated " + mfile.FileName)
	}

	err := m.DB.Create(&Migration{
		Migration: mfile.FileName,
		Batch:     batch,
	}).Error

	console.ExitIf(err)
}

func (m Migrator) rollbackMigrations(migrations []Migration) bool {
	runed := false

	for _, _migration := range migrations {
		console.Warning("rollback " + _migration.Migration)

		mfile := getMigrationFile(_migration.Migration)
		if mfile.Down != nil {
			mfile.Down(database.DB.Migrator(), database.SQLDB)
		}

		runed = true

		m.DB.Delete(&_migration)

		console.Success("finish " + mfile.FileName)
	}
	return runed
}
