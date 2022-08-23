import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import PrimeVue from "/node_modules/primevue/config"
import VueCookie from "/node_modules/vue-cookie";

import "../node_modules/primeflex/primeflex.scss";
import "../node_modules/primeicons/primeicons.css";
import "../node_modules/primevue/resources/themes/luna-blue/theme.css"
import "./assets/fonts.css";

const app = createApp(App);

app.use(PrimeVue);
app.use(VueCookie);
app.use(router);

app.mount("#app");
