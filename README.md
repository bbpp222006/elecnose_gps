# docker_elecnose_gps

车载电子鼻的后端

使用方式
```
docker run  -d --name gps -p 18081:8081 varitia/elecnose_gps:latest
```

## 测试路径

### 获取数据

curl -X POST -i 'http://127.0.0.1:18081/test/get_data' --data '{"begin":1,"end":10,"interval":1}'

返回的都是随机值，返回三个地方坐标,分别是光电大楼，北京某高尔夫球场，深圳的老三和人才市场

{114.440429,30.513163},{116.448569,40.012752},{114.029224,22.666757}

### 传数据

curl -X POST -i 'http://127.0.0.1:18081/test/push_data' --data '{"board_name":"board1","E":114.440429,"N":30.513163,"data":[22,31,23,19,1,2]}'

返回当前时间戳
