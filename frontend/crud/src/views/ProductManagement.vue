<script setup>
import ProductManagementMenuBar from "../components/products/ProductManagementMenuBar.vue"
import ProductManagementTopContent from "../components/products/ProductManagementTopContent.vue"
import LoginDialog from "../components/auth/LoginDialog.vue"
import ProductManagementMainContent from "../components/products/ProductManagementMainContent.vue"

import axios from 'axios'
</script>

<template>
  <ProductManagementMenuBar @interface="getChildInterface"/>
  <ProductManagementTopContent />
  <ProductManagementMainContent />
</template>

<script>
export default {
    //data() {
    //  return {
    //    isCalculated: 0
    //  }
    //},
    components: {
      ProductManagementMenuBar
    },
    childInterface: {
      showRes: () => {}
    },
    created() {
      const fetchData = async() => {
        this.authenticate()
      }
      fetchData();
    },
    methods: {
      authenticate() {
        // get local storage 
        if(localStorage.getItem('app_token') == null) {
          axios.defaults.headers.common['Authorization'] = "";
        }else {
          axios.defaults.headers.common['Authorization'] = localStorage.getItem('app_token')
        }

        // get url
        let url = import.meta.env.VITE_BACKEND_URL

        // fetch data
        axios.get(url + "/v1/api/auth/refresh_token")
          .then(resp => {
            if(resp.data.status_type == 0) {
              this.$store.state.isAuthenticated = true
              //this.$store.commit('toggleAuthenticated', true)
              this.$store.commit('changeUsername', resp.data.username)
              //console.log(this.$store.state.isAuthenticated)
              this.$options.childInterface.showRes()
              return
            }
            this.$options.childInterface.showRes()
          })
          .catch(function(error) {
            console.log(error)
          })
      },
      getChildInterface(childInterface) {
        this.$options.childInterface = childInterface;
      }
    }
}
</script>

<style>
.body {

  }
</style>
