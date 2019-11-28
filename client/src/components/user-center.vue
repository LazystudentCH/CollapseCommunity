<template>
    <div>
        <v-header></v-header>
        <el-container>
            <el-aside class="side-bar" width="400px">
                <el-row class="tac">
                    <el-col :span="12">
                        <el-menu :default-active="active" @select="selectTag" class="el-menu-vertical-demo">
                            <el-menu-item index="1">
                                <i class="el-icon-menu"></i>
                                <span slot="title">个人信息</span>
                            </el-menu-item>
                            <el-menu-item index="2">
                                <i class="el-icon-menu"></i>
                                <span slot="title">我发布的主题</span>
                            </el-menu-item>
                            <el-menu-item index="3">
                                <i class="el-icon-menu"></i>
                                <span slot="title">发布新主题</span>
                            </el-menu-item>
                        </el-menu>
                    </el-col>
                </el-row>
            </el-aside>
            <div class="user-info">
                <div v-if="index === '1'">
                    <v-userinfo></v-userinfo>
                </div>

                <div v-else-if="index === '2'">
                    <v-user-subjects></v-user-subjects>
                </div>

                <div v-else-if="index === '3'">
                    <v-addSubject></v-addSubject>
                </div>

                <div v-else>

                </div>
            </div>
        </el-container>
    </div>
</template>

<script>
    import header from "../components/header"
    import userinfo from "./userinfo"
    import subjects from "./all-user-subject"
    import addSubject from "./addSubject"
    import VueEvent from "../utils/VueEvent";
    import axios from "../route/intercept"

    export default {
        data() {
            return {
                index: "1",
                user:{},
                active:"1",
                subObj:{}
            };
        },
        components: {
            "v-header": header,
            "v-userinfo": userinfo,
            "v-user-subjects":subjects,
            "v-addSubject":addSubject,
        },
        methods: {
            selectTag(index) {
                this.index = index
            },
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
                }).catch((error)=>{
                    console.log(error)
                })
            }
            VueEvent.$on("subjectObj",(res)=>{
                this.subObj = res;
                this.index = "3";
                this.active = "3";
            });
            VueEvent.$on("sign",()=>{
                this.$bus.$emit('editSub', this.subObj);
            });


        },

    }
</script>

<style scoped>
    .user-info{
        margin-top: 20px;
        margin-right: 300px;
        width: 100%;
    }

</style>