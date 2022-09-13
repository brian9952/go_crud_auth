import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import ProductManagement from "../views/ProductManagement.vue"
import axios from "axios";

import "../../node_modules/primeflex/primeflex.css";

axios.defaults.baseURL = 'http://107.102.183.168:5073/v1/api';
// headers
axios.defaults.headers.common['Authorization'] = 'woe';
axios.defaults.headers.post['Content-Type'] = 'application/json';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: ProductManagement,
    },
    {
      path: "/about",
      name: "about",
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/AboutView.vue"),
    },
  ],
});

export default router;
