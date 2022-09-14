<script setup>
import Button from "/node_modules/primevue/button";
</script>

<template>
  <div class="flex flex-column">
    <div class="flex align-items-center justify-content-center font-roboto text-blue-100 pt-4">
      Welcome {{ getUsername }}!
    </div>
  </div>
  <div class="card">
    <div class="flex pt-3 text-blue-100 align-items-center justify-content-center text-xl font-semibold font-roboto">
      Product Management
    </div>
  </div>

  <div class="flex flex-row-reverse flex-wrap card-container pr-6">
    <Button class="p-button-primary" label="Add" icon="pi pi-plus" />
  </div>
</template>

<script>
  import { computed } from 'vue'
  export default {
    computed: {
      getUsername() {
        return this.$store.state.username
      }
    }
  }
</script>

<style>
.font-roboto {
    font-family: 'Roboto', sans-serif;
  }
</style>

