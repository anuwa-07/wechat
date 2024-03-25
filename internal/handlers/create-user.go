package handlers
import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateEmployee struct {
	FullName string `json:"fullname"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	City string `json:"city"`
}
//
func (cr *CreateEmployee) validate() (bool, map[string]interface{}) {
  var isValid bool = true;
  var errors = map[string]interface{}{};
  datamap := map[string]interface{}{
    "fullname": cr.FullName,
    "email": cr.Email,
    "phone": cr.Phone,
    "city": cr.City,
  }
  //
  for key, value := range datamap {
    if value == "" {
      isValid = false;
      errors[key] = fmt.Sprintf("%s is required", key);
    }
  }
  //
  return isValid, errors;
}

// CreateUser is a handler for the sign in page
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
	  http.Error(w, "Please send a request body", http.StatusBadRequest)
	  return
	}
	defer r.Body.Close()
  
	// Read request body and decode JSON
	var user CreateEmployee  // Use CreateEmployee for consistency
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
	  http.Error(w, "Can't parse the request body", http.StatusBadRequest)
	  return
	}
  //
  // validate input
  isValid, errors := user.validate();
  if !isValid {
    response := map[string]interface{}{
      "stauts": "error",
      "message": "Input validation error",
      "info": errors,
    }
    err = json.NewEncoder(w).Encode(response)
    if err != nil {
      http.Error(w, "Can't encode response", http.StatusInternalServerError);
      return;
    }
    return;
  }
  //
	// TODO: Create user in database (not shown here)
	// ... connect to database, insert user data ...
  
	// Send JSON response (example)
	response := map[string]interface{}{
	  "status": "success",
    "message": "User created successfully",
    "info": user,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
	  http.Error(w, "Can't encode response", http.StatusInternalServerError)
	  return
	}
}
