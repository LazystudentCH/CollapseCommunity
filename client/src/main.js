import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import './plugins/element.js'
import router from "./route/route";
import VueQuillEditor from 'vue-quill-editor'
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'

Vue.use(ElementUI);
Vue.config.productionTip = false;
Vue.prototype.$bus = new Vue;
Vue.use(VueQuillEditor, {
  placeholder: '请输入内容',
});

new Vue({
    router,
    render: h => h(App),
}).$mount('#app');
