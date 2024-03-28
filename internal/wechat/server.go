package wechat
import (
	"fmt"
	"context"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gofiber/fiber/v2"
	"github.com/anuwa-07/wechat/internal/wechat/service"
	
)

type EmployeeInfo struct {
	FullName string `json:"fullname"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	City string `json:"city"`
}

type Server struct {
	listenAddr string
	employeeService *service.EmployeeService
}

func (s *Server) Start() error {
	app := fiber.New();
	// TODO: Set middleware for the application...
	// TODO: Set routes for the application...

	s.routes(app);
	//
	// Start the server...
	if err := app.Listen(s.listenAddr); err != nil {
		return fmt.Errorf("failed to start the server: %w", err)
	}
	return nil;
}

func (s *Server) routes(ap *fiber.App) {
	routes := ap.Group("/api/v1");
	//
	routes.Get("/employee/:email", s.getEmployeeByEmail);
	routes.Post("/employee/create", s.createEmployee);
	routes.Put("/employee/update", s.updateEmployee);
	routes.Delete("/employee/delete", s.deleteEmployee);
	//
	routes.Get("/health", s.healthCheck);
}

func (s *Server) getEmployeeByEmail(c *fiber.Ctx) error {
	// TODO: Implement the logic here...
	// send a test employee data...
	return c.JSON(fiber.Map{
		"email": "anuruddha@gmail.com",
		"first_name": "Anuruddha",
		"last_name": "Bandara",
		"age": 26,
		"address": "Colombo, Sri Lanka",
	});
	// Here after validating the user data. Then we need to call on the related method from the service layer.
	// s.employeeService.GetEmployeeByEmail(email); - something like this...
}

func (s *Server) createEmployee(c *fiber.Ctx) error {
	// get the `fullname`, `email`, `phone`, `city` from the request body...
	emp := new(EmployeeInfo);
	if err := c.BodyParser(emp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
			"message": "Invalid Request Body",
		});
	}
	// TODO: Check is there better way to do this ...
	detail := map[string]interface{}{
		"fullname": emp.FullName,
		"email": emp.Email,
		"phone": emp.Phone,
		"city": emp.City,
	};
	err := s.employeeService.CreateEmployee(detail);
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
			"message": "Failed to create the employee",
		});
	}
	//
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Employee created successfully!",
	});
}

func (s *Server) updateEmployee(c *fiber.Ctx) error {
	// TODO: Implement the logic here...
	return nil;
}

func (s *Server) deleteEmployee(c *fiber.Ctx) error {
	// TODO: Implement the logic here...
	return nil;
}

func (s *Server) healthCheck(c *fiber.Ctx) error {
	return c.SendString("Application is Running Bro... :)");
}
//
// NewServer creates a new server instance
func NewServer(ctx context.Context, listenAddr string, db *sql.DB) (*Server, error) {
	return &Server{
		listenAddr: listenAddr,
		employeeService: service.NewEmployeeService(db), // TODO: Update the service.EmployeeService function to accept the db connection and other scenarios.
	}, nil
}
