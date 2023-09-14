package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	mutex := http.NewServeMux()
	mutex.Handle("/google/userinfo", GoogleUserInfo())

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", "localhost", 8081),
		Handler: mutex,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("google proxy server is exited ! err = %v", err)
		}
	}()

	log.Printf("google user info api is %s\n", "/google/userinfo?access_token={access_token}")
	log.Printf("google proxy running port is %s\n", "http://localhost:8081")

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", time.Now().Format("2006-01-02 15:04:05"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Shutdown Server Failed %s\n", err.Error())
	}

	select {
	case <-ctx.Done():
		fmt.Println("timeout of 5 seconds.")
		return
	default:
		return
	}

}

func GoogleUserInfo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.URL.Query().Get("access_token")
		if len(accessToken) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("accessToken 参数错误"))
			return
		}

		data, err := getUserinfoFromThird("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(errors.Wrap(err, "user google login verify failed").Error()))
			return
		}
		_, _ = w.Write(data)
		return
	})
}

func getUserinfoFromThird(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	return ioutil.ReadAll(resp.Body)
}
