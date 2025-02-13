package main

import (
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/18/pathfinder"
	"github.com/stretchr/testify/assert"
	//"reflect"
	"testing"
)

var maze1 = `#########
#b.A.@.a#
#########
`

var maze2 = `########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################
`

var maze3 = `########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################
`

var maze4 = `#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################
`

var maze5 = `########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################
`

var maze_input = `#################################################################################
#.........#..............p..#.....#.....#....l#...........#.....#.........#.....#
###.#####.#######.#########.#Q###.#.###.#.###.#.###.#######.###.#.#####.#.#####.#
#...#...#.........#.......#...#...#...#.#.#...#...#.#.....#g..#.#...#.#.#...#...#
#.###.#.###########.#####.#####.#######.#.#.#####.###.###.###.#.###.#.#.###.#.#.#
#...#.#...#.....#.#...#.#.#...#...#.....#.#.....#...#...#...#.#.#...#.#.#.....#.#
###.#.###.###.#.#.###.#.#.#.#.###.#.#####.#####.#.#.#.#####.#.#.#.###.#.#######.#
#.K.#...#.#...#.....#...#...#...#.#.....#...#.#.#.#...#.....#.#.#.....#...#.#...#
#.#####.#.#.###########.#######.#.###.#.###.#.#.#######.###.#.#.#####.###.#.#.###
#.#...#.#.#.......#.....#.....#.#...#.#.#...#.#.......#.#...#.#...#.#.#.#.#.#...#
#.#.#.#.#.#####.#.#.#####.#.###.###.#.#.#.###.#######.#.#.###.###.#.#.#.#.#.###.#
#...#...#.....#.#.#.......#.#.....#.#.#.#.#.........#...#.#.#...#.#.....#.#...#k#
#.###########.###.#######.###.#####.#.#.#.#######.#.#####.#.###.#.#######.#.#W#.#
#...Z.......#...#.....#...#...#.....#.#.#.....#...#.....#...#...#.........#.#.#.#
###########.###.#####.#.###.#.#.#######.#####.#.###.#######.#.#############.#.#.#
#....f..#...#.#.#...#...#...#.#...#.....#.....#.#...#.......#x#.......#.....#...#
#.#####.#.###.#.#.#.#####.###.###.#.###.#.#####.###.#.#######.#.#.#####.#########
#.#...#.#.#...#...#.....#...#...#...#...#.#.......#.#.#.#...#.#.#.....#.....#...#
#.#.#.#.#.#.#####.#####.#.#.###.#####.###.#######.###.#.#.#.#.#.#####.#####.#.#.#
#...#.#.....#...#.#.....#.#.#.......#...#.#.......#...#.#.#...#.#.........#.#.#.#
#############.#.###.#######.#######.###.#D#.###.###.###.#.#####.#.#######.#.#.#.#
#y....C.....#.#.....#n#.....#.....#.#...#.#...#.#.......#.#...#.#.#...#...#.#.#.#
#.#########.#.#####.#.#.#.#.#.###.###.###.#.#.#.#.#######.#.###.#.#.#.#####.#.#.#
#.......#.#.#.....#.#.#.#.#.#.#.#...#...#.#.#.#.#.#.#.....#.#.#.#...#.....#...#.#
#######.#.#.#####.###U#.#.#.#.#.###.###.#.#.#.#.#.#.#.#.###.#.#.#######.#.#####.#
#...#...#.#.......#...#.#.#.#...#.#...#.#.#.#.#.#.#.#.#.#...#.....#...#.#.....#.#
#.#.#.###.#####.###.#####.#.###.#.###.#.#.###.###.#.#.###.#########.#.#####.###.#
#.#.#...#.......#...#.....#...#.....#.#.#...#...#.#.#.....#.......#.#.....#...#.#
#.#####.#########.#####.###.#####.###.#.#.#.###.#.#.#########.###.#.#####.###.#.#
#.....#.......#...#.....#.#.#.....#...#.#.#...#.#.......#.....#...#...H.#...#.#.#
#.###.#######T#.###.#####.#.#.#####.###.#.###.#.###.#####.#.###.#####.#.###.#.#.#
#.#.#...#...#...#.#.....#...#.....#.....#.#...#.....#.....#.#.......#.#.#.....#.#
#.#.###.#.#.#####.#.###.#.#######.#######.#.#########.#####.#######.#.#.#.#####.#
#.#...#j..#.#.....#.#.#.#.#.....#.#.....#.#...#.....#.#...#...#.....#.#.#.#.....#
#.#.#.#####.#.#.###.#.#.#.#.###.#.###.#.#.###.#.###.#.###.###.#.#######.#.#.#####
#...#.#.#...#.#.....#.#.#.#.#...#...#.#.#.#...#...#...#...#.#.#.........#.#.#...#
#####.#.#.#.#.#######.#.#.#.#.#####.###.#.#.#####.#####.#.#.#.#.###########.#.#.#
#.....#.#.#e#.#.......#.#.#.#.....#.....#.#......i#.....#..a#.#.#.....#...J.#.#.#
#.#####.#.###.###.###.#.###.###.#######.#.#########.#######.#.###.###.#.#######.#
#.......#.........#...#.......#.....................#.......#.....E.#...........#
#######################################.@.#######################################
#.....#.#...........#.......#.............#...............#.#...........#.#.....#
#.#.#.#.#.#####.#####.#.###.#####.#####.#.#.#####.#######.#.#.#######.#.#.#.###.#
#.#h#...#.#...#...#...#...#.......#...#.#.#.#...#.#...#.....#.#...#z..#.#...#...#
#.#.#####.#.#.###.#.#####.#####.###.#.#.#.#.#.#.#.#.#.#####.#.#.###.###.#####.#.#
#.#.#...#.#.#.....#.#.....#...#.#...#.#.#...#.#.#...#...#.#.#.#...#...#.....#.#.#
###.#.#.#.#I#######.#A#####.#.#.#.###.###.###.#.#######.#.#.#.#.#.###.#####.#.#M#
#...#.#.#.#...#...#.#.#...#.#.#.#.#.#...#.#...#...#.#...#...#.#.#.#...#...#.#.#.#
#.###.#.#.###.#.#.#.#.#.#.#.#.#.#.#.###.#.###.###.#.#.###.###.###.#.###.#.#.#.#.#
#.....#...#.#...#...#...#...#.#.#.#.#...#.....#.....#...#...#.....#.#...#.#...#.#
#.#########.#####.#######.###.###.#.#.#####.###########.#########.#.#.###.#####.#
#.#...#.....#...#.#.....#...#.......#...#.#.#.....#...#.#.........#.....#.....#.#
#.#.#.#.#.###.#.###.###.###.###########.#.#.#.###.#.#.#.#.#################.#.###
#...#...#.#...#.....#.#...#.....#.#...#.#.#d#.#...#.#...#...........#.....#.#...#
#########.#.#########.###.#####.#.#.#.#.#.#.#.#.###.#.#############.#.###.#####.#
#.......#.#.#...#.......#.....#...#.#...#.#.#.#.....#.#.............#.#.#...#...#
#.#####.#.#.#.#.#.#####.#####.###.#.###.#.#.#.#######.#.#############.#.###.#.###
#.#...#.#.#...#.#.#.........#...#.#.#.#.#...#.#.....#.#.#.....#.......#...#.#.F.#
#.#.#.#.#######.#.#.#######P###.###.#.#.#####.#.###.#.#.#.###.#.#########.#.###.#
#.#.#...........#.#.......#...#.#...#.#.#.....#.#.#...#...#...#.#.........#.....#
#.###############.#######.#####.#.###.#.#.#####.#.#####.#######.#.###.#########.#
#.#...#...#.......#.#...#.....#.#.#...#.#.#.....L.#.#...#.....#.#...#.#.....#...#
#.#.#.###.#.#######.###.#####.#.#.#.###.#.#######.#.#.###R#####.#####.#.#.###.###
#...#...#...#.....#...#.....#.....#.....#.......#...#...#...#.......#...#...#.#.#
#.#####.###.#.###.#.#.#####.#######.#####.#####.#######.###.###.###.#######.#.#.#
#.....#...#...#.#...#.#...#.....#.#.....#.#...#...........#..r#.#.#.......#...V.#
#####.###.#####.#####.#.#######.#.#####.#.#.#.###########.###.#.#.#######.#######
#...#...#...#.#.....#.#.#.......#.....#.#...#.#......o..#.#...#.........#.......#
#.#.###.###.#.#.#.###.#.#.#######.#.###.#####.#B#######.###O###########.#######.#
#.#.#...#...#m..#.....#.#.#.....#.#.....#...#.#.......#.....#........u#....c#...#
#.#.###.#.#.#####.#####.#.#.#.###.#######.#.#.#######.#######.#.#####.#####.#.###
#.#...#.#.#.#...#.....#.#.#.#.....#.....#.#...#.#.S.#.#.#...#.#.#...#...#...#...#
#.###.###.#.#.#.#####.#.#.#.#######.###.#.#####.#.#.#.#.#.#.#.#.#.#.###.###.###.#
#...#v....#.#.#.#...#...#.#...#.......#.#.#.#...#.#..b#...#...#.#.#...#...#...#.#
#.#.#########.#.#.#####.#.###.#.#####.#.#.#.#.#.#.#####.#######.###.#####.#####.#
#.#...........#.#.......#...#.#...#...#.#.#...#.#...#.......#...#...#...#....t#.#
#.#############X###.#######.#.#####.###.#.#####.###.#######.#.###.###.#.#####.#.#
#...#.........#...#.#.....#.#..w....#...#...#.....#s..#.....#.#.#...#.#.#...#.#.#
###.#####.###.###.###.###.#.#########.#####.#.###.###.#######.#.#.#.#.#.#.#.#Y#.#
#.G.......#.....#.......#...#...........#q....#.....#...N.....#...#...#...#.....#
#################################################################################
`

