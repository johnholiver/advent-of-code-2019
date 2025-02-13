package main

import (
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/20/pathfinder"
	"github.com/stretchr/testify/assert"
	"testing"
)

var input1 = `         A           
         A           
  #######.#########  
  #######.........#  
  #######.#######.#  
  #######.#######.#  
  #######.#######.#  
  #####  B    ###.#  
BC...##  C    ###.#  
  ##.##       ###.#  
  ##...DE  F  ###.#  
  #####    G  ###.#  
  #########.#####.#  
DE..#######...###.#  
  #.#########.###.#  
FG..#########.....#  
  ###########.#####  
             Z       
             Z      
`

var input2 = `                   A               
                   A               
  #################.#############  
  #.#...#...................#.#.#  
  #.#.#.###.###.###.#########.#.#  
  #.#.#.......#...#.....#.#.#...#  
  #.#########.###.#####.#.#.###.#  
  #.............#.#.....#.......#  
  ###.###########.###.#####.#.#.#  
  #.....#        A   C    #.#.#.#  
  #######        S   P    #####.#  
  #.#...#                 #......VT
  #.#.#.#                 #.#####  
  #...#.#               YN....#.#  
  #.###.#                 #####.#  
DI....#.#                 #.....#  
  #####.#                 #.###.#  
ZZ......#               QG....#..AS
  ###.###                 #######  
JO..#.#.#                 #.....#  
  #.#.#.#                 ###.#.#  
  #...#..DI             BU....#..LF
  #####.#                 #.#####  
YN......#               VT..#....QG
  #.###.#                 #.###.#  
  #.#...#                 #.....#  
  ###.###    J L     J    #.#.###  
  #.....#    O F     P    #.#...#  
  #.###.#####.#.#####.#####.###.#  
  #...#.#.#...#.....#.....#.#...#  
  #.#####.###.###.#.#.#########.#  
  #...#.#.....#...#.#.#.#.....#.#  
  #.###.#####.###.###.#.#.#######  
  #.#.........#...#.............#  
  #########.###.###.#############  
           B   J   C               
           U   P   P               
`

