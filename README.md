# go-studying


DI:
cd cmd
~/go/bin/wire

ls
openapi:
cd api
~/go/bin/swag init -g ./handler/center_handler.go -o ./docs -parseDependency
http://localhost:8080/swagger/index.html#/default/

[**readmeの書き方**](https://style.potepan.com/articles/33682.html)


振り返し：  
1. 基本環境環境構築
   1. フォルダ構成を決める  
   2. modelsにドメインを設計する  
   3. repoにデータベース層  
   4. serviceにサービス層  
   5. apiにapi層  
       1. handler
       2. router  
   6. appにginの追加  
2. wireの使用  
3. openapiの導入
4. karateの導入 
5. middlewareの導入
    
