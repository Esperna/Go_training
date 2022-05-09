package params

import (
	"log"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestPack(t *testing.T) {
	type queryParam struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}

	data := queryParam{Labels: []string{"golang", "programming"}, MaxResults: 100, Exact: true}
	url, err := Pack(&data)
	if err != nil {
		log.Fatalf("Pack failed. data: %v err: %s", data, err)
	}

	want := "http://localhost:12345/search?l=golang&l=programming&max=100&x=true"
	actual := url.String()
	if actual != want {
		t.Errorf("url is %s, want %s", actual, want)
	}
}

func TestUnpack(t *testing.T) {
	type queryParam struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	var tests = []struct {
		want  queryParam
		given string
	}{
		{queryParam{Labels: []string(nil), MaxResults: 0, Exact: false}, "http://localhost:12345/search"},
		{queryParam{Labels: []string{"golang", "programming"}, MaxResults: 10, Exact: false}, "http://localhost:12345/search?l=golang&l=programming&max=10"},
		{queryParam{Labels: []string{"programming", "golang"}, MaxResults: 0, Exact: true}, "http://localhost:12345/search?&x=true&l=programming&q=hello&l=golang"},
	}
	for _, test := range tests {
		var req http.Request
		var actual queryParam
		url, err := url.Parse(test.given)
		if err != nil {
			t.Errorf("url parse failed:%s", err.Error())
		}
		req.URL = url
		if err := Unpack(&req, &actual); err != nil {
			t.Errorf("Unpacked failed:%s", err.Error())
		}
		if !reflect.DeepEqual(actual, test.want) {
			t.Errorf("actual:%v, want:%v", actual, test.want)
		}
	}
}

func TestUnpackFailed(t *testing.T) {
	type queryParam struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	var tests = []struct {
		want  string
		given string
	}{
		{`x: strconv.ParseBool: parsing "123": invalid syntax`, "http://localhost:12345/search?q=hello&x=123"},
		{`max: strconv.ParseInt: parsing "lots": invalid syntax`, "http://localhost:12345/search?q=hello&max=lots"},
	}
	for _, test := range tests {
		var req http.Request
		var data queryParam
		url, err := url.Parse(test.given)
		if err != nil {
			t.Errorf("url parse failed:%s", err.Error())
		}
		req.URL = url
		err = Unpack(&req, &data)
		actual := err.Error()
		if actual != test.want {
			t.Errorf("actual:%s, want:%s", actual, test.want)
		}
	}
}

func TestUnpackExtentionWithInvalidParam(t *testing.T) {
	type queryParam struct {
		mailAddr   string `http:"ma,mail"`
		cardNo     string `http:"cn,number"`
		postalCode string `http:"pc,code"`
	}
	var tests = []struct {
		want  string
		given string
	}{
		//{"", "http://localhost:12345/search?ma=abc@gmail.com"},
		{"invalid mail:abc.com", "http://localhost:12345/search?ma=abc.com"},
		{"invalid number:abcdef", "http://localhost:12345/search?cn=abcdef"},
		{"invalid code:123456", "http://localhost:12345/search?pc=123456"},
	}
	for _, test := range tests {
		var req http.Request
		var data queryParam
		url, err := url.Parse(test.given)
		if err != nil {
			t.Errorf("url parse failed:%s", err.Error())
		}
		req.URL = url
		err = Unpack(&req, &data)
		actual := err.Error()
		if actual != test.want {
			t.Errorf("actual:%s, want:%s", actual, test.want)
		}
	}

}

func TestUnpackExtentionWithValidParam(t *testing.T) {
	type queryParam struct {
		MailAddr   string `http:"ma,mail"`
		CardNo     string `http:"cn,number"`
		PostalCode string `http:"pc,code"`
	}
	var tests = []struct {
		want  queryParam
		given string
	}{
		{queryParam{MailAddr: "abc@gmail.com", CardNo: "", PostalCode: ""}, "http://localhost:12345/search?ma=abc@gmail.com"},
	}
	for _, test := range tests {
		var req http.Request
		var actual queryParam
		url, err := url.Parse(test.given)
		if err != nil {
			t.Errorf("url parse failed:%s", err.Error())
		}
		req.URL = url
		if err := Unpack(&req, &actual); err != nil {
			t.Errorf("Unpacked failed:%s", err.Error())
		}
		if !reflect.DeepEqual(actual, test.want) {
			t.Errorf("actual:%v, want:%v", actual, test.want)
		}
	}
}
