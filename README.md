# 调用方法：

### 定义调用nacos的常量，例如命名空间id、DataId、Group


var (


	NamespaceId = "dcc3ab24-6834-4cd1-bd57-2c24d61ffe5f"
	
	DataId      = "ak"
	Group       = "ops"  
)


### 创建结构体来反序列化其中的配置文件，nacos中的配置文件为json格式


type AKConfig struct {


	AKID string `mapstructure:"akid" json:"akid"`
	AKSK string `mapstructure:"aksk" json:"aksk"`
	
	
}



var ak AKConfig



func AK() {  
	app := utils.InitNacos(NamespaceId, DataId, Group)  
	json.Unmarshal([]byte(app), &ak)  
}  
