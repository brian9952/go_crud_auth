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
  <ProductDialog :display="showIsVisible" @hide="showIsVisible = false" @closeDialog="showIsVisible = false"  @interface="getChildInterface">
  </ProductDialog>

  <!-- delete dialog -->
  <DeleteDialog :product="delProductId" :display="deleteIsVisible" @hide="deleteIsVisible = false" @delProd="deleteRow" @closeDialog="deleteIsVisible = false" />

  <!-- edit dialog -->
  <EditDialog :product="editProductId" :display="editIsVisible" @hide="editIsVisible = false" @closeDialog="editIsVisible = false" />

</template>

<script>
import ProductDialog from "./dialogs/ShowProductDialog.vue"
import DeleteDialog from "./dialogs/DeleteProductDialog.vue"
import EditDialog from "./dialogs/EditProductDialog.vue"

export default {
    components: {
      ProductDialog,
      DeleteDialog,
      EditDialog
    },
    childInterface: {
      fetchProduct: () => {}
    },
    data() {
      return {
        columns: null,
        products: null,
        errorMessage: null,
        showIsVisible: false,
        deleteIsVisible: false,
        editIsVisible: false,
        delProductId: -1,
        editProductId: -1
      }
    },
    watch: {
      '$store.state.newData': function() {
        let newData = this.$store.state.newData
        newData["num"] = Object.keys(this.products).length + 1;
        this.products.push(newData)
      },
      '$store.state.editedData': function() {
        // search index
        let editedData = this.$store.state.editedData
        var idx = this.products.findIndex(function(item, i) {
            return item.product_id === editedData.product_id
        })

        // change data on index 
        this.products[idx].product_name = editedData.product_name
        this.products[idx].product_value = editedData.product_value
        this.products[idx].product_description = editedData.product_description

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
            // set data
            this.products = resp.data;
          })
          .catch(function(error) {
            console.log(error.toJSON());
          })
      },
      deleteProduct(product_id) {
        this.delProductId = product_id
        if(this.deleteIsVisible == false) {
          this.deleteIsVisible = true
        }
      },
      deleteRow(product_id) {
        // remove table
        this.products = this.products.filter(function(product) {
          return product.product_id !== product_id
        });

        // update number
        for(var i = 0; i < Object.keys(this.products).length; i++) {
          this.products[i].num = i + 1;
        }
      },

      // dialog interaction
      getChildInterface(childInterface) {
        this.$options.childInterface = childInterface;
      },
      Product(product_name, product_value, product_description) {
        this.products.product_name = product_name
        this.products.product_value = product_value
        this.products.product_description = product_description
      },
      
      // interface communcation
      // show product
      fetchProduct(productId) {
        if(this.showIsVisible == false) {
          this.showIsVisible = true;
        }
        this.$options.childInterface.fetchProduct(productId);
      },

      // edit product
      editProduct(productId) {
        this.editProductId = productId
        if(this.editIsVisible == false) {
          this.editIsVisible = true;
        }

      }
    }
  }

</script>

<style>

</style>
