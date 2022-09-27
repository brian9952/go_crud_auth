<script setup>
import Dialog from "/node_modules/primevue/dialog"
import Button from "/node_modules/primevue/button"
</script>

<template>
<Dialog v-model:visible="display" v-bind:modal="true">
  <div class="flex flex-column mx-6">
    <h3>Are you sure ?</h3>
  </div>
  <div class="flex align-items-center justify-content-center gap-3 mb-4">
    <Button class="p-button-success" label="Yes" @click="deleteProduct"/>
    <Button class="p-button-danger" label="Cancel" @click="$emit('closeDialog')" />
  </div>
</Dialog>
</template>

<script>
import axios from 'axios'

export default {
    props: [
      'display',
      'product',
    ],
    emits: [
      'delProd',
      'closeDialog'
    ],
    methods: {
      deleteProduct() {
        let url = import.meta.env.VITE_BACKEND_URL

        axios.post(url + "/v1/api/product/delete/" + this.product)
          .then(resp => {
            if(resp.data["status_type"] == 0){
              this.$emit('delProd', this.product)
              this.$emit('closeDialog')
            }
          })
          .catch(function(error) {
            console.log(error)
          })
      }
    }

}
</script>
