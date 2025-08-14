package MapLearning

import "errors"

type Dictionary map[string]string

func Search(dictonary map[string]string, SearchString string) string {
	result := dictonary[SearchString]

	return result

}
func (d Dictionary) dicSearch(SearchString string)(string,error){

    definition, ok := d[SearchString]
    if !ok {
        return "", errors.New("could not find the word you were looking for")
    }

    return definition, nil
}
