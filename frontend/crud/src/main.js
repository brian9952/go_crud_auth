import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import PrimeVue from "/node_modules/primevue/config"
import VueCookies from "vue-cookies";

import "../node_modules/primeflex/primeflex.scss";
import "../node_modules/primeicons/primeicons.css";
import "../node_modules/primevue/resources/themes/luna-blue/theme.css"
import "./assets/fonts.css";

const app = createApp(App);

app.use(VueCookies);
app.use(PrimeVue);
app.use(router);

app.mount("#app");
