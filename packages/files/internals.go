package files

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"paintings-shop/packages/databases"
	"paintings-shop/packages/setup"
	"paintings-shop/packages/shared"
	"path/filepath"
	"strings"

	"github.com/bakape/thumbnailer/v2"
)

// fileUpload - выполняет загрузку файла на сервер и сохранение в файловой системе и информации в базе данных
func fileUpload(w http.ResponseWriter, req *http.Request, role string) (databases.File, error) {

	log.Println("Начинаем получение файла...")

	var NewFile databases.File
	// Читаем из формы двоичные данные
	f, fh, err := req.FormFile("file")

	if shared.HandleInternalServerError(w, err) {
		return NewFile, err
	}
	defer f.Close()

	// Подключаемся к базе данных
	dbc := setup.ServerSettings.SQL.Connect(w, role)
	if dbc == nil {
		shared.HandleOtherError(w, databases.ErrNoConnection.Error(), databases.ErrNoConnection, http.StatusServiceUnavailable)
		return NewFile, databases.ErrNoConnection
	}
	defer dbc.Close()

	// Проверяем тип файла
	buff := make([]byte, 512)
	_, err = f.Read(buff)

	if shared.HandleInternalServerError(w, err) {
		return NewFile, err
	}

	filetype := http.DetectContentType(buff)

	if filetype == "image/jpeg" || filetype == "image/jpg" || filetype == "image/gif" ||
		filetype == "image/png" || filetype == "application/pdf" {

		// На всякий случай сохраняем расширение
		ext := strings.Split(fh.Filename, ".")[1]

		fn := sha1.New()

		io.Copy(fn, f)

		basename := fmt.Sprintf("%x", fn.Sum(nil))

		filename := basename + "." + ext

		linktofile := strings.Join([]string{"uploads", filename}, "/")

		pathtofile := filepath.Join(".", "public", "uploads", filename)

		nf, err := os.Create(pathtofile)

		if shared.HandleInternalServerError(w, err) {
			return NewFile, err
		}

		defer nf.Close()

		_, err = f.Seek(0, 0)

		if shared.HandleInternalServerError(w, err) {
			return NewFile, err
		}

		_, err = io.Copy(nf, f)

		if shared.HandleInternalServerError(w, err) {
			return NewFile, err
		}

		log.Printf("Файл получен и сохранён под именем %s", filename)

		NewFile.Filename = fh.Filename
		NewFile.Filesize = int(fh.Size)
		NewFile.Filetype = ext
		NewFile.FileID = filename

		// Создаем превьюшку
		_, thumb, err := thumbnailer.Process(f, thumbnailer.Options{})

		previewname := basename + "pv.png"

		linktopreview := strings.Join([]string{"uploads", previewname}, "/")

		pathtopreview := filepath.Join(".", "public", "uploads", previewname)

		nfp, err := os.Create(pathtopreview)

		if shared.HandleInternalServerError(w, err) {
			return NewFile, err
		}

		defer nfp.Close()

		png.Encode(nfp, thumb)

		NewFile.PreviewID = previewname

		NewFile.ID, err = databases.PostgreSQLFileChange(NewFile, dbc)

		if err != nil {
			if errors.Is(databases.ErrFirstNotUpdate, err) {
				shared.HandleOtherError(w, err.Error(), err, http.StatusBadRequest)
				return NewFile, err
			}

			if shared.HandleInternalServerError(w, err) {
				return NewFile, err
			}
		}

		NewFile.FileID = linktofile
		NewFile.PreviewID = linktopreview

	} else {

		shared.HandleOtherError(w, ErrUnsupportedFileType.Error(), ErrUnsupportedFileType, http.StatusBadRequest)
		return NewFile, ErrUnsupportedFileType
	}

	return NewFile, nil

}
