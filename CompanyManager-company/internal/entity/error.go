package entity
import "errors"

//ErrNotFound not found
var ErrNotFound = errors.New("Not found")

//ErrInvalidEntity invalid presenter
var ErrInvalidEntity = errors.New("Invalid presenter")

//ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("Cannot Be Deleted")

//ErrCannotBeCreated cannot be created
var ErrCannotBeCreated = errors.New("Cannot Be Created")

