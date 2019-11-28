<template>
    <div>
        <ul class="card-list">
            <li class="card" v-for="item in list" :key="item.ID">
                <div class="card-info">
                    <div class="title"><router-link :to="'/subjectDetail/'+item.ID">{{item.title}}</router-link></div>
                    <div class="metadata">
                        <span class="author">{{user.name}}</span>
                        <time>{{item.CreatedAt|formatDate}}</time>

                    </div>
                </div>
                <span class="option">
                            <el-button type="warning" @click="editSubject(item)" size="mini">编辑</el-button>
                            <el-button type="danger" size="mini" @click="deleteSub(item.ID)">删除</el-button>
                </span>
            </li>
        </ul>
        <el-pagination @current-change="changePage" class="page" background layout="prev, pager, next" :total="totalPage"></el-pagination>
    </div>
</template>
<script>
    import axios from "../route/intercept"
    import { formatDate } from '../utils/formatTime.js'
    import VueEvent from "../utils/VueEvent";
    export default {
        data() {
            return {
                totalPage:1000,
                list:[],
                user:{}
            };
        },
        methods: {
            deleteSub(id){
                axios.get("/api/deleteSub",{params:{
                        "sub_id":id
                    }}).then((response)=>{
                        if(response.data.Code == 0 ){
                            this.$alert('删除成功', {
                                confirmButtonText: '确定',
                                callback:() => {
                                    location.reload()
                                }
                            });
                        }else{
                            this.$message({
                                message: "删除失败",
                                type: 'error',
                                center: true
                            });
                        }
                }).catch((error)=>{
                    console.log(error)
                })
            },
            changePage(page){
                axios.get("/api/user/subject",{params:{
                        "page":page
                    }}).then((response)=>{
                    this.list = response.data.Data;
                })
            },
            getTotalPage(){
                axios.get("/api/user/getSubjectCount").then((response)=>{
                    this.totalPage = response.data.Data * 1;
                })
            },
            editSubject(item){
                VueEvent.$emit("subjectObj",item);
            }
        },
        mounted() {
            this.changePage(1);
            this.getTotalPage();
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
    ul{
        margin-left: auto;
        border-radius:10px;
        border: solid 1px #dbdbdb;
        list-style:none;
    }
    .card{
        overflow: hidden;
    }
    ul li {
        border-bottom: solid 1px #dbdbdb;
        margin-top: 5px;
        margin-right: 5%;
    }

    .card-info,.block{
        float: left;
        margin-top: 15px;
        margin-bottom: 15px;
    }

    .card-info{
        text-align: left;
        margin-left: 10px;
    }

    .option{
        float: right;
        margin-top: 2%;
    }

    .title{
        margin-top: 0;
        font-weight: bold;
        font-size: 16.8px;
    }

    .metadata{
        font-size: 14px;
        color: #969696;
    }

    .page{
        float: right;
    }

    .author{
        margin-right: 20px;
    }

    a{
        text-decoration:none;
        color: inherit;
    }
</style>

