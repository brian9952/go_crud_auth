<script setup>
import Dialog from "/node_modules/primevue/dialog"
import Button from "/node_modules/primevue/button"
import InputText from "/node_modules/primevue/inputtext"

</script>

<template>
  <Dialog v-model:visible="display" v-bind:modal="true">
    <div class="flex flex-column mx-8">
      <h2 class="flex">Edit Product</h2>
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
        <Button class="flex align-items-center p-button-primary" label="Edit" @click="onEdit()" />
        <Button class="flex align-items-center p-button-primary" label="Cancel" @click="$emit('hide')" />
      </div>

    </div>
  </Dialog>
</template>

<script>
import axios from "axios"

export default {
  props: [
    'display',
    'product'
  ],
  data() {
    return {
      message: '',
      message_show: false,
      message_color: 'bg-red-400',
      product_id: -1,
      product_name: '',
      product_name_class: '',
      product_value: -1,
      product_value_class: '',
      product_description: '',
      product_description_class: ''
    }
  },
  watch: {
    product(newId) {
      // fetch product id
      let url = import.meta.env.VITE_BACKEND_URL
      axios.get(url + "/v1/api/product/show/" + newId)
        .then(resp => {
          if(resp.data["product_id"] != -1) {
            this.product_name = resp.data["product_name"]
            this.product_value = resp.data["product_value"]
            this.product_description = resp.data["product_description"]
          }
        })
        .catch(function(error) {
          console.log(error.toJSON())
        })
    }
  },
  methods: {
    checkIntegrity() {
      if(this.product_name == '' || this.product_value == -1 || this.product_description == '') {
        this.message = 'required field is empty!'
      }else {
        return 1
      }

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
      } else {
        this.product_description_class = ''
      }

      return 0

    },
    onEdit() {

      if(checkIntegrity()) {
        let url = import.meta.env.VITE_BACKEND_URL
        let data = {
          product_id: this.product_id,
          product_name: this.product_name,
          product_value: this.product_value,
          product_description: this.product_description
        }
        axios.post(url + "/v1/api/product/edit_product", data)
          .then(resp => {

          })
          .catch(function(error) {
            console.log(error)
          })
      }

    }
  }
}

</script>
