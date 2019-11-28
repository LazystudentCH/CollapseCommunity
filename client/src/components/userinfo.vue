<template>
    <div class="from">
        <el-form :label-position="labelPosition" label-width="80px" :model="user">
            <el-form-item label="头像">
                <el-upload
                        class="avatar-uploader"
                        :show-file-list="false"
                        action="/api/uploadAvatar"
                        :headers="headers"
                        name="avatar"
                        :on-success="handleAvatarSuccess"
                        :before-upload="beforeAvatarUpload">
                    <img v-if="users.Avatar" :src="users.Avatar" class="avatar">
                    <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                </el-upload>
            </el-form-item>
            <el-form-item label="名称">
                <el-input v-model="users.name"></el-input>
            </el-form-item>
            <el-form-item label="电话号码">
                <el-input disabled v-model="users.phone"></el-input>
            </el-form-item>
            <el-form-item label="性别">
                <el-radio-group v-model="users.gender">
                    <el-radio label="0">男</el-radio>
                    <el-radio label="1">女</el-radio>
                </el-radio-group>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" size="medium" @click="modifyUserInfo">修改</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>
<script>
    import axios from "../route/intercept"

    export default {
        data() {
            return {
                labelPosition: 'right',
                phone:"",
                avatar:"",
                users:{}
            };
        },
        methods: {
            modifyUserInfo(){
                axios.post("/api/user",this.users).then((response)=>{
                    if(response.data.Code == 0) {
                        this.$message({
                            message: "修改个人信息成功",
                            type: 'success',
                            center: true
                        });
                        var token = response.data.Data;
                        localStorage.setItem('token', token);
                        location.reload()
                    }
                }).catch(()=>{
                    this.$message({
                        message: "出错了，请稍后再试",
                        type: 'error',
                        center: true
                    });
                })
            },
            handleAvatarSuccess(response) {
                this.users.Avatar = response.Data
            },
            beforeAvatarUpload(file) {
                const isJPG = file.type === 'image/jpeg';
                const isLt2M = file.size / 1024 / 1024 < 2;

                if (!isJPG) {
                    this.$message.error('上传头像图片只能是 JPEG 格式!');
                    return false
                }
                if (!isLt2M) {
                    this.$message.error('上传头像图片大小不能超过 2MB!');
                    return false
                }

            },
        },

        computed:{
            headers(){
                return {
                    'token':localStorage.getItem("token")
                }
            }
        },

        mounted() {
            if (localStorage.getItem('token')) {
                (()=>{
                    var token =  localStorage.getItem('token');
                    axios.get("/api/parsejwt",{
                        params:{
                            'token':token,
                        }
                    }).then((res)=>{
                        this.phone = res.data.Data.phone;
                        axios.get("/api/user",{params:{
                                "phone":this.phone
                            }}).then((response)=>{
                            if (response.data.Code != 0){
                                this.$message({
                                    message: "网络繁忙，请稍后再试",
                                    type: 'error',
                                    center: true
                                });
                            }else{
                                this.users = response.data.Data;
                            }
                        }).catch((error)=>{
                            this.$message({
                                message: error,
                                type: 'error',
                                center: true
                            });
                        })
                    }).catch((error)=>{
                        console.log(error)
                    })
                })()

            }


        }
    }
</script>

<style scoped>
    .avatar-uploader .el-upload {
        border: 1px dashed #d9d9d9;
        border-radius: 6px;
        cursor: pointer;
        position: relative;
        overflow: hidden;
    }
    .avatar-uploader .el-upload:hover {
        border-color: #409EFF;
    }
    .avatar-uploader-icon {
        font-size: 28px;
        color: #8c939d;
        width: 178px;
        height: 178px;
        line-height: 178px;
        text-align: center;
    }
    .avatar {
        width: 178px;
        height: 178px;
        display: block;
    }

    .from{
        width: 400px;
    }
</style>
<style>
    input::-webkit-outer-spin-button,input::-webkit-inner-spin-button{
        -webkit-appearance:textfield;
    }
    input[type="number"]{
        -moz-appearance:textfield;
    }
</style>

