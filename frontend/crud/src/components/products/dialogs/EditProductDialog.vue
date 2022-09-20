<script setup>
import Dialog from "/node_modules/primevue/dialog"
import Button from "/node_modules/primevue/button"

import axios from "axios"
</script>

<template>
  <Dialog v-model:visible="display" v-bind:modal="true">
  </Dialog>
</template>
