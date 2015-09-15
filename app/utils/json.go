package utils

import (
    "encoding/json"
    "io"
    "io/ioutil"
//    "log"
)

func JsonDecode(i io.Reader, s interface{}) error {
    bytes, err := ioutil.ReadAll(i)
    if err != nil {
        return nil
    }
    if len(bytes) == 0 {
        return nil
    }
    //log.Println(string(bytes))
    return json.Unmarshal(bytes, s)
}