var input_aoc = `                                       L N           Q     O     W   M       F                                         
                                       W G           V     W     F   Q       F                                         
  #####################################.#.###########.#####.#####.###.#######.#######################################  
  #.......#.#...........#...#.#.#.......#.....#.....#.....#.....#.#...#.....#.........#...#.....#.#.#.#...#.....#.#.#  
  #.#####.#.###########.###.#.#.###.#########.###.#####.#######.#.###.#.#.#####.#######.###.#####.#.#.#.###.#####.#.#  
  #...#.#.#...#.#.#.#...#...............#.....#.........#.....#.#.#.....#.#.............#.......#.#...#.#.....#.#...#  
  ###.#.#####.#.#.#.###.###.#.#########.###.#####.#########.#.#.#.#######.###.#####.#######.#####.#.###.###.#.#.#.###  
  #.#.#.........#...........#...#.......#.......#...#.......#...#.#...#.....#.#.#.#.#.#.#.....#.#.#...#...#.#.#.#.#.#  
  #.#.###.#.###.###.#.#.#####.#####.###.#######.###.#######.#.###.#.#######.#.#.#.###.#.###.###.#.#.###.#.#.###.#.#.#  
  #.....#.#...#.....#.#.#...#.#.....#.....#...#.#...#...#...#.#.......#.....#.........................#.#.#...#.#...#  
  #.#.#########.###.###.###.###.###.#######.#.#.###.#.#.#.#########.#####.###.###.#.#.###.#.#######.###.###.###.#.###  
  #.#.#.#.......#...#...#.......#.......#...#.....#...#.#.#.......#.#...#...#.#.#.#.#...#.#.#.#...#.........#.....#.#  
  ###.#.#.#.###.###########.#######.#.###.###.#########.#.#.###.###.###.###.#.#.#############.###.#####.###.###.###.#  
  #.#.#.#.#...#.#...........#...#...#.#.....#.#.#...#...#.....#.#...#.......#.......#.#.#.....#.#.....#.#.#.#.#.#...#  
  #.#.#.###.###.#.#####.###.###.###.#####.#####.#.#.#.#########.#.#####.###.###.#####.#.#.#####.###.#####.###.#.#.###  
  #.#...#...#...#.....#.#...#.........#.#.......#.#.....#.......#.#.....#.#...#.....#.#.#...........#...#.#.#.....#.#  
  #.#.#.###########.###########.#.#.###.###.#####.###.###.#######.###.###.#.###.#####.#.#.###.###.#.#.###.#.###.###.#  
  #...#.#.......#.....#.........#.#.....#.....#.#.#...#.......#...#.#.#.#.#.#.....#...#...#.....#.#.......#.#.....#.#  
  ###.###.#.#######.###.###.#.###.###.###.#####.#.###.#####.#.###.#.#.#.#.###.#####.###.#.#######.#########.###.###.#  
  #...#...#.#.......#...#...#.#.....#.#.#...#.....#.....#...#.#.....#...#...#...#.#.....#.......#...#.#.#.#.#.#.#.#.#  
  ###.###.#####.###.#########.#######.#.#.#######.#######.#.#######.#.#####.#.###.#.#################.#.#.#.#.#.#.#.#  
  #...#.#...#.#.#.#.#...#.#.#.#.#.......#.#.#.....#.#...#.#...#.....#.#.#.............#...#.#.....#.......#.#.#.#...#  
  ###.#.#.###.###.#.###.#.#.###.###.###.#.#.#####.#.###.###.#####.#.#.#.###.#######.###.###.#.###.#.#.#.###.#.#.#.###  
  #.#.#.#.#.#.....#.#.......#...#.#.#...#.#.#.#.......#.#...#.....#.#...#.#.#...#.....#.#.......#.#.#.#...#.....#.#.#  
  #.#.#.#.#.###.###########.###.#.#.#####.#.#.#.###.###.#.#######.###.###.#.###.#.###.#.#.#.###.#######.###.#.#.#.#.#  
  #.....#.........#.#.#...............#.....#.#.#...#...........#...#...#.#.#.....#.#.#.#.#.#.........#...#.#.#...#.#  
  #.#####.#.#######.#.#######.###.###.#####.#.###.#######.#.#.###.###.###.#####.#.#.###.#.#.###.###.#####.#####.###.#  
  #...#...#.....#.......#.....#...#.#...#...#.........#.#.#.#.#.....#.........#.#.........#...#...#.#.......#...#.#.#  
  #.#######.#######.###.###.###.###.#.#####.#######.###.#.#######.###.#############.#######.#########.#######.###.#.#  
  #...#.....#.#.....#...#.#.#...#.#.#.....#.#.......#.#.........#.#.#.....#.......#.#.#...#...#.#...#...#.#...#.#.#.#  
  ###.###.###.#####.#####.#######.#.#.###.#.#####.###.#####.#.###.#.###.#####.###.#.#.#.#######.###.#.#.#.###.#.#.#.#  
  #.....#.#...#.....#.....#.............#.#.....#.........#.#.#.....#.....#.....#.......#.#...#...#.#.#...#...#.#...#  
  ###.#.#.#.#.#####.#####.#.###.#######.#####.#######.#######.#####.###.#####.###########.#.#####.#.###.#.#.###.###.#  
  #...#.#...#.....#.#.#...#.#.#.#      P     Q       O       P     F   A     M        #...#...#.......#.#.#.#.......#  
  ###.#####.###.###.#.#.#####.###      C     V       W       B     F   Y     P        ###.#.#####.#######.#.###.###.#  
  #...#.#...#.#.#...#.#.#...#...#                                                     #.....#.....#.#.#...#...#.#....PB
  ###.#.#.###.#####.#.#.###.#.###                                                     #.###.###.###.#.#.###.#####.#.#  
  #.#.........#.#.#.............#                                                   MQ..#...........#.#.....#.....#..AA
  #.#.#####.###.#.#####.#.#######                                                     #.#.###.#.#.###.#####.#.###.###  
GG....#.#...#...#.....#.#.....#.#                                                     #.#.#...#.#.....#.#.#.#.#.#.#.#  
  #.###.#####.#####.#####.#####.#                                                     ###.#######.#####.#.#.###.#.#.#  
  #...#...#.#.#.....#.#...#......LE                                                   #.#.#...#.#...................#  
  #.#####.#.#.###.###.#.#.#.#.#.#                                                     #.###.###.#####################  
SS..#.....#...#...#.....#...#.#..EC                                                 LJ..#...........#...#...#.......#  
  #.#.#.###.#.###.###.###.#######                                                     #.#.#.#####.#.#.#.#.###.#.#####  
  #...#.....#.........#.#.#.....#                                                     #...#.#.#...#...#...#...#.#....PC
  #####################.###.#.###                                                     #.#####.#####.#.###.#.#.#.###.#  
  #.......#...#.........#.#.#...#                                                     #.....#...#...#.#...#.#.#.#...#  
  ###.###.#.#.#.###.#.#.#.###.#.#                                                     #.#####.#.#####.###.#.###.#.###  
  #.#.#...#.#...#...#.#.#...#.#.#                                                     #...#.#.#...#...#.....#.....#.#  
  #.#.#.###.#.#.#####.###.###.#.#                                                     #####.#.#######.#.###.#####.#.#  
NW....#...#.#.#.....#.#...#.#.#.#                                                   NW....#.#...#.....#.#...#.#...#.#  
  #.#.###.#.#.###.###.###.#.#.#.#                                                     #.###.#.#.###.#.#.#.#.#.#.#.#.#  
  #.#.#.....#.#.#.#.#.........#..BB                                                   #.#.#.#.#.#.#.#.#.#.#.#.#.#.#..LE
  ###.#########.###.#############                                                     #.#.#.#.###.###########.#####.#  
  #.#.#...................#......LV                                                   #...........#.#...#.#...#.#.#.#  
  #.#####.#######.#####.###.#####                                                     #####.#####.#.###.#.#.###.#.#.#  
LB............#...#.#...#...#...#                                                     #...#.#.#.....................#  
  #.#####.#.#.###.#.###.###.#.#.#                                                     #.#####.###.###################  
  #...#...#.#...#.#.#.#.#.#...#.#                                                   UJ........#.#.#.....#...#.......#  
  #.#####.###.###.#.#.#.#.#####.#                                                     #.###.#.#.###.###.#.#.#.###.#.#  
  #...#.....#.#...#.#.#.........#                                                     #...#.#.........#...#.....#.#.#  
  ###.#.#######.###.#.#######.#.#                                                     #.###.#####.###.#.#.#######.#.#  
  #...#.#.#...#.#.#.......#...#.#                                                     #...#.#.#...#...#.#.#...#.#.#..BB
  #.#.###.#.#####.#.#.###.#######                                                     #######.#######.#.#####.#.#####  
  #.#...#.#.#.#.#.#.#...#...#....GG                                                 JY..#...........#.#...#.....#....CD
  #######.#.#.#.#.#.#####.#####.#                                                     #.###.###.###.#########.#.###.#  
AY..#.#...#.#.#.......#...#.#...#                                                     #.#.#.#.....#.#.#.#.....#.#...#  
  #.#.###.#.#.#.#.#.#####.#.#.###                                                     #.#.#.#####.###.#.###.#.#.#.#.#  
  #.............#.#...#.#.....#.#                                                     #.......#...#.....#.#.#.#...#.#  
  ###########.#####.###.#.#.###.#                                                     #####.###.#####.###.###.###.###  
  #.........#...#...#...#.#.#.#.#                                                     #.#...#.................#.#.#.#  
  #####.###.#########.#######.#.#                                                     #.#######.#.#.#####.#####.#.#.#  
  #.....#...#.#.#.#.....#...#.#..LB                                                   #...#.#.#.#.#...#.....#...#.#.#  
  #.#.#.#.###.#.#.###.#####.#.#.#                                                     #.#.#.#.#####.#####.###.#.###.#  
DD..#.#.#.#...#...#.#.#.......#.#                                                     #.#.........#.#.#...#.#.#.#...#  
  #.#####.#.#.#.###.#.###.###.#.#                                                     #.###.#########.#####.#.#.#.#.#  
  #...#.....#...............#...#                                                   DD..#.....#.#.#.#.#.....#.#.#.#..UJ
  ###.###########################                                                     ###.#####.#.#.#.#.#.###.#.#.###  
  #.#.#...#.........#...........#                                                     #.................#.....#.....#  
  #.###.#.#.#######.#.#####.#####                                                     #.#.#.#.###############.#######  
  #.....#...#...#...#.#.#...#....GK                                                   #.#.#.#.#.........#.#.#.#.....#  
  #.#.#########.###.#.#.#.###.###                                                     #########.#####.###.#.###.###.#  
  #.#.....#.............#.....#.#                                                   LW..#...#.....#.........#...#....LV
  #.###.#.###.#.#.#.###.#.#.#.#.#                                                     #.#.###.#.#.###.#.###.#.###.###  
JY..#.#.#...#.#.#.#...#.#.#.#...#                                                     #.....#.#.#...#.#.#...#.#...#..ZZ
  #.#.###.###.#.###.###.###.#####                                                     #.###.#.###.#######.###.#.###.#  
  #.....#.#.#.#...#...#...#.....#                                                     #...#.....#.#...........#.....#  
  #.#.#####.#.#####.###.###.#.###      A       N           C       S X     W          #.#####.#######.#.###.#.#####.#  
  #.#...#.....#.......#...#.#.#.#      L       G           D       S Z     F          #.#...........#.#.#...#.#.....#  
  ###.###.#.#.###.#.###.#.###.#.#######.#######.###########.#######.#.#####.###########.###.#.#.#.###.#.###.#######.#  
  #.....#.#.#.#...#.#...#.#.....#...........#.........#...#.#.......#.....#.....#...#.#...#.#.#.#.#.#.#.#.....#.#...#  
  ###.#.#####.#.###.#.#.#.###.#.#.#.###.###.###.###.###.###.#.###########.#.#####.###.###.###.#####.###.###.#.#.###.#  
  #...#...#...#...#.#.#.#.#...#.#.#...#...#.#.#.#...#.......#...........#.#.#.#.......#.....#...#.......#...#...#...#  
  ###.#######.#.#.###.###.###.#########.###.#.#####.###.###.#.#.#.###.###.#.#.#.#.#.#.#######.#########.###.#.#####.#  
  #.#...#...#.#.#.#...#.#.#.........#.#.#.....#.#.#.#.....#.#.#.#.#.#...#.#.....#.#.#.#.......#...........#.#.#.....#  
  #.#.#####.#.###.#.#.#.#############.#####.###.#.#.#.#######.#####.#####.#####.#########.#.#.#########.###.#######.#  
  #.......#.....#.#.#.#.#...#.#.#.#...#.#.......#...#...#.#.#...#.........#.............#.#.#.#...#.......#...#.#...#  
  ###.#####.#####.#####.###.#.#.#.###.#.#.###.#####.#.#.#.#.#.###########.#.#############.#.#.###.###.###.###.#.#.#.#  
  #.....#.......#.#.......#.......#.........#.#.#...#.#...#.........#.....#...........#.#.#.#.#.#.....#.#.#...#...#.#  
  ###.###.#.###.#.###.#######.###.#.#######.###.###.###.#########.#######.###.#########.#####.#.#.#.#.#.###.###.#.#.#  
  #.....#.#.#.#.#.#.#...#.#.#.#.....#...#.....#.......#...#.#.......#.#.#.#.#...#...#.#...#...#...#.#...#.#...#.#.#.#  
  ###.#######.#####.###.#.#.#######.###.###.#######.###.#.#.#.#######.#.#.#.###.#.###.#.#####.#####.#.#.#.#####.#.#.#  
  #.......#.................#.....#.#.#.#.....#.....#.#.#...#.......#.......#.#...#.......#.....#...#.#.......#.#.#.#  
  #.#.#############.###.#.#.#.#.#.###.#.#.#####.#.###.###.#.#.#.#####.###.#.#.#.###.#####.###.###.#####.#.###.#.###.#  
  #.#.........#...#...#.#.#...#.#...#.#.#.....#.#.....#...#.#.#.#...#.#...#.#.......#...#.#.....#...#...#...#.#.#...#  
  #.###.#######.#.#.###############.#.#.#.#.#######.#######.#.###.#######.#####.#.###.#################.#########.#.#  
  #.#.....#.....#.#.#...#...#.............#.#.........#.....#.....#.........#.#.#.....#.#.......#.......#.#.#.#.#.#.#  
  ###.###########.#####.###.###.#.#.#.###########.#.#######.###.#.###.#.#####.###.#.###.#####.#.#######.#.#.#.#.#####  
  #.......#...#.#...#.....#.#...#.#.#...#.....#...#.....#...#...#.#...#...#...#.#.#.#.#...#...#.#.....#...#.#.....#.#  
  #.#.#######.#.#.#######.#.#######.#.#.#.#.#.#.###.#######.#.###########.#.###.#.###.#.#####.###.#.#.###.#.#.#.###.#  
  #.#.....#.......#...#.#...#.#.#...#.#...#.#.#.#.....#...#.#.....#.#.........#...#.....#.....#.#.#.#.#.......#.....#  
  #.#.###.#######.#.###.###.#.#.###########.#####.#####.###.###.#.#.#.#####.###.#####.#######.#.#.#######.#####.###.#  
  #.#...#.....#.......#...........#.#.....#.#.......#...#...#...#.#.......#.#.......#.#.#.#...#...#.#...#.#...#.#...#  
  #.#####.#######.###.#.#.###.###.#.###.#.#.#.###.###.#.#.###.#########.#.#######.###.#.#.###.#.###.#.#####.#####.#.#  
  #.#.......#.....#.....#.#...#.........#...#...#.#.#.#.#...#.......#...#.......#.....#.#.......#...#.#.#...#...#.#.#  
  #.#########.#.#.#############.#.###.#####.###.#.#.#.#.#.#.#####.#######.#.#####.#.#.#.#.###.#.###.#.#.###.#.#####.#  
  #.#.#.#.....#.#...#.....#.....#.#.......#.#...#.#...#...#.#.........#...#.#...#.#.#...#...#.#...........#.......#.#  
  ###.#.#########.###.#####.###.#####.###.###.#######.#############.#####.###.#.#.#####.#.#.###.#######.###.#.###.###  
  #...............#.........#...#.....#...#.......#...........#.....#.....#...#.......#...#...#.....#.......#...#...#  
  #####################################.#########.#####.#########.#####.#####.#######################################  
                                       X         E     A         M     L     G                                         
                                       Z         C     L         P     J     K                                         
`

