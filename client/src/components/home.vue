<template>
    <div>
        <el-container>
            <el-header>
                <v-header></v-header>
            </el-header>
            <el-container>

                <el-container>
                    <el-main class="discuss">
                        <v-discuss></v-discuss>
                    </el-main>
                </el-container>
                <el-aside class="side-bar" width="200px">
                    <v-friend-link class="friend-link"></v-friend-link>
                    <v-community-status></v-community-status>
                </el-aside>
            </el-container>
        </el-container>
    </div>

</template>

<script>
    import header from "./header"
    import discuss from "./discuss-area.vue"
    import friendLink from "./side-friend-link"
    import communityStatus from "./community-status.vue"

    export default {
        data() {
            return {};
        },
        components: {
            'v-header': header,
            'v-discuss': discuss,
            'v-friend-link':friendLink,
            'v-community-status':communityStatus,
        },
    }
</script>

<style scoped>
    .discuss{
        margin-left: 10%;
        margin-right: 5%;
    }
    .side-bar{
        margin-top: 2%;
        margin-right: 11%;
    }
    .friend-link{
        padding-bottom: 20%;
        border: none;
    }
</style>