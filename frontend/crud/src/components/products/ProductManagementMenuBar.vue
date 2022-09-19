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
      <div class="w-8rem" v-show="skeleton_show">
        <Skeleton height="1.5rem" />
      </div>
      <div v-show="button_show">
        <Button v-if="authShow" class="mx-2 p-button-text" v-for="item of items" :label="item.label" @click="toggleVisible()" />
        <Button v-if="welcomeShow" class="mx-2 p-button-text">Welcome {{ this.$store.getters.getUsername }}</Button>
      </div>
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
    emits: [
      'interface'
    ],
    data() {
      return {
        isVisible: false,
        isLoggedIn: false,
        skeleton_show: true,
        button_show: false,
        username: '',
        items: [
          { label: 'Login' },
          { label: 'Register' }
        ]
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
      }
    },
    mounted() {
      this.emitInterface()
      //this.button_show = true
      //this.a = this.authShow
      //this.b = this.welcomeShow
      //console.log(this.$store.state.isAuthenticated)
      //if (this.$store.state.isAuthenticated == true) {
      //  console.log("WES AUTHENTICATED COK")
      //  this.welcome_show = true
      //  this.login_register_show = false
      //  this.$store.state.username = this.username
      //  //console.log(this.login_register_show)
      //}else{
      //  this.welcome_show = false
      //  this.login_register_show = true
      //}
    },
    methods: {
      toggleVisible() {
        if(this.isVisible == false) {
          this.isVisible = true
          return
        }
        this.isVisible = false
      },
      showRes() {
        this.button_show = true
        this.skeleton_show = false
      },
      emitInterface() {
        this.$emit("interface", {
          showRes: () => this.showRes()
        });
      }
    }
}
</script>
<style>

</style>
