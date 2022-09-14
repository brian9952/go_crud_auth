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
      isAuthenticated: false
    }
  },
  mutations: {
    toggleAuthenticated(state) {
      if (state.isAuthenticated == false) {
        state.isAuthenticated = true
        return
      }
      state.isAuthenticated = false
    },
    changeUsername(state, username) {
      state.username = username
    }
  }
});

app.use(VueCookies);
app.use(PrimeVue);
app.use(router);
app.use(store);

app.mount("#app");
