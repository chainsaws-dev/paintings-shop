// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

// File - тип для хранения информации о файле в базе данных
type File struct {
	ID        int
	Filename  string
	Filesize  int
	Filetype  string
	FileID    string
	PreviewID string
}

// FilesList - тип для хранения списка файлов
type FilesList []File

// FilesResponse - тип для возврата с ответом,
// описывающий список файлов для постраничной разбивки
type FilesResponse struct {
	Files  FilesList
	Edit   bool
	Delete bool
	Total  int
	Offset int
	Limit  int
}

// Country - тип для хранения информации
// о стране в базе данных
type Country struct {
	ID              int
	Name            string
	FullName        string
	English         string
	Alpha2          string
	Alpha3          string
	ISO             string
	Location        string
	LocationPrecise string
}

// CountriesList - тип для хранения списка стран
type CountriesList []Country

// CountriesResponse - тип для возврата с ответом,
// описывающий список стран для постраничной разбивки
type CountriesResponse struct {
	Countries CountriesList
	Total     int
	Offset    int
	Limit     int
}
