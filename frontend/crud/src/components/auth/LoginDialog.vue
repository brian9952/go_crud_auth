<script setup>
import Dialog from "/node_modules/primevue/dialog";
import InputText from "/node_modules/primevue/inputtext";
import Button from "/node_modules/primevue/button";
import ProgressSpinner from "/node_modules/primevue/progressspinner";

import axios from 'axios';
</script>

<template>
  <Dialog v-model:visible="display" v-bind:modal="true">
    <div class="flex flex-column mx-8">
      <h2 class="flex">Log In</h2>
      <!-- forms -->
      <label class="flex mb-2" for="Username">Username</label>
      <InputText type="text" v-model="username" class="flex w-20rem" />
      <div class="mb-4">
        <p v-show="username_error_show">{{ username_error }}</p>
      </div>

      <label class="flex mb-2" for="Password">Password</label>
      <InputText type="password" v-model="password" class="flex w-20rem" />
      <div class="mb-5">
        <p v-show="password_error_show">{{ password_error }}</p>
      </div>

      <!-- buttons -->
      <div class="flex gap-4 mb-6">
        <Button class="flex-grow-1" :icon="spinner_computed" :label="submit_label" @click="onSubmit" />
      </div>

    </div>
  </Dialog>
</template>

<script>
export default {
  props: ['display'],
  data() {
    return {
      isLoading: 0,
      username_error: "",
      password_error: "",
      username_error_show: false,
      password_error_show: false,
      username: '',
      password: '',

      spinner_computed: '',
      submit_label: 'Submit'
    }
  },
  methods: {
    checkIntegrity(){
      if(this.username == '') {
        this.username_error = "Username field is empty!";
        this.username_error_show = true;
      }else {
        this.username_error = '';
      }

      if(this.password == '') {
        this.password_error = "Password field is empty!";
        this.password_error_show = true;
      }else {
        this.password_error = '';
      }

      if(this.username_error == '' && this.password_error == '') {
        this.username_error_show = false;
        this.password_error_show = false;
        return 0;
      }
      return 1;
    },
    convertBase64(){
      var rawStr = this.username + ':' + this.password + ':' + "user"
      return btoa(rawStr)
    },
    postInput(dataStr) {
      // change login to loading icon
      this.spinner_computed = "pi pi-spin pi-spinner";

      // get backend url
      let url = import.meta.env.VITE_BACKEND_URL

      // fetch api
      axios.post(url + "/v1/api/auth/login",{
          data: dataStr
        })
        .then(resp => {
          if(!this.processResponse(resp.data)){
            this.changeButton()
            this.insertUserData(resp.data)
          }
        })
        .catch(function(err) {
          console.log(err)
        });
    },
    changeButton() {
      this.spinner_computed = '';
      this.submit_label = 'Login Success';
    },
    insertUserData(data) {
      // insert cookies
      this.$cookies.set('isLoggedIn', 'true', { expires: '10m' })
      this.$cookies.set('username', data["username"], { expires: '10m' })

      // insert authorization header
      let bearerStr = 'Bearer ' + data["access_token"] + ';' + data["refresh_token"]
      localStorage.setItem("app_token", bearerStr)
      axios.defaults.headers.common['Authorization'] = bearerStr


      // change state
      this.$store.commit('toggleAuthenticated')
      this.$store.commit('changeUsername', data["username"])

      // close dialog
      this.$emit('hide')
    },
    processResponse(data) {
      if(data["status_type"] == 1) { // username error
        this.username_error = "Username is incorrect!";
        this.username_error_show = true;
        this.spinner_computed = '';
        return 1;
      }else if(data["status_type"] == 2) { // password incorrect
        this.password_error = "Password is incorrect!";
        this.password_error_show = true;
        this.spinner_computed = '';
        return 1;
      }else if(data["status_type"] == 3){ // internal error
        console.log("Internal error")
        this.spinner_computed = '';
        return 1;
      }else if(data["status_type"] == 4){ // input error
        console.log("User input error")
        this.spinner_computed = '';
        return 1;
      }
      return 0;
    },
    onSubmit() {
      console.log("Clicked");
      console.log("HERE")
      this.$store.commit('changeUsername', 'WOOEY')

      // check user input
      var integrityStatus = this.checkIntegrity();
      if(integrityStatus != 0) {
        return;
      }
      console.log("HERE")

      // convert to base64
      var b64Str = this.convertBase64();

      // fetch data
      this.postInput(b64Str);

    },
  }
}

</script>
