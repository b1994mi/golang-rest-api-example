package util

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/uptrace/bunrouter"
)

type StatusError struct {
	HTTPCode     int
	InternalCode string
	Err          error
}

func New400Res(format string, a ...any) StatusError {
	internalCode, format := findInternalCode(format)
	return StatusError{
		InternalCode: internalCode,
		HTTPCode:     http.StatusBadRequest,
		Err:          fmt.Errorf(format, a...),
	}
}

func New401Res(format string, a ...any) StatusError {
	internalCode, format := findInternalCode(format)
	return StatusError{
		InternalCode: internalCode,
		HTTPCode:     http.StatusUnauthorized,
		Err:          fmt.Errorf(format, a...),
	}
}

func New403Res(format string, a ...any) StatusError {
	internalCode, format := findInternalCode(format)
	return StatusError{
		InternalCode: internalCode,
		HTTPCode:     http.StatusUnauthorized,
		Err:          fmt.Errorf(format, a...),
	}
}

func New404Res(format string, a ...any) StatusError {
	internalCode, format := findInternalCode(format)
	return StatusError{
		InternalCode: internalCode,
		HTTPCode:     http.StatusUnauthorized,
		Err:          fmt.Errorf(format, a...),
	}
}

func New409Res(format string, a ...any) StatusError {
	internalCode, format := findInternalCode(format)
	return StatusError{
		InternalCode: internalCode,
		HTTPCode:     http.StatusConflict,
		Err:          fmt.Errorf(format, a...),
	}
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

type StatusErrorLogged struct {
	HTTPCode int
	Err      error
}

func New500Err(format string, a ...any) StatusErrorLogged {
	return StatusErrorLogged{
		HTTPCode: http.StatusUnauthorized,
		Err:      fmt.Errorf(format, a...),
	}
}

func (sel StatusErrorLogged) Error() string {
	return sel.Err.Error()
}

// findInternalCode finds a parseable int before a ":" char,
// example: "007: unable to find something: %v" will return
// internalCode = "007" and format = "unable to find something: %v"
func findInternalCode(s string) (internalCode, format string) {
	internalCode = "000" // this is an arbitrary number

	strstr := strings.Split(s, ":")
	for i, str := range strstr {
		_, err := strconv.Atoi(str)
		if i == 0 && err == nil {
			internalCode = str
			continue
		}

		if format == "" {
			format = strings.TrimLeft(str, " ")
			continue
		}

		format = fmt.Sprintf("%v:%v", format, str)
	}

	return // don't be surprised for this naked return bcs it uses named return value ;)
	// but make sure to avoid naked returns + named return values in this code base bcs it's confusing
}

type HandlerFunc[T any] func(*T) (any, error)
type BindFunc func(any, bunrouter.Request) error

// Handler wrapper to abstract away json/uri/form binding and error response boilerplate.
// I specifically make bind funcs as variadic because it's better syntax than using slice.
func MakeHandler[T any](h HandlerFunc[T], binds ...BindFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, r bunrouter.Request) error {
		var reqBody T
		for _, b := range binds {
			err := b(&reqBody, r)
			if err != nil {
				switch e := err.(type) {
				case StatusError:
					w.WriteHeader(e.HTTPCode)
					bunrouter.JSON(w, bunrouter.H{
						"internalCode": e.InternalCode,
						"message":      e.Error(),
					})
					return err
				case StatusErrorLogged:
					log.Printf("makeHandler StatusErrorLogged: %v", err)
					w.WriteHeader(e.HTTPCode)
					bunrouter.JSON(w, bunrouter.H{
						"message": "Uh oh no, something went wrong :(",
					})
					return err
				default:
					log.Printf("makeHandler error: %v", err)
					w.WriteHeader(http.StatusInternalServerError)
					bunrouter.JSON(w, bunrouter.H{
						"message": "Uh oh no, something went wrong :(",
					})
					return err
				}
			}
		}

		obj, err := h(&reqBody)
		if err != nil {
			switch e := err.(type) {
			case StatusError:
				w.WriteHeader(e.HTTPCode)
				bunrouter.JSON(w, bunrouter.H{
					"internalCode": e.InternalCode,
					"message":      e.Error(),
				})
				return err
			case StatusErrorLogged:
				log.Printf("makeHandler StatusErrorLogged: %v", err)
				w.WriteHeader(e.HTTPCode)
				bunrouter.JSON(w, bunrouter.H{
					"message": "Uh oh no, something went wrong :(",
				})
				return err
			default:
				log.Printf("makeHandler error: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				bunrouter.JSON(w, bunrouter.H{
					"message": "Uh oh no, something went wrong :(",
				})
				return err
			}
		}

		w.WriteHeader(http.StatusOK)
		bunrouter.JSON(w, bunrouter.H{
			"status": "SUCCESS",
			"result": obj,
		})

		return err
	}
}
