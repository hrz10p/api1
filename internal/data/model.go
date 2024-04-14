package data

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Models struct {
	*sql.DB
}

type ModelsI interface {
	InsertModuleInfo(info ModuleInfo) (int, error)
	RetrieveModuleInfo(id int) (ModuleInfo, error)
	UpdateModuleInfo(info ModuleInfo) error
	DeleteModuleInfo(id int) error
}

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/y.kuanyshDB?sslmode=disable"
)

func New() ModelsI {
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	return &Models{
		db,
	}
}

func (d *Models) InsertModuleInfo(info ModuleInfo) (int, error) {
	var id int
	err := d.QueryRow("INSERT INTO module_info (created_at, updated_at, module_name, module_duration, exam_type, version) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		time.Now(), time.Now(), info.ModuleName, info.ModuleDuration, info.ExamType, info.Version).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil

}

func (d *Models) RetrieveModuleInfo(id int) (ModuleInfo, error) {
	var info ModuleInfo
	err := d.QueryRow("SELECT id, created_at, updated_at, module_name, module_duration, exam_type, version FROM module_info WHERE id = $1", id).Scan(
		&info.ID, &info.CreatedAt, &info.UpdatedAt, &info.ModuleName, &info.ModuleDuration, &info.ExamType, &info.Version)
	if err != nil {
		return ModuleInfo{}, err
	}
	return info, nil
}

func (d *Models) UpdateModuleInfo(info ModuleInfo) error {
	_, err := d.Exec("UPDATE module_info SET updated_at = $1, module_name = $2, module_duration = $3, exam_type = $4, version = $5 WHERE id = $6",
		time.Now(), info.ModuleName, info.ModuleDuration, info.ExamType, info.Version, info.ID)
	if err != nil {
		return err
	}
	return nil
}

func (d *Models) DeleteModuleInfo(id int) error {
	_, err := d.Exec("DELETE FROM module_info WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
