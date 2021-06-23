// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

// File - тип для хранения информации о файле в базе данных
type File struct {
	ID        int
	FileName  string
	FileSize  int
	FileType  string
	FileID    string
	PreviewID string
}

// FilesList - тип для хранения списка файлов
type FilesList []File

// FilesResponse - тип для возврата с ответом,
// описывающий список файлов для постраничной разбивки
type FilesResponse struct {
	Files  FilesList
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

// Currency - тип для хранения информации
// о валюте в базе данных
type Currency struct {
	ID      int
	RusName string
	EngName string
	LatISO  string
	DigISO  string
	Symbol  string
}

// CurrenciesList - тип для хранения списка валют
type CurrenciesList []Currency

// CurrenciesResponse - тип для возврата с ответом,
// описывающий список стран для постраничной разбивки
type CurrenciesResponse struct {
	Currencies CurrenciesList
	Total      int
	Offset     int
	Limit      int
}

// Term - тип для хранения информации
// об условиях продажи, доставки и возврата
type Term struct {
	ID           int
	Name         string
	DeliveryTime string
	DeliveryCost string
	Returns      string
	Currency     Currency
}

// Terms тип для хранения списка условий
type TermsList []Term

// TermsResponse - тип для возврата с ответом,
// описывающий список условий продажи, доставки и возврата
// с данными для постраничной разбивки
type TermsResponse struct {
	SaleTerms TermsList
	Total     int
	Offset    int
	Limit     int
}

// Author - тип для хранения информации об авторе картины
type Author struct {
	ID            int
	FirstName     string
	MiddleName    string
	LastName      string
	EngName       string
	Bio           string
	OriginCity    string
	Photo         File
	OriginCountry Country
	User          User
}

// Authors - тип для хранения списка авторов картин
type AuthorsList []Author

// AuthorsResponse - тип для возврата с ответом,
// описывающий список авторов картин
// с данными для постраничной разбивки
type AuthorsResponse struct {
	ArtAuthors AuthorsList
	Total      int
	Offset     int
	Limit      int
}

// ArtworkType - тип для хранения информации о типе картины
type ArtworkType struct {
	ID      int
	Name    string
	EngName string
}

// ArtworkTypesList - тип для хранения списка типов картин
type ArtworkTypesList []ArtworkType

// ArtworkTypesResponse - тип для возврата с ответом,
// описывающий список типов картин
// с данными для постраничной разбивки
type ArtworkTypesResponse struct {
	ArtTypes ArtworkTypesList
	Total    int
	Offset   int
	Limit    int
}
