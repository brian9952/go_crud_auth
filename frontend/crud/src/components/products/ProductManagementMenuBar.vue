<script setup>
import Menubar from "/node_modules/primevue/menubar";
import Button from "/node_modules/primevue/button";
</script>

<template>
  <Menubar>
    <template #start>
      <Button label="Product Management" icon="pi pi-list" class="p-button-text" />
    </template>
  
    <template #end>
      <Button class="mx-2 p-button-text" v-for="item of items" :label="item.label" @click="toggleVisible()" />
    </template>
  </Menubar>

  <LoginDialog :display="isVisible" @hide="isVisible = false"></LoginDialog>
</template>

<script>
import LoginDialog from "../auth/LoginDialog.vue";

export default {
    components: {
        LoginDialog
    },
    data() {
      return {
        isVisible: false,
        items: [
          { label: 'Login' },
          { label: 'Register' }
        ]
      }
    },
    methods: {
      toggleVisible() {
        if(this.isVisible == false) {
          this.isVisible = true
          return
        }
        this.isVisible = false
      }
    }
}
</script>
<style>

</style>
