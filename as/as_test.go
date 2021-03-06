package as

import (
	"strconv"
	"testing"
)

func Test_Bool(t *testing.T) {
	if Bool("t") != true {
		t.Fatalf("Bool Test failed")
	}
	if Bool("true") != true {
		t.Fatalf("Bool Test failed")
	}
	if Bool(1) != true {
		t.Fatalf("Bool Test failed")
	}
	if Bool("1") != true {
		t.Fatalf("Bool Test failed")
	}
	if Bool("FALSE") != false {
		t.Fatalf("Bool Test failed")
	}
}

func Test_Bytes(t *testing.T) {
	if String(Bytes([]byte{'l', 'o', 'r', 'e', 'm'})) != String(Bytes("lorem")) {
		t.Fatalf("Bytes Test failed")
	}
	if String([]byte("true")) != String(Bytes(true)) {
		t.Fatalf("Bytes Test failed")
	}
	if String([]byte("false")) != String(Bytes(false)) {
		t.Fatalf("Bytes Test failed")
	}
	if Int(Bytes(nil)) != Int([]byte{}) {
		t.Fatalf("Bytes Test failed")
	}
	if String(Bytes(float64(23.42))) != "23.42" {
		t.Fatalf("Bytes Test failed")
	}
}

func Test_Duration(t *testing.T) {
	if String(Duration(42)) != "42ns" {
		t.Fatalf("Duration Test failed")
	}
	if String(Duration(uint32(42))) != "42ns" {
		t.Fatalf("Duration Test failed")
	}
	if String(Duration(float64(42))) != "42ns" {
		t.Fatalf("Duration Test failed")
	}
	if String(Duration("42ns")) != "42ns" {
		t.Fatalf("Duration Test failed")
	}
}

func Test_Float(t *testing.T) {
	if Float("13.37") != float64(13.37) {
		t.Fatalf("Float Test failed")
	}
	if Float(false) != float64(0) {
		t.Fatalf("Float Test failed")
	}
	if Float(Bool(1)) != float64(1) {
		t.Fatalf("Float Test failed")
	}
	if Float(int(42)) != float64(42) {
		t.Fatalf("Float Test failed")
	}
	if Float(int8(43)) != float64(43) {
		t.Fatalf("Float Test failed")
	}
	if Float(int16(44)) != float64(44) {
		t.Fatalf("Float Test failed")
	}
	if Float(int32(45)) != float64(45) {
		t.Fatalf("Float Test failed")
	}
	if Float(int64(4500000)) != float64(4500000) {
		t.Fatalf("Float Test failed")
	}
	if Float(uint(42)) != float64(42) {
		t.Fatalf("Float Test failed")
	}
	if Float(uint8(43)) != float64(43) {
		t.Fatalf("Float Test failed")
	}
	if Float(uint16(44)) != float64(44) {
		t.Fatalf("Float Test failed")
	}
	if Float(uint32(45)) != float64(45) {
		t.Fatalf("Float Test failed")
	}
	if Float(uint64(4500000)) != float64(4500000) {
		t.Fatalf("Float Test failed")
	}
	if Float(float32(45)) != float64(45) {
		t.Fatalf("Float Test failed")
	}
	if Float(float64(4500000)) != float64(4500000) {
		t.Fatalf("Float Test failed")
	}
	if Float(Time("01.01.1970")) != 0 {
		t.Fatalf("Float Test failed")
	}
}

func Test_FloatFromXString(t *testing.T) {
	if FloatFromXString("13.000.000,00") != float64(13000000) {
		t.Fatalf("FloatFromXString Test failed")
	}
	if FloatFromXString("13.000,00") != float64(13000) {
		t.Fatalf("FloatFromXString Test failed")
	}
	if FloatFromXString("13,000,000.00") != float64(13000000) {
		t.Fatalf("FloatFromXString Test failed")
	}
	if FloatFromXString("13,00") != float64(13) {
		t.Fatalf("FloatFromXString Test failed")
	}
	if FloatFromXString("13.00") != float64(13) {
		t.Fatalf("FloatFromXString Test failed")
	}
	if FloatFromXString("13") != float64(13) {
		t.Fatalf("FloatFromXString Test failed")
	}
	if FloatFromXString("13,000.00") != float64(13000) {
		t.Fatalf("FloatFromXString Test failed")
	}
	if FloatFromXString("13,0.0,0.00") != float64(0) {
		t.Fatalf("FloatFromXString Test failed")
	}
}

func Test_Int(t *testing.T) {
	if Int(int8(1)) != int64(1) {
		t.Fatalf("Int Test failed")
	}
	if Int(int16(2)) != int64(2) {
		t.Fatalf("Int Test failed")
	}
	if Int(int32(3)) != int64(3) {
		t.Fatalf("Int Test failed")
	}
	if Int(int64(4)) != int64(4) {
		t.Fatalf("Int Test failed")
	}
	if Int(uint(5)) != int64(5) {
		t.Fatalf("Int Test failed")
	}
	if Int(uint8(6)) != int64(6) {
		t.Fatalf("Int Test failed")
	}
	if Int(uint16(7)) != int64(7) {
		t.Fatalf("Int Test failed")
	}
	if Int(uint32(8)) != int64(8) {
		t.Fatalf("Int Test failed")
	}
	if Int(uint64(9)) != int64(9) {
		t.Fatalf("Int Test failed")
	}
	if Int(float32(32.23)) != int64(32) {
		t.Fatalf("Int Test failed")
	}
	if Int(float64(32.23)) != int64(32) {
		t.Fatalf("Int Test failed")
	}
	if Int("32.73") != int64(33) {
		t.Fatalf("Int Test failed")
	}
	if Int(true) != int64(1) {
		t.Fatalf("Int Test failed")
	}
	if Int(false) != int64(0) {
		t.Fatalf("Int Test failed")
	}
	if Int(Time("01.01.1970")) != int64(0) {
		t.Fatalf("Int Test failed")
	}
}

