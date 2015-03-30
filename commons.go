package web

import (
    "strings"
)

func in(key, list string) bool {
    if key == "" || list == "" {
        return false
    }
    for _, i := range strings.Split(list, ",") {
        if key == i {
            return true
        }
    }
    return false
}