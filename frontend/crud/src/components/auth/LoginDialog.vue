<script setup>
import Dialog from "/node_modules/primevue/dialog";
import InputText from "/node_modules/primevue/inputtext";
import Button from "/node_modules/primevue/button";
import ProgressSpinner from "/node_modules/primevue/progressspinner";
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
        <Button class="flex-grow-1" label="Submit" @click="onSubmit()" />
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
      password: ''
    }
  },
  methods: {
    checkIntegrity(username, password){
      if(username == '') {
        this.username_error = "Username field is empty!";
        this.username_error_show = true;
      }

      if(password == '') {
        this.password_error = "Password field is empty!";
        this.password_error_show = false;
      }

      if(this.username_error == '' && this.password_error == '') {
        return 0;
      }
      return 1;
    },
    convertBase64(username, password){
      var rawStr = username + ':' + password + ':' + "user"
      console.log(rawStr);
      return btoa(rawStr)
    },
    postInput(dataStr) {
      const requestOptions = {
        method: "POST",
        mode: 'cors',
        credentials: 'same-origin',
        cache: 'no-cache',
        headers: { 
          "Content-Type": "application/json"
          },
        body: JSON.stringify({ data: dataStr })
      };
      fetch("http://107.102.183.168:8081/v1/api/auth/login", requestOptions)
        .then(async response => {
          const data = await response.json();

          if(!response.ok) {
            const error = (data && data.message) || response.statusText;
            return Promise.reject(error);
          }

          return data;
        });
    },
    onSubmit() {
      console.log("Submitted!");

      // check user input
      let integrityStatus = this.checkIntegrity(this.username, this.password);
      if(integrityStatus != 0) {
        return;
      }

      // convert to base64
      let b64Str = this.convertBase64(this.username, this.password);
      console.log(b64Str);

      // fetch data
      var response = this.postInput(b64Str);

      // process data
      console.log(response);

    },
  }
}

//export default {
//  data() {
//    return {
//      display: true
//    }
//  }
//}
</script>
