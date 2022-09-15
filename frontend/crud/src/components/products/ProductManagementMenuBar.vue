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
      <Button v-show="login_register_show" class="mx-2 p-button-text" v-for="item of items" :label="item.label" @click="toggleVisible()" />
      <Button v-show="welcome_show" class="mx-2 p-button-text" label="Welcome {{ getUsername }}"></Button>
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
        welcome_show: false,
        login_register_show: true,
        isLoggedIn: false,
        username: '',
        items: [
          { label: 'Login' },
          { label: 'Register' }
        ]
      }
    },
    mounted() {
      if (this.$store.state.isAuthenticated == true) {
        this.welcome_show = true
        this.login_register_show = false
        this.$store.state.username = 
      } else {
        this.welcome_show = false
        this.login_register_show = true
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
