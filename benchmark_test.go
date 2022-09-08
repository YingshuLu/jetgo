package jetgo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
	"time"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"

	"github.com/yingshulu/jetgo/vm"
)

func Benchmark_Map(b *testing.B) {
	m := map[string]interface{}{
		"Good": 12,
		"Bad":  "start",
	}
	data := &struct {
		Puin  uint64
		Store map[string]interface{}
	}{
		Puin:  uint64(12345),
		Store: m,
	}

	expr := `rule test_map {
				if Fact.Store.Good == 12 && Fact.Store.Bad == "start" && (Fact.Puin << 1) == 12345 * 2 {
					return true
				}
				return false
			}`

	v := vm.New()
	rule, err := vm.Compile(expr)
	assert.Nil(b, err)

	b.ResetTimer()
	now := time.Now()
	p := vm.Fact(data)
	for n := 0; n < b.N; n++ {
		_, err = v.Run(rule, p)
	}
	assert.Nil(b, err)
	b.Logf("N: %v, cost: %v", b.N, time.Since(now))
	b.StopTimer()
}

func Benchmark_access(b *testing.B) {
	expr := `rule test {
				return Root.E == 0 && Root.E == 0 && Root.E == 0&& Root.E == 0&& Root.E == 0 && Root.E == 0 && Root.E == 0 
			}`
	rule, err := vm.Compile(expr)
	assert.Nil(b, err)

	type P struct {
		D int
		F bool
		G string
	}

	data := struct {
		A string
		B []int
		C P
		E int
	}{
		A: "testabc",
		B: []int{1, 2, 4},
		C: P{
			D: 5,
			F: false,
			G: "testbcd",
		},
	}

	root := struct {
		Root interface{}
	}{
		Root: data,
	}

	v := vm.New()
	p := structs.Map(root)

	b.ResetTimer()
	now := time.Now()

	for n := 0; n < b.N; n++ {
		_, err = v.Run(rule, p)
	}
	assert.Nil(b, err)
	b.Logf("N: %v, cost: %v", b.N, time.Since(now))
	b.StopTimer()
}

func Benchmark_multi_access(b *testing.B) {
	expr := `rule test {
				return Root.C.D == 5 && Root.C.D == 5 &&
					Root.C.D == 5 && Root.C.D == 5 &&
					Root.C.D == 5 && Root.C.D == 5 &&
					Root.C.D == 5 && Root.C.D == 5 &&
					in(Root.C.Uin, array(655345, 544356, 65536, 778899))
			}`
	rule, err := vm.Compile(expr)
	assert.Nil(b, err)

	type P struct {
		D   int
		F   bool
		G   string
		Uin uint64
	}

	data := struct {
		A string
		B []int
		C P
	}{
		A: "testabc",
		B: []int{1, 2, 4},
		C: P{
			D:   5,
			F:   false,
			G:   "testbcd",
			Uin: 655345,
		},
	}

	root := struct {
		Root interface{}
	}{
		Root: data,
	}

	v := vm.New()
	p := structs.Map(root)

	b.ResetTimer()
	now := time.Now()

	for n := 0; n < b.N; n++ {
		_, err = v.Run(rule, p)
	}
	assert.Nil(b, err)
	b.Logf("N: %v, cost: %v", b.N, time.Since(now))
	b.StopTimer()
}

func Benchmark_local_multi_access(b *testing.B) {
	expr := `rule test {
				valD = Root.C.D
				valC = Root.C
				return valD == 5 && valD == 5 && valD == 5 &&
						valD == 5 && valD == 5 && valD == 5 &&
						valD == 5 && valD == 5 && 
						in(valC.Uin, array(655345, 544356, 65536, 778899))
			}`
	rule, err := vm.Compile(expr)
	assert.Nil(b, err)

	type P struct {
		D   int
		F   bool
		G   string
		Uin uint64
	}

	data := struct {
		A string
		B []int
		C P
	}{
		A: "testabc",
		B: []int{1, 2, 4},
		C: P{
			D:   5,
			F:   false,
			G:   "testbcd",
			Uin: 655345,
		},
	}

	root := struct {
		Root interface{}
	}{
		Root: data,
	}

	v := vm.New()
	p := structs.Map(root)
	b.Log(v.Run(rule, p))
	b.ResetTimer()
	now := time.Now()
	for n := 0; n < b.N; n++ {
		_, err = v.Run(rule, p)
	}
	assert.Nil(b, err)
	b.Logf("N: %v, cost: %v", b.N, time.Since(now))
	b.StopTimer()
}