var input_jan = `                                 C X   J         X     B         U       N     J                                 
                                 N Z   I         Y     U         D       U     M                                 
  ###############################.#.###.#########.#####.#########.#######.#####.###############################  
  #...#.#.#.......#...#...#...........#...#.#...#...#...#.#...#...#.........#.....#...#.#.......#.......#.#.#.#  
  ###.#.#.#.#.###.###.###.#####.#######.###.#.#####.#.###.#.#.###.#.#.#.#.#######.#.###.#.###.###.#######.#.#.#  
  #.#...#...#.#.#...#...#.#.......#...#.#.#...#.....#.#...#.#...#.#.#.#.#.#.#.....#...#...#.#...#.#...#...#...#  
  #.###.#.#####.#.###.#.#.#.#####.#.#.#.###.#.###.###.###.#.###.#.#.#######.###.###.###.###.#####.#.#####.#.###  
  #.......#...#.#...#.#.....#.#...#.#.#.#...#.#.....#.....#...#...#.#.....#...#...#.....#.#.....#.#...#.#.....#  
  #.###.#.###.#.###.#####.#.#.#.###.#.#.#.###.#.###.#####.#.#####.#.#.###.#.#.###.#.###.#.#.###.#.#.###.###.###  
  #.#...#.#.....#.#...#.#.#.#.......#.#...#...#.#.#.#.#...#...#...#.....#.#.#.......#.....#...#.....#.#.......#  
  #.###########.#.###.#.#############.#####.#.#.#.###.#.###.#####.#.#.###.#.###.###.###.#.#.#########.###.#####  
  #.#...........#...#.............#...#.#...#.#.....#...#.#.....#.#.#.#.#.#.#...#.#.#.#.#.....#.#...#.........#  
  #########.###.###.#######.###.###.###.#.#####.#######.#.###.#########.#.#.###.#.###.#.###.###.###.###.###.#.#  
  #...#.#.#.#...#.#.#.#.....#.#.......#.......#...#.#.....#.....#.....#...#...#.....#.....#.#...#.#.......#.#.#  
  #.###.#.#####.#.#.#.###.###.###.#.#.###.#####.###.#.#####.###.#####.#.#.#.#################.###.###.#.#.#####  
  #.............#.....#...#.#.....#.#.#.#.....#...#.....#...#.....#...#.#.#.....#.#...#.....#.#...#...#.#...#.#  
  #############.###.#.#.###.###.#.#####.#.#####.#.###.###.#.###.#####.###.#.###.#.#.#####.#.#.###.###.#######.#  
  #.#.....#.....#.#.#.#.#.......#.#...#.......#.#.#...#.#.#...#.#...#.#...#...#.#.#...#.#.#.....#.#.#.#.#.#...#  
  #.#####.#####.#.###.#######.#.###.#.#.#######.#####.#.###.#.#.###.#.#.###.#####.###.#.###.#.#.#.#.#.#.#.#.#.#  
  #.....#.#.............#.#...#...#.#.......#.#...#.....#...#.#.#.#.......#.#.......#.......#.#...........#.#.#  
  #####.#.#.###.###.#.###.#####.#########.###.###.###.###.#.#####.#######.#.###.#####.###.#.#########.#######.#  
  #.#.#.......#.#...#.................#.....#.......#...#.#.......#.#.....#...#...#.....#.#.....#.......#...#.#  
  #.#.###.#.#####.###.#.#.###.#.###.###.###.#.#.#####.#####.###.###.#.###.#.###.###.###############.#.#.###.#.#  
  #.#...#.#.#.....#...#.#...#.#.#.#.#...#.#.#.#.....#.....#.#...#.#.....#.#.....................#...#.#.#.#...#  
  #.###.#############.###########.#.###.#.#######.#####.#######.#.###.#.#########.#######.###.#.#.#####.#.#.###  
  #.#...#...#.#.#...#.....#...........#...#.#.........#.....#.#.....#.#.#.#.........#.#.#.#...#.#.#.#.#.#.....#  
  #.###.#.###.#.#.#####.###.#.#######.#.#.#.###.###.###.#.###.###.#.###.#.###.#.###.#.#.###########.#.#######.#  
  #...#.#.......#.........#.#.#.......#.#.....#...#...#.#.#.......#.#.....#...#...#.#...#.....#...#...........#  
  #.###.#.#########.#.###.#.#######.#.#######.###.#######.#.#############.#.###########.###.###.#############.#  
  #.#.........#.....#.#.#.#.#      M E       U   H       U D             V X        #.#.#.....#.....#.#.#.#...#  
  #.###.###############.#####      F Y       D   L       I R             E Z        #.#.#.#######.###.#.#.###.#  
  #.#.....#.#...#.....#.#.#.#                                                       #.......#.#.#.#...#.#...#.#  
  #.#####.#.###.#.#####.#.#.#                                                       ###.###.#.#.#.#.###.###.#.#  
  #.#.#...#...#...#...#.#...#                                                       #...#.#.....#...#.#.....#.#  
  #.#.###.#.#####.###.#.#.###                                                       ###.#.#.#######.#.#####.#.#  
  #.............#.......#...#                                                       #.#...#...#.#.#...#.#......VE
  #.#######.#####.###.###.###                                                       #.###.#.###.#.###.#.#.#.###  
  #...#.........#...#...#.#.#                                                     TJ......#...............#...#  
  ###.#.#########.#.#####.#.#                                                       #.###.#########.###########  
DR..#.#...#.....#.#.#...#.#..HP                                                     #.#...#.#.#.#.#.#.........#  
  #.#.###.#.###.#.#.###.#.#.#                                                       #.###.#.#.#.#.###.#####.###  
  #.....#...#.#...#.........#                                                       #...#.#...#.#.......#...#.#  
  #######.###.###############                                                       #######.###.#####.###.###.#  
  #.....#.#.......#.#.....#.#                                                       #.#...#.#.#.#...#...#.#....TJ
  #.###.###.###.#.#.#.###.#.#                                                       #.###.#.#.#.#.###.###.#.###  
HP..#.....#...#.#...#...#...#                                                     BU..................#.....#.#  
  #####.###.#####.#.###.#.#.#                                                       #########.#########.#####.#  
  #.........#.....#...#.#.#.#                                                     TC........#.#.#...#.#.#.....#  
  ###########.###.#.#.#.###.#                                                       ###.###.###.#.#.#.#####.#.#  
  #.#.#.....#.#.#.#.#...#.#..FN                                                     #...#.........#...#.....#.#  
  #.#.###.#####.#########.###                                                       #########.#####.#.#.#####.#  
MF..#.#...#.....#...#...#....XY                                                     #...#.#...#.#.#.#.....#.#..FN
  #.#.#.#.#####.#.###.#.#.###                                                       ###.#.#####.#.#####.###.###  
  #.#...#.#.#.........#.#.#.#                                                       #.................#.#.#...#  
  #.#.#.###.###.#.###.#.#.#.#                                                       #.###.###.#.#.###.###.#.#.#  
  #...#.........#...#.#.....#                                                       #.#.....#.#.#.#.#.....#.#..WI
  ###############.#########.#                                                       #.#####.#.#.#.#.#####.#.###  
  #...........#...#...#.#...#                                                     EH..#.#...#.#.#.#...#...#...#  
  #.#####.###.#####.###.#####                                                       ###.#############.#.#####.#  
  #.#...#.#.....#.#.......#.#                                                       #.....#.#...#...#.........#  
  #.#######.#.###.#.###.###.#                                                       ###.###.###.###.#########.#  
  #.......#.#.#.....#...#....YW                                                     #.......#...#...........#.#  
  #####.###.#####.#.###.###.#                                                       #.###.###.#.#.###.#.#######  
TA........#.......#.#.......#                                                     MR..#.#...#.#.....#.#...#....EY
  ###.#.#####.#####.###.#####                                                       #.#.#.###.###########.###.#  
  #...#.#...#.#.....#.......#                                                       #.#.#...#.....#.#.#.......#  
  #######.#.#############.#.#                                                       ###.#.#.###.###.#.#####.###  
  #...#...#.#.#.....#...#.#.#                                                       #...#.#.....#...#.#.#...#.#  
  #.###.###.#.#.#######.#####                                                       ###.###########.#.#.#####.#  
MR....#...#.......#.....#.#.#                                                       #.......#...............#..FM
  #.#####.#####.###.#####.#.#                                                       #.#.#.###.###.#.#####.###.#  
  #.........#................NU                                                   JI..#.#...#.#.#.#...#...#...#  
  ###########################                                                       #.#.#.###.#.#.#######.#.###  
  #...#.#...........#.#.....#                                                       #.#.#...#...#.....#...#...#  
  #.#.#.###.#####.#.#.###.#.#                                                       #######.#.###########.###.#  
ZZ..#.....#.....#.#.#.....#.#                                                       #.#.#.......#...#...#.....#  
  #.#####.#.#######.###.###.#                                                       #.#.#########.#####.#######  
  #...#.....#.#.#.#.....#.#..TA                                                     #.#.#.............#.....#..BW
  #.###.###.#.#.#.#.#.#.#.#.#                                                       #.#.#.#.#.#.#.###.#.###.#.#  
UI....#.#.......#...#.#.#...#                                                       #.....#.#.#.#.#.......#...#  
  #.###.#.#####.#####.#.#####                                                       #.#.###.#####.#####.###.#.#  
  #.#...#.#.#.....#...#.....#                                                     HM..#.#.....#.....#...#...#.#  
  #.#######.#.#######.#######                                                       #.#####.#####.#####.###.#.#  
  #.....#.......#.....#...#.#                                                       #...#...#...#...#.....#.#.#  
  #.#######.#.###.#.#.#.#.#.#      C       F         W         J   B     F          #.#########.#.###.#####.###  
  #.#.#.#.#.#.#...#.#.#.#.#.#      N       M         I         M   W     A          #...#.........#...#.......#  
  #.#.#.#.#####.#####.###.#.#######.#######.#########.#########.###.#####.###########.#######.###.###.#####.###  
  #.........#.....#...#.#.....#.#.#.....#...........#...#.....#.#...#...............#.......#...#.#.....#.....#  
  ###.#.#.#.#.#.###.###.#.#.#.#.#.#.###.#.#.###.#######.#.###.#.#.#######.###.#.#.#.###.#####.#.###.#####.#.#.#  
  #...#.#.#.#.#.#.....#...#.#.........#.#.#...#.#.#.#...#...#.#.#.#...#.#.#.#.#.#.#...#...#...#.#...#.....#.#.#  
  ###.#.#.#.#.#.#######.#.#######.#####.#######.#.#.###.###.#.#.#.#.#.#.###.###.#####.#.###.#######.#.###.#.#.#  
  #...#.#.#.#.#.....#...#.#.#.....#.#.....#.#...#.......#...#.#.#...#.......#.....#.#.#.#.......#...#...#.#.#.#  
  #######.#####.###########.#####.#.###.###.#.#.#.#######.###.#.###.###.#######.###.###########.#####.###.###.#  
  #...........#.....#.........#.#.#.......#...#.#.....#...#...#.#.#.#.#.#.#.#...#.#.......#.#.#.#.#...#.....#.#  
  #####.#.###.#.#.###########.#.#######.###.#######.#.###.#.###.#.###.#.#.#.###.#.#.#.#.###.#.###.#####.#.#.#.#  
  #.....#...#.#.#.#...#.#.#.......#.......#.....#.#.#...#.#.......#.......#...#.....#.#.............#.#.#.#.#.#  
  #.#.#.###.###.#####.#.#.#######.###.#####.#.#.#.#####.#.#.###.#####.#######.#.###########.#.#####.#.###.#.#.#  
  #.#.#...#.#.#.#.....#...............#.#.#.#.#...#.#.#.#.#.#...#...........#.....#.#.#...#.#.....#.....#.#.#.#  
  #.###.###.#.#######.#.#.#####.#.###.#.#.#####.#.#.#.#.#.#####.#######.#####.#####.#.#.#.###.#.#########.#.#.#  
  #...#.#.........#.....#...#.#.#...#.......#.#.#...#...#.....#.#...#...#.#.#.#.#...#.#.#...#.#.......#...#.#.#  
  #.###.###.#.###.#.###.#####.#########.#####.#.#######.#.#######.###.#.#.#.#.#.###.#.#.#######.###.#.###.#.#.#  
  #.#.....#.#.#...#.#...#.#.#.#.#.#...#.....#.....#...#.#.......#...#.#.#...............#.#...#.#.#.#.#...#.#.#  
  #.#.#####.#.#####.#####.#.#.#.#.###.#.#######.###.###.#.#########.###.#.#.#.#.###.#.###.#.#####.###.###.#.###  
  #.#.....#.#...#...#.#.#.#.....#.#.#.#...#.....#.......#.....#.#.#.#.#.#.#.#.#...#.#.#.............#...#.#...#  
  ###.#####.#.#####.#.#.#.#.###.#.#.#.#.#######.#.###.###.#####.#.#.#.#.#.#########.###.###.###.#########.###.#  
  #...#.....#.....#.#.......#...........#.#.#...#...#.#.....#...........#.#.#.#.........#...#.#.#...#.#.#.#...#  
  #.#####.###.#####.#############.#.#.###.#.#.#.#.#.#.###.#######.#####.#.#.#.#####.#########.###.###.#.#.#####  
  #...#.....#...#...#.............#.#...#.....#.#.#.#.#.......#.#.....#.#.#.....#...#.....#.#...#.....#.#.....#  
  #.#####.#.###########.#.#.###.#######.#.#######.###.###.#####.###.#####.#.#####.#.#.#.#.#.#.###.#.###.#######  
  #...#...#...#.........#.#.#...#.......#.....#...#...#.#...#...#.......#.#.#...#.#...#.#.........#...........#  
  #.#######.#####.#.#####.#.#####.#####.#####.###.#####.#.###.#.#.###.###.#.###.###.#.#####.###.#.###.#.###.#.#  
  #.#.........#...#.....#.#.#.....#.......#.....#.....#.......#.#...#.#...........#.#.....#...#.#.#...#...#.#.#  
  #####################################.###.#######.#########.###.#.#######.###################################  
                                       H   Y       H         F   A E       T                                     
                                       M   W       L         A   A H       C                                     
`

