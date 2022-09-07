# GoBasicAPI
Go的基礎API Web Server

## 說明
* 從簡單到進階以方便學習。
* 接收來自80port的呼叫，意味著你可以使用瀏覽器訪問。
* 使用`go run ./version`來測試不同版本的GoBasicAPI

### version
最基礎的版本，監聽80port並回傳一段文字  
在Web中測試: http://localhost/hi  

### version2
在這個版本中以GET的方式接收job和method方法來執行不同的API  
在Web中測試: http://localhost/?job=business&method=one  

### version3
在這個版本中每個API不局限在main.go裡面，可以歸類於不同的資料夾中  
在Web中測試: http://localhost/?job=business&method=one  

### version4
在這個版本中不再透過job和method的方式，改由路徑直接對應函式
在Web中測試: http://localhost/?job=business&method=one  
