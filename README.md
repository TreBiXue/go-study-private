# go-studying

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
    [**wire**](https://farer.org/2021/04/21/go-dependency-injection-wire/)  
    1. providerset作成(appフォルダに)  
    2. wire build設定(cmdフォルダに)  
    3. wire　実行  
    例：  
    ```cd cmd```  
    ```~/go/bin/wire```  
    4. wire_gen.go自動生成

3. openapiの導入  
    [**swaggerドキュメント**](https://qiita.com/KNR109/items/7e094dba6bcf37ed73cf)  
    1. 各APIにドキュメントを記入（@Descriptionなど）  
    2. 標準ドキュメント生成  
       例：  
       ```cd api```  
       ```~/go/bin/swag init -g ./handler/center_handler.go -o ./docs -parseDependency```
    3. ドキュメント引用  
       例：  
       server.goに追加  
       ```s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))```　　
    4. Openapi確認  
    　　``` http://localhost:8080/swagger/index.html#/default/ ```
4. karateの導入 
5. middlewareの導入
    
