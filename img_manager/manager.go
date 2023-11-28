package img_manager

import (
	"bytes"
	"github.com/h2non/filetype"
	"github.com/joho/godotenv"
	"github.com/nfnt/resize"
	"github.com/unrolled/render"
	"github.com/xbsoftware/wfs"
	local "github.com/xbsoftware/wfs-local"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var imagePath string
var uploadLimit int
var drive wfs.Drive
var format = render.New()

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Fatalf("Error read enveroment %v", e)
	}
	imagePath = os.Getenv("upload_image_path")
	uploadLimit, _ = strconv.Atoi(os.Getenv("upload_image_limit"))
	driveConfig := wfs.DriveConfig{Verbose: true}
	driveConfig.Operation = &wfs.OperationConfig{PreventNameCollision: true}
	temp := wfs.Policy(&wfs.AllowPolicy{})
	driveConfig.Policy = &temp
	var err error
	drive, err = local.NewLocalDrive(imagePath, &driveConfig)
	if err != nil {
		log.Fatal(err)
	}
}

func ListImages(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		id = "/"
	}
	config := &wfs.ListConfig{
		Nested:  true,
		Exclude: func(name string) bool { return strings.HasPrefix(name, ".") },
	}
	files, err := drive.List(id, config)
	if err != nil {
		log.Fatalf("Error listing images %v", err)
	}
	err = format.JSON(w, http.StatusOK, files)
	if err != nil {
		log.Fatalf("Error render files %v", err)
	}
}

func ImagePreview(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		id = "/"
	}
	path := imagePath + id
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Cannot read file %s, %v", path, err)
		_ = format.Text(w, http.StatusInternalServerError, err.Error())
		return
	}
	kind, _ := filetype.Match(buf)
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Cannot open image %s, %v", path, err)
		_ = format.Text(w, http.StatusInternalServerError, err.Error())
		return
	}
	var img image.Image
	switch kind.Extension {
	case "jpg":
		img, _ = jpeg.Decode(file)
		break
	case "png":
		img, _ = png.Decode(file)
		break
	}
	thumb := resize.Thumbnail(200, 200, img, resize.Lanczos3)
	writeImage(w, &thumb)
}

func writeImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func ImageUpload(w http.ResponseWriter, r *http.Request) {
	var limit = int64(32 << 20) // default is 32MB
	if int64(uploadLimit) < limit {
		limit = int64(uploadLimit)
	}
	r.Body = http.MaxBytesReader(w, r.Body, limit)
	err := r.ParseMultipartForm(limit)
	if err != nil {
		_ = format.JSON(w, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}
	file, handler, err := r.FormFile("upload")
	if err != nil {
		_ = format.JSON(w, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			_ = format.JSON(w, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}
	}(file)
	base := r.URL.Query().Get("id")
	parts := strings.Split(handler.Filename, "/")
	fileID, err := drive.Make(base, parts[len(parts)-1], false)
	if err != nil {
		_ = format.JSON(w, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}
	err = drive.Write(fileID, file)
	if err != nil {
		_ = format.JSON(w, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}
	info, err := drive.Info(fileID)
	_ = format.JSON(w, 200, info)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		_ = format.JSON(w, http.StatusNotFound, map[string]interface{}{"error": "Invalid method"})
		return
	}
	err := r.ParseForm()
	if err != nil {
		_ = format.JSON(w, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}
	ids := r.Form["ids"]
	for _, id := range ids {
		err = drive.Remove(id)
		if err != nil {
			_ = format.JSON(w, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}
	}
	_ = format.JSON(w, 200, ids)
}
