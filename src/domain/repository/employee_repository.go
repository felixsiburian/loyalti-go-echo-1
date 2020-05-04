package repository

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

type EmployeeRepository interface {
	CreateEmployee(newemployee *model.Employee) error
	UpdateEmployee(newemployee *model.Employee) error
	DeleteEmployee(newemployee *model.Employee) error
}

type employee_repo struct {
	DB *gorm.DB
}

func (p employee_repo) CreateEmployee(newemployee *model.Employee) error {
	fmt.Println("masuk fungsi")
	employee := model.Employee{
		Created:       time.Now(),
		CreatedBy:     "",
		Modified:      time.Now(),
		ModifiedBy:    "",
		Active:        false,
		IsDeleted:     false,
		Deleted:       time.Time{},
		Deleted_by:    "",
		EmployeeName:  newemployee.EmployeeName,
		EmployeeEmail: newemployee.EmployeeEmail,
		EmployeePin:   newemployee.EmployeePin,
		EmployeeRole:  newemployee.EmployeeRole,
		OutletId:      newemployee.OutletId,
		OutletName:    newemployee.OutletName,
	}

	db := database.ConnectionDB()
	err := db.Create(&employee).Error
	if err != nil {
		fmt.Println("Tak ada error")
	}
	return err
}

func CreateEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employee_repo{
		DB: db,
	}
}

func (p *employee_repo) UpdateEmployee(updateemployee *model.Employee) error {
	db := database.ConnectionDB()
	employee := model.Employee{
		Created:       time.Now(),
		CreatedBy:     "",
		Modified:      time.Now(),
		ModifiedBy:    "",
		Active:        true,
		IsDeleted:     false,
		Deleted:       time.Time{},
		Deleted_by:    "",
		EmployeeName:  updateemployee.EmployeeName,
		EmployeeEmail: updateemployee.EmployeeEmail,
		EmployeePin:   updateemployee.EmployeePin,
		EmployeeRole:  updateemployee.EmployeeRole,
		OutletId:      updateemployee.OutletId,
		OutletName:    updateemployee.OutletName,
	}
	err := db.Model(&employee).Where("outlet_id = ?", employee.OutletId).Update(&employee).Error
	return err
}

func (p *employee_repo) DeleteEmployee(deleteemployee *model.Employee) error {
	db := database.ConnectionDB()

	err := db.Model(&deleteemployee).Where("id = ?", deleteemployee.Id).Update("active", true).Error
	if err == nil {
		fmt.Println("tidak ada error")
	}
	return err
}

func CreateEmployee(employee *model.Employee) string {
	db := database.ConnectionDB()
	emp := model.Employee{
		Created:       time.Now(),
		CreatedBy:     "Admin",
		Modified:      time.Now(),
		ModifiedBy:    "Admin",
		Active:        true,
		IsDeleted:     false,
		Deleted:       time.Now(),
		Deleted_by:    "",
		EmployeeName:  employee.EmployeeName,
		EmployeeEmail: employee.EmployeeEmail,
		EmployeePin:   employee.EmployeePin,
		EmployeeRole:  employee.EmployeeRole,
		OutletId:      employee.OutletId,
		OutletName:    "a",
	}
	db.Create(&emp)
	defer db.Close()
	return employee.EmployeeEmail
}

func UpdateEmployee(employee *model.Employee) string {
	db := database.ConnectionDB()
	db.Model(&employee).Where("employee_email = ?", employee.EmployeeEmail).Update(&employee)
	employee.Modified = time.Now()
	defer db.Close()
	return employee.EmployeeEmail
}

func DeleteEmployee(employee *model.Employee) string {
	db := database.ConnectionDB()
	employee.Deleted = time.Now()
	db.Model(&employee).Where("employee_email = ?", employee.EmployeeEmail).Update("active", false)
	db.Model(&employee).Where("employee_email = ?", employee.EmployeeEmail).Update("is_deleted", true)

	//db.Model(&employee).Where("employee_email = ?",employee.EmployeeEmail).Update("active", false)
	defer db.Close()
	return "berhasil dihapus"
}

func GetEmployee(page *int, size *int, sort *int) []model.Employee {
	db := database.ConnectionDB()
	var employee []model.Employee
	var rows *sql.Rows
	var err error
	var total int

	if page == nil && size == nil && sort == nil {
		rows, err = db.Find(&employee).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil {
		fmt.Println("masuk kesini")
		rows, err = db.Order("employee_name asc").Find(&employee).Count(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil && sort != nil {
		fmt.Println("masuk sort")
		switch *sort {
		case 1:
			rows, err = db.Find(&employee).Order("employee_name asc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		case 2:
			rows, err = db.Find(&employee).Order("employee_name desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	result := make([]model.Employee, 0)

	for rows.Next() {
		e := &model.Employee{}
		fmt.Println(e)

		err = rows.Scan(
			&e.Id,
			&e.Created,
			&e.CreatedBy,
			&e.Modified,
			&e.ModifiedBy,
			&e.Active,
			&e.IsDeleted,
			&e.Deleted,
			&e.Deleted_by,
			&e.EmployeeName,
			&e.EmployeeEmail,
			&e.EmployeePin,
			&e.EmployeeRole,
			&e.OutletId,
			&e.OutletName,
		)

		outlet := new(model.Outlet2)
		db.Table("outlet2").
			Select("outlet2.outlet_name").
			Where("id = ? ", e.OutletId).
			First(&outlet)
		e.OutletName = outlet.OutletName

		if err != nil {
			log.Fatal(err)
		}
		result = append(result, *e)
	}

	db.Close()
	return result
}