//func Test_part1(t *testing.T) {
//	r := algo(maze1)
//	assert.Equal(t, 8, r.Steps)
//	assert.Equal(t, []string{"a","b"}, r.Path)
//
//	r = algo(maze2)
//	assert.Equal(t, 86, r.Steps)
//	assert.Equal(t,[]string{"a","b","c","d","e","f"}, r.Path)
//
//	r = algo(maze3)
//	assert.Equal(t, 132, r.Steps)
//	assert.Equal(t,[]string{"b", "a", "c", "d", "f", "e", "g"}, r.Path)
//
//	r = algo(maze4)
//	assert.Equal(t, 136, r.Steps)
//	assert.Equal(t,[]string{"a","f","b","j","g","n","h","d","l","o","e","p","c","i","k","m"}, r.Path)
//
//	r = algo(maze5)
//	assert.Equal(t, 81, r.Steps)
//	assert.Equal(t,[]string{"a","c","f","i","d","g","b","e","h"}, r.Path)
//
//
//	assert.Fail(t, "Implement me")
//}

func Test_fetchKeys(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"1", args{maze1}, 2, 1},
		{"2", args{maze2}, 6, 5},
		{"3", args{maze3}, 7, 5},
		{"4", args{maze4}, 16, 8},
		{"5", args{maze5}, 9, 5},
		{"input", args{maze_input}, 26, 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got, got1 := fetchKeys(buildGrid(tt.args.input))
			if len(got) != tt.want {
				t.Errorf("fetchKeys() got = %v, want %v", len(got), tt.want)
			}
			if len(got1) != tt.want1 {
				t.Errorf("fetchKeys() got1 = %v, want %v", len(got1), tt.want1)
			}
		})
	}
}

