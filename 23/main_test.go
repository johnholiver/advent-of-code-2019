package main

import (
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/machine/network"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var program = `3,62,1001,62,11,10,109,2259,105,1,0,2187,1043,604,1333,1434,1898,977,775,1366,946,1793,1500,1595,713,1218,1962,1119,1865,744,1929,1463,1760,2121,907,1628,1564,1078,1694,2156,845,680,2092,1659,1148,808,876,2228,1830,1531,641,1993,1247,1284,1399,2024,1177,1729,571,2059,1008,0,0,0,0,0,0,0,0,0,0,0,0,3,64,1008,64,-1,62,1006,62,88,1006,61,170,1105,1,73,3,65,21002,64,1,1,21002,66,1,2,21101,0,105,0,1106,0,436,1201,1,-1,64,1007,64,0,62,1005,62,73,7,64,67,62,1006,62,73,1002,64,2,132,1,132,68,132,1002,0,1,62,1001,132,1,140,8,0,65,63,2,63,62,62,1005,62,73,1002,64,2,161,1,161,68,161,1102,1,1,0,1001,161,1,169,1001,65,0,0,1102,1,1,61,1101,0,0,63,7,63,67,62,1006,62,203,1002,63,2,194,1,68,194,194,1006,0,73,1001,63,1,63,1105,1,178,21101,210,0,0,105,1,69,2101,0,1,70,1101,0,0,63,7,63,71,62,1006,62,250,1002,63,2,234,1,72,234,234,4,0,101,1,234,240,4,0,4,70,1001,63,1,63,1106,0,218,1105,1,73,109,4,21101,0,0,-3,21101,0,0,-2,20207,-2,67,-1,1206,-1,293,1202,-2,2,283,101,1,283,283,1,68,283,283,22001,0,-3,-3,21201,-2,1,-2,1106,0,263,22101,0,-3,-3,109,-4,2106,0,0,109,4,21101,0,1,-3,21102,1,0,-2,20207,-2,67,-1,1206,-1,342,1202,-2,2,332,101,1,332,332,1,68,332,332,22002,0,-3,-3,21201,-2,1,-2,1105,1,312,22102,1,-3,-3,109,-4,2106,0,0,109,1,101,1,68,359,20101,0,0,1,101,3,68,367,20101,0,0,2,21101,376,0,0,1106,0,436,21202,1,1,0,109,-1,2105,1,0,1,2,4,8,16,32,64,128,256,512,1024,2048,4096,8192,16384,32768,65536,131072,262144,524288,1048576,2097152,4194304,8388608,16777216,33554432,67108864,134217728,268435456,536870912,1073741824,2147483648,4294967296,8589934592,17179869184,34359738368,68719476736,137438953472,274877906944,549755813888,1099511627776,2199023255552,4398046511104,8796093022208,17592186044416,35184372088832,70368744177664,140737488355328,281474976710656,562949953421312,1125899906842624,109,8,21202,-6,10,-5,22207,-7,-5,-5,1205,-5,521,21102,1,0,-4,21101,0,0,-3,21102,1,51,-2,21201,-2,-1,-2,1201,-2,385,470,21002,0,1,-1,21202,-3,2,-3,22207,-7,-1,-5,1205,-5,496,21201,-3,1,-3,22102,-1,-1,-5,22201,-7,-5,-7,22207,-3,-6,-5,1205,-5,515,22102,-1,-6,-5,22201,-3,-5,-3,22201,-1,-4,-4,1205,-2,461,1106,0,547,21101,-1,0,-4,21202,-6,-1,-6,21207,-7,0,-5,1205,-5,547,22201,-7,-6,-7,21201,-4,1,-4,1106,0,529,21201,-4,0,-7,109,-8,2106,0,0,109,1,101,1,68,564,20101,0,0,0,109,-1,2106,0,0,1102,1297,1,66,1101,0,2,67,1101,0,598,68,1102,1,302,69,1101,1,0,71,1101,0,602,72,1106,0,73,0,0,0,0,8,17398,1102,43159,1,66,1102,4,1,67,1102,631,1,68,1102,1,302,69,1101,0,1,71,1102,1,639,72,1105,1,73,0,0,0,0,0,0,0,0,30,83914,1102,92459,1,66,1101,5,0,67,1101,0,668,68,1102,302,1,69,1101,1,0,71,1102,1,678,72,1106,0,73,0,0,0,0,0,0,0,0,0,0,37,33314,1102,41957,1,66,1102,1,2,67,1102,1,707,68,1101,302,0,69,1101,1,0,71,1102,1,711,72,1105,1,73,0,0,0,0,10,114159,1101,0,89393,66,1101,1,0,67,1101,740,0,68,1101,0,556,69,1102,1,1,71,1101,742,0,72,1106,0,73,1,81,44,59667,1102,28871,1,66,1102,1,1,67,1101,771,0,68,1102,1,556,69,1101,0,1,71,1102,773,1,72,1105,1,73,1,1559,32,39079,1102,1,32063,66,1102,2,1,67,1102,802,1,68,1102,302,1,69,1101,1,0,71,1101,0,806,72,1106,0,73,0,0,0,0,17,84914,1102,1559,1,66,1102,1,4,67,1102,1,835,68,1102,1,302,69,1102,1,1,71,1101,843,0,72,1106,0,73,0,0,0,0,0,0,0,0,10,76106,1101,0,54679,66,1102,1,1,67,1102,872,1,68,1102,556,1,69,1102,1,1,71,1101,0,874,72,1106,0,73,1,77489,44,19889,1101,42577,0,66,1102,1,1,67,1101,0,903,68,1101,556,0,69,1102,1,1,71,1102,1,905,72,1105,1,73,1,461,20,154653,1101,88873,0,66,1102,5,1,67,1102,1,934,68,1101,253,0,69,1102,1,1,71,1102,1,944,72,1105,1,73,0,0,0,0,0,0,0,0,0,0,3,5281,1101,33377,0,66,1101,1,0,67,1102,1,973,68,1102,556,1,69,1102,1,1,71,1101,0,975,72,1105,1,73,1,2395871,21,109966,1101,8779,0,66,1101,1,0,67,1102,1,1004,68,1101,0,556,69,1102,1,1,71,1101,1006,0,72,1106,0,73,1,1756,1,53871,1101,0,15761,66,1102,1,1,67,1101,1035,0,68,1101,556,0,69,1102,3,1,71,1101,1037,0,72,1106,0,73,1,3,43,147166,34,4677,39,277377,1101,0,17957,66,1102,1,3,67,1102,1,1070,68,1101,0,302,69,1102,1,1,71,1101,0,1076,72,1105,1,73,0,0,0,0,0,0,23,266619,1102,59369,1,66,1102,1,1,67,1102,1105,1,68,1101,556,0,69,1101,6,0,71,1101,0,1107,72,1106,0,73,1,7,20,206204,3,10562,43,220749,34,1559,2,172636,39,184918,1102,1,67763,66,1101,1,0,67,1102,1146,1,68,1102,556,1,69,1102,1,0,71,1101,1148,0,72,1105,1,73,1,1414,1102,1,8693,66,1102,1,1,67,1102,1175,1,68,1102,1,556,69,1102,1,0,71,1102,1177,1,72,1106,0,73,1,1808,1102,1,67057,66,1101,0,6,67,1101,0,1204,68,1101,302,0,69,1102,1,1,71,1102,1,1216,72,1106,0,73,0,0,0,0,0,0,0,0,0,0,0,0,38,88994,1101,97789,0,66,1101,1,0,67,1102,1245,1,68,1101,556,0,69,1101,0,0,71,1101,1247,0,72,1105,1,73,1,1630,1102,69653,1,66,1102,4,1,67,1102,1274,1,68,1102,1,302,69,1102,1,1,71,1101,1282,0,72,1106,0,73,0,0,0,0,0,0,0,0,45,335285,1102,45061,1,66,1101,0,1,67,1102,1,1311,68,1101,556,0,69,1101,10,0,71,1101,0,1313,72,1105,1,73,1,1,1,35914,20,51551,44,39778,32,117237,7,32063,17,42457,47,1297,8,8699,2,43159,39,92459,1101,0,5281,66,1102,1,2,67,1101,1360,0,68,1101,302,0,69,1101,0,1,71,1102,1,1364,72,1106,0,73,0,0,0,0,43,73583,1101,8699,0,66,1101,2,0,67,1102,1,1393,68,1102,302,1,69,1101,1,0,71,1102,1397,1,72,1105,1,73,0,0,0,0,2,129477,1101,0,73583,66,1102,3,1,67,1101,1426,0,68,1101,302,0,69,1102,1,1,71,1101,0,1432,72,1106,0,73,0,0,0,0,0,0,34,3118,1102,1,2161,66,1101,1,0,67,1102,1,1461,68,1101,0,556,69,1102,1,0,71,1102,1,1463,72,1106,0,73,1,1802,1102,51551,1,66,1101,4,0,67,1101,0,1490,68,1102,1,302,69,1102,1,1,71,1102,1,1498,72,1105,1,73,0,0,0,0,0,0,0,0,23,177746,1102,1,4969,66,1101,1,0,67,1102,1,1527,68,1102,556,1,69,1101,1,0,71,1101,0,1529,72,1105,1,73,1,18,21,54983,1102,44497,1,66,1101,2,0,67,1101,1558,0,68,1101,351,0,69,1102,1,1,71,1102,1562,1,72,1105,1,73,0,0,0,0,255,47969,1102,11251,1,66,1102,1,1,67,1102,1,1591,68,1101,0,556,69,1101,1,0,71,1102,1,1593,72,1106,0,73,1,9817,1,17957,1102,1,47981,66,1102,1,1,67,1102,1,1622,68,1101,0,556,69,1102,2,1,71,1102,1,1624,72,1105,1,73,1,10,41,208959,45,134114,1101,0,34261,66,1102,1,1,67,1102,1,1655,68,1101,556,0,69,1101,1,0,71,1101,1657,0,72,1105,1,73,1,-147,39,462295,1102,39079,1,66,1101,3,0,67,1102,1,1686,68,1101,0,302,69,1102,1,1,71,1101,0,1692,72,1106,0,73,0,0,0,0,0,0,23,444365,1101,31151,0,66,1101,1,0,67,1102,1721,1,68,1102,556,1,69,1102,3,1,71,1102,1723,1,72,1106,0,73,1,5,41,69653,41,278612,45,268228,1101,0,75083,66,1101,1,0,67,1102,1756,1,68,1101,0,556,69,1101,0,1,71,1102,1758,1,72,1105,1,73,1,125,41,139306,1101,0,54983,66,1102,1,2,67,1102,1787,1,68,1101,0,302,69,1102,1,1,71,1101,0,1791,72,1105,1,73,0,0,0,0,23,355492,1101,38053,0,66,1101,4,0,67,1101,1820,0,68,1101,253,0,69,1101,1,0,71,1102,1828,1,72,1106,0,73,0,0,0,0,0,0,0,0,38,44497,1102,16657,1,66,1101,3,0,67,1102,1857,1,68,1102,302,1,69,1101,0,1,71,1101,1863,0,72,1105,1,73,0,0,0,0,0,0,10,38053,1101,42457,0,66,1102,2,1,67,1102,1892,1,68,1101,302,0,69,1102,1,1,71,1102,1896,1,72,1106,0,73,0,0,0,0,47,2594,1101,91291,0,66,1102,1,1,67,1101,1925,0,68,1102,1,556,69,1102,1,1,71,1102,1927,1,72,1106,0,73,1,79481,7,64126,1101,82193,0,66,1101,0,1,67,1102,1,1956,68,1102,556,1,69,1102,1,2,71,1102,1958,1,72,1105,1,73,1,2,45,67057,45,201171,1102,1,65777,66,1102,1,1,67,1101,1989,0,68,1102,556,1,69,1101,1,0,71,1102,1,1991,72,1106,0,73,1,276,32,78158,1101,97453,0,66,1102,1,1,67,1102,1,2020,68,1102,556,1,69,1101,1,0,71,1102,1,2022,72,1105,1,73,1,160,45,402342,1102,19889,1,66,1102,3,1,67,1102,2051,1,68,1102,302,1,69,1101,1,0,71,1101,0,2057,72,1105,1,73,0,0,0,0,0,0,23,88873,1102,93151,1,66,1101,1,0,67,1101,0,2086,68,1102,556,1,69,1102,2,1,71,1102,2088,1,72,1106,0,73,1,11,34,6236,39,369836,1101,40759,0,66,1101,0,1,67,1102,2119,1,68,1101,556,0,69,1101,0,0,71,1101,0,2121,72,1106,0,73,1,1387,1101,0,93283,66,1101,3,0,67,1102,2148,1,68,1101,302,0,69,1102,1,1,71,1102,1,2154,72,1106,0,73,0,0,0,0,0,0,10,152212,1102,1,78977,66,1102,1,1,67,1101,2183,0,68,1101,0,556,69,1101,0,1,71,1102,1,2185,72,1105,1,73,1,-3030,20,103102,1101,0,47969,66,1101,1,0,67,1101,0,2214,68,1101,0,556,69,1101,0,6,71,1102,2216,1,72,1105,1,73,1,19153,30,41957,37,16657,37,49971,22,93283,22,186566,22,279849,1101,0,76081,66,1101,0,1,67,1101,2255,0,68,1102,1,556,69,1101,1,0,71,1101,2257,0,72,1106,0,73,1,1949,2,86318`

func Test_OneNic(t *testing.T) {
	nic := network.NewController(0, program)
	nic.SetDebugMode(true)

	go func() {
		nic.Exec()
	}()
	time.Sleep(5 * time.Second)
	assert.Equal(t, 30, nic.P.Output.ReadAt(0))
	assert.Equal(t, 41957, nic.P.Output.ReadAt(1))
	assert.Equal(t, 19153, nic.P.Output.ReadAt(2))
	assert.Equal(t, 37, nic.P.Output.ReadAt(3))
	assert.Equal(t, 16657, nic.P.Output.ReadAt(4))
	assert.Equal(t, 19153, nic.P.Output.ReadAt(5))
	assert.Equal(t, 37, nic.P.Output.ReadAt(6))
	assert.Equal(t, 49971, nic.P.Output.ReadAt(7))
	assert.Equal(t, 19153, nic.P.Output.ReadAt(8))
	assert.Equal(t, 22, nic.P.Output.ReadAt(9))
	assert.Equal(t, 93283, nic.P.Output.ReadAt(10))
	assert.Equal(t, 19153, nic.P.Output.ReadAt(11))
	assert.Equal(t, 22, nic.P.Output.ReadAt(12))
	assert.Equal(t, 186566, nic.P.Output.ReadAt(13))
	assert.Equal(t, 19153, nic.P.Output.ReadAt(14))
	assert.Equal(t, 22, nic.P.Output.ReadAt(15))
	assert.Equal(t, 279849, nic.P.Output.ReadAt(16))
	assert.Equal(t, 19153, nic.P.Output.ReadAt(17))
}

func Test_OneNic1(t *testing.T) {
	nic := network.NewController(1, program)
	nic.SetDebugMode(true)

	go func() {
		nic.Exec()
	}()
	time.Sleep(5 * time.Second)
}

func Test_OneNicOneMessage(t *testing.T) {
	nic := network.NewController(30, program)
	nic.SetDebugMode(true)

	fakeMessage := network.NewIntPairPacket(0, 30, 41957, 19153)
	nic.QueuePush(fakeMessage)
	go func() {
		nic.Exec()
	}()
	time.Sleep(5 * time.Second)
	assert.Equal(t, 30, nic.P.Output.ReadAt(0))
	assert.Equal(t, 41957, nic.P.Output.ReadAt(1))
	assert.Equal(t, 19153, nic.P.Output.ReadAt(2))
	assert.Equal(t, 37, nic.P.Output.ReadAt(3))
	assert.Equal(t, 16657, nic.P.Output.ReadAt(4))
	assert.Equal(t, 19153, nic.P.Output.ReadAt(5))
	assert.Equal(t, 37, nic.P.Output.ReadAt(6))
	assert.Equal(t, 49971, nic.P.Output.ReadAt(7))
	assert.Equal(t, 19153, nic.P.Output.ReadAt(8))
	assert.Equal(t, 22, nic.P.Output.ReadAt(9))
	assert.Equal(t, 93283, nic.P.Output.ReadAt(10))
	assert.Equal(t, 19153, nic.P.Output.ReadAt(11))
	assert.Equal(t, 22, nic.P.Output.ReadAt(12))
	assert.Equal(t, 186566, nic.P.Output.ReadAt(13))
	assert.Equal(t, 19153, nic.P.Output.ReadAt(14))
	assert.Equal(t, 22, nic.P.Output.ReadAt(15))
	assert.Equal(t, 279849, nic.P.Output.ReadAt(16))
	assert.Equal(t, 19153, nic.P.Output.ReadAt(17))
}

func Test_OneNicTwoMessage(t *testing.T) {
	nic := network.NewController(0, program)
	nic.SetDebugMode(true)

	fakeMessage := network.NewIntPairPacket(0, 0, 0, 0)
	nic.QueuePush(fakeMessage)
	fakeMessage = network.NewIntPairPacket(0, 0, 3, 75)
	nic.QueuePush(fakeMessage)
	go func() {
		nic.Exec()
	}()

	time.Sleep(5 * time.Second)
	assert.Equal(t, 30, nic.P.Output.ReadAt(0))
	assert.Equal(t, 41957, nic.P.Output.ReadAt(1))
	assert.Equal(t, 19153, nic.P.Output.ReadAt(2))
	assert.Equal(t, 37, nic.P.Output.ReadAt(3))
	assert.Equal(t, 16657, nic.P.Output.ReadAt(4))
	assert.Equal(t, 19153, nic.P.Output.ReadAt(5))
}

func Test_TwoNicOneRouter(t *testing.T) {
	router := network.NewRouter(nil)
	router.SetDebugMode(true)

	nic0 := network.NewController(0, program)
	nic0.SetDebugMode(true)
	router.AddNic(nic0)

	nic30 := network.NewController(30, program)
	nic30.SetDebugMode(true)
	router.AddNic(nic30)

	fakeMessage := network.NewIntPairPacket(0, 0, 0, 0)
	nic0.QueuePush(fakeMessage)
	go func() {
		nic0.Exec()
	}()
	go func() {
		nic30.Exec()
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("done")
}
