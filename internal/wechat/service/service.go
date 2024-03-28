package service
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// custom package
	"github.com/anuwa-07/wechat/internal/wechat/model"
)

type EmployeeService struct {
	empModle *model.EmployeeModel
}

// GetEmployeeByEmail method will be used to get the employee details by email.
func (e *EmployeeService) GetEmployeeByEmail(email string) (string, error) {
	// Do some logic here...
	return "Employee", nil;
};

// CreateEmployee method will be used to create a new employee.
func (e *EmployeeService) CreateEmployee(details map[string]interface{}) error {
	fmt.Println("Goging to validate the data...")
	for key, value := range details {
		fmt.Println(key, value);
	};
	//
	err := e.empModle.CreateEmployee(details);
	if err != nil {
		fmt.Println("Opps Error!: ", err);
		return err;
	}
	return nil;
};

func NewEmployeeService(db *sql.DB) *EmployeeService {
	model := model.NewEmployee(db);
	return &EmployeeService{
		empModle: model,
	};
}