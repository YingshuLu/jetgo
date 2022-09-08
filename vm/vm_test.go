package vm

import (
	"go/token"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/yingshulu/jetgo/native"
	"github.com/yingshulu/jetgo/value/number"
)

func Test_table_get_set(t *testing.T) {
	expr := `rule test_rule {
			ma = table()
			tset(ma, string(123),  number("456"))
			val = tget(ma, string(123))
			printf("table key: 123, value: %v, type: %T", val, val)
			return val
		}`
	v := New()
	r, err := Compile(expr)
	assert.Nil(t, err)
	res, err := v.Run(r, nil)
	assert.Nil(t, err)
	equal, err := number.Eval(res, 456, token.EQL)
	assert.Nil(t, err)
	assert.True(t, equal.(bool))

	expr = `
		rule table_test {
			cnt = 1
			if Fact.PoJo - cnt > 10 {
				cnt = Fact.PoJo
				tset(Fact, "GMiss", "miss")
			}

			printf("cnt: %v", cnt == 12)

			if Fact.GMiss != nil {
				tset(Fact, "Name", "goo_ren")
			}

			if Fact.Tee == nil {
				printf("fact tee is nil")
			}

			mt = table()
			tset(mt, "Test", 123)

			printf("mt key test, value: %v", mt.Test)

			return Fact.Name
		}
	`
	pojo := 12
	name := "rengo"
	data := struct {
		PoJo *int
		Name *string
	}{
		PoJo: &pojo,
		Name: &name,
	}

	r, err = Compile(expr)
	assert.Nil(t, err)
	res, err = v.Run(r, Fact(data))
	assert.Nil(t, err)
	assert.Equal(t, "goo_ren", res.(string))
}

func Test_table_get_set2(t *testing.T) {
	expr := `
rule puin_credit {
  
  beidou = Fact.ScData.Beidou

  machineinfo = Fact.ScData.UinMachineinfo

  if beidou == nil || machineinfo == nil {
    return false
  }

  login = machineinfo.GuidInfo
  if login == nil {
    return false
  }

  if beidou.IpcDays == 0 && beidou.CityDays == 0 &&
     beidou.ProvDays == 0 && login.MachineLoginDays == 0 &&
     InviteCount1Hour > 10 {
    return true
  }
  
  return false
}
`
	r, err := Compile(expr)
	assert.Nil(t, err)
	assert.NotNil(t, r)
}

func Test_table_get_set3(t *testing.T) {
	expr := `
		rule group_credit {

  groupInviteCount2H = Fact.GroupInviteCount2Hour

  // 直接放过
  if groupInviteCount2H <= 0 {
    return false
  }

  // 直接打击
  if groupInviteCount2H > 300 {
    return true
  }

  creditScore = Fact.GroupCreditScore

  if creditScore > 0 && creditScore < 20 && groupInviteCount2H > 0 {
    return true
  }

  if creditScore >= 20 && creditScore < 40 && groupInviteCount2H >= 40 {
    return true
  }

  if creditScore >= 40 && creditScore < 60 && groupInviteCount2H >= 60 {
    return true
  }

  if creditScore >= 60 && creditScore < 80 && groupInviteCount2H >= 100 {
    return true
  }

  if creditScore >= 80 && groupInviteCount2H > 100 {
    return true
  }
  
  return false
}
`
	r, err := Compile(expr)
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func Test_local_rule(t *testing.T) {
	expr := `rule test_rule {
			a = 2
			b = 5
			c = a + b * 3
			if c == 17 {
				d = 12
			}
			e = d + 2
			f = e + c + a / 2
			g = f - b
			return g
		}`
	v := New()
	r, err := Compile(expr)
	assert.Nil(t, err)
	res, err := v.Run(r, nil)
	assert.Nil(t, err)
	assert.Equal(t, res.(float64), float64(27))
}

func TestRun(t *testing.T) {
	expr := `rule test base 1 1001 {
				a = "testa"
				b = 1
				if Fact.A == a {
					return "true"
				}
				if b == Fact.B[2] {
					return false
				}

				mymap = table(16)
				key = string("test")
				tset(mymap, string("test"), 12)
				v = tget(mymap, "test")
				printf(v, len(mymap))

				tset(Fact, "Global_", "testvalue")

				val = tget(Fact, "Global_")

				printf("split global value: %v in printf", val)
				if v == 12 {
					return true
				}
				
				test = Fact.C.G
				return test == "testbcd"
			}`
	rule, err := Compile(expr)
	assert.Nil(t, err)

	type P struct {
		D int
		F bool
		G string
	}

	data := struct {
		A string
		B []int
		C P
	}{
		A: "testabc",
		B: []int{1, 2, 4},
		C: P{
			D: 5,
			F: false,
			G: "testbcd",
		},
	}

	fact := Fact(data)
	v := New()
	x, err := v.Run(rule, fact)
	assert.True(t, x.(bool))
	assert.Nil(t, err)
}

func Test_growLen(t *testing.T) {
	var s stack
	tests := []struct {
		expect int
		grow   int
	}{
		{3, 4},
		{4, 4},
		{8, 8},
		{9, 16},
		{16, 16},
		{17, 32},
		{42, 64},
		{111, 128},
		{199, 256},
		{333, 512},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.grow, s.growLen(tt.expect))
	}
}