func Test_part1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{input1}, 23},
		{"2", args{input2}, 58},
		{"input", args{input_aoc}, 560},
		{"jan", args{input_jan}, 482},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, pp := buildWorld(tt.args.input)
			allpaths := pathfinder.NewAllPaths(pp)
			graph := buildGraph(pp, allpaths)

			vAA, _ := graph.GetMapping("AAe")
			vZZ, _ := graph.GetMapping("ZZe")
			bestPath, _ := graph.Shortest(vAA, vZZ)
			if int(bestPath.Distance) != tt.want {
				t.Errorf("part1() got = %v, want %v", bestPath.Distance, tt.want)
			}
		})
	}
}

func Test_debug_part1(t *testing.T) {
	g, pp := buildWorld(input1)
	fmt.Println(g)
	fmt.Println(pp)

	allpaths := pathfinder.NewAllPaths(pp)
	fmt.Println(allpaths)
	graph := buildGraph(pp, allpaths)

	vAA, _ := graph.GetMapping("AAe")
	vZZ, _ := graph.GetMapping("ZZe")
	bestPath, _ := graph.Shortest(vAA, vZZ)
	fmt.Println(bestPath)
	s := "["
	for _, p := range bestPath.Path {
		vertex, _ := graph.GetMapped(p)
		s += fmt.Sprintf("%v ", vertex)
	}
	s += "]"
	fmt.Println(s)

	assert.Equal(t, 23, int(bestPath.Distance))
}

