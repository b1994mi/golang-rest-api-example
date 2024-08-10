package util

import (
	"encoding/json"
	"fmt"
	"io"
	// "os"
	// "path"
	"reflect"
	"strconv"

	"github.com/uptrace/bunrouter"
)

func GetLimitOffset(page, size int) (limit int, offset int) {
	if page <= 0 || size <= 0 {
		// using -1 to disable gorm size and offset in case page and size not set
		size = -1
		offset = -1
		return size, offset
	}
	offset = (page - 1) * size
	return size, offset
}

// ShouldBindJSON is a shortcut io.ReadAll(); json.Unmarshal()
func ShouldBindJSON(obj any, r bunrouter.Request) error {
	rv := reflect.ValueOf(obj)
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("please pass a pointer")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, obj)
	if err != nil {
		return err
	}

	return nil
}

// ShouldBindQuery is a shortcut for bunrouter.ParamsFromContext(); params.ByName() on passed obj every field with tag "uri".
// For now, those fields with tag "uri" must be either string or int.
func ShouldBindUri(obj any, r bunrouter.Request) error {
	rv := reflect.ValueOf(obj)
	if rv.Kind() == reflect.Ptr {
		rv = reflect.Indirect(rv)
	}

	t := rv.Type()
	params := r.Params()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		val, ok := field.Tag.Lookup("uri")
		if !ok {
			continue
		}

		extractedParam := params.ByName(val)

		rvField := rv.FieldByName(field.Name)
		switch rvField.Kind() {
		case reflect.Int:
			parsedInt, err := strconv.Atoi(extractedParam)
			if err != nil {
				return err
			}
			rvField.Set(reflect.ValueOf(parsedInt))
		case reflect.String:
			rvField.Set(reflect.ValueOf(extractedParam))
		default:
			return fmt.Errorf("can not set %v", rvField.Kind())
		}
	}

	return nil
}

// ShouldBindQuery is a shortcut for bunReq.URL.Query().Get() on passed obj every field with tag "form".
// If you want to parse multipart/form-data, then just use bunReq.ParseMultipartForm(); bunReq.MultipartForm.
func ShouldBindQuery(obj any, r bunrouter.Request) error {
	rv := reflect.ValueOf(obj)
	if rv.Kind() == reflect.Ptr {
		rv = reflect.Indirect(rv)
	}

	t := rv.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		val, ok := field.Tag.Lookup("form")
		if !ok {
			continue
		}

		rv.FieldByName(field.Name).Set(reflect.ValueOf(
			r.URL.Query().Get(val),
		))
	}

	return nil
}

// TODO: create should bind multipart form, maybe using tag `mpffield`, `mpfjson`, and `mpffile`
// func ShouldBindForm(obj any, r bunrouter.Request ) error {
// 	err := r.ParseMultipartForm(64)
// 	if err != nil {
// 		return err
// 	}

// 	rv := reflect.ValueOf(obj)
// 	if rv.Kind() == reflect.Ptr {
// 		rv = reflect.Indirect(rv)
// 	}

// 	t := rv.Type()
// 	for i := 0; i < t.NumField(); i++ {
// 		field := t.Field(i)
// 		val, ok := field.Tag.Lookup("multipartform")
// 		if !ok {
// 			continue
// 		}

// 		rv.FieldByName(field.Name).Set(reflect.ValueOf(
// 			r.URL.Query().Get(val),
// 		))
// 	}

// 	var req complaint.CreateComplaintRequest
// 	for _, v := range r.MultipartForm.Value["issue"] {
// 		json.Unmarshal([]byte(v), &req)
// 		break
// 	}

// 	req.AuthenticatedUser = r.Context().Value("user_id").(int)

// 	wd, err := os.Getwd()
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnprocessableEntity)
// 		bunrouter.JSON(w, bunrouter.H{
// 			"code":    "007",
// 			"message": err,
// 		})

// 		return nil
// 	}

// 	files := r.MultipartForm.File["image"]
// 	for _, f := range files {
// 		dir := path.Join(wd, "/tmp/media")
// 		if _, err := os.Stat(dir); os.IsNotExist(err) {
// 			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
// 				w.WriteHeader(http.StatusUnprocessableEntity)
// 				bunrouter.JSON(w, bunrouter.H{
// 					"code":    "007",
// 					"message": err,
// 				})

// 				return nil
// 			}
// 		}

// 		tmpFile := path.Join(dir, fmt.Sprintf("%v-%v", time.Now().Unix(), f.Filename))
// 		if err = helper.SaveUploadedFile(f, tmpFile); err != nil {
// 			w.WriteHeader(http.StatusUnprocessableEntity)
// 			bunrouter.JSON(w, bunrouter.H{
// 				"code":    "007",
// 				"message": err,
// 			})

// 			return nil
// 		}

// 		req.Media = append(req.Media, tmpFile)
// 	}
// }
