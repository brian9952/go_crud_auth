import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";

import "../node_modules/primeflex/primeflex.scss";
import "../node_modules/primeicons/primeicons.css";
import "./assets/saga-blue.css";

const app = createApp(App);

app.use(router);

app.mount("#app");
