package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// EnemyData   敵のデータ
type EnemyData struct {
	uid       int        `json:"uid"`
	tribe     int        `json:"tribe"`
	position  [3]float32 `json:"position"`
	spawnTime float32    `json:"spawn-time"`
}

var enemies = []EnemyData{
	// wave1 チュートリアル
	EnemyData{
		uid:       1,
		tribe:     1,
		position:  [3]float32{8, 0, 10},
		spawnTime: 2.0,
	},
	EnemyData{
		uid:       2,
		tribe:     1,
		position:  [3]float32{-4, 0, 12},
		spawnTime: 6.0,
	},
	EnemyData{
		uid:       3,
		tribe:     1,
		position:  [3]float32{6, 0, -5},
		spawnTime: 8.0,
	},
	EnemyData{
		uid:       4,
		tribe:     1,
		position:  [3]float32{11, 0, 10},
		spawnTime: 11.0,
	},
	EnemyData{
		uid:       5,
		tribe:     1,
		position:  [3]float32{12, 0, 13},
		spawnTime: 12,
	},
	EnemyData{
		uid:       6,
		tribe:     1,
		position:  [3]float32{0, 0, 14},
		spawnTime: 13,
	},
	EnemyData{
		uid:       7,
		tribe:     1,
		position:  [3]float32{3, 0, 10},
		spawnTime: 18,
	},
	EnemyData{
		uid:       8,
		tribe:     1,
		position:  [3]float32{-3, 0, -10},
		spawnTime: 19,
	},
	EnemyData{
		uid:       9,
		tribe:     1,
		position:  [3]float32{11, 0, 1},
		spawnTime: 20,
	},
	// wave2 ラッシュ
	EnemyData{
		uid:       10,
		tribe:     1,
		position:  [3]float32{13, 0, 0},
		spawnTime: 26,
	},
	EnemyData{
		uid:       11,
		tribe:     1,
		position:  [3]float32{12, 0, 1},
		spawnTime: 26.9,
	},
	EnemyData{
		uid:       12,
		tribe:     1,
		position:  [3]float32{11, 0, 2},
		spawnTime: 27.8,
	},
	EnemyData{
		uid:       13,
		tribe:     1,
		position:  [3]float32{10, 0, 3},
		spawnTime: 28.6,
	},
	EnemyData{
		uid:       14,
		tribe:     1,
		position:  [3]float32{9, 0, 4},
		spawnTime: 29.4,
	},
	EnemyData{
		uid:       15,
		tribe:     1,
		position:  [3]float32{8, 0, 5},
		spawnTime: 30.1,
	},
	EnemyData{
		uid:       16,
		tribe:     1,
		position:  [3]float32{7, 0, 5},
		spawnTime: 30.7,
	},
	EnemyData{
		uid:       17,
		tribe:     1,
		position:  [3]float32{6, 0, 4},
		spawnTime: 31.3,
	},
	EnemyData{
		uid:       18,
		tribe:     1,
		position:  [3]float32{-6, 0, -4},
		spawnTime: 31.8,
	},
	EnemyData{
		uid:       19,
		tribe:     1,
		position:  [3]float32{-7, 0, -5},
		spawnTime: 32.3,
	},
	EnemyData{
		uid:       20,
		tribe:     1,
		position:  [3]float32{-8, 0, -5},
		spawnTime: 32.8,
	},
	EnemyData{
		uid:       21,
		tribe:     1,
		position:  [3]float32{-12, 0, 1},
		spawnTime: 33.3,
	},
	EnemyData{
		uid:       22,
		tribe:     1,
		position:  [3]float32{0, 0, -11},
		spawnTime: 33.8,
	},
	EnemyData{
		uid:       23,
		tribe:     1,
		position:  [3]float32{4, 0, 6},
		spawnTime: 34.2,
	},
	EnemyData{
		uid:       24,
		tribe:     1,
		position:  [3]float32{-6, 0, 5},
		spawnTime: 34.6,
	},
	EnemyData{
		uid:       25,
		tribe:     1,
		position:  [3]float32{7, 0, -4},
		spawnTime: 35,
	},
	// wave3 軍隊行進
	EnemyData{
		uid:       26,
		tribe:     1,
		position:  [3]float32{10, 0, 0},
		spawnTime: 41,
	},
	EnemyData{
		uid:       27,
		tribe:     1,
		position:  [3]float32{11, 0, 1},
		spawnTime: 41.3,
	},
	EnemyData{
		uid:       28,
		tribe:     1,
		position:  [3]float32{11, 0, -1},
		spawnTime: 41.6,
	},
	EnemyData{
		uid:       29,
		tribe:     1,
		position:  [3]float32{12, 0, 2},
		spawnTime: 41.9,
	},
	EnemyData{
		uid:       30,
		tribe:     1,
		position:  [3]float32{12, 0, 0},
		spawnTime: 42.2,
	},
	EnemyData{
		uid:       31,
		tribe:     1,
		position:  [3]float32{12, 0, -2},
		spawnTime: 42.5,
	},
	EnemyData{
		uid:       32,
		tribe:     1,
		position:  [3]float32{3, 0, 8},
		spawnTime: 43.2,
	},
	EnemyData{
		uid:       33,
		tribe:     1,
		position:  [3]float32{4, 0, 9},
		spawnTime: 43.5,
	},
	EnemyData{
		uid:       34,
		tribe:     1,
		position:  [3]float32{2, 0, 9},
		spawnTime: 43.8,
	},
	EnemyData{
		uid:       35,
		tribe:     1,
		position:  [3]float32{5, 0, 10},
		spawnTime: 44.1,
	},
	EnemyData{
		uid:       36,
		tribe:     1,
		position:  [3]float32{3, 0, 10},
		spawnTime: 44.4,
	},
	EnemyData{
		uid:       37,
		tribe:     1,
		position:  [3]float32{1, 0, 10},
		spawnTime: 44.7,
	},
	EnemyData{
		uid:       38,
		tribe:     1,
		position:  [3]float32{-6, 0, 1},
		spawnTime: 45.3,
	},
	EnemyData{
		uid:       39,
		tribe:     1,
		position:  [3]float32{-7, 0, 0},
		spawnTime: 45.6,
	},
	EnemyData{
		uid:       40,
		tribe:     1,
		position:  [3]float32{-7, 0, 2},
		spawnTime: 45.9,
	},
	EnemyData{
		uid:       41,
		tribe:     1,
		position:  [3]float32{-8, 0, -1},
		spawnTime: 46.2,
	},
	EnemyData{
		uid:       42,
		tribe:     1,
		position:  [3]float32{-8, 0, 1},
		spawnTime: 46.5,
	},
	EnemyData{
		uid:       43,
		tribe:     1,
		position:  [3]float32{-8, 0, 3},
		spawnTime: 46.8,
	},
	// wave4 ラストラッシュ
	EnemyData{
		uid:       44,
		tribe:     1,
		position:  [3]float32{18, 0, 0},
		spawnTime: 50,
	},
	EnemyData{
		uid:       45,
		tribe:     1,
		position:  [3]float32{16, 0, 2},
		spawnTime: 50,
	},
	EnemyData{
		uid:       46,
		tribe:     1,
		position:  [3]float32{14, 0, 4},
		spawnTime: 50,
	},
	EnemyData{
		uid:       47,
		tribe:     1,
		position:  [3]float32{12, 0, 6},
		spawnTime: 50,
	},
	EnemyData{
		uid:       48,
		tribe:     1,
		position:  [3]float32{6, 0, -12},
		spawnTime: 50.5,
	},
	EnemyData{
		uid:       49,
		tribe:     1,
		position:  [3]float32{4, 0, -14},
		spawnTime: 50.5,
	},
	EnemyData{
		uid:       50,
		tribe:     1,
		position:  [3]float32{2, 0, -16},
		spawnTime: 50.5,
	},
	EnemyData{
		uid:       51,
		tribe:     1,
		position:  [3]float32{0, 0, -18},
		spawnTime: 50.5,
	},
	EnemyData{
		uid:       52,
		tribe:     1,
		position:  [3]float32{-18, 0, 0},
		spawnTime: 51,
	},
	EnemyData{
		uid:       53,
		tribe:     1,
		position:  [3]float32{-16, 0, -2},
		spawnTime: 51,
	},
	EnemyData{
		uid:       54,
		tribe:     1,
		position:  [3]float32{-14, 0, -4},
		spawnTime: 51,
	},
	EnemyData{
		uid:       55,
		tribe:     1,
		position:  [3]float32{-12, 0, -6},
		spawnTime: 51,
	},
	EnemyData{
		uid:       56,
		tribe:     1,
		position:  [3]float32{0, 0, 12},
		spawnTime: 51.5,
	},
	EnemyData{
		uid:       57,
		tribe:     1,
		position:  [3]float32{-2, 0, 14},
		spawnTime: 51.5,
	},
	EnemyData{
		uid:       58,
		tribe:     1,
		position:  [3]float32{-4, 0, 16},
		spawnTime: 51.5,
	},
	EnemyData{
		uid:       59,
		tribe:     1,
		position:  [3]float32{-5, 0, 18},
		spawnTime: 51.5,
	},
	EnemyData{
		uid:       60,
		tribe:     1,
		position:  [3]float32{11, 0, 1},
		spawnTime: 52,
	},
	EnemyData{
		uid:       61,
		tribe:     1,
		position:  [3]float32{-11, 0, -1},
		spawnTime: 52,
	},
	EnemyData{
		uid:       62,
		tribe:     1,
		position:  [3]float32{1, 0, 11},
		spawnTime: 52,
	},
	EnemyData{
		uid:       63,
		tribe:     1,
		position:  [3]float32{-1, 0, -11},
		spawnTime: 52,
	},
	EnemyData{
		uid:       64,
		tribe:     1,
		position:  [3]float32{6, 0, 6},
		spawnTime: 52.5,
	},
	EnemyData{
		uid:       65,
		tribe:     1,
		position:  [3]float32{6, 0, -6},
		spawnTime: 52.5,
	},
	EnemyData{
		uid:       66,
		tribe:     1,
		position:  [3]float32{-6, 0, -6},
		spawnTime: 52.5,
	},
	EnemyData{
		uid:       67,
		tribe:     1,
		position:  [3]float32{-6, 0, 6},
		spawnTime: 52.5,
	},
	EnemyData{
		uid:       68,
		tribe:     1,
		position:  [3]float32{8, 0, 0},
		spawnTime: 53,
	},
	EnemyData{
		uid:       69,
		tribe:     1,
		position:  [3]float32{0, 0, 8},
		spawnTime: 53,
	},
	EnemyData{
		uid:       70,
		tribe:     1,
		position:  [3]float32{-8, 0, 0},
		spawnTime: 53,
	},
	EnemyData{
		uid:       71,
		tribe:     1,
		position:  [3]float32{0, 0, -8},
		spawnTime: 53,
	},
}

// ReadEnemies   GET "/enemy"
func ReadEnemies(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.JSON(200, generateMap(enemies))
}

// KillEnemy   POST "/enemy/kill/:id"
func KillEnemy(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(200, gin.H{"is_enable": false})
		return
	}
	var isReject bool
	enemies, isReject = rejectMap(id, enemies)
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.JSON(200, gin.H{"is_enable": isReject})
}

func generateMap(s []EnemyData) []gin.H {
	ans := make([]gin.H, 0)
	for _, x := range s {
		ans = append(ans, gin.H{
			"uid":        x.uid,
			"tribe":      x.tribe,
			"position":   x.position,
			"spawn_time": x.spawnTime,
		})
	}
	return ans
}

// rejectMap   fがtrueとなる要素を削除する
func rejectMap(id int, s []EnemyData) ([]EnemyData, bool) {
	ans := make([]EnemyData, 0)
	var isReject bool
	for _, x := range s {
		if x.uid != id {
			ans = append(ans, x)
		} else {
			isReject = true
		}
	}
	return ans, isReject
}
