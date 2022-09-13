<script setup>
import Dialog from "/node_modules/primevue/dialog"
import axios from "axios"
</script>

<template>
  <Dialog v-model:visible="display" v-bind:modal="true">
    <div class="flex flex-column">
      <h2 class="font-bold">Product Name</h2>
      <p>{{ product_name }}</p>
      <h2 class="font-bold">Price</h2>
      <p>{{ product_value }}</p>
      <h2 class="font-bold">Description</h2>
      <p>{{ product_description }}</p>
    </div>
  </Dialog>
</template>

<script>
  export default {
    props: [
      'display'
    ],
    data() {
      return {
        product_name: '',
        product_value: '',
        product_description: ''
      }
    },
    mounted() {
      // emits on mount
      this.emitInterface();
    },
    methods: {
      fetchProduct(productId) {
        axios.get("http://107.102.183.168:8081/v1/api/product/show/" + productId)
          .then(resp => {
            this.product_name = resp.data.product_name
            this.product_value = resp.data.product_value
            this.product_description = resp.data.product_description
          })
          .catch(function(error) {
            console.log(error.toJSON());
          })
      },
      emitInterface() {
        this.$emit("interface", {
          fetchProduct: (productId) => this.fetchProduct(productId)
        });
      }
    }
  }
</script>
