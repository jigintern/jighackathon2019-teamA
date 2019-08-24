# API仕様書

## 初めに以下を行ってください
`goの導入`
`go get github.com/gin-gonic/gin`

## ステージデータ(全ての敵のデータ)を取得する
### GET "/api/v1/enemy/"

以下のjsonでレスポンスが送られます

```JSON
[
	{
		"position": [
			出現位置のx座標(float32),
			出現位置のy座標(float32),
			出現位置のz座標(float32),
		],
        "spawn_time": ゲーム開始後何秒後に出現するか,
        "tribe": 敵の種類ID,
        "uid": 敵のユニークID
	}
]
```

ex)

```JSON
[
    {
        "position": [
            0,
            0,
            1
        ],
        "spawn_time": 12,
        "tribe": 0,
        "uid": 1
    },
    {
        "position": [
            0,
            0,
            1
        ],
        "spawn_time": 12,
        "tribe": 0,
        "uid": 2
    },
    {
        "position": [
            0,
            0,
            1
        ],
        "spawn_time": 12,
        "tribe": 1,
        "uid": 3
    }
]
```



## 敵を倒したことにする
### GET "/api/v1/enemy/:uid/kill"

ユニークid(uid)が:uidである敵を削除し、倒したことにします
ex) "/api/v1/enemy/1/kill" -> uid: 1の敵を倒したことにする