/*
Jan edge weights
CN-05-XZ
CN-59-EY
CN-61-MF
XZ-61-EY
XZ-63-MF
MF-05-EY
JI-57-UD
XY-43-HL
BU-41-UI
UD-57-DR
NU-55-VE
JM-55-XY
DR-67-HP
HP-59-FN
MF-47-XY
TA-61-YW
MR-45-NU
UI-53-TA
ZZ-49-TA
ZZ-07-UI
TJ-45-VE
BU-45-TJ
TC-43-FN
EH-53-WI
MR-55-EY
JI-57-FM
HM-39-BW
CN-39-HM
YW-51-FM
HL-49-WI
FA-61-JM
EH-55-BW
AA-05-EH
AA-53-BW
*/

var input3 = `             Z L X W       C                 
             Z P Q B       K                 
  ###########.#.#.#.#######.###############  
  #...#.......#.#.......#.#.......#.#.#...#  
  ###.#.#.#.#.#.#.#.###.#.#.#######.#.#.###  
  #.#...#.#.#...#.#.#...#...#...#.#.......#  
  #.###.#######.###.###.#.###.###.#.#######  
  #...#.......#.#...#...#.............#...#  
  #.#########.#######.#.#######.#######.###  
  #...#.#    F       R I       Z    #.#.#.#  
  #.###.#    D       E C       H    #.#.#.#  
  #.#...#                           #...#.#  
  #.###.#                           #.###.#  
  #.#....OA                       WB..#.#..ZH
  #.###.#                           #.#.#.#  
CJ......#                           #.....#  
  #######                           #######  
  #.#....CK                         #......IC
  #.###.#                           #.###.#  
  #.....#                           #...#.#  
  ###.###                           #.#.#.#  
XF....#.#                         RF..#.#.#  
  #####.#                           #######  
  #......CJ                       NM..#...#  
  ###.#.#                           #.###.#  
RE....#.#                           #......RF
  ###.###        X   X       L      #.#.#.#  
  #.....#        F   Q       P      #.#.#.#  
  ###.###########.###.#######.#########.###  
  #.....#...#.....#.......#...#.....#.#...#  
  #####.#.###.#######.#######.###.###.#.#.#  
  #.......#.......#.#.#.#.#...#...#...#.#.#  
  #####.###.#####.#.#.#.#.###.###.#.###.###  
  #.......#.....#.#...#...............#...#  
  #############.#.#.###.###################  
               A O F   N                     
               A A D   M                     
`

