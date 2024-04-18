package main

import (
	"encoding/binary"
	"log"
	"os"
)

// Package main provides a demonstration of binary read and write operations with different endianness.
// It creates two files, LittleEndian.txt and BigEndian.txt, and writes two 4-byte integers (1 and 8) to each file.
// It then reads an 8-byte integer from each file using the corresponding endianness.
// The read values are printed in hexadecimal format.
// The output should show that the LittleEndian value is 800000001 and the BigEndian value is 100000008000000.
// 2024/04/18 19:38:57 fLittleEndian value:800000001
// 2024/04/18 19:38:57 fBigEndian value:100000008000000
func main() {
	logger := log.New(os.Stdout, "prefix", log.LstdFlags|log.Lshortfile)

	// Open the file in write mode
	fLittleEndian, err := os.OpenFile("./file/LittleEndian.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	defer fLittleEndian.Close()

	// Write a 4-byte integer of value 0 to the file
	var zeroLittleEndian int32 = 1
	err = binary.Write(fLittleEndian, binary.LittleEndian, zeroLittleEndian)
	zeroLittleEndian = 8
	err = binary.Write(fLittleEndian, binary.LittleEndian, zeroLittleEndian)

	if err != nil {
		panic(err)
	}
	// Read a 4 byte integer from a file
	// show me code
	fLittleEndian, err = os.Open("./file/LittleEndian.txt")
	if err != nil {
		panic(err)
	}
	var length int64

	err = binary.Read(fLittleEndian, binary.LittleEndian, &length)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Printf("fLittleEndian value:%x", length)

	// Open the file in write mode
	fBigEndian, err := os.OpenFile("./file/BigEndian.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer fBigEndian.Close()

	// Write a 4-byte integer of value 0 to the file
	var zeroBigEndian int32 = 1
	err = binary.Write(fBigEndian, binary.LittleEndian, zeroBigEndian)
	if err != nil {
		logger.Fatal(err)
	}
	zeroBigEndian = 8
	err = binary.Write(fBigEndian, binary.LittleEndian, zeroBigEndian)
	if err != nil {
		logger.Fatal(err)
	}
	if err != nil {
		panic(err)
	}
	// Read a 4 byte integer from a file
	// show me code
	fBigEndian, err = os.Open("./file/BigEndian.txt")
	if err != nil {
		panic(err)
	}

	err = binary.Read(fBigEndian, binary.BigEndian, &length)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Printf("fBigEndian value:%x", length)
	var twice int64
	err = binary.Read(fBigEndian, binary.BigEndian, &twice)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Printf("fBigEndian value:%x", twice)

	var again int64
	err = binary.Read(fBigEndian, binary.BigEndian, &again)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Printf("fBigEndian value:%x", again)

}