func Test_map(t *testing.T) {
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
				if Fact.Store.Good == 12 &&
					Fact.Store.Bad == "start" &&
					(Fact.Puin << 1) == 12345 * 2 {
					return true
				}
				return false
			}`

	v := New()
	r, err := Compile(expr)
	assert.Nil(t, err)
	res, err := v.Run(r, Fact(data))
	assert.Nil(t, err)
	assert.Equal(t, res.(bool), true)
}

func Test_fork(t *testing.T) {
	expr := `
		rule fork_rule {
			arr = array(1, 2, 3, 4, 5)
			res = [100]fork {
				len(arr),
				len(arr),
				len(arr),
			}
			
			printf("len: %v", res)
		}
	`

	rule, err := Compile(expr)
	assert.Nil(t, err)

	v := New()
	_, err = v.Run(rule, nil)
	assert.Nil(t, err)

	expr = `
		rule fork_rule {
			arr = array(1, 2, 3, 4, 5)
			timeout = 100
			res = [timeout]fork {
				len(arr),
				len(arr),
				len(arr),
			}
			
			printf("len: %v", res)
		}
	`
	rule, err = Compile(expr)
	assert.Nil(t, err)

	_, err = v.Run(rule, nil)
	assert.Nil(t, err)
}

func Test_fork_timeout(t *testing.T) {
	timeout := 120 * time.Millisecond
	sleep := func(args []interface{}) (interface{}, error) {
		time.Sleep(timeout)
		return args, nil
	}
	native.Register("sleep", sleep)

	expr := `
		rule fork_rule {
			arr = array(1, 2, 3, 4, 5)
			timeout = 100
			a = 1
			b = 2
			c = 3
			d = 4
			res = [timeout]fork {
				sleep(a,b),
				len(arr),
				sleep(c,d),
				len(arr),
				sleep(a,c),
				len(arr),
				sleep(b,d),
			}
			
			s = printf("len: %v", res)
			sleep(a, 2)
			return res
		}
	`
	rule, err := Compile(expr)
	assert.Nil(t, err)

	times := 1
	v := New()
	for range make([]struct{}, times) {
		r, err := v.Run(rule, nil)
		assert.Nil(t, err)
		t.Log(r)
	}

	t.Log("@parent1: ", v.(*vm).debugInfo())

	timeout = 90 * time.Millisecond

	for range make([]struct{}, times) {
		r, err := v.Run(rule, nil)
		assert.Nil(t, err)
		t.Log(r)
	}

	t.Log("@parent2: ", v.(*vm).debugInfo())
}

func Test_fork_error(t *testing.T) {
	expr := `
		rule fork_rule {
			arr = array(1, 2, 3, 4, 5)
			res = [100]fork {
				1 + 1,
				2 + 2,
				len(arr),
			}
			
			printf("len: %v", res)
		}
	`
	_, err := Compile(expr)
	t.Log(err)
	assert.NotNil(t, err)

	expr = `
		rule fork_rule {
			arr = array(1, 2, 3, 4, 5)
			res = [100] conc {
				1 + 1,
				2 + 2,
				len(arr),
			}
			
			printf("len: %v", res)
		}
	`
	_, err = Compile(expr)
	t.Log(err)
	assert.NotNil(t, err)
}

func Test_fork_callStacks(t *testing.T) {
	f := func(args []interface{}) (interface{}, error) {
		n := len(args)
		time.Sleep(time.Duration(n) * 10 * time.Millisecond)
		return n, nil
	}
	native.Register("call_len", f)

	expr := `
		rule test {
			a0 = 1
			a1 = 2
			a2 = 3
			a3 = 4
			a4 = 5
			a5 = 6

			res = [100] fork {
				call_len(a0),
				call_len(a0, a1),
				call_len(a0, a1, a2),
				call_len(a0, a1, a2, a3),
				call_len(a0, a1, a2, a3, a4),
				call_len(a0, a1, a2, a3, a4, a5),
				call_len(a0, a1, a2, a3, a4),
				call_len(a0, a1, a2, a3),
				call_len(a0, a1, a2),
				call_len(a0, a1),
				call_len(a0),
				call_len(Fact.Foo, a3, Fact.Goo),
				call_len(a0, Fact.Foo, Fact.Goo),
			}
			return res
		}
	`
	rule, err := Compile(expr)
	assert.Nil(t, err)
	v := New()

	data := &struct {
		Foo int
		Goo int
	}{
		Foo: 12,
		Goo: 13,
	}
	fact := Fact(data)
	var val interface{}
	for i := 0; i < 20; i++ {
		t.Run("test", func(t *testing.T) {
			val, err = v.Run(rule, fact)
			assert.Nil(t, err)
		})
	}
	t.Log(val)
}
