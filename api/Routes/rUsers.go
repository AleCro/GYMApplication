package Routes

import (
	Db "Svelgok-API/Database"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func EnumerateUsers(c *gin.Context) {
	user, session := GetAuthenticatedRequest(c)
	if user == nil || session == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid session",
		})
		return
	}
	if user.Group != Db.GroupAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	limitStr := c.Query("limit")
	pageStr := c.Query("page")
	search := c.Query("search")

	if limitStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Need 'limit' parameter",
		})
		return
	} else if pageStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Need 'limit' parameter",
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Paramter 'limit' must be a number",
		})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Paramter 'offset' must be a number",
		})
		return
	}

	res, err := Db.Connection.GetUsers(int64(page), int64(limit), search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	res.RemovePasswords()

	c.JSON(http.StatusOK, res)
}

func CreateUser(c *gin.Context) {
	var form RegisterRequestForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body received"})
		return
	}

	if form.Username == "" || form.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body received"})
		return
	}

	pwd, err := Db.Argon2.HashEncoded([]byte(form.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error encoding password " + err.Error()})
		return
	}

	_, found, err := Db.Connection.FilterOneUser(bson.D{{"username", form.Username}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error validating username uniqueness " + err.Error()})
		return
	}
	if found {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "User with that username already exists"})
		return
	}

	user, err := Db.Connection.InsertOneUser(&Db.User{
		Username: form.Username,
		Password: string(pwd),
		Group:    Db.GroupUser,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error " + err.Error()})
		return
	}

	session, err := Db.Connection.CreateSession(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error creating session " + err.Error(),
		})
		return
	}

	token, err := Db.CreateJWTToken(user.Claims(session.ID.Hex()))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error while creating JWT token: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"session": token,
	})

}

func EditUsers(c *gin.Context) {
	id := strings.TrimSpace(c.Param("id"))
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID provided"})
		return
	}

	var changes map[string]any
	if err := c.BindJSON(&changes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON body: " + err.Error()})
		return
	}

	// Fetch the user
	user, found, err := Db.Connection.FilterOneUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error: " + err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	// Handle reflect value safely
	uVal := reflect.ValueOf(user)
	if uVal.Kind() == reflect.Pointer {
		uVal = uVal.Elem()
	}

	uType := uVal.Type()

	for i := 0; i < uVal.NumField(); i++ {
		field := uType.Field(i)
		jsonTag := strings.Split(field.Tag.Get("json"), ",")[0]
		if jsonTag == "" || jsonTag == "-" || jsonTag == "_id" {
			continue
		}

		if newValue, ok := changes[jsonTag]; ok && newValue != nil {
			fieldVal := uVal.Field(i)
			if fieldVal.CanSet() {
				if err := setFieldFromValue(fieldVal, newValue); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"message": fmt.Sprintf("Error setting field %s: %v", jsonTag, err),
					})
					return
				}
			}
		}
	}

	user.Update()

	c.JSON(http.StatusOK, user)
}

func setFieldFromValue(field reflect.Value, value interface{}) error {
	val := reflect.ValueOf(value)

	if val.Type().ConvertibleTo(field.Type()) {
		field.Set(val.Convert(field.Type()))
		return nil
	}

	// Handle specific type conversions
	switch field.Kind() {
	case reflect.String:
		field.SetString(fmt.Sprintf("%v", value))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch v := value.(type) {
		case float64:
			field.SetInt(int64(v))
		case int:
			field.SetInt(int64(v))
		default:
			return fmt.Errorf("cannot convert %T to int", value)
		}
	case reflect.Bool:
		if boolVal, ok := value.(bool); ok {
			field.SetBool(boolVal)
		} else {
			return fmt.Errorf("cannot convert %T to bool", value)
		}
	default:
		return fmt.Errorf("unsupported field type: %s", field.Type())
	}

	return nil
}

func GetUsers(c *gin.Context) {
	user, session := GetAuthenticatedRequest(c)
	if user == nil || session == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid session",
		})
		return
	}
	if user.Group != Db.GroupAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	id := c.Param("id")
	if id == "" || id == " " {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID provided",
		})
		return
	}
	user, found, err := Db.Connection.FilterOneUser(bson.D{{"_id", id}})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Internal Server Error " + err.Error(),
		})
		return
	}
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
		})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, user)
}
