<template>
    <div>
        <el-card class="box-card">
            <el-input v-model="title" placeholder="请输入标题"></el-input>
        </el-card>
        <div style="margin-top: 2%"></div>
        <el-card class="box-card" style="height: 500px">
            <v-editor ref="editors"></v-editor>
            <el-button class="submit" type="success" @click="addSubject" round>提交</el-button>
        </el-card>
    </div>

</template>

<script>
    import editor from "./editor"
    import axios from "../route/intercept"
    import VueEvent from "../utils/VueEvent";
    export default {
        data(){
            return {
                title:"",
                content:"",
                users:{},
                subid:0,
            }
        },
        components:{
            "v-editor":editor
        },
        methods:{
            addSubject(){
                this.content = this.$refs.editors.detailContent;
                axios.post("/api/subject",{
                    "title":this.title,
                    "content":this.content,
                    "user_id":this.users.ID,
                    "id":this.subid
                }).then((response)=>{
                    if(response.data.Code == 0) {
                        this.$alert('发布成功', {
                            confirmButtonText: '确定',
                            callback:() => {
                                location.reload()
                            }
                        });
                    }
                }).catch((error)=>{
                    this.$message({
                        message: "发布失败",
                        type: 'error',
                        center: true
                    });
                    console.log(error)
                })
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
                    this.users = res.data.Data;
                }).catch((error)=>{
                    console.log(error)
                })
            }
            VueEvent.$emit("sign","1");
        },
        created(){
            this.$bus.$on('editSub', (subObj) => {
                this.title = subObj.title;
                this.$refs.editors.detailContent = subObj.content;
                this.subid = subObj.ID
            });
        }

    }
</script>

<style scoped>
    .submit{
        margin-top: 8%;
        margin-left: 93%;
    }
</style>
