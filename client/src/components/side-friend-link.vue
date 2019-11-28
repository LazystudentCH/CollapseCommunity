<template>
    <div>
        <el-table :header-cell-style="{background:'#f5f5f5'}" :data="tableData" class="friend-link" style="width: 100%">
            <el-table-column prop="address" label="站内公告" width="" >

            </el-table-column>
        </el-table>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                tableData: [{
                    address: '严禁发布含有广告色情信息',
                }, {
                    address: '严禁使用不雅词汇',
                }, {
                    address: '严禁传播恐怖或恶心内容',
                }, {
                    address: '严禁发表地域歧视等言论'
                }]
            };
        },
    }
</script>

<style scoped>
    .friend-link {
        border: solid 1px #dbdbdb;
        border-radius: 10px;
    }

</style>