func Test_String(t *testing.T) {
	if String(nil) != "" {
		t.Fatalf("String Test failed")
	}
	if String(1) != "1" {
		t.Fatalf("String Test failed")
	}
	if String(true) != "true" {
		t.Fatalf("String Test failed")
	}
	if String([]byte{'i', 'p', 's', 'u', 'm'}) != "ipsum" {
		t.Fatalf("String Test failed")
	}
	if String(42.0001) != "42.0001" {
		t.Fatalf("String Test failed")
	}
	if String(false) != "false" {
		t.Fatalf("String Test failed")
	}
	if String(strconv.ParseInt("42", 10, 0)) != "42" {
		t.Fatalf("String Test failed")
	}
}

func Test_Time(t *testing.T) {
	if String(Time("11.01.2015")) != "2015-01-11T00:00:00Z" {
		t.Fatalf("Time Test failed")
	}
	if String(Time("foobar")) != "0001-01-01T00:00:00Z" {
		t.Fatalf("Time Test failed")
	}
}

func Test_Uint(t *testing.T) {
	if Uint(int(0)) != uint64(0) {
		t.Fatalf("Uint Test failed")
	}
	if Uint(int8(1)) != uint64(1) {
		t.Fatalf("Uint Test failed")
	}
	if Uint(int16(2)) != uint64(2) {
		t.Fatalf("Uint Test failed")
	}
	if Uint(int32(3)) != uint64(3) {
		t.Fatalf("Uint Test failed")
	}
	if Uint(int64(4)) != uint64(4) {
		t.Fatalf("Uint Test failed")
	}
	if Uint(uint(5)) != uint64(5) {
		t.Fatalf("Uint Test failed")
	}
	if Uint(uint8(6)) != uint64(6) {
		t.Fatalf("Uint Test failed")
	}
	if Uint(uint16(7)) != uint64(7) {
		t.Fatalf("Uint Test failed")
	}
	if Uint(uint32(8)) != uint64(8) {
		t.Fatalf("Uint Test failed")
	}
	if Uint(uint64(9)) != uint64(9) {
		t.Fatalf("Uint Test failed")
	}
	if Uint(float32(32.23)) != uint64(32) {
		t.Fatalf("Uint Test failed")
	}
	if Uint(float64(32.23)) != uint64(32) {
		t.Fatalf("Uint Test failed")
	}
	if Uint("32.73") != uint64(33) {
		t.Fatalf("Uint Test failed")
	}
	if Uint(true) != uint64(1) {
		t.Fatalf("Uint Test failed")
	}
	if Uint(false) != uint64(0) {
		t.Fatalf("Uint Test failed")
	}
	if Uint(Time("01.01.1970")) != uint64(0) {
		t.Fatalf("Uint Test failed")
	}
}

func println_type(t *testing.T, str, type1, type2 string) {
	t1, err := Type(str)

	if err != nil {
		t.Fatalf("Type Test for %v failed with: %v", str, err)
	}

	t2 := DBType(str)

	if t1 != type1 {
		t.Fatalf("Type Test for %v failed with wrong value: %v != %v", str, t1, type1)
	}
	if t2 != type2 {
		t.Fatalf("DBType Test for %v failed with wrong value: %v != %v", str, t2, type2)
	}
}

func Test_Type(t *testing.T) {
	println_type(t, "true", "bool", "bool")
	println_type(t, "1000", "int", "int")
	println_type(t, "2^16", "int", "int")
	println_type(t, "10E24", "int", "int")
	println_type(t, "10 EUR", "price", "string")
	println_type(t, "10.23 £", "price", "string")
	println_type(t, "mail@example.tld", "email", "string")
	println_type(t, "978-3499626005", "isbn", "string")
	println_type(t, "127.0.0.1", "ipv4", "string")
	println_type(t, "2a01:4f8:192:34aa::10", "ipv6", "string")
	println_type(t, "b8:c7:5d:c6:6c:c6", "mac", "string")
	println_type(t, "#fefefe", "color", "string")
	println_type(t, "THk4Z1JFSlVlWEJsSUhKbGRIVnlibk1nWVNCRVlYUmhZbUZ6WlNCVWVYQmxJRzltSUdFZ2MzUnlhVzVuTGcwS1puVnVZeUJFUWxSNWNHVW9jM1J5SUhOMGNtbHVaeWtnYzNSeWFXNW5JSHNOQ2dsMExDQmxjbklnT2owZ1ZIbHdaU2h6ZEhJcERRb0phV1lnWlhKeUlDRTlJRzVwYkNCN0RRb0pDWEpsZEhWeWJpQWljM1J5YVc1bklnMEtDWDBOQ2dsemQybDBZMmdnZENCN0RRb0pZMkZ6WlNBaVltOXZiQ0lzSUNKcGJuUWlMQ0FpYzNSeWFXNW5JaXdnSW1ac2IyRjBJam9OQ2drSmNtVjBkWEp1SUhRTkNnbGtaV1poZFd4ME9nMEtDUWx5WlhSMWNtNGdJbk4wY21sdVp5SU5DZ2w5RFFwOQ==", "base64", "string")
	println_type(t, "!\"#$%&'()*+,-./0123456789:;<=>?@[\\]^_`{|}~ ", "ascii", "string")
}
