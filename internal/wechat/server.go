package wechat
import (
	"fmt"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/anuwa-07/wechat/internal/wechat/service"
	"github.com/anuwa-07/wechat/pkg/sql"
)

type Server struct {
	listenAddr string
	employeeService *service.EmployeeService
}

func (s *Server) Start() error {
	app := fiber.New();

	// TODO: Set middleware for the application...
	// TODO: Set routes for the application...

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
	return nil;

	// Here after validating the user data. Then we need to call on the related method from the service layer.
	// s.employeeService.GetEmployeeByEmail(email); - something like this...
}

func (s *Server) createEmployee(c *fiber.Ctx) error {
	// TODO: Implement the logic here...
	return nil;
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
	return c.SendString("OK");
}
//
// NewServer creates a new server instance
func NewServer(ctx context.Context, listenAddr string, db *sql.DBConfig) (*Server, error) {
	return &Server{
		listenAddr: listenAddr,
		employeeService: service.EmployeeService(db), // TODO: Update the service.EmployeeService function to accept the db connection and other scenarios.
	}, nil
}