func Test_buildGrid(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{maze1}, maze1},
		{"2", args{maze2}, maze2},
		{"3", args{maze3}, maze3},
		{"4", args{maze4}, maze4},
		{"5", args{maze5}, maze5},
		{"input", args{maze_input}, maze_input},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildGrid(tt.args.input).String(); got != tt.want {
				t.Errorf("buildGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AStar(t *testing.T) {
	g := buildGrid(maze1)
	at, k, _ := fetchKeys(g)

	kPlusAt := append(k, at[0])

	paths := pathfinder.NewAllPaths(kPlusAt)

	assert.Equal(t, float64(2), paths['@']['a'].Distance)
	assert.Len(t, paths['@']['a'].Dependencies, 0)
	assert.Equal(t, float64(4), paths['@']['b'].Distance)
	assert.Equal(t, []rune{'a'}, paths['@']['b'].Dependencies)
	assert.Equal(t, float64(6), paths['a']['b'].Distance)
	assert.Equal(t, []rune{'a'}, paths['a']['b'].Dependencies)
}

func Test_DependencyGraph(t *testing.T) {
	at := '@'
	ks := []rune{'a', 'b', 'c', 'd'}
	pSize := 1 + len(ks)

	paths := make(pathfinder.AllPaths, pSize)
	paths['@'] = make(map[rune]*pathfinder.OnePath, pSize)
	paths['@']['a'] = &pathfinder.OnePath{Distance: 1, Dependencies: []rune{}}
	paths['@']['b'] = &pathfinder.OnePath{Distance: 2, Dependencies: []rune{'a'}}
	paths['@']['c'] = &pathfinder.OnePath{Distance: 2, Dependencies: []rune{'a'}}
	paths['@']['d'] = &pathfinder.OnePath{Distance: 3, Dependencies: []rune{'a', 'c'}}

	paths['a'] = make(map[rune]*pathfinder.OnePath, pSize)
	paths['a']['@'] = &pathfinder.OnePath{Distance: 1, Dependencies: []rune{}}
	paths['a']['b'] = &pathfinder.OnePath{Distance: 1, Dependencies: []rune{}}
	paths['a']['c'] = &pathfinder.OnePath{Distance: 1, Dependencies: []rune{}}
	paths['a']['d'] = &pathfinder.OnePath{Distance: 2, Dependencies: []rune{'c'}}

	paths['b'] = make(map[rune]*pathfinder.OnePath, pSize)
	paths['b']['@'] = &pathfinder.OnePath{Distance: 2, Dependencies: []rune{'a'}}
	paths['b']['a'] = &pathfinder.OnePath{Distance: 1, Dependencies: []rune{}}
	paths['b']['c'] = &pathfinder.OnePath{Distance: 2, Dependencies: []rune{'a'}}
	paths['b']['d'] = &pathfinder.OnePath{Distance: 3, Dependencies: []rune{'a', 'c'}}

	paths['c'] = make(map[rune]*pathfinder.OnePath, pSize)
	paths['c']['@'] = &pathfinder.OnePath{Distance: 2, Dependencies: []rune{'a'}}
	paths['c']['a'] = &pathfinder.OnePath{Distance: 1, Dependencies: []rune{}}
	paths['c']['b'] = &pathfinder.OnePath{Distance: 2, Dependencies: []rune{'a'}}
	paths['c']['d'] = &pathfinder.OnePath{Distance: 1, Dependencies: []rune{}}

	paths['d'] = make(map[rune]*pathfinder.OnePath, pSize)
	paths['d']['@'] = &pathfinder.OnePath{Distance: 3, Dependencies: []rune{'a', 'c'}}
	paths['d']['a'] = &pathfinder.OnePath{Distance: 2, Dependencies: []rune{'c'}}
	paths['d']['b'] = &pathfinder.OnePath{Distance: 3, Dependencies: []rune{'a', 'c'}}
	paths['d']['c'] = &pathfinder.OnePath{Distance: 1, Dependencies: []rune{}}

	depTree, _ := buildDependencyTree(at, ks, paths)
	fmt.Println(depTree)
}

func Test_debug_DependencyGraph(t *testing.T) {
	g := buildGrid(maze2)
	at, ks, _ := fetchKeys(g)

	ksPlusAt := append(ks, at)

	paths := pathfinder.NewAllPaths(ksPlusAt)

	atK, leftKs := simplifyDependencyTreeInput(at, ks)

	depTree, _ := buildDependencyTree(atK, leftKs, paths)
	fmt.Println(depTree)
}

func Test_buildDependencyTree(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []rune
	}{
		{"1", args{maze1}, 8, []rune{'@', 'a', 'b'}},
		{"2", args{maze2}, 86, []rune{'@', 'a', 'b', 'c', 'd', 'e', 'f'}},
		{"3", args{maze3}, 132, []rune{'@', 'b', 'a', 'c', 'd', 'f', 'e', 'g'}},
		{"4", args{maze4}, 136, []rune{'@', 'a', 'f', 'b', 'j', 'g', 'n', 'h', 'd', 'l', 'o', 'e', 'p', 'c', 'i', 'k', 'm'}},
		{"5", args{maze5}, 81, []rune{'@', 'a', 'c', 'f', 'i', 'd', 'g', 'b', 'e', 'h'}},
		{"input", args{maze_input}, 5288, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := buildGrid(tt.args.input)
			at, ks, _ := fetchKeys(g)
			ksPlusAt := append(ks, at[0])
			paths := pathfinder.NewAllPaths(ksPlusAt)

			atK, leftKs := simplifyDependencyTreeInput(at, ks)

			_, root := buildDependencyTree(atK, leftKs, paths)
			if root.Value.(*DependencyNode).MinCost != tt.want {
				t.Errorf("fetchKeys() got = %v, want %v", root.Value.(*DependencyNode).MinCost, tt.want)
			}
			//if !reflect.DeepEqual(p, tt.want1) {
			//	t.Errorf("fetchKeys() got1 = %v, want %v", p, tt.want1)
			//}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"input", args{maze_input}, 2082},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := buildGrid(tt.args.input)
			fixGrid(g)
			ats, ks, _ := fetchKeys(g)
			ksPlusAt := append(ks, ats...)
			paths := pathfinder.NewAllPaths(ksPlusAt)

			atK, leftKs := simplifyDependencyTreeInput(ats, ks)

			_, root := buildDependencyTree(atK, leftKs, paths)
			if root.Value.(*DependencyNode).MinCost != tt.want {
				t.Errorf("fetchKeys() got = %v, want %v", root.Value.(*DependencyNode).MinCost, tt.want)
			}
		})
	}
}
