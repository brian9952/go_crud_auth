<script setup>
import Button from "/node_modules/primevue/button";
</script>

<template>
  <div class="card">
    <div class="flex pt-5 text-blue-100 align-items-center justify-content-center text-xl font-semibold font-roboto">
      Product Management
    </div>
  </div>

  <div class="flex flex-row-reverse flex-wrap card-container pr-6" v-show="">
    <Button class="p-button-primary" label="Add" icon="pi pi-plus" @click="addProduct()" v-if="addButtonVisible"/>
  </div>

  <AddProductDialog :display="add_product_visible" @hide="add_product_visible = false" @closeDialog="add_product_visible = false" />
</template>

<script>
import AddProductDialog from "../products/dialogs/AddProductDialog.vue"

export default {
  components: {
    AddProductDialog
  },
  data() {
    return {
      add_product_visible: false,
    }
  },
  computed: {
    addButtonVisible() {
      if(this.$store.getters.authenticated) {
        return true
      }
      return false
    }
  },
  methods: {
    addProduct() {
      if(this.add_product_visible == true) {
        this.add_product_visible = false
        return
      }
      this.add_product_visible = true
    }
  }
}
</script>

<style>
.font-roboto {
    font-family: 'Roboto', sans-serif;
  }
</style>