func Test_part2(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{input1}, 26},
		{"3", args{input3}, 396},
		{"input", args{input_aoc}, 6642},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, pp := buildWorld(tt.args.input)
			allpaths := pathfinder.NewAllPaths(pp)
			graph := buildGraph2(pp, allpaths)

			vAA, _ := graph.GetMapping("AAe_0")
			vZZ, _ := graph.GetMapping("ZZe_0")
			bestPath, _ := graph.Shortest(vAA, vZZ)
			if int(bestPath.Distance) != tt.want {
				t.Errorf("part2() got = %v, want %v", bestPath.Distance, tt.want)
			}
		})
	}
}

func Test_debug_part2(t *testing.T) {
	g, pp := buildWorld(input1)
	fmt.Println(g)
	fmt.Println(pp)

	allpaths := pathfinder.NewAllPaths(pp)
	fmt.Println(allpaths)
	graph := buildGraph2(pp, allpaths)

	vAA, _ := graph.GetMapping("AAe_0")
	vZZ, _ := graph.GetMapping("ZZe_0")
	bestPath, _ := graph.Shortest(vAA, vZZ)
	fmt.Println(bestPath)
	s := "["
	for _, p := range bestPath.Path {
		vertex, _ := graph.GetMapped(p)
		s += fmt.Sprintf("%v ", vertex)
	}
	s += "]"
	fmt.Println(s)

	assert.Equal(t, 26, int(bestPath.Distance))
}
