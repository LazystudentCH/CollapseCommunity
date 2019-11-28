<template>
    <div>
        <link rel="stylesheet" href="http://at.alicdn.com/t/font_711214_93mm591lbta.css">
        <v-header></v-header>
        <div style="margin-top: 2%"></div>
        <div class="subject-body">
            <el-card class="box-card">
                <div slot="header" class="clearfix">
                    <span>{{subject.title}}</span>
                </div>
                <div class="author-info">
                    <div class="block">
                        <el-avatar class="avatar" :src="author.Avatar"></el-avatar>
                    </div>
                    <div class="metadata">
                        <span class="author">{{author.name}}</span>
                        <div v-if="!hasPrise" class="prise">
                            <el-button size="mini" class="primary" @click="Prise">
                                <i class="icon iconfont icon-thumbup"></i>
                            </el-button>
                        </div>
                        <div v-else class="prise">
                            <el-button size="mini" class="primary" @click="Prise">
                                <i class="icon iconfont icon-dianzan1"></i>
                            </el-button>
                        </div>
                    </div>
                    <div class="tag">
                        <el-tag size="small" type="info">发布</el-tag>
                        <el-tag size="small" color="#00d1b2" style="margin-right: 2%">{{subject.CreatedAt|formatDate}}
                        </el-tag>

                        <el-tag size="small" type="info">浏览</el-tag>
                        <el-tag size="small" color="#00d1b2" style="margin-right: 2%">{{totalReview}}</el-tag>

                        <el-tag size="small" type="info">回复</el-tag>
                        <el-tag size="small" color="#00d1b2" style="margin-right: 2%">{{totalComment}}</el-tag>
                    </div>

                </div>
                <div style="margin-top: 4%"></div>
                <div v-html="subject.content" class="text item"></div>

            </el-card>

            <div style="margin-top: 2%"></div>
            <el-card class="box-card">
                <div slot="header">
                    <span>共{{totalComment}}个回复</span>
                </div>
                <div v-for="item in commentList" :key="item.ID">
                    <div class="block">
                        <el-avatar class="avatar" :src="item.User.Avatar"></el-avatar>
                    </div>
                    <div class="metadata">
                        <span class="author">{{item.User.name}}</span>
                        <time style="margin-left: 2%">{{item.CreatedAt|formatDate}}</time>
                    </div>
                    <div v-html="item.content" class="text item"></div>
                    <hr>
                </div>
                <el-pagination v-if="totalComment != 0" @current-change="changePage" class="page" background
                               layout="prev, pager, next"
                               :total="totalPage"></el-pagination>
            </el-card>

            <div style="margin-top: 2%"></div>
            <hr>
            <div class="box-comment">
                <div v-show="isLogin" class="avatar">
                    <el-avatar :src="user.Avatar"></el-avatar>
                </div>
                <el-card v-if="!isLogin" class="tips-login">
                    <el-button class="tips-btn" type="primary" round @click="toLogin">登录</el-button>
                    后发表评论
                </el-card>
                <div class="box-editor" v-else>
                    <v-editor ref="editors"></v-editor>
                    <el-button class="submit" type="success" @click="submitComment" round>提交</el-button>
                </div>

            </div>

        </div>

    </div>
</template>

