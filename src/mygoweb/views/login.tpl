<html>
	<head>
		<title>登录</title>
		<link rel="stylesheet" href="static/css/font-awesome.css" />
	    <link rel="stylesheet" href="static/css/element.css" />
	    <link rel="stylesheet" href="static/css/base.css" />
	    <script src="static/js/vue.js"></script>
	    <script src="static/js/vue-resource.min.js"></script>
	    <script src="static/js/element.js"></script>
	    <script src="static/js/store.modern.min.js"></script>
	    <script src="static/js/base.js"></script>
	</head>
	<body>
		<div id="app">
			<div style="width: 350px; height: 320px; float:right; position:relative; top:70px;">
		 	<el-tabs type="border-card" style="height: 100%">
	            <el-tab-pane label="账户登录(ajax)">
					
	                <el-input size="large" v-model="username" placeholder="用户名" :value="username"
	                          style="margin: 20px auto;">
	                    <template slot="prepend">
	                        <i class="fa fa-user"></i>
	                    </template>
	                </el-input>
	                <el-input size="large" v-model="password" placeholder="密码" type="password" :value="password">
	                    <template slot="prepend">
	                        <i class="fa fa-lock"></i>
	                    </template>
	                </el-input>
	
	                <el-button @click="tologin" size="large" type="danger" style="margin-top:30px; width: 100%">
	                    登 录(ajax)
	                </el-button>
					
	            </el-tab-pane>
				<el-tab-pane label="账户登录(submit)">
					<form action="/home" method="post">
		                <input placeholder="用户名" id="LogonName" name="LogonName" class="el-input__inner"
		                          style="margin: 20px auto;">
		                </input>
		                <input type="password" id="PassWord" name="PassWord" class="el-input__inner">
		                </input>
		
						<el-button native-type="submit" size="large" type="danger" style="margin-top:30px; width: 100%">
		                    登 录（submit）
		                </el-button>
					</form>
	            </el-tab-pane>
			</el-tabs>
			</div>
		</div>
	</body>
	<script type="text/javascript">
    Vue.http.options.emulateJSON = true;
    var vo = new Vue({
        el: "#app",
        data: function () {
            return {
                username: "changxb",
                password: "123456"
            }
        },
        methods: {
            tologin: function () {
				//前端传过去的是JSON，后端用map接（json解析到map里） 
				//前端传过去的是{"LogonName":"changxb","PassWord":"123456"}
				//Vue.http.post('/home',JSON.stringify({ LogonName: this.username, PassWord: this.password })).then(function (res) {
				
				//前端传过去的是LogonName=changxb&PassWord=123456
				// 后端直接用controller对象的GetString("LogonName")取
				Vue.http.post('/home',{ LogonName: this.username, PassWord: this.password }).then(function (res) {
                    console.log(res)
                });
            },


            loginok: function (data, status, xhr) {
                var token = store.get("token");
                if (token) {
                    //store.set('token', xhr.getResponseHeader('token'));
                    window.location = "../mainview/main?token=" + token;

                }
            }
        }
    });

    Hp.createNew(vo);
	</script>
</html>