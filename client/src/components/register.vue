<template>
    <div>
        <v-header></v-header>
        <div class="from">
            <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="80px" class="demo-ruleForm">
                <el-form-item label="昵称" prop="name">
                    <el-input prefix-icon="el-icon-user" small v-model="ruleForm.name"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="password">
                    <el-input prefix-icon="el-icon-lock" type="password" v-model="ruleForm.password"
                              autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item label="手机号码" prop="phone">
                    <el-input class="phone-input" type='number' prefix-icon="el-icon-phone-outline" v-model="ruleForm.phone"></el-input>
                </el-form-item>
                <el-form-item label="验证码" prop="verifyCode">
                    <el-col :span="14">
                        <el-input prefix-icon="el-icon-s-promotion" v-model="ruleForm.verifyCode"></el-input>
                    </el-col>
                    <el-button class="send-msg" :class="{disabled: !this.canClick}" :disabled="!canClick" type="primary" @click="sendSmsCode(ruleForm.phone,'ruleForm')">{{content}}</el-button>
                </el-form-item>
                <el-form-item label="性别" prop="gender">
                    <el-radio-group v-model="ruleForm.gender">
                        <el-radio label="0">汉子</el-radio>
                        <el-radio label="1">妹子</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="submitForm('ruleForm')">立即注册</el-button>
                    <el-button @click="resetForm('ruleForm')">重置</el-button>
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
                totalTime: 60,
                canClick: true,
                content:"发送验证码",
                ruleForm: {
                    name: '',
                    gender: '',
                    phone: '',
                    password: '',
                    verifyCode: ''
                },
                rules: {
                    name: [
                        {required: true, message: '请输入昵称', trigger: 'blur'},
                        {min: 3, message: '最短3个字符', trigger: 'blur'}
                    ],
                    phone: [
                        {required: true, message: '请输入手机号码', trigger: 'blur'},
                        {min: 11, max: 11, message: '手机号码只能为11位', trigger: 'blur'}
                    ],
                    password: [
                        {required: true, message: '请输入密码', trigger: 'blur'},
                        {min: 6, message: '最短6个字符', trigger: 'blur'}
                    ],
                    verifyCode: [
                        {required: true, message: '请输入验证码', trigger: 'blur'},
                        {min: 4, max: 4, message: '验证码只能为4位', trigger: 'blur'}
                    ],
                    gender: [
                        {required: true, message: '请选择性别', trigger: 'change'}
                    ],
                }
            };
        },
        methods: {
            submitForm(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        Axios.post("/api/register", this.ruleForm).then((response) => {
                            if (response.data.Code == 2) {
                                this.$message({
                                    message: response.data.Msg,
                                    type: 'error',
                                    center: true
                                });
                            } else {
                                this.$message({
                                    message: "注册成功,即将跳转",
                                    type: 'success',
                                    center: true
                                });
                                this.$router.push({path:'/login'})
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
            resetForm(formName) {
                this.$refs[formName].resetFields();
            },
            sendSmsCode(phone) {
                if (phone.length == 11){
                    Axios.post("/api/sms", this.ruleForm.phone).then((response) => {
                        console.log(response.data.Msg)
                    }).catch((error) => {
                        console.log(error)
                    });
                    this.countDown()
                }else{
                    this.$message({
                        message: "手机号码有误",
                        type: 'error',
                        center: true
                    });
                }
            },
            countDown () {
                if (!this.canClick) return;
                this.canClick = false;
                this.content = this.totalTime + 's后重新发送';
                let clock = window.setInterval(() => {
                    this.totalTime--;
                    this.content = this.totalTime + 's后重新发送';
                    if (this.totalTime < 0) {
                        window.clearInterval(clock);
                        this.content = '重新发送验证码';
                        this.totalTime = 60;
                        this.canClick = true
                    }
                },1000)
            }
        },
        components: {
            "v-header": header
        }
    }
</script>

<style scoped>
    .from {
        margin-top: 5%;
        margin-left: 35%;
        margin-right: 40%;
        padding: 20px;
        border: 1px solid #dbdbdb;
        border-radius: 6px;
    }
    .send-msg {
        float: right;
    }

    .disabled{
        background-color: #ddd;
        border-color: #ddd;
        color:#57a3f3;
        cursor: not-allowed;
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