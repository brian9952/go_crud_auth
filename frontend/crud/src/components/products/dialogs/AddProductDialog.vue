<script>
import Dialog from "/node_modules/primevue/dialog"
import Button from "/node_modules/primevue/button"
</script>

<template>
  <Dialog v-model:visible="display" v-bind:model="true">
    <div class="flex flex-column px-8 pb-4">
      <h1 class="font-bold">Edit Product</h2>
      <h2>Product Name</h2><br>
      {{ product_name }}
      <h2>Product Value</h2><br>

    </div>
  </Dialog>
</template>

<script>
  export default {

  }
</script>
