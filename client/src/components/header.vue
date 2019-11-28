<template>
    <div class="header">
        <el-menu class="el-menu-demo" mode="horizontal">
            <el-menu-item @click="jump">
                <img class="headImage" :src='imgurl' alt="">
            </el-menu-item>
            <el-menu-item @click="jump">主题</el-menu-item>
            <el-menu-item @click="toHome">官网</el-menu-item>

            <el-menu-item v-show="isHome">
                <el-input placeholder="请输入内容" v-model="keyWord" class="input-with-select">
                    <el-button slot="append" icon="el-icon-search" @click="searchSubject(1)"></el-button>
                </el-input>
            </el-menu-item>

            <el-menu-item v-if="!isLogin" class="login-btn">
                <el-button type="primary" class="login-btn2" @click="login">登录</el-button>
                <el-button @click="register">注册</el-button>
            </el-menu-item>
            <el-menu-item v-else class="login-btn">
                <el-dropdown @command="handleCommand">
                      <span class="el-dropdown-link">
                         <el-avatar  :src="user.Avatar"></el-avatar>
                      </span>
                    <el-button type="primary" round size="mini" @click="addSub" style="color: white;margin-left: 7%">发布主题</el-button>
                    <el-dropdown-menu slot="dropdown">
                        <el-dropdown-item command="usercenter">个人中心</el-dropdown-item>
                        <el-dropdown-item command="logout">退出登录</el-dropdown-item>
                    </el-dropdown-menu>
                </el-dropdown>
            </el-menu-item>
        </el-menu>
        <div class="line"></div>
    </div>
</template>

<script>
    import imgPath from "../assets/logo.png"
    import axios from "../route/intercept"
    import VueEvent from "../utils/VueEvent";
    export default {
        data() {
            return {
                imgurl: imgPath,
                keyWord: "",
                isLogin: false,
                user: {},
                isHome:false,
            };
        },
        methods: {
            login() {
                this.$router.push({path: '/login'})
            },
            toHome(){
                window.location.href = 'http://www.mihayo.com/';
            },
            register() {
                this.$router.push({path: '/register'})
            },
            handleCommand(command) {
                switch (command) {
                    case "usercenter":
                        this.$router.push('/userCenter');
                        break;
                    case "logout":
                        localStorage.removeItem('token');
                        this.$router.push('/login');
                        break
                }
            },
            jump(){
                if(this.$route.path == '/'){
                    location.reload()
                }else{
                    this.$router.push("/")
                }
            },
            addSub(){
                this.$router.push("/userCenter")
            },
            jumpToMemberList(){
                this.$router.push("/memberList")
            },
            searchSubject(page){
                axios.get("/api/search",{
                    params:{
                        "keyWord":this.keyWord,
                        "page":page
                    }
                }).then((response)=>{
                    VueEvent.$emit("searchResult",response.data.Data);
                    VueEvent.$emit("keyWord",this.keyWord)
                }).catch((error)=>{
                    console.log(error)
                })
            }
        },
        mounted() {
            if (localStorage.getItem('token')) {
                var token =  localStorage.getItem('token');
                axios.get("/api/parsejwt",{
                    params:{
                        'token':token,
                    }
                }).then((res)=>{
                    this.user = res.data.Data;
                    this.isLogin = true;
                }).catch((error)=>{
                    console.log(error)
                })
            }
            if (this.$route.path == '/'){
                this.isHome = true
            }
        }

    }
</script>

<style scoped>
    img {
        height: 40px;
    }

    .search-input {
        float: left;
    }

    .el-menu-demo {
        margin: auto 10% auto 10%;
        border-bottom: none;
    }

    .header {
        border-bottom: 1px solid #dddddd;
    }

    .login-btn {
        float: right;
    }

</style>