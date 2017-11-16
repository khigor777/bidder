# bidder



0.Install Redis

1.cd ./cmd

2.edit config: config.json

3.go run main.go

4.http://localhost:8080

```GO
r.POST("/", handlers.IncrementIfaStat) - Ifa increment 600 sec
r.GET("/stats", handlers.GetStat) - Get http network stat
r.GET("/stats/sort", handlers.GetSortedStat) - Get http network stat sorted
```
  
 