<script>
    import {formatDate} from '../utils/formatTime.js'
    import header from "../components/header"
    import axios from "../route/intercept"
    import editor from "../components/editor"

    export default {
        data() {
            return {
                subject: {},
                author: {},
                user: {},
                isLogin: false,
                commentList: [],
                content: "",
                totalPage: 1,
                totalComment: 1,
                totalReview: "",
                hasPrise: false,
            };
        },
        components: {
            "v-header": header,
            "v-editor": editor
        },
        methods: {
            submitComment() {
                this.content = this.$refs.editors.detailContent;
                if (this.content == "") {
                    this.$message({
                        message: "评论不能为空",
                        type: 'error',
                        center: true
                    });
                } else {
                    axios.post("/api/comment", {
                        "user_id": this.user.ID,
                        "sub_id": this.subject.ID,
                        "content": this.content
                    }).then(() => {
                        this.$message({
                            message: "评论成功",
                            type: 'success',
                            center: true
                        });
                        location.reload()
                    })
                }
            },
            Prise() {
                var sid = this.$route.params;
                axios.post("/api/prise", {
                    "sub_id": sid.sid * 1,
                    "has_prise": this.hasPrise
                }).then((response) => {
                    if (response.data.Code == 0) {
                        if (this.hasPrise == false) {
                            this.$message({
                                message: "点赞成功",
                                type: 'success',
                                center: true
                            });
                        } else {
                            this.$message({
                                message: "取消点赞成功",
                                type: 'success',
                                center: true
                            });
                        }
                        this.hasPrise = !this.hasPrise
                    } else {
                        this.$message({
                            message: "操作失败,请先登录",
                            type: 'error',
                            center: true
                        });
                    }
                }).catch((error) => {
                    console.log(error)
                })
            },
            changePage(page) {
                var sid = this.$route.params;
                axios.get("/api/comment", {
                    params: {
                        "page": page,
                        "sub_id": sid.sid
                    }
                }).then((response) => {
                    this.commentList = response.data.Data
                }).catch((error) => {
                    console.log(error)
                })
            },
            getTotalComment() {
                var sid = this.$route.params;
                axios.get("/api/commentCount", {
                    params: {
                        "sub_id": sid.sid
                    }
                }).then((response) => {
                    this.totalComment = response.data.Data * 1
                }).catch((error) => {
                    console.log(error)
                })
            },
            getTotalPage() {
                var sid = this.$route.params;
                axios.get("/api/commentCount", {
                    params: {
                        "sub_id": sid.sid
                    }
                }).then((response) => {
                    this.totalPage = response.data.Data * 1
                }).catch((error) => {
                    console.log(error)
                })
            },
            toLogin() {
                this.$router.push("/login")
            },
            queryHasPrise(){
                var sid = this.$route.params;
                if(this.isLogin == true){
                    axios.post("/api/queryHasPrise", {
                        "sub_id": sid.sid*1
                    }).then((response) => {
                        this.hasPrise = response.data.Data
                    })
                }
            },
            queryReview(){
                var sid = this.$route.params;
                axios.get("/api/getSubjectReview", {
                    params: {
                        "sub_id": sid.sid
                    }
                }).then((response) => {
                    this.totalReview = response.data.Data
                })
            },
            getDetail(){
                var sid = this.$route.params;

                //获取详情
                axios.get("/api/subject", {
                    params: {
                        "sub_id": sid.sid
                    }
                }).then((response) => {
                    if (response.data.Code != 0) {
                        this.$message({
                            message: "获取详情失败，请稍后再试",
                            type: 'error',
                            center: true
                        });
                    } else {
                        this.subject = response.data.Data
                        this.author = response.data.Data.User
                    }
                }).catch((error) => {
                    console.log(error)
                });

            }
        },
        mounted(){
            if (localStorage.getItem('token')) {
                var token =  localStorage.getItem('token');
                axios.get("/api/parsejwt",{
                    params:{
                        'token':token,
                    }
                }).then((res)=>{
                    this.user = res.data.Data;
                    this.isLogin = true;
                    this.getDetail();
                    //展示评论
                    this.changePage(1);
                    this.getTotalComment();
                    this.getTotalPage();
                    //获取浏览量
                    this.queryReview();
                    //查询是否已经点赞
                    this.queryHasPrise()
                }).catch((error)=>{
                    console.log(error)
                })
            }

        },
        filters: {
            formatDate(time) {
                var date = new Date(time);
                return formatDate(date, 'yyyy-MM-dd hh:mm');
            }
        },

    }
</script>

<style scoped>
    .metadata {
        margin-left: 5%;
        font-size: 14px;
        color: #969696;
    }

    .tag {
        margin-top: 1%;
        margin-left: 5%;
        font-size: 14px;
    }

    .author {
        font-size: 16px;
    }

    .block {
        margin-top: 15px;
        margin-bottom: 15px;
    }

    .item {
        margin-left: 5%;
    }

    .subject-body {
        margin-left: 10%;
        margin-right: 30%;
    }

    .box-card {
        border-radius: 10px;
    }

    .clearfix {
        font-weight: bold;
        font-size: 20px;
        font-family: "Microsoft YaHei";
    }

    .prise {
        float: right;

    }

    .avatar {
        float: left;
    }

    .tips-login {
        width: 100%;
        float: right;
        background-color: #f5f5f5;
    }

    .tips-btn {
        margin-left: 400px;
    }

    .box-editor {
        width: 95%;
        float: right;
    }

    .box-comment {
        height: 400px;
    }

    .submit {
        margin-top: 8%;
        margin-left: 93%;
    }


    hr {
        border: none;
        height: 1px;
        background-color: #EBEEF5;
    }
</style>

<style>
    hr {
        border: none;
        height: 1px;
        background-color: #f5f5f5;
    }
</style>