func BenchmarkRule_Bool(b *testing.B) {
	params := &struct {
		Origin  string
		Country string
		Adults  int
		Value   int
	}{
		"MOW",
		"RU",
		1,
		100,
	}

	rule, err := vm.Compile(`rule test {
				// comment test
				if (Origin == "MOW" || Country == "RU") && (Value >= 100 && Adults == 1) {
					return true
				}
				return false
			}`)
	if err != nil {
		b.Fatal(err)
	}
	v := vm.New()

	var out interface{}

	b.ResetTimer()
	p := structs.Map(params)
	for n := 0; n < b.N; n++ {
		out, err = v.Run(rule, p)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fatal(out)
	}
}

func Benchmark_array(b *testing.B) {
	expr := `rule array_test {
				return in("test", array(1, "test", 10.0, 12))
			}`
	rule, err := vm.Compile(expr)
	if err != nil {
		b.Log("compile error: ", err)
		return
	}
	v := vm.New()
	b.ResetTimer()
	now := time.Now()
	for n := 0; n < b.N; n++ {
		_, err = v.Run(rule, nil)
	}
	assert.Nil(b, err)
	b.Logf("N: %v, cost: %v", b.N, time.Since(now))
	b.StopTimer()
}

func Benchmark_BigFieldAccess(b *testing.B) {
	type Env struct {
		Inner struct {
			Data  [1024 * 1024 * 10]byte
			Field int
		}
	}

	expr := `rule nested_access {
		if Inner.Field >= 0 && Inner.Field > 1 && Inner.Field < 2 {
			return true
		}
		return false
		}`

	rule, err := vm.Compile(expr)
	if err != nil {
		b.Log("compile error: ", err)
		return
	}

	env := &Env{}
	env.Inner.Field = 21

	p := structs.Map(env)
	v := vm.New()

	b.ResetTimer()
	now := time.Now()
	for n := 0; n < b.N; n++ {
		_, err = v.Run(rule, p)
	}
	assert.Nil(b, err)
	b.Logf("N: %v, cost: %v", b.N, time.Since(now))
}

func Benchmark_RealWorld(b *testing.B) {
	expr := `rule real_world {
		if UserAgentDevice == "DESKTOP" && 
			(OriginCountry == "RU" && DestinationCountry == "RU") &&
			in(Market, array("ru","kz", "by","uz","ua","az","am")) {
			return true
		}
		return false
	}
	`
	rule, err := vm.Compile(expr)
	if err != nil {
		b.Log("compile error: ", err)
		return
	}
	env := createEnv()
	v := vm.New()

	b.ResetTimer()
	fact := structs.Map(env)
	now := time.Now()

	for n := 0; n < b.N; n++ {
		_, err = v.Run(rule, fact)
	}
	assert.Nil(b, err)
	b.Logf("N: %v, cost: %v", b.N, time.Since(now))
}

func createEnv() interface{} {
	type DirectFlightsDays struct {
		Start string
		Days  string
	}
	type RouteSegment struct {
		Origin                string
		OriginName            string
		Destination           string
		DestinationName       string
		Date                  string
		OriginCountry         string
		DestinationCountry    string
		TranslatedOrigin      string
		TranslatedDestination string
		UserOrigin            string
		UserDestination       string
		DirectFlightsDays     *DirectFlightsDays
	}
	type Passengers struct {
		Adults   uint32
		Children uint32
		Infants  uint32
	}
	type UserAgentFeatures struct {
		Assisted     bool
		TopPlacement bool
		TourTickets  bool
	}
	country := "RU"
	agentDevice := "DESKTOP"
	type SearchParamsEnv struct {
		Segments           []*RouteSegment
		OriginCountry      *string
		DestinationCountry string
		SearchDepth        int
		Passengers         *Passengers
		TripClass          string
		UserIP             string
		KnowEnglish        bool
		Market             string
		Marker             string
		CleanMarker        string
		Locale             string
		ReferrerHost       string
		CountryCode        string
		CurrencyCode       string
		IsOpenJaw          bool
		Os                 string
		OsVersion          string
		AppVersion         string
		IsAffiliate        bool
		InitializedAt      int64
		Random             float32
		TravelPayoutsAPI   bool
		Features           *UserAgentFeatures
		GateID             int32
		UserAgentDevice    *string
		UserAgentType      string
		IsDesktop          bool
		IsMobile           bool
	}

	return SearchParamsEnv{
		Segments: []*RouteSegment{
			{
				Origin:      "VOG",
				Destination: "SHJ",
			},
			{
				Origin:      "SHJ",
				Destination: "VOG",
			},
		},
		OriginCountry:      &country,
		DestinationCountry: "RU",
		SearchDepth:        44,
		Passengers:         &Passengers{1, 0, 0},
		TripClass:          "Y",
		UserIP:             "::1",
		KnowEnglish:        true,
		Market:             "ru",
		Marker:             "123456.direct",
		CleanMarker:        "123456",
		Locale:             "ru",
		ReferrerHost:       "www.aviasales.ru",
		CountryCode:        "",
		CurrencyCode:       "usd",
		IsOpenJaw:          false,
		Os:                 "",
		OsVersion:          "",
		AppVersion:         "",
		IsAffiliate:        true,
		InitializedAt:      1570788719,
		Random:             0.13497187,
		TravelPayoutsAPI:   false,
		Features:           &UserAgentFeatures{},
		GateID:             421,
		UserAgentDevice:    &agentDevice,
		UserAgentType:      "WEB",
		IsDesktop:          true,
		IsMobile:           false,
	}
}

