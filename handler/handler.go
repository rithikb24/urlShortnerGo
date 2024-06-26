package urlshort

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"net/http"
	"reflect"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...

	redirect := func(w http.ResponseWriter, r *http.Request) {
		path, ok := pathsToUrls[r.URL.Path]
		fmt.Println(path, ok)
		fmt.Println(reflect.TypeOf(path), reflect.TypeOf(ok))
		if ok {
			http.Redirect(w, r, path, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
	return redirect
}

type Urlmap []struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYAML(yml []byte) Urlmap {
	var parsedYaml Urlmap
	err := yaml.Unmarshal([]byte(yml), &parsedYaml)
	if err != nil {
		panic(err)
	}
	return parsedYaml

}

func buildMap(parsedYaml Urlmap) map[string]string {
	builtMap := make(map[string]string)
	for i := 0; i < len(parsedYaml); i++ {
		fmt.Println(parsedYaml[i].Path)
		builtMap[parsedYaml[i].Path] = parsedYaml[i].URL
	}
	// fmt.Println(builtMap)
	return builtMap
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	parsedYaml := parseYAML(yml)
	pathMap := buildMap(parsedYaml)
	fmt.Println(pathMap)
	return MapHandler(pathMap, fallback), nil

}
