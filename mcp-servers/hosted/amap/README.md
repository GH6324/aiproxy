# AMAP MCP Server

https://lbs.amap.com/api/mcp-server/gettingstarted

产品介绍
在AI时代，随着AI技术的迅猛发展，各种出行助手应用如雨后春笋，受限于大模型的数据孤岛、能力边界限制，始终未能发挥其在应用层面落地价值。 MCP的出现统一了大模型与外部数据、工具间的通讯协议。而在出行服务领域，数据的时效性、工具的便捷性尤为重要，高德MCP Server旨在为大模型在出行领域的应用落地高效赋能。

2025年3月，高德地图MCP首发，为开发者提供了基于位置服务、地点信息搜索、路径规划、天气查询等12大核心高鲜度数据，让用户在出行规划、位置信息检索场景下轻松获取即时信息。

2025年5月，高德地图MCP全新升级，通过高德MCP Server 与高德地图APP无缝打通，用户可将大模型产出的攻略与高德地图APP无缝衔接。实现一键生成专属地图，将攻略中的点位、描述、行程规划等个性化信息自动导入到高德地图APP，生成一张独属于用户的私有地图，实际出行中可实现由攻略到一键导航、打车、 订票的丝滑体验。 

真正让高德贯穿你的行前-行中-行后始终，让每个人都能轻松拥有一个“真正懂你的出行秘书”。

能力介绍
生成专属地图
将出行规划导入高德地图，生成专属地图

输入

行程名称、行程详情（每日行程描述、行程途径点位）

输出

专属地图唤端链接

导航到目的地
根据用户传入经纬度，启动导航

输入

目的地经纬度

输出

高德导航唤端链接

打车
根据用户输入起终经纬度坐标，发起打车请求

输入

origin (起点经纬度)，destination (终点经纬度)

输出

高德打车唤端链接

地理编码
将详细的结构化地址转换为经纬度坐标。

输入

address (位置信息)，city (城市信息，非必须)

输出

location (位置经纬度)

逆地理编码
将一个高德经纬度坐标转换为行政区划地址信息。

输入

location (位置经纬度)

输出

addressComponent (位置信息，包括省市区等信息)


IP 定位
IP 定位根据用户输入的 IP 地址，定位 IP 的所在位置。

输入

IP

输出

province (省)，city (城市)，adcode (城市编码)

天气查询
根据城市名称或者标准adcode查询指定城市的天气。

输入

city (城市名称或城市adcode)

输出

forecasts (预报天气)

骑行路径规划
用于规划骑行通勤方案，规划时会考虑天桥、单行线、封路等情况。最大支持 500km 的骑行路线规划。

输入

origin (起点经纬度)，destination (终点经纬度)

输出

distance (规划距离)，duration (规划时间)，steps (规划步骤信息)

步行路径规划
可以根据输入起点终点经纬度坐标，规划100km 以内的步行通勤方案，并且返回通勤方案的数据。

输入

origin (起点经纬度)，destination (终点经纬度)

输出

origin (起点信息)，destination (终点信息)，paths (规划具体信息)

驾车路径规划
根据用户起终点经纬度坐标规划以小客车、轿车通勤出行的方案，并且返回通勤方案的数据。

输入

origin (起点经纬度)，destination (终点经纬度)

输出

origin (起点信息)，destination (终点信息)，paths (规划具体信息)

公交路径规划
根据用户起终点经纬度坐标规划综合各类公共（火车、公交、地铁）交通方式的通勤方案，并且返回通勤方案的数据，跨城场景下必须传起点城市与终点城市。

输入

origin (起点经纬度)，destination (终点经纬度)，city (起点城市)，cityd (终点城市)

输出

origin (起点信息)，destination (终点信息)，distance (规划距离)，transits (规划具体信息)

距离测量
测量两个经纬度坐标之间的距离。

输入

origin (起点经纬度)，destination (终点经纬度)

输出

origin_id (起点信息)，dest_id (终点信息)，distance (规划距离)，duration (时间)

关键词搜索
根据用户传入关键词，搜索出相关的POI地点信息。

输入

keywords (搜索关键词)，city (查询城市，非必须)

输出

suggestion (搜索建议)，pois (地点信息列表)


周边搜索
根据用户传入关键词以及坐标location，搜索出radius半径范围的POI地点信息。

输入

keywords (搜索关键词)，location (中心点经度纬度)，radius (搜索半径，非必须)

输出

pois (地点信息列表)


详情搜索
查询关键词搜或者周边搜获取到的POI ID的详细信息。

输入

id (关键词搜或周边搜获取的poiid)

输出

地点详情信息

location (地点经纬度)，address (地址)，business_area (商圈)，city(城市)，type (地点类型) 等