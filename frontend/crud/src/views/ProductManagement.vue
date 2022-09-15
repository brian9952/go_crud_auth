<script setup>
import ProductManagementMenuBar from "../components/products/ProductManagementMenuBar.vue"
import ProductManagementTopContent from "../components/products/ProductManagementTopContent.vue"
import LoginDialog from "../components/auth/LoginDialog.vue"
import ProductManagementMainContent from "../components/products/ProductManagementMainContent.vue"

import axios from 'axios'
</script>

<template>
  <ProductManagementMenuBar :isAuthenticated="isAuthenticated"/>
  <ProductManagementTopContent />
  <ProductManagementMainContent />
</template>

<script>
export default {
    created() {
      this.$watch (
        () => this.$route.params,
        () => {
          this.authenticate();
        },

        { immediate: true }
      )
    },
    methods: {
      authenticate() {
        // get local storage 
        axios.defaults.headers.common['Authorization'] = localStorage.getItem('app_token')

        // get url
        let url = import.meta.env.VITE_BACKEND_URL

        // fetch data
        axios.get(url + "/v1/api/auth/refresh_token")
          .then(resp => {
            if(resp.data["status_type"] == 0) {
              this.$store.commit('toggleAuthenticated')
              this.$store.commit('changeUsername', data["username"])
              return
            }
          })
          .catch(function(error) {
            console.log(error.toJSON())
          })
      }
    }
}
</script>

<style>
.body {

  }
</style>
