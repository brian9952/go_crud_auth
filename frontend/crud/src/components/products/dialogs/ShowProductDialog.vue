<script setup>
import Dialog from "/node_modules/primevue/dialog"
import Button from "/node_modules/primevue/button"
import axios from "axios"
</script>

<template>
  <Dialog v-model:visible="display" v-bind:modal="true">
    <div class="flex flex-column px-8 pb-4">
      <h2 class="font-bold">Product Name</h2><br>
      {{ product_name }}
      <h2 class="font-bold">Price</h2><br>
      {{ product_value }}
      <h2 class="font-bold">Description</h2><br>
      {{ product_description }}
    </div>
    <div class="flex flex-column align-items-center">
      <Button @click="$emit('closeDialog')">Close</Button>
    </div>
  </Dialog>
</template>

<script>
  export default {
    props: [
      'display'
    ],
    emits: [
      'closeDialog',
      'interface'
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
        let url = import.meta.env.VITE_BACKEND_URL
        axios.get(url + "/v1/api/product/show/" + productId)
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
