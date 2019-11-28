<template>
    <div>
        <el-table :header-cell-style="{background:'#f5f5f5'}" :data="tableData" class="friend-link" style="width: 100%">
            <el-table-column prop="key" label="社区状态" width="" >
                <template slot-scope="scope">
                    <el-button size="mini" class="key-btn">{{scope.row.key}}</el-button>
                </template>
            </el-table-column>
            <el-table-column prop="value" label="" width="" >
                <template slot-scope="scope">
                    <el-button size="mini" class="value-btn">{{scope.row.value}}</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
    import axios from "../route/intercept"
    export default {
        data() {
            return {
                tableData: [{
                    key: '会员',
                    value:""
                }, {
                    key: '主题',
                    value:""
                },{
                    key: '回复',
                    value:""
                }
                ]
            };
        },
        mounted() {
            axios.get("/api/count").then((response)=>{
                var res = response.data.Data;
                this.tableData[0].value = res.MemberCount;
                this.tableData[1].value = res.SubjectCount;
                this.tableData[2].value = res.CommentCount;
            })
        }
    }
</script>

<style scoped>
    .friend-link {
        border: solid 1px #dbdbdb;
        border-radius: 10px;
    }

    .key-btn{
        background-color: #f5f5f5;
        border: none;
    }

    .value-btn{
        background-color: #00d1b2;
        border: none;
    }

</style>