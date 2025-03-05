package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"os/exec"
)

type CreationInterfaceHandler struct {
	templ *template.Template
}

func NewCreationInterfaceHandler(templ *template.Template) *CreationInterfaceHandler {
	return &CreationInterfaceHandler{templ: templ}
}

func (h *CreationInterfaceHandler) Handle(w http.ResponseWriter, r *http.Request) {
	err := h.templ.ExecuteTemplate(w, "base", struct{}{})
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
}

func download(url *url.URL, path string) {
	cmd := exec.Command(
		"yt-dlp",
		"-P", path,
		url.String(),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("Error: ", err)
	}
}
