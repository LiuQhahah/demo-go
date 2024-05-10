package main

import "fmt"

func main() {
	web := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := len(web)
	for i := 0; i < length; i++ {
		fmt.Printf("char %s : uint8: %d,uint16: %d, uint32: %d, uint64: %d, float32: %f, float64:%f \n", string(web[i]), web[i], uint16(web[i]), uint32(web[i]), uint64(web[i]), float32(web[i]), float64(web[i]))
	}

}

// Output:
//char a : uint8: 97,uint16: 97, uint32: 97, uint64: 97, float32: 97.000000, float64:97.000000
//char b : uint8: 98,uint16: 98, uint32: 98, uint64: 98, float32: 98.000000, float64:98.000000
//char c : uint8: 99,uint16: 99, uint32: 99, uint64: 99, float32: 99.000000, float64:99.000000
//char d : uint8: 100,uint16: 100, uint32: 100, uint64: 100, float32: 100.000000, float64:100.000000
//char e : uint8: 101,uint16: 101, uint32: 101, uint64: 101, float32: 101.000000, float64:101.000000
//char f : uint8: 102,uint16: 102, uint32: 102, uint64: 102, float32: 102.000000, float64:102.000000
//char g : uint8: 103,uint16: 103, uint32: 103, uint64: 103, float32: 103.000000, float64:103.000000
//char h : uint8: 104,uint16: 104, uint32: 104, uint64: 104, float32: 104.000000, float64:104.000000
//char i : uint8: 105,uint16: 105, uint32: 105, uint64: 105, float32: 105.000000, float64:105.000000
//char j : uint8: 106,uint16: 106, uint32: 106, uint64: 106, float32: 106.000000, float64:106.000000
//char k : uint8: 107,uint16: 107, uint32: 107, uint64: 107, float32: 107.000000, float64:107.000000
//char l : uint8: 108,uint16: 108, uint32: 108, uint64: 108, float32: 108.000000, float64:108.000000
//char m : uint8: 109,uint16: 109, uint32: 109, uint64: 109, float32: 109.000000, float64:109.000000
//char n : uint8: 110,uint16: 110, uint32: 110, uint64: 110, float32: 110.000000, float64:110.000000
//char o : uint8: 111,uint16: 111, uint32: 111, uint64: 111, float32: 111.000000, float64:111.000000
//char p : uint8: 112,uint16: 112, uint32: 112, uint64: 112, float32: 112.000000, float64:112.000000
//char q : uint8: 113,uint16: 113, uint32: 113, uint64: 113, float32: 113.000000, float64:113.000000
//char r : uint8: 114,uint16: 114, uint32: 114, uint64: 114, float32: 114.000000, float64:114.000000
//char s : uint8: 115,uint16: 115, uint32: 115, uint64: 115, float32: 115.000000, float64:115.000000
//char t : uint8: 116,uint16: 116, uint32: 116, uint64: 116, float32: 116.000000, float64:116.000000
//char u : uint8: 117,uint16: 117, uint32: 117, uint64: 117, float32: 117.000000, float64:117.000000
//char v : uint8: 118,uint16: 118, uint32: 118, uint64: 118, float32: 118.000000, float64:118.000000
//char w : uint8: 119,uint16: 119, uint32: 119, uint64: 119, float32: 119.000000, float64:119.000000
//char x : uint8: 120,uint16: 120, uint32: 120, uint64: 120, float32: 120.000000, float64:120.000000
//char y : uint8: 121,uint16: 121, uint32: 121, uint64: 121, float32: 121.000000, float64:121.000000
//char z : uint8: 122,uint16: 122, uint32: 122, uint64: 122, float32: 122.000000, float64:122.000000
//char A : uint8: 65,uint16: 65, uint32: 65, uint64: 65, float32: 65.000000, float64:65.000000
//char B : uint8: 66,uint16: 66, uint32: 66, uint64: 66, float32: 66.000000, float64:66.000000
//char C : uint8: 67,uint16: 67, uint32: 67, uint64: 67, float32: 67.000000, float64:67.000000
//char D : uint8: 68,uint16: 68, uint32: 68, uint64: 68, float32: 68.000000, float64:68.000000
//char E : uint8: 69,uint16: 69, uint32: 69, uint64: 69, float32: 69.000000, float64:69.000000
//char F : uint8: 70,uint16: 70, uint32: 70, uint64: 70, float32: 70.000000, float64:70.000000
//char G : uint8: 71,uint16: 71, uint32: 71, uint64: 71, float32: 71.000000, float64:71.000000
//char H : uint8: 72,uint16: 72, uint32: 72, uint64: 72, float32: 72.000000, float64:72.000000
//char I : uint8: 73,uint16: 73, uint32: 73, uint64: 73, float32: 73.000000, float64:73.000000
//char J : uint8: 74,uint16: 74, uint32: 74, uint64: 74, float32: 74.000000, float64:74.000000
//char K : uint8: 75,uint16: 75, uint32: 75, uint64: 75, float32: 75.000000, float64:75.000000
//char L : uint8: 76,uint16: 76, uint32: 76, uint64: 76, float32: 76.000000, float64:76.000000
//char M : uint8: 77,uint16: 77, uint32: 77, uint64: 77, float32: 77.000000, float64:77.000000
//char N : uint8: 78,uint16: 78, uint32: 78, uint64: 78, float32: 78.000000, float64:78.000000
//char O : uint8: 79,uint16: 79, uint32: 79, uint64: 79, float32: 79.000000, float64:79.000000
//char P : uint8: 80,uint16: 80, uint32: 80, uint64: 80, float32: 80.000000, float64:80.000000
//char Q : uint8: 81,uint16: 81, uint32: 81, uint64: 81, float32: 81.000000, float64:81.000000
//char R : uint8: 82,uint16: 82, uint32: 82, uint64: 82, float32: 82.000000, float64:82.000000
//char S : uint8: 83,uint16: 83, uint32: 83, uint64: 83, float32: 83.000000, float64:83.000000
//char T : uint8: 84,uint16: 84, uint32: 84, uint64: 84, float32: 84.000000, float64:84.000000
//char U : uint8: 85,uint16: 85, uint32: 85, uint64: 85, float32: 85.000000, float64:85.000000
//char V : uint8: 86,uint16: 86, uint32: 86, uint64: 86, float32: 86.000000, float64:86.000000
//char W : uint8: 87,uint16: 87, uint32: 87, uint64: 87, float32: 87.000000, float64:87.000000
//char X : uint8: 88,uint16: 88, uint32: 88, uint64: 88, float32: 88.000000, float64:88.000000
//char Y : uint8: 89,uint16: 89, uint32: 89, uint64: 89, float32: 89.000000, float64:89.000000
//char Z : uint8: 90,uint16: 90, uint32: 90, uint64: 90, float32: 90.000000, float64:90.000000
//char 0 : uint8: 48,uint16: 48, uint32: 48, uint64: 48, float32: 48.000000, float64:48.000000
//char 1 : uint8: 49,uint16: 49, uint32: 49, uint64: 49, float32: 49.000000, float64:49.000000
//char 2 : uint8: 50,uint16: 50, uint32: 50, uint64: 50, float32: 50.000000, float64:50.000000
//char 3 : uint8: 51,uint16: 51, uint32: 51, uint64: 51, float32: 51.000000, float64:51.000000
//char 4 : uint8: 52,uint16: 52, uint32: 52, uint64: 52, float32: 52.000000, float64:52.000000
//char 5 : uint8: 53,uint16: 53, uint32: 53, uint64: 53, float32: 53.000000, float64:53.000000
//char 6 : uint8: 54,uint16: 54, uint32: 54, uint64: 54, float32: 54.000000, float64:54.000000
//char 7 : uint8: 55,uint16: 55, uint32: 55, uint64: 55, float32: 55.000000, float64:55.000000
//char 8 : uint8: 56,uint16: 56, uint32: 56, uint64: 56, float32: 56.000000, float64:56.000000
//char 9 : uint8: 57,uint16: 57, uint32: 57, uint64: 57, float32: 57.000000, float64:57.000000
