package main

import "github.com/ysanson/AdventOfCode/pkg/execute"

var Tests = execute.TestCases{
	{
		Input:         Test1,
		ExpectedPart1: 52,
		ExpectedPart2: 0,
	},
	// {
	// 	Input:         Puzzle,
	// 	ExpectedPart1: 41859,
	// 	ExpectedPart2: 30842,
	// },
}

const Test1 = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

const Puzzle = `\..-..../.....................................-....|.....|....\./|.............|...\.....|....................
...............\...|.................\.......\.......\.....-...............-........\\........................
.....|./...................................../.........|...-\.............-....\..............................
..........-........|.\.-...............\....-.....-...................................................../.....
......-.......//......|......-...-..././.........\...............-.............../..||.|.-.....-..............
./.....\..........\...........................\.....-.../....................-................\............../
......./../..................\.........-.............|......................../......\..../..........\........
................|............../................../.....|.....................................-......\........
\..............|.....././-.-...............\................................\....\../.............\..|........
..../...\................|...................\...........|-........//..|\................./............\....-.
..../....-.......................\.../............../...-../-.........../-.........................\./.\......
.....................|..\\\...................\.-..........\................-......./.............-...........
..\..|........|..|..\.|.............|..........\....|./...............\............../.................|......
......\./........./....-.....\................\..-.................................\.......|\.................
.....-............................./......|........-......................./............-.-.\.................
.\.......|.............\.......|....................-....\./..................\......./....................|..
......................|.......|........./....\..............\....\....................|..............-...-...\
.\............../......|................................/............../..............-.........//....\..../..
/..\...................\../....\..\............-........-..../........../...\|................................
.................../\..................-..|................./................|...............|../....../../...
.....\-.\\............-.|.........-....\.......\.......\......\.........................|..........\||...\....
....\..........//...............|...-..../...|..|......../..-......................|.........................-
............-....\..............................|......../|.......|.............|..../..........-.............
......................./..........-...........|-.\.\.................-......-.....\...............\...........
......................|...-\.......\/.....-.|.......|.....-............../....../...........|.................
....|.......-.............................../................./............/.........../.....................|
...../|...|.......-................|.............\./......./.................|-........................-......
..\\../../........-../......|.................../............../................\................../..........
.\...-.............-..-..............\.............|\..-.............|-.......\...........\..../....../.......
....\...../.........|...........................................................................\..\..........
................|.................../....-.../...............|....-.............../.....\..................|..
...-...........|...........\./...|............../.--.................\.............../......|....-|...........
...............-..|............/....|..|-..................-......\....................-...-........\.........
........-.....................................|.................................................-.............
............/......|..................../......../............................................................
-.........../...............................\....../\..../...........-................||............/../....|.
.........|....................\.......-.\.......-|.....|........|............-.../........|.|.................
..\........\....|.......-..........................|......-....-...-...../..........\-.........\..............
......./...\..-.-.......|./.............../...|..../.........\....-./.........\.....................||.....-..
./........../................-.......\.................|\.................................\\-........../..\..|
........................./...../.|\..\......./....................................\.\..|\............../......
.-.......|.|..................................../.....\/\|.......-.|............\.........................|...
............../.-........................|-.........\.....|............/........|...............|.........|...
......................................\..|.............|...................-............-........-//..........
..|/.\\....................\....-......../.............--.../........\.........................|....|./.......
....\/.....|/....|/............/.....|.\........../......./.../...............................|...............
...........\....|.../............../...............-.../................|\.............................-.../..
.....-................../......../...........\............../......./......\.-...//..../.\.../................
............||..............--.......\.\........................./....|........./...../........../...\/.....\.
......-...........-...................-...............|.....................-.../.|.........|.................
....\............./......-.........................\..........-......|./../..............-....................
..........\....................|....../................\..............\.\............................\/.......
....../......-.......|................/................-............|....|......-\......|..........\..........
./................................................................|....../....--........./.........|\.........
...../..\.....-...-..........-.........../-......................-..............|...-........../...|..........
....-..../...|............-...|........-./..-...............................|......|........................|.
.............|..../.................................................-/....../......../.............|.........\
....................|//......../.\...../................................-...|.........|..||..................|
.|....././......../.-........................|.................\...............|.\........-...................
.......................................-................-...../....../.........../......................./../.
.|......\................|........|......\.\-....................................................|........-...
.\|..............|......-............./.......\........\../....../../.......|.................................
....|............................|.|......|.......................|..-......................./..|..-........-.
......\....................-..........\-.|........................|.....\.....-...............................
./.........|..........|..../......|..........................................................|........|\......
...../..../../........-..............................|.|..\..\...............-.........../.|..................
.................|-...|................|......../............-..\../|....................../..................
.../...............|........-\..../....|..../...../...../.../\...........................|......-..|\.........
|......|..........\|.-./.../..\.........\...............././........./......./|................-.........|.-..
.......--...........\...|.....\..-...../.......\...../.....\..............\............|/...................-.
........\........................|..\......-.........|-..............|........................................
....../.-..-........-.\...................\//.......\.............-\./....|\......-....-.........\..../......|
........./../..\..............-/......-...|/.........../...............................\..........\.|...-.....
.\.................|...\-./......-...........-..|....--..../.-.....-..............|..............-............
................|.-/.....|.|................\.....|/....-.../...|..|..........|.......-.......................
........./....-..........|.........................................|\..\........................-.\.......-..\
...-....|..../..............................\................................-.....|............-/.\..........
.........\..........\../|...../................/.....-.\....-........-....................|...............\...
.|........................|............-.........................................|.......................-....
.//............|......../.............|../..\......|.../...-.......|...|\....|...|............................
.../..........-................\.\...../......../...|........-.................\...................\..........
.............-|........|.\.....................|.....................|.................|........../...........
.....\......-........-..........................\.................\...................|.-.....\.........\/....
........-|........................-......................................................\/....../........../.
.....\..../.......|../...........-....|.............|../.........\..-..............................-..........
.\|......................./...|.......|................-.....-.....-.............\....|.....\.-.....|.........
.|......................\.............../.........|..............-...................|.\....\........../..../.
...\......-.-|.............-...........|..............\..../.\.......|/....\....\................/............
.\...........|.|...\...-...\...........-......../-/.............\.....................|...........|./.........
..|./...\.......\..........|.....\.|............./.../.....\....../...-........................./.............
../.....-....../.\.......................................|.........-.\...-.|..........................-.......
....-...-.......-..................../..-..................|.............../........\...../...|..-......../...
................/....-.......-.............................../..........-......\.......................\....-.
........|....-............................../|......................................................\.../.....
...................................................\.../................/..\........\................/........
.......|....|...........................|...../..........-/....|...................../..........\......./.....
.............|....................\......|..../...........|.....\.....|.......................|......\../...|.
......\...........\.................\/..-....../.......-...\.-..........|\.../........\.|..........-..\.\-....
......../................\-..../...................|...-...-.........................\........................
............................|..............\........\.......................|.....-...........................
..--................-/..-/\.\.-|.........../.\.\...-.....-.-./.....\.....-../.....\..............\...|...../..
.......\........\....\...|.\....-.\|..................|.............../\..........|.......\.......\.|.../.....
/............\......-............................\..../...............................\...../........-.-..|../
..............\..\.-......\..........\\..-...........-..\....................................\.....-.......|..
.............-.......................-........\......\.......||....\..........-.../..........-.\..............
..-.............../...\...../|..../.........................-.\..../-....../.......\../.....-.................
...-.............\.....-..|...\...|................|...........\-...........|.\...................../.\.......
....\.............................|..........-...../........-.....|..............-.........\.........-........
..../.............-//......./..../.........\..\../....--...............................\/..............-......
................|.....-......|.............../....................|...../..../.|..........................|...`
