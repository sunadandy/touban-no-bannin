import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import axios from 'axios'
import VueAxios from 'vue-axios'
import '@mdi/font/css/materialdesignicons.css'  // vuetifyでアイコンを表示するためのモジュール
import store from './store'

createApp(App)
  .use(router)
  .use(vuetify)
  .use(store)
  .use(VueAxios, axios)
  .mount('#app')
