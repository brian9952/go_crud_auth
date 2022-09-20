<script setup>
import Button from "/node_modules/primevue/button";
</script>

<template>
  <div class="card">
    <div class="flex pt-5 text-blue-100 align-items-center justify-content-center text-xl font-semibold font-roboto">
      Product Management
    </div>
  </div>

  <div class="flex flex-row-reverse flex-wrap card-container pr-6">
    <Button class="p-button-primary" label="Add" icon="pi pi-plus" @click="addProduct()"/>
  </div>

  <AddProductDialog />
</template>

<script>
import AddProductDialog from "../products/dialogs/AddProductDialog.vue"

export default {
  data() {
    return {

    }
  }
}
</script>

<style>
.font-roboto {
    font-family: 'Roboto', sans-serif;
  }
</style>

