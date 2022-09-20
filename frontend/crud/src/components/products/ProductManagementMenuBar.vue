<script setup>
import Menubar from "/node_modules/primevue/menubar";
import Button from "/node_modules/primevue/button";
import Skeleton from "/node_modules/primevue/skeleton";
</script>

<template>
  <Menubar>
    <template #start>
      <Button label="Product Management" icon="pi pi-list" class="p-button-text" />
    </template>
  
    <template #end>
      <div class="w-12rem" v-show="skeleton_show">
        <Skeleton height="1.5rem" />
      </div>
      <div v-show="button_show">
        <div v-if="authShow" id="auth">
          <!-- login button -->
          <Button class="mx-2 p-button-text" label="Login" @click="openLoginDialog()" />
          <!-- register button -->
          <Button class="mx-2 p-button-text" label="Register" @click="openRegisterDialog()" />
        </div>
        <div v-if="welcomeShow" id="loggedin">
          <Button class="mx-2 p-button-text">Welcome {{ this.$store.getters.getUsername }}</Button>
          <Button class="mx-2 p-button-text" @click="logout()">Logout</Button>
        </div>
        <!--
        <Button v-if="authShow" class="mx-2 p-button-text" v-for="item of auth_items" :label="item.label" @click="toggleVisible()" />
        <Button v-if="welcomeShow" class="mx-2 p-button-text" v-for="item of loggedin_items" :label="item.label"></Button>
        -->
      </div>
    </template>
  </Menubar>

  <LoginDialog :display="loginVisible" @hide="loginVisible = false"></LoginDialog>
  <RegisterDialog :display="registerVisible" @hide="registerVisible = false"></RegisterDialog>
</template>
<script>
import axios from 'axios';
import LoginDialog from "../auth/LoginDialog.vue";
import RegisterDialog from "../auth/RegisterDialog.vue";

export default {
    components: {
      LoginDialog,
      RegisterDialog
    },
    emits: [
      'interface'
    ],
    data() {
      return {
        loginVisible: false,
        registerVisible: false,
        skeleton_show: true,
        button_show: false,
        username: '',
        auth_items: [
          { label: 'Login' },
          { label: 'Register' },
        ],
        loggedin_items: [
          { label: 'Welcome ' + this.username },
          { label: 'Logout' },
        ],
      }
    },
    computed: {
      authShow() {
        if (!this.$store.getters.authenticated) {
          return true
        }
        return false
      },
      welcomeShow() {
        if (this.$store.getters.authenticated) {
          return true
        }
        return false
      },
    },
    mounted() {
      this.emitInterface()
    },
    methods: {
      openLoginDialog() {
        if(this.loginVisible == false) {
          this.loginVisible = true
          return
        }
        this.loginVisible = false
      },
      openRegisterDialog() {
        if(this.registerVisible == false) {
          this.registerVisible = true
          return 
        }
        this.registerVisible = false
      },
      getUsername() {
        return this.$store.getters.getUsername
      },
      showRes() {
        this.username = this.$store.getters.getUsername
        this.button_show = true
        this.skeleton_show = false
      },
      emitInterface() {
        this.$emit("interface", {
          showRes: () => this.showRes()
        });
      },
      logout() {
        // change header
        localStorage.setItem("app_token", '')
        // change state
        this.$store.commit('toggleAuthenticated', 0)
        this.$store.commit('changeUsername', 'guest')
      }
    }
}
</script>
<style>

</style>
