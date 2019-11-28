import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter);
import home from '../components/home'
import register from "../components/register";
import login from "../components/login"
import subject from "../components/subject-detail"
import userCenter from "../components/user-center"


const routes = [
    {path:'/',component:home},
    {path:'/login',component:login},
    {path:'/register',component:register},
    {path:'/subjectDetail/:sid',component:subject},
    {path:'/userCenter',component:userCenter},
    {path:'*',component:home}
];

const router = new VueRouter({
    mode: 'history',
    routes
});
router.beforeEach((to, from, next) => {
    if (to.path === '/login' || to.path === '/register' || to.path === '/'|| to.path.startsWith('/subjectDetail/')) {
        next();
    } else {
        let token = localStorage.getItem('token');
        if (token === null || token === '') {
            next('/login');
        } else {
            next();
        }
    }
});
export default router;