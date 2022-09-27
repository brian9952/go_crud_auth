<script setup>
import Dialog from "/node_modules/primevue/dialog"
import Button from "/node_modules/primevue/button"
import InputText from "/node_modules/primevue/inputtext"
</script>

<template>
  <Dialog v-model:visible="display" v-bind:modal="true">
    <div class="flex flex-column mx-8">
      <h2 class="flex">Add Product</h2>
      <!-- message -->
      <div v-if="message_show" class="flex card-container p-2 mt-2 mb-4 border-round" :class="message_color">
        {{ message }}
      </div>

      <!-- forms -->
      <label class="flex mb-2" for="product_name">Product Name</label>
      <InputText type="text" v-model="product_name" class="flex w-20rem mb-4" :class="product_name_class" />

      <label class="flex mb-2" for="product_value">Product Value</label>
      <InputText type="number" v-model="product_value" class="flex w-20rem mb-4" :class="product_value_class" />

      <label class="flex mb-2" for="product_description">Product Description</label>
      <InputText type="text" v-model="product_description" class="flex w-20rem mb-5" :class="product_description_class" />

      <!-- buttons -->
      <div class="flex justify-content-center flex-wrap gap-4 mb-5">
        <Button class="flex align-items-center p-button-primary" label="Submit" @click="onSubmit()"/>
        <Button class="flex align-items-center p-button-secondary" label="Cancel" @click="$emit('closeDialog')" />
      </div>

    </div>
  </Dialog>
</template>

<script>
import axios from 'axios'

export default {
  props: [
    'display'
  ],
  emits: [
    'closeDialog'
  ],
  data() {
    return {
      // text input
      isLoading: 0,
      product_name: '',
      product_value: -1,
      product_description: '',

      // text input class
      product_name_class: '',
      product_value_class: '',
      product_description_class: '',
      message_show: false,
      message: '',
      message_color: 'bg-red-400',

      spinner_computed: '',
    }
  },
  methods: {
    checkIntegrity() {
      if(this.product_name == '' || this.product_value == '' || this.product_description == '')
        this.message = "Field is empty!";
      else
        return 0

      if(this.product_name == '') {
        this.product_name_class = 'p-invalid'
      }else{
        this.product_name_class = ''
      }

      if(this.product_value == '') {
        this.product_value_class = 'p-invalid'
      }else {
        this.product_value_class = ''
      }

      if(this.product_description == '') {
        this.product_description_class = 'p-invalid'
      }else {
        this.product_description_class = ''
      }

      this.message_show = true;
      return 1
    },
    postData() {
      // change login icon
      this.spinner_computed = "pi pi-spin pi-spinner";

      // get backend url 
      let url = import.meta.env.VITE_BACKEND_URL

      let data = {
        product_name: this.product_name,
        product_value: parseInt(this.product_value),
        product_description: this.product_description
      }

      // fetch api
      axios.post(url + "/v1/api/product/create_product", data)
      .then(resp => {
        console.log(resp.data)
        if(resp.data["status_type"] == 0) {
          // close dialog
          this.$emit('closeDialog')

          // clear input
          this.clearInputs()

          // get created product_id
          data["product_id"] = resp.data["product_id"]

          // send data to vuex state
          this.$store.commit('setNewData', data)
        }
      })
      .catch(function(err) {
        console.log(err)
      });

    },
    clearInputs() {
      this.product_name = ''
      this.product_value = -1
      this.product_description = ''
    },
    onSubmit() {
      if(!this.checkIntegrity()){
        this.postData()
      }else{
        return
      }
    }
  }

}
</script>

