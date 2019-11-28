<template>
    <div class="discuss-area">
        <link rel="stylesheet" href="http://at.alicdn.com/t/font_711214_93mm591lbta.css">
        <ul class="card-list">
            <li class="card" v-for="(item,index) in list" :key="item.ID">
                <div class="block">
                    <el-avatar class="avatar" :src="item.User.Avatar"></el-avatar>
                </div>
                <div class="card-info">
                    <div class="title">
                        <span>
                            <router-link :to="'/subjectDetail/'+item.ID">{{item.title}}</router-link>
                            <el-tag v-show="item.is_hot" size="mini" style="background-color: #ff6547;margin-left: 20px;color: #EBEEF5">
                                <i class="icon iconfont icon-huoyanjiare"></i>
                                烫
                            </el-tag>
                        </span>
                    </div>
                    <div class="metadata">
                        <span class="author">{{item.User.name}}</span>
                        <time>{{item.CreatedAt|formatDate}}</time>
                    </div>
                </div>
                <div>
                    <span class="discuss-count">
                        <i class="el-icon-chat-dot-round"></i>
                        {{commentCountList[index]}}
                    </span>
                    <span class="discuss-count">
                        <i class="icon iconfont icon-thumbup"></i>
                        {{priseList[index]}}|
                    </span>
                </div>
            </li>
        </ul>
        <el-pagination v-if="!this.searched" @current-change="changePage" class="page" background layout="prev, pager, next" :total="totalPage"></el-pagination>
        <el-pagination v-else @current-change="nextSearchRes" class="page" background layout="prev, pager, next" :total="searchTotalPage"></el-pagination>
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
                searched:false,
                commentCountList:[],
                priseList:[],
                keyWord:"",
                searchTotalPage:1000,
            };
        },
        methods:{
           changePage(page){
               axios.get("/api/subject",{params:{
                   "page":page
                   }}).then((response)=>{
                       this.list = response.data.Data;
                       this.getCommentCount(response.data.Data);
                       this.getPriseCount(response.data.Data)
           }).catch((error)=>{
                   console.log(error)
               })
           },
           getTotalPage(){
               axios.get("/api/indexTotalPage").then((response)=>{
                   this.totalPage = response.data.Data * 1
               }).catch((error)=>{
                   console.log(error)
               })
           },
            getSearchTotalPage(){
                axios.get("/api/searchTotalPage",{
                    params:{
                        "keyWord":this.keyWord
                    }
                }).then((response)=>{
                    this.searchTotalPage = response.data.Data * 1
                }).catch((error)=>{
                    console.log(error)
                })
            },
           getCommentCount(dataList){
               var temList = new Array();
               for(let i = 0;i<dataList.length;i++){
                   axios.get("/api/commentCount", {
                       params: {
                           "sub_id": this.list[i].ID
                       }
                   }).then((response) => {
                       this.$set(temList,i,response.data.Data)
                   }).catch((error)=>{
                       console.log(error)
                   })
               }
               this.commentCountList = temList
            },
            getPriseCount(dataList){
                var temList = new Array(dataList.length);
                for(let i = 0;i<dataList.length;i++){

                        axios.get("/api/queryPriseCount", {
                            params: {
                                "sub_id": this.list[i].ID
                            }
                        }).then((response) => {
                            this.$set(temList,i,response.data.Data)
                        }).catch((error)=>{
                            console.log(error)
                        })
                }
                this.priseList = temList;
            },
            nextSearchRes(page){
                axios.get("/api/search",{
                    params:{
                        "keyWord":this.keyWord,
                        "page":page
                    }
                }).then((response)=>{
                    this.list = response.data.Data
                }).catch((error)=>{
                    console.log(error)
                })
            }
        },
        mounted() {
            this.changePage(1);
            this.getTotalPage();
            VueEvent.$on("searchResult",(res)=>{
                this.list = res;
                this.searched = true;
            });
            VueEvent.$on("keyWord",(res)=>{
                this.keyWord = res;
                this.getSearchTotalPage();
            });

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
    .icon-thumbup{

    }
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

    .title{
        margin-top: 0;
        font-weight: bold;
        font-size: 16.8px;
        font-family: "Microsoft YaHei";
    }

    .metadata{
        font-size: 14px;
        color: #969696;
    }
    .discuss-count{
        float: right;
        margin-top: 50px;
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