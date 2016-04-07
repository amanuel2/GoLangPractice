package session;

import(
    "errors"
    "time"
    "net/http"
    )

type Cookie struct{
    Name string
    Expires time.Time
    Value string
    Secure bool
    HTTPONLY bool
}

func(c Cookie) setCookie(r *http.Request, w http.ResponseWriter){
    cookieSet,err := r.Cookie(c.Name)
    if err!=nil{
        err = errors.New("Cant Set Cookie. Invalid Key")
        panic(err)
    }
    http.SetCookie(w, cookieSet)
}

func check(e error){
    if e!=nil{
        panic(e)
    }
}
