import { createApp } from "vue";
import { createStore } from "vuex";
import App from "./App.vue";
import router from "./router";
import PrimeVue from "/node_modules/primevue/config"
import VueCookies from "vue-cookies";

import "../node_modules/primeflex/primeflex.scss";
import "../node_modules/primeicons/primeicons.css";
import "../node_modules/primevue/resources/themes/luna-blue/theme.css"
import "./assets/fonts.css";

const app = createApp(App);
app.config.unwrapInjectedRef = true;

// create store instance
const store = createStore({
  state() {
    return {
      username: 'guest',
      isAuthenticated: false,
      newData: {
        product_name: '',
        product_value: 0,
        product_description: ''
      }
    }
  },
  getters: {
    authenticated: state => {
      return state.isAuthenticated
    },
    getUsername: state => {
      return state.username
    },
    getData: state => {
      return state.newData
    }
  },
  mutations: {
    toggleAuthenticated(state, arg) {
      state.isAuthenticated = arg
    },
    changeUsername(state, username) {
      state.username = username
    },
    setNewData(state, data) {
      state.newData = data
    }
  }
});

app.use(VueCookies);
app.use(PrimeVue);
app.use(router);
app.use(store);

app.mount("#app");
