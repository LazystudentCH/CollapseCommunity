<template>
    <div>
        <v-header></v-header>

        <div class="from">
            <div class="title">
                登录
            </div>
            <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="0" class="demo-ruleForm">
                <el-form-item  prop="phone">
                    <el-input class="phone-input" type='number' prefix-icon="el-icon-phone-outline" v-model="ruleForm.phone"></el-input>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input prefix-icon="el-icon-lock" type="password" v-model="ruleForm.password"
                              autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="success" style="width: 100%" @click="submitForm('ruleForm')">登录</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>
<script>
    import header from "./header.vue"
    import Axios from "axios"

    export default {
        data() {
            return {
                ruleForm: {
                    phone: '',
                    password: '',
                },
                rules: {
                    phone: [
                        {required: true, message: '请输入手机号码', trigger: 'blur'},
                        {min: 11, max: 11, message: '手机号码只能为11位', trigger: 'blur'}
                    ],
                    password: [
                        {required: true, message: '请输入密码', trigger: 'blur'},
                        {min: 6, message: '最短6个字符', trigger: 'blur'}
                    ],
                }
            };
        },
        methods: {
            submitForm(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        Axios.post("/api/login", this.ruleForm).then((response) => {
                            if (response.data.Code != 0) {
                                this.$message({
                                    message: response.data.Msg,
                                    type: 'error',
                                    center: true
                                });
                            } else {
                                //登录成功
                                var token = response.data.Data;
                                localStorage.setItem('token', token);
                                // 跳转到 / 路由
                                location.replace('/')
                            }
                        }).catch((error) => {
                            this.$message({
                                message: "出错了" + error,
                                type: 'error',
                                center: true
                            });
                        })
                    } else {
                        return false;
                    }
                });
            },
        },
        components: {
            "v-header": header
        }
    }
</script>

<style scoped>

    .from{
        width:20%;margin:10% auto;
        padding: 10px;
        border: 1px solid #dbdbdb;
        border-radius: 6px;
    }

    .title{
        color: #7a7a7a;
        font-size: 2rem;
        font-weight: 600;
        line-height: 1.125;
        padding-left: 40%;
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