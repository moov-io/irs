package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
	"github.com/moov-io/irs/pkg/config"
	"github.com/moov-io/irs/pkg/file"
)

func parseInputFromRequest(r *http.Request) (file.File, error) {
	src, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer src.Close()

	var input bytes.Buffer
	if _, err = io.Copy(&input, src); err != nil {
		return nil, err
	}

	space := regexp.MustCompile(`\s+`)
	buf := space.ReplaceAllString(input.String(), " ")
	mf, err := file.CreateFile([]byte(buf))
	if err != nil {
		return nil, err
	}
	return mf, nil
}

func outputString(w http.ResponseWriter, output string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}

func outputJson(w http.ResponseWriter, output interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

// title: validate irs file
// path: /validator
// method: POST
// produce: multipart/form-data
// responses:
//   200: OK
//   400: Bad Request
//   501: Not Implemented
func validator(w http.ResponseWriter, r *http.Request) {
	mf, err := parseInputFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = mf.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		return
	}

	outputString(w, "valid file")
}

// title: print irs file
// path: /print
// method: POST
// produce: multipart/form-data
// responses:
//   200: OK
//   400: Bad Request
//   501: Not Implemented
func print(w http.ResponseWriter, r *http.Request) {
	mf, err := parseInputFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	format := r.FormValue("format")
	if strings.EqualFold(format, config.OutputIrsFormat) {
		outputString(w, string(mf.Ascii()))
	} else if strings.EqualFold(format, config.OutputJsonFormat) || len(format) == 0 {
		outputJson(w, mf)
	} else {
		http.Error(w, "invalid print format", http.StatusBadRequest)
	}
}

// title: convert irs file
// path: /convert
// method: POST
// produce: multipart/form-data
// responses:
//   200: OK
//   400: Bad Request
//   501: Not Implemented
func convert(w http.ResponseWriter, r *http.Request) {
	mf, err := parseInputFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	format := r.FormValue("format")
	buf, err := json.Marshal(mf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		return
	}

	filename := "irs.json"
	output := string(buf)
	if strings.EqualFold(format, config.OutputIrsFormat) {
		output = string(mf.Ascii())
		filename = "irs"
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}

// title: health server
// path: /health
// method: GET
// responses:
//   200: OK
func health(w http.ResponseWriter, r *http.Request) {
	outputJson(w, map[string]bool{"health": true})
}

// configure handlers
func ConfigureHandlers(r *mux.Router) error {
	r.HandleFunc("/health", health).Methods("GET")
	r.HandleFunc("/print", print).Methods("POST")
	r.HandleFunc("/validator", validator).Methods("POST")
	r.HandleFunc("/convert", convert).Methods("POST")
	return nil
}

// configure test handlers
func TestConfigureHandlers() (http.Handler, error) {
	r := mux.NewRouter()
	r.HandleFunc("/health", health).Methods("GET")
	r.HandleFunc("/print", print).Methods("POST")
	r.HandleFunc("/validator", validator).Methods("POST")
	r.HandleFunc("/convert", convert).Methods("POST")
	return r, nil
}
