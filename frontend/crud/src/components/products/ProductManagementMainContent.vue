<script setup>
import DataTable from "/node_modules/primevue/datatable";
import Column from "/node_modules/primevue/column";
import Button from "/node_modules/primevue/button";
import axios from 'axios';
</script>

<template>

  <!-- table -->
  <div class="flex flex-column pt-3 pl-6 pr-6">
    <div class="align-items-center justify-content-center">
      <DataTable :value="products" responsiveLayout="scroll">
        <Column header="No" field="num"></Column>
        <Column header="Product Name" field="product_name"></Column>
        <Column header="Price" field="product_value"></Column>
        <Column header="Description" field="product_description"></Column>
        <!-- <Column v-for="col of columns" :field="col.field" :header="col.header" :key="col.field"></Column> -->
        <Column>
          <template #body="buttonProps">
            <div class="flex justify-content-center flex-wrap gap-3">
              <Button class="flex align-items-center justify-content-center p-button-warning" @click="fetchProduct(buttonProps.data.product_id)">Show</Button>
              <Button class="flex align-items-center justify-content-center p-button-info" v-if="checkAuth" @click="editProduct(buttonProps.data.product_id)">Edit</Button>
              <Button class="flex align-items-center justify-content-center p-button-danger" v-if="checkAuth" @click="deleteProduct(buttonProps.data.product_id)">Delete</Button>
            </div>
          </template>
        </Column>
      </DataTable>
    </div>
  </div>

  <!-- dialog -->
  <ProductDialog :display="isVisible" @hide="isVisible = false" @closeDialog="isVisible = false"  @interface="getChildInterface">
  </ProductDialog>

</template>

<script>
import ProductDialog from "./dialogs/ShowProductDialog.vue"

export default {
    components: {
      ProductDialog
    },
    childInterface: {
      fetchProduct: () => {}
    },
    data() {
      return {
        columns: null,
        products: null,
        errorMessage: null,
        isVisible: false
        }
    },
    computed :{
      checkAuth() {
        if (this.$store.getters.authenticated) {
          return true
        }
        return false
      }
    },
    created() {
      this.columns = [
        {field: 'num', header: 'No'},
        {field: 'product_name', header: 'Product Name'},
        {field: 'product_value', header: 'Product Value'},
        {field: 'product_description', header: 'Product Description'},
      ];

      // async functions
      this.$watch (
        () => this.$route.params,
        () => {
          this.fetchData();
        },

        { immediate: true }
      )

    },
    methods: {
      fetchData() {
        let url = import.meta.env.VITE_BACKEND_URL
        axios.get(url + "/v1/api/product/show_products")
          .then(resp => {
            // get numbering
            for (var i = 0; i < Object.keys(resp.data).length; i++) {
              resp.data[i].num = i + 1;
            }
            console.log(resp.data);
            // set data
            this.products = resp.data;
          })
          .catch(function(error) {
            console.log(error.toJSON());
          })
      },
      showProduct(product_id) {
        console.log("Product show = " + product_id);
      },

      // dialog interaction
      getChildInterface(childInterface) {
        this.$options.childInterface = childInterface;
      },
      
      // interface communcation
      // show product
      fetchProduct(productId) {
        if(this.isVisible == false) {
          this.isVisible = true;
        }
        this.$options.childInterface.fetchProduct(productId);
      },

      // edit product
      editProduct(productId) {

      }
    }
  }

</script>

<style>

</style>
