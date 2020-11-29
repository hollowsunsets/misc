import (
	"net/http"
)


func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	// TODO: implement
	return nil
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: implement
	return nil, nil
}