type TestPortrait struct {
	Metric int
	Count  int
}

func Benchmark_compile(b *testing.B) {
	expr := `
	rule test {
		if true == true && 1 == 1 && 2 == 2 &&
			3==3 && "test" == "test" && false == false &&
			nil != nil {
			return true
		}
		return false
	}
	`
	rule, err := vm.Compile(expr)
	assert.Nil(b, err)
	v := vm.New()

	b.ResetTimer()
	now := time.Now()

	for n := 0; n < b.N; n++ {
		_, err = v.Run(rule, nil)
	}
	assert.Nil(b, err)
	b.Logf("N: %v, cost: %v", b.N, time.Since(now))

}

func TestMap(t *testing.T) {
	req := struct {
		Name      string
		Portraits map[string]interface{}
	}{
		Name: "test",

		Portraits: map[string]interface{}{
			"TestPortrait": structs.Map(&TestPortrait{}),
			"good":         true,
		},
	}

	fact := vm.Fact(req)
	expr := `
		rule test {
			if Fact.Portraits.TestPortrait.Metric == 0 {
				return printf(Fact.Portraits.TestPortrait.Metric)
			}

			return false
		}
	`
	rule, err := vm.Compile(expr)
	assert.Nil(t, err)
	v, err := vm.New().Run(rule, fact)
	t.Log(v)
	assert.Nil(t, err)
}

func Test_exam(t *testing.T) {
	src := `
package main

func main() {
	results = [100] fork {
		harbor_call1(),
		harbor_call2(),
		harbor_call3(),
		harbor_call4(),
	}

	t = table {
		"123": abc,
	}
}
`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)
}

func Test_env(t *testing.T) {
	expr := `rule real_world {
		printf("test: %v", Segments[0].Origin)
		return false
	}
	`
	rule, err := vm.Compile(expr)
	if err != nil {
		t.Log("compile error: ", err)
		return
	}
	env := createEnv()
	v := vm.New()

	_, err = v.Run(rule, structs.Map(env))
	assert.Nil(t, err)
}

type Base struct {
	Name string
	Age  uint32
}

type Elem struct {
	Judge  bool
	Wedget int
}

type Derive struct {
	Base
	Action int
	Table  map[string]interface{}
}

func Test_Nested(t *testing.T) {
	d := &Derive{
		Base: Base{
			Name: "derive",
			Age:  12,
		},
		Action: 34,
		Table: map[string]interface{}{
			"Foo": Elem{
				Judge:  true,
				Wedget: 12,
			},
		},
	}

	result := map[string]interface{}{}
	err := mapstructure.Decode(d, &result)
	assert.Nil(t, err)

	t.Log(result)
}

func BenchmarkRun(b *testing.B) {
	expr := `
	rule test {
		sum = 0
		for i = 0; i < 1000000; i = i + 1 {
			sum = sum + i
		}
		return sum
	}`

	rule, err := vm.Compile(expr)
	if err != nil {
		b.Log("compile error: ", err)
		return
	}

	v := vm.New()

	b.ResetTimer()
	now := time.Now()
	for n := 0; n < b.N; n++ {
		_, err = v.Run(rule, nil)
	}
	assert.Nil(b, err)
	b.Logf("N: %v, cost: %v", b.N, time.Since(now))
}
