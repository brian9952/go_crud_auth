import { createRouter, createWebHistory } from 'vue-router'
import ProductsManagement from '../views/ProductsManagement.vue'
import UserLogin from '../views/UserLogin.vue'
import PageNotFound from '../views/PageNotFound.vue'

const routes = [
  {
    path: '/',
    name: 'products',
    component: ProductsManagement
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  },
  {
    path: '/user',
    name: 'user',
    component: UserLogin
  },
  {
    path: "/:pathMatch(.*)*",
    name: '404',
    component: PageNotFound
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
