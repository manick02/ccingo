package main
import ("fmt"
	"io/ioutil"
	"log"
	"net/http")

func main() {
	fmt.Println("Example of http client")
	client:=&http.Client{}
	resp, err :=client.Get("https://example.com")

	if err!=nil {
		log.Fatal("this is  fatal")
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if (err!=nil) {
			fmt.Println(err)
		} else {
			fmt.Println(string(body))
		}
	}
}
