package magicgate

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

const content = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
	<Say>Hello Gopher how you doing?</Say>
	<Record timeout="10"/>
</Response>
`

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/xml")
	fmt.Fprint(w, content)
}

type speechReq struct {
	Config struct {
	} `json:"config"`
	Audio struct {
	} `json:"audio"`
}

// func fetchTranscription(c context.Context, b []byte) (string, error) {

// }
