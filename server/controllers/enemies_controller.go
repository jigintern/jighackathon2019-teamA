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
	spawnTime int        `json:"spawn-time"`
}

var enemies = []EnemyData{
	EnemyData{
		uid:       1,
		tribe:     0,
		position:  [3]float32{0, 0, 1},
		spawnTime: 12,
	},
	EnemyData{
		uid:       2,
		tribe:     0,
		position:  [3]float32{0, 0, 1},
		spawnTime: 12,
	},
	EnemyData{
		uid:       3,
		tribe:     1,
		position:  [3]float32{0, 0, 1},
		spawnTime: 12,
	},
}

// ReadEnemies   GET "/enemy"
func ReadEnemies(ctx *gin.Context) {
	ctx.JSON(200, generateMap(enemies))
}

// KillEnemy   POST "/enemy/kill/:id"
func KillEnemy(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}
	enemies = rejectMap(id, enemies)
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
func rejectMap(id int, s []EnemyData) []EnemyData {
	ans := make([]EnemyData, 0)
	for _, x := range s {
		if x.uid != id {
			ans = append(ans, x)
		}
	}
	return ans
}
