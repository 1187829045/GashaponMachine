# GashaponMachine
用于练手的一个扭蛋机系统

## 扭蛋机

1. 消耗5砖石玩一轮；
2. 商品：（五个等级）；
3. 每10轮必中商品；
4. 没60轮一个循环；第60轮安排一个最高价值商品；
5. 可以用内存来替代Redis;

监听端口号 8080  

POST   /gash/v1/login

POST   /gash/v1/lottery

GET    /gash/v1